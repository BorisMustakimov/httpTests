package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=11&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("expected status ok, got: %d", responseRecorder.Code)
	}

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))

}
