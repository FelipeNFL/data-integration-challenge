package services

import (
	"log"
	"bytes"
	"strings"
	"net/http"
	"encoding/json"	
	"gorm.io/gorm"
	"yawoen/core"
)

const separatorCsv = ';'

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
    db := r.Context().Value("DB").(*gorm.DB)

	ignoreHeader := r.URL.Query().Get("ignore_header") == "true"

	buffer := new(bytes.Buffer)
    buffer.ReadFrom(r.Body)
    csvContent := buffer.String()

	reader := core.GetCsvReaderFromString(csvContent, separatorCsv)

	registersChanged := []core.Company{}
	
	core.IterateCsv(reader, ignoreHeader, func(record []string) {
		name := strings.ToUpper(record[0])
		zipcode := strings.ToUpper(record[1])
		website := strings.ToLower(record[2])

		result := db.Model(&core.Company{}).Where("name = ? AND zipcode = ?", name, zipcode).Update("website", website)

		if result.RowsAffected > 0 {
			companyUpdated := core.Company{}
			db.Model(&core.Company{}).Where("name = ? AND zipcode = ?", name, zipcode).First(&companyUpdated)
			registersChanged = append(registersChanged, companyUpdated)	
			
			log.Println("Company "+name+" just have your website updated!")
		} else {
			log.Println("Company "+name+" not found")
		}
	})

	if len(registersChanged) == 0 {
		w.WriteHeader(http.StatusNotModified)
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registersChanged)
}

func GetCompanyByNameAndZipCode(w http.ResponseWriter, r *http.Request) {
    db := r.Context().Value("DB").(*gorm.DB)

	queryString := r.URL.Query()
	name := queryString.Get("name")
	zipcode := queryString.Get("zipcode")

	likeQueryName := strings.ToUpper(name) + "%"
	company := core.Company{}
	result := db.Where("zipcode = ? AND name LIKE ?", zipcode, likeQueryName).Find(&company)
	
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(company)
	}
}