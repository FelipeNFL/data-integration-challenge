package services

import (
	"log"
	"bytes"
	"strings"
	"net/http"
	"yawoen/core"
	"gorm.io/gorm"
)

const separatorCsv = ';'

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
    db := r.Context().Value("DB").(*gorm.DB)

	buffer := new(bytes.Buffer)
    buffer.ReadFrom(r.Body)
    csvContent := buffer.String()

	reader := core.GetCsvReaderFromString(csvContent, separatorCsv)
	
	core.IterateCsv(reader, func(record []string) {
		name := strings.ToUpper(record[0])
		zipcode := strings.ToUpper(record[1])
		result := db.Model(&core.Company{}).Where("name = ? AND zipcode = ?", name, zipcode).Update("website", record[2])

		if result.RowsAffected == 0 {
			w.WriteHeader(http.StatusNotModified)
		} else {
			log.Println("Company "+name+" just have your website updated!")
		}
	})
}