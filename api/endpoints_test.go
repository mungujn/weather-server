package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestWeatherHandler: test the location endpoint/handler
func TestWeatherHandler(t *testing.T) {
	expectedCode := http.StatusOK
	expectedBody := `{"location":"kampala","date":"today","temparature":28}`
	testEndpoint(t, "GET", "/weather?location=kampala&date=today", expectedCode, expectedBody)
}

// TestBadRequest: test the location endpoint/handler returns bad request if no location/date specified
func TestBadRequest(t *testing.T) {
	expectedCode := http.StatusBadRequest
	expectedBody := `{"message":"No location/date specified"}`
	testEndpoint(t, "GET", "/weather?location=kampala", expectedCode, expectedBody)
}

// function type
type fn func(w http.ResponseWriter, r *http.Request)

// testEndpoint : Helper function to reduce test verbosity
func testEndpoint(t *testing.T, method, url string, expectedCode int, expectedBody string) {
	// create request to pass to handler
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create response recoder
	recorder := httptest.NewRecorder()

	// get router
	router := createRouter()

	// serve
	router.ServeHTTP(recorder, request)

	// confirm status code
	code := recorder.Code
	if code != expectedCode {
		t.Errorf("Handler returned wrong code: got %v instead of %v", code, expectedCode)
	}

	//confirm body
	body := recorder.Body.String()
	if body != expectedBody {
		t.Errorf("Handler returned wrong body: got %v instead of %v", body, expectedBody)
	}
}
