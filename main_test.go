package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestIsValidLogin(t *testing.T) {
    // Kasus pengujian dengan login yang valid
    validUsername := "admin"
    validPassword := "admin"
    if !isValidLogin(validUsername, validPassword) {
        t.Errorf("Expected isValidLogin(%s, %s) to be true, got false", validUsername, validPassword)
    }

    // Kasus pengujian dengan login yang tidak valid
    invalidUsername := "user"
    invalidPassword := "password"
    if isValidLogin(invalidUsername, invalidPassword) {
        t.Errorf("Expected isValidLogin(%s, %s) to be false, got true", invalidUsername, invalidPassword)
    }
}

func TestLoginHandler(t *testing.T) {
    // Membuat permintaan palsu ke handler login
    req, err := http.NewRequest("POST", "/login", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Menguji handler login menggunakan rekaman respons HTTP palsu
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(loginHandler)

    // Menjalankan handler dengan permintaan palsu
    handler.ServeHTTP(rr, req)

    // Memeriksa apakah respons yang diterima adalah respons yang diharapkan
    expectedStatus := http.StatusOK
    if rr.Code != expectedStatus {
        t.Errorf("Expected HTTP status %d, but got %d", expectedStatus, rr.Code)
    }

    // Memeriksa apakah halaman login telah dibuat dengan benar
    expectedBody := "Login"
    if rr.Body.String() != expectedBody {
        t.Errorf("Expected body to contain %s, but got %s", expectedBody, rr.Body.String())
    }
}
