// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWeatherByLatLongAndDateHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/weatherByLatLongAndDate/34.0836851/-117.7432015/1497944369", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(WeatherByLatLongAndDate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `currently`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// New test
	req, err = http.NewRequest("GET", "/weatherByLatLongAndDate/34.0836851/-117.7432015/fail", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(WeatherByLatLongAndDate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected = `invalid syntax`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
