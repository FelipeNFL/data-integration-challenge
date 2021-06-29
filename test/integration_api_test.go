package test

import (
	"os"
	"strings"
	"strconv"
	"testing"
	"net/http"
	"yawoen/core"
)

const HOST = "http://localhost:8080"

func setup() {
	core.CleanDb()
	core.InitializeDb()
	row := core.Company{Name: "TEST", Zipcode: "12345"}
	db, _ := core.GetDb()
	db.Create(&row)
}

func shutdown() {
	core.CleanDb()
}

func TestMain(m *testing.M) {
    setup()
    code := m.Run() 
    shutdown()
    os.Exit(code)
}

func DoPut(url string, data *strings.Reader, t *testing.T) (*http.Response) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)

	if err != nil {
		t.Errorf(err.Error())
	}

	resp, err := client.Do(req)

	if err != nil {
		t.Errorf(err.Error())
	}

	return resp
}

func TestUpdateCompanyExisting(t *testing.T) {
	url := HOST + "/company"
	data := strings.NewReader("test;12345;www.google.com")
	resp := DoPut(url, data, t)

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}

	company := core.Company{}
	db, err := core.GetDb()

	if err != nil {
		t.Errorf(err.Error())
	}

	db.Where("zipcode = '12345' AND name = 'TEST'").Find(&company)

	if company.Website != "www.google.com" {
		t.Errorf("website wrong: "+company.Website)
	}
}

func TestUpdateCompanyExistingWithUpperCaseWebSite(t *testing.T) {
	url := HOST + "/company"
	data := strings.NewReader("test;12345;WWW.GOOGLE.COM")
	resp := DoPut(url, data, t)

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}

	company := core.Company{}
	db, err := core.GetDb()

	if err != nil {
		t.Errorf(err.Error())
	}

	db.Where("zipcode = '12345' AND name = 'TEST'").Find(&company)

	if company.Website != "www.google.com" {
		t.Errorf("website wrong: "+company.Website)
	}
}

func TestUpdateCompanyNotFound(t *testing.T) {
	url := HOST + "/company"
	data := strings.NewReader("wrong;12345;WWW.GOOGLE.COM")
	resp := DoPut(url, data, t)

	if resp.StatusCode != 304 {
		t.Errorf("status code wrong: "+strconv.Itoa(resp.StatusCode))
	}
}

func TestHealthIntegrationApi(t *testing.T) {
	url := HOST + "/health"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}
}