package main

import (
    "os"
	"log"
	"fmt"
	"strings"
	"yawoen/core"
	"gorm.io/gorm"
)

const fileNameData = "data/q1_catalog.csv"
const separator = ';'

func openFile(filename string) (*os.File) {
	file, err := os.Open(filename)
	
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func getDb() (*gorm.DB) {
	db, _ := core.GetDb()
	core.InitializeDb()
	return db
}

func main() {
	db := getDb()
	file := openFile(fileNameData)
	reader := core.GetCsvReaderFromFile(file, separator)

	core.IterateCsv(reader, func (record []string) {
		row := core.Company{Name: strings.ToUpper(record[0]), Zipcode: record[1]}
		result := db.Create(&row)

		if result.Error == nil {
			fmt.Println(row.Name + " inserted!")
		} else {
			fmt.Println("Error to insert company "+row.Name)
		}
	})

	defer file.Close()
}