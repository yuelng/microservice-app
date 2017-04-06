package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"api/services"
)

func add_account(name string, tel string, pass string, team_type int) {
	db := services.InitDB()
	defer db.Close()

	if db.Where("tel = ?", tel).First(&Account{}).RecordNotFound() {
		// account
		account := Account{Tel: tel, Password: pass}
		db.NewRecord(account)
		db.Create(&account)
		// health
	}
}

func InitAccount() {
	p := fmt.Println
	pwd, err := os.Getwd()
	if err != nil {
		p(err)
		os.Exit(1)
	}
	lines := readCsv(filepath.Join(pwd, "src/api/utils/contact_list.csv"))
	for _, value := range lines {
		password := "password"
		name, tel, team_type := parseLine(value)
		add_account(name, tel, password, team_type)
	}
}

func readCsv(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()
	lines := [][]string{}
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		lines = append(lines, record)
	}
	return lines

}

func parseLine(line []string) (name string, tel string, team_type int) {
	i, err := strconv.Atoi(line[0])
	if err != nil {
		fmt.Println(err)
	}
	name, tel, team_type = line[3], line[6], i
	return
}
