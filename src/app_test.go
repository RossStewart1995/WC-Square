package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"
)

type TestData struct {
	Error		bool
	String		string
	Answer		int
	Author		string
}

func TestSquare(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/?x=18", nil)
	if err != nil {
		t.Fatal(err)
	}else{
		fmt.Println("Found URL - Test 1")
	}

	res := httptest.NewRecorder()
	squareNum(res, req)

	test_data := TestData{false, "18^2=324", 324, "Ross Stewart"}

	js, err := json.Marshal(test_data)
	if err != nil {
		return
	}
	
	exp := string(js)

	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}

func TestStatus200(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/?x=18", nil)
	if err != nil {
		t.Fatal(err)
	}else{
		fmt.Println("Found URL - Test 2")
	}

	res := httptest.NewRecorder()
	squareNum(res, req)
	
	exp := 200

	act := res.Code
	if exp != act {
		t.Errorf("handler returned wrong status code: got %v want %v", exp, act)
	}
}

func TestStatus400(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/?x=", nil)
	if err != nil {
		t.Fatal(err)
	}else{
		fmt.Println("Found URL - Test 3")
	}

	res := httptest.NewRecorder()
	squareNum(res, req)
	
	exp := 400

	act := res.Code
	if exp != act {
		t.Errorf("handler returned wrong status code: got %v want %v", exp, act)
	}
}



