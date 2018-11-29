package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPastWeatherHandler: test the location endpoint/handler
func TestPastWeatherHandler(t *testing.T) {
	expectedCode := http.StatusOK
	expectedBody := `{"message":"success"}`
	testEndpoint(t, "GET", "/api/v1/past-weather?location=kampala", PastWeatherHandler, expectedCode, expectedBody)
}

// TestCurrentWeatherHandler: test the location endpoint/handler
func TestCurrentWeatherHandler(t *testing.T) {
	expectedCode := http.StatusOK
	expectedBody := `{"message":"success"}`
	testEndpoint(t, "GET", "/api/v1/current-weather?location=kampala", PastWeatherHandler, expectedCode, expectedBody)
}

// TestFutureWeatherHandler: test the endpoint that retrieves the next days
func TestFutureWeatherHandler(t *testing.T) {
	expectedCode := http.StatusOK
	expectedBody := `{"message":"success"}`
	testEndpoint(t, "GET", "/api/v1/future-weather?location=kampala", PastWeatherHandler, expectedCode, expectedBody)
}

// function type
type fn func(w http.ResponseWriter, r *http.Request)

// testEndpoint : Helper function to reduce test verbosity
func testEndpoint(t *testing.T, method, url string, handlerFunc fn, expectedCode int, expectedBody string) {
	// create request to pass to handler
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create response recoder
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFunc)

	// serve request
	handler.ServeHTTP(recorder, request)

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
