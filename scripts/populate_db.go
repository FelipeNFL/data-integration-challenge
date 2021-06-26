package main

import (
	"io"
    "os"
    "fmt"
	"log"
	"strings"
    "encoding/csv"
	"yawoen/core"
	"gorm.io/gorm"
)

const fileNameData = "data/q1_catalog.csv"
const separator = ';'

func getCsv(filename string, separator rune) (*csv.Reader, *os.File) {
	csvFile, err := os.Open(filename)
	
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(csvFile)
	reader.Comma = separator
	return reader, csvFile
}

func getDb() (*gorm.DB) {
	db, _ := core.GetDb()
	core.InitializeDb()
	return db
}

func main() {
	db := getDb()
	reader, csvFile := getCsv(fileNameData, separator)

	for {
		record, err := reader.Read()
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			log.Fatal(err)
		}

		row := core.Company{Name: strings.ToUpper(record[0]), Zipcode: record[1]}
		result := db.Create(&row)
	}

	defer csvFile.Close()
}