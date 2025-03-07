package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseResponseHandler(t *testing.T) {
	w := httptest.NewRecorder()
	handler := &BaseResponseHandler{}
	handler.Handle(w, "Hello, World!", http.StatusOK)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Success {
		t.Errorf("Expected result to be '%s', got '%v'", Success, response["result"])
	}

	if response["data"] != "Hello, World!" {
		t.Errorf("Expected data to be 'Hello, World!', got '%v'", response["data"])
	}
}

func TestSuccessResponseDecorator(t *testing.T) {
	w := httptest.NewRecorder()
	decorator := &SuccessResponseDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, "Hello, World!", http.StatusOK)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be '*', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Success {
		t.Errorf("Expected result to be '%s', got '%v'", Success, response["result"])
	}

	if response["data"] != "Hello, World!" {
		t.Errorf("Expected data to be 'Hello, World!', got '%v'", response["data"])
	}
}

func TestErrorResponseDecorator(t *testing.T) {
	w := httptest.NewRecorder()
	decorator := &ErrorResponseDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, "Error Message", http.StatusInternalServerError)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Failure {
		t.Errorf("Expected result to be '%s', got '%v'", Failure, response["result"])
	}

	if response["data"].(map[string]interface{})["message"] != "Error Message" {
		t.Errorf("Expected data to be {message: 'Error Message'}, got '%v'", response["data"])
	}
}

func TestSuccessResponseListDecorator(t *testing.T) {
	w := httptest.NewRecorder()
	decorator := &SuccessResponseListDecorator{&BaseResponseHandler{}}
	decorator.Handle(w, []string{"Item1", "Item2"}, "0", "10", "2")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Success {
		t.Errorf("Expected result to be '%s', got '%v'", Success, response["result"])
	}

	if response["data"].(map[string]interface{})["count"] != "2" {
		t.Errorf("Expected count to be '2', got '%v'", response["data"].(map[string]interface{})["count"])
	}

	if response["data"].(map[string]interface{})["offset"] != "0" {
		t.Errorf("Expected offset to be '0', got '%v'", response["data"].(map[string]interface{})["offset"])
	}

	if response["data"].(map[string]interface{})["limit"] != "10" {
		t.Errorf("Expected limit to be '10', got '%v'", response["data"].(map[string]interface{})["limit"])
	}

	if response["data"].(map[string]interface{})["list"].([]interface{})[0] != "Item1" {
		t.Errorf("Expected first item to be 'Item1', got '%v'", response["data"].(map[string]interface{})["list"].([]interface{})[0])
	}

	if response["data"].(map[string]interface{})["list"].([]interface{})[1] != "Item2" {
		t.Errorf("Expected second item to be 'Item2', got '%v'", response["data"].(map[string]interface{})["list"].([]interface{})[1])
	}
}

func TestErrorResponseHelper(t *testing.T) {
	w := httptest.NewRecorder()
	ErrorResponseHelper("exampleMethod", w, "Error Message", http.StatusInternalServerError)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Failure {
		t.Errorf("Expected result to be '%s', got '%v'", Failure, response["result"])
	}

	if response["data"].(map[string]interface{})["message"] != "Error Message" {
		t.Errorf("Expected data to be {message: 'Error Message'}, got '%v'", response["data"])
	}
}

func TestSuccessResponseHelper(t *testing.T) {
	w := httptest.NewRecorder()
	SuccessResponseHelper(w, "Hello, World!", http.StatusOK)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected Access-Control-Allow-Origin header to be '*', got '%s'", w.Header().Get("Access-Control-Allow-Origin"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Success {
		t.Errorf("Expected result to be '%s', got '%v'", Success, response["result"])
	}

	if response["data"] != "Hello, World!" {
		t.Errorf("Expected data to be 'Hello, World!', got '%v'", response["data"])
	}
}

func TestSuccessResponseList(t *testing.T) {
	w := httptest.NewRecorder()
	SuccessResponseList(w, []string{"Item1", "Item2"}, "0", "10", "2")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json;charset=utf-8" {
		t.Errorf("Expected Content-Type header to be 'application/json;charset=utf-8', got '%s'", w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response["result"] != Success {
		t.Errorf("Expected result to be '%s', got '%v'", Success, response["result"])
	}

	if response["data"].(map[string]interface{})["count"] != "2" {
		t.Errorf("Expected count to be '2', got '%v'", response["data"].(map[string]interface{})["count"])
	}

	if response["data"].(map[string]interface{})["offset"] != "0" {
		t.Errorf("Expected offset to be '0', got '%v'", response["data"].(map[string]interface{})["offset"])
	}

	if response["data"].(map[string]interface{})["limit"] != "10" {
		t.Errorf("Expected limit to be '10', got '%v'", response["data"].(map[string]interface{})["limit"])
	}

	if response["data"].(map[string]interface{})["list"].([]interface{})[0] != "Item1" {
		t.Errorf("Expected first item to be 'Item1', got '%v'", response["data"].(map[string]interface{})["list"].([]interface{})[0])
	}

	if response["data"].(map[string]interface{})["list"].([]interface{})[1] != "Item2" {
		t.Errorf("Expected second item to be 'Item2', got '%v'", response["data"].(map[string]interface{})["list"].([]interface{})[1])
	}
}
