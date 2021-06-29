package core

import (
	"io"
	"os"
	"log"
	"strings"
    "encoding/csv"
)

type functionToBeExecutedOnRow func([]string)

func GetCsvReaderFromFile(file *os.File, separator rune) (*csv.Reader) {
	reader := csv.NewReader(file)
	reader.Comma = separator
	return reader
}

func GetCsvReaderFromString(text string, separator rune) (*csv.Reader) {
	reader := csv.NewReader(strings.NewReader(text))
	reader.Comma = separator
	return reader
}

func IterateCsv(csvReader *csv.Reader, ignoreHeader bool, toDo functionToBeExecutedOnRow) {
	counter := 0

	for {
		counter++

		if ignoreHeader && counter == 1 {
			continue
		}

		record, err := csvReader.Read()
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			log.Fatal(err)
		}

		toDo(record)
	}
}