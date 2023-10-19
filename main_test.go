package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLoginHandler(t *testing.T) {
    req, err := http.NewRequest("POST", "/login", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(loginHandler)

    handler.ServeHTTP(rr, req)

    expectedStatus := http.StatusFound
    if rr.Code != expectedStatus {
        t.Errorf("Expected HTTP status %d, but got %d", expectedStatus, rr.Code)
    }

    expectedLocation := "/"
    location, _ := rr.Result().Location()
    if location.String() != expectedLocation {
        t.Errorf("Expected redirect location: %s, but got: %s", expectedLocation, location)
    }
}

func TestIsValidLogin(t *testing.T) {
    validUsername := "admin"
    validPassword := "admin"
    if !isValidLogin(validUsername, validPassword) {
        t.Errorf("Expected isValidLogin(%s, %s) to be true, but got false", validUsername, validPassword)
    }

    invalidUsername := "user"
    invalidPassword := "password"
    if isValidLogin(invalidUsername, invalidPassword) {
        t.Errorf("Expected isValidLogin(%s, %s) to be false, but got true", invalidUsername, invalidPassword)
    }
}
