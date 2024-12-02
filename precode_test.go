package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenValidRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	actual := responseRecorder.Body.String()
	assert.NotEmpty(t, actual)
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/?count=2&city=amsterdam", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	actual := responseRecorder.Body.String()
	expected := "wrong city value"
	assert.Equal(t, expected, actual)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/?count=8&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	// require.Equal(t, http.StatusOK, responseRecorder.Code)

	expected := strings.Join(cafeList["moscow"], ",")
	actual := responseRecorder.Body.String()
	cafes := strings.Split(actual, ",")
	assert.Len(t, cafes, totalCount)

	assert.Equal(t, expected, actual)
}
