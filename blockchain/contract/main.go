package main

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(ActionContract)
	contract.TransactionContextHandler = new(ActionContext)
	contract.Name = "test.test"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)
	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err))
	}

	chaincode.Info.Title = "ExamCertificateChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()
	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err))
	}
}
