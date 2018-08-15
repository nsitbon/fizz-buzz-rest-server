package main

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
	"reflect"
)

func TestValueIfTrue(t *testing.T) {
	if v := valueIfTrue(true, 1); v != 1 {
		t.Error("Expected 1, got ", v)
	}

	if v := valueIfTrue(false, 1); v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestIsMultipleOf(t *testing.T) {
	tables := []struct {
		value    uint
		multiple uint
		expected bool
	}{
		{15, 1, true},
		{15, 2, false},
		{15, 3, true},
		{15, 4, false},
		{15, 5, true},
		{15, 6, false},
		{15, 7, false},
		{15, 8, false},
		{15, 9, false},
		{15, 10, false},
		{15, 11, false},
		{15, 12, false},
		{15, 13, false},
		{15, 14, false},
		{15, 15, true},
	}

	for _, set := range tables {
		if isMultipleOf(set.value, set.multiple) != set.expected {
			t.Errorf("Expected %t, got %t", set.expected, !set.expected)
		}
	}
}

func TestGenerateElement(t *testing.T) {
	tables := []struct {
		value uint
		int1 uint
		int2 uint
		string1 string
		string2 string
		expected string
	}{
		{15, 2, 13, "foo", "bar", "15"},
		{15, 3, 13, "foo", "bar", "foo"},
		{15, 2, 5,  "foo", "bar", "bar"},
		{15, 3, 5,  "foo", "bar", "foobar"},
	}

	for _, set := range tables {
		if actual := generateElement(set.value, set.int1, set.int2, set.string1, set.string2); actual != set.expected {
			t.Errorf("Expected %s, got %s", set.expected, actual)
		}
	}
}

func TestCreateHttpServer(t *testing.T) {
	request := httptest.NewRequest("GET", "/generate?limit=10&int1=2&int2=3&string1=foo&string2=bar", nil)
	response := httptest.NewRecorder()
	createHttpServer().ServeHTTP(response, request)
	expectedStatus := 200
	expectedContentType := "application/json; charset=utf-8"
	expectedJson := `["1","foo","bar","foo","5","foobar","7","foo","bar","foo"]`

	checkStatus(response, expectedStatus, t)
	checkContentType(response, expectedContentType, t)
	checkJson(expectedJson, response, t)
}

func checkJson(expectedJson string, response *httptest.ResponseRecorder, t *testing.T) {
	if !isEqualJson(expectedJson, response.Body.String()) {
		t.Errorf("Expected '%s', got '%s'", expectedJson, response.Body.String())
	}
}

func checkContentType(response *httptest.ResponseRecorder, expectedContentType string, t *testing.T) {
	if ct := response.Header().Get("Content-Type"); ct != expectedContentType {
		t.Errorf("Expected Content-Type '%s', got %s", expectedContentType, ct)
	}
}

func checkStatus(response *httptest.ResponseRecorder, expectedStatus int, t *testing.T) {
	if response.Code != expectedStatus {
		t.Errorf("Expected code %d, got %d", expectedStatus, response.Code)
	}
}

func isEqualJson(left, right string) bool {
	var leftJson, rightJson interface{}
	json.Unmarshal([]byte(left), &leftJson)
	json.Unmarshal([]byte(right), &rightJson)

	return reflect.DeepEqual(leftJson, rightJson)
}