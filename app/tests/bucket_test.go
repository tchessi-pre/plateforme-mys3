package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"mys3/handlers"
)

func TestCreateBucket(t *testing.T) {
	req, err := http.NewRequest("GET", "/create-bucket?bucket=testbucket", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateBucket)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
