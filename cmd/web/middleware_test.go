package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"snippetbox.umaralfaruq/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(response, request)

	result := response.Result()
	defer result.Body.Close()

	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, result.Header.Get("Content-Security-Policy"), expectedValue)

	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, result.Header.Get("Referrer-Policy"), expectedValue)

	expectedValue = "nosniff"
	assert.Equal(t, result.Header.Get("X-Content-Type-Options"), expectedValue)

	expectedValue = "deny"
	assert.Equal(t, result.Header.Get("X-Frame-Options"), expectedValue)

	expectedValue = "0"
	assert.Equal(t, result.Header.Get("X-XSS-Protection"), expectedValue)

	assert.Equal(t, result.StatusCode, http.StatusOK)

	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace([]byte(body))

	assert.Equal(t, string(body), "OK")
}
