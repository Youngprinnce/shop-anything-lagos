package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProductHandler(t *testing.T) {
	t.Run("CreateProduct", func(t *testing.T) {
		body := `{"sku": "1234abcd","name": "All around the world","description": "demo description","price": 100}`
		req, err := http.NewRequest("POST", "/api/products/1", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateProduct)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
	t.Run("GetAllProducts", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/products/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetAllProducts)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
	t.Run("GetProduct", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/products/1/sku1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetProduct)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
	t.Run("UpdateProduct", func(t *testing.T) {
		body := `{"name": "All around the world","description": "demo description","price": 100}`
		req, err := http.NewRequest("PUT", "/api/products/1/sku1", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(UpdateProduct)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	// t.Run("DeleteProduct", func(t *testing.T) {
	// 	req, err := http.NewRequest("DELETE", "/api/products/1/sku1", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	rr := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(DeleteProduct)
	// 	handler.ServeHTTP(rr, req)
	// 	if status := rr.Code; status != http.StatusOK {
	// 		t.Errorf("handler returned wrong status code: got %v want %v",
	// 			status, http.StatusOK)
	// 	}
	// })
}