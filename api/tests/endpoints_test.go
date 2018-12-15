package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mungujn/weather-server/api/backend"
)

// TestBadRequest: test the location endpoint/handler returns bad request if no location/date specified
func TestBadRequest(t *testing.T) {
	expectedCode := http.StatusBadRequest
	body := testEndpoint(t, "GET", "/weather?location=kampala", expectedCode)

	expectedBody := `{"message":"No location/date specified"}`

	if body != expectedBody {
		t.Errorf("Handler returned wrong body: got %v instead of %v", body, expectedBody)
	}
}

// function type
type fn func(w http.ResponseWriter, r *http.Request)

// testEndpoint : Helper function to reduce test verbosity
func testEndpoint(t *testing.T, method, url string, expectedCode int) string {
	// create request to pass to handler
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// create response recoder
	recorder := httptest.NewRecorder()

	// get router
	router := backend.CreateRouter()

	// serve
	router.ServeHTTP(recorder, request)

	// confirm status code
	code := recorder.Code
	if code != expectedCode {
		t.Errorf("Handler returned wrong code: got %v instead of %v", code, expectedCode)
	}

	// return body string
	return recorder.Body.String()
}
