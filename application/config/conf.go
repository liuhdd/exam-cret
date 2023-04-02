package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)



var v = viper.New()



func LoadConfig() {
	
	v.SetConfigType("yaml")
	v.SetConfigFile(getCurrentAbPathByCaller()+"/../app.yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
}



func GetRootPath() (root string) {
	return getCurrentAbPath()+"/../"
}

func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir,getTmpDir())  {
		return getCurrentAbPathByCaller()
	}
	return dir
}

func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}


func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}