package config

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path"
	"time"
)

type config struct {
	MspID         string `mapstructure:"mspID"`
	CryptoPath    string `mapstructure:"cryptoPath"`
	CertPath      string `mapstructure:"certPath"`
	KeyPath       string `mapstructure:"keyPath"`
	TlsCertPath   string `mapstructure:"tlsCertPath"`
	PeerEndpoint  string `mapstructure:"peerEndpoint"`
	GatewayPeer   string `mapstructure:"gatewayPeer"`
	ChaincodeName string `mapstructure:"chaincodeName"`
	ChannelName   string `mapstructure:"channelName"`
}

var contract *client.Contract
var c *config

func GetContract() *client.Contract {
	if contract != nil {
		return contract
	}
	initConfig()
	clientConnection := newGrpcConnection()
	id := newIdentity()
	sign := newSign()
	gw, err := client.Connect(id, client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	network := gw.GetNetwork(c.ChannelName)
	contract = network.GetContract(c.ChaincodeName)
	return contract
}

func initConfig() {
	if v == nil {
		LoadConfig()
	}
	vb := v.Sub("blockchain")
	c = &config{}
	err := vb.Unmarshal(c)
	if err != nil {
		panic(err)
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		c.ChannelName = cname
	}
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		c.ChaincodeName = ccname
	}
}
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(c.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, c.GatewayPeer)

	connection, err := grpc.Dial(c.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(c.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(c.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

func newSign() identity.Sign {
	files, err := os.ReadDir(c.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(c.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}
