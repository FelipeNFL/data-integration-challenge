package test

import (
	"os"
	"testing"
	"net/http"
	"yawoen/core"
)

const HOST = "http://localhost:8081"

func createRow(row core.Company) {
	db, _ := core.GetDb()
	db.Create(&row)
}

func setup() {
	core.CleanDb()
	core.InitializeDb()
	createRow(core.Company{Name: "TEST FULLNAME", Zipcode: "12345"})	
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

func TestGetCompanySuccessfulFullname(t *testing.T) {
	url := HOST + "/company?name=Test%20Fullname&zipcode=12345"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}
}

func TestGetCompanySuccessfulPartialName(t *testing.T) {
	url := HOST + "/company?name=Test&zipcode=12345"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}
}

func TestGetCompanyNotFound(t *testing.T) {
	url := HOST + "/company?name=Wrong&zipcode=123"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != 404 {
		t.Errorf("status code wrong")
	}
}

func TestHealthMatchingApi(t *testing.T) {
	url := HOST + "/health"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("status code wrong")
	}
}

