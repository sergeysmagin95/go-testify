package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenResponseIsCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafes?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafes?city=london&count=3", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	city := "moscow"
	totalCount := len(cafeList[city])
	req := httptest.NewRequest(http.MethodGet, "/cafes?city=moscow&count=100", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, strings.Join(cafeList["moscow"][:totalCount], ","), responseRecorder.Body.String())
}
