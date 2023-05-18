package services

import (
	"fmt"
	"github.com/liuhdd/exam-cret/application/config"
	"os"
	"os/exec"
	"strings"
)

func ResumeData() error {
	db := config.GetDB()

	s, err := os.ReadFile("/tmp/exam.sql")
	if err != nil {
		return err
	}

	sql := strings.ReplaceAll(string(s), "CREATE TABLE", "CREATE TABLE IF not EXISTS")
	tx := db.Exec(sql)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func BackupData() error {
	dir, _ := os.Getwd()
	cmd := exec.Command(dir+"/backup.sh", dir)
	str := cmd.String()
	fmt.Println(str)
	return cmd.Run()

}
