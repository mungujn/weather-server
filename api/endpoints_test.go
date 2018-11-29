package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// function type
type fn func(w http.ResponseWriter, r *http.Request)

// accessEndpoint : Helper function to reduce test verbosity
func accessEndpoint(t *testing.T, method, url string, handlerFunc fn) (code int, body string) {
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

	return recorder.Code, recorder.Body.String()
}

// test the location endpoint/handler
func testLocationHandler(t *testing.T) {

	code, body := accessEndpoint(t, "GET", "/api/v1/locations/kampala", LocationHandler)

	// confirm status code
	expectedCode := http.StatusOK
	if code != expectedCode {
		t.Errorf("Handler returned wrong code: got %v instead of %v", code, expectedCode)
	}

	//confirm body
	expectedBody := `{"message": "success"}`
	if body != expectedBody {
		t.Errorf("Handler returned wrong body: got %v instead of %v", body, expectedBody)
	}
}
