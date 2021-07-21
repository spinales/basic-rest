package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	responseBody       string
	observedStatusCode int
}

func TestGetProducts(t *testing.T) {
	url := "http://localhost:1323/product"

	tests := []TestStruct{
		{`{}`, 200, `[{"ID":1,"CreatedAt":"2021-07-19T22:24:00.3530518-04:00","UpdatedAt":"2021-07-19T22:24:00.3530518-04:00","DeletedAt":null,"name":"Basic computer","provider":"Prov Inc.","quantity":20,"price":200,"description":"Basic product."}]`, 0},
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("GET", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	CaseResults("GET Products", tests, t)

}

func TestGetProduct(t *testing.T) {
	url := "http://localhost:1323/product/1"

	tests := []TestStruct{
		{`{}`, 200, `{"ID":1,"CreatedAt":"2021-07-19T22:24:00.3530518-04:00","UpdatedAt":"2021-07-19T22:24:00.3530518-04:00","DeletedAt":null,"name":"Basic computer","provider":"Prov Inc.","quantity":20,"price":200,"description":"Basic product."}`, 0},
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("GET", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	CaseResults("GET Product", tests, t)

}

func TestAddProduct(t *testing.T) {
	url := "http://localhost:1323/product"

	tests := []TestStruct{
		{`{"name":"Basic monitor","provider":"Prov Inc.","quantity":5,"price":500,"description":"Basic product."}`, 200, `[{"ID":1,"CreatedAt":"2021-07-19T22:24:00.3530518-04:00","UpdatedAt":"2021-07-19T22:24:00.3530518-04:00","DeletedAt":null,"name":"Basic monitor","provider":"Prov Inc.","quantity":5,"price":500,"description":"Basic product."}]`, 0},
	}

	for i, testCase := range tests {

		var reader io.Reader
		reader = strings.NewReader(testCase.requestBody) //Convert string to reader

		request, err := http.NewRequest("POST", url, reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err) //Something is wrong while sending request
		}
		body, _ := ioutil.ReadAll(res.Body)

		tests[i].responseBody = strings.TrimSpace(string(body))
		tests[i].observedStatusCode = res.StatusCode

	}

	CaseResults("POST Product", tests, t)

}

func CaseResults(functionalityName string, tests []TestStruct, t *testing.T) {

	for _, test := range tests {

		if test.observedStatusCode == test.expectedStatusCode {
			t.Logf("Passed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		} else {
			t.Errorf("Failed Case:\n  request body : %s \n expectedStatus : %d \n responseBody : %s \n observedStatusCode : %d \n", test.requestBody, test.expectedStatusCode, test.responseBody, test.observedStatusCode)
		}
	}
}
