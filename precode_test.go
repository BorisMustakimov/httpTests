package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)


func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    // здесь нужно создать запрос к сервису
    req :=  httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
    
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // Если в параметре `count` указано больше, 
    // чем есть всего, должны вернуться все доступные кафе.
    require.Equal(t, http.StatusOK, responseRecorder.Code)
    list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, list, totalCount) 
}

func TestMainHandlerWhenOk(t *testing.T) {
    // запрос к сервису
    req :=  httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
    
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // здесь нужно добавить необходимые проверки
    // Запрос сформирован корректно, 
    // сервис возвращает код ответа 200 и тело ответа не пустое
    require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenBad(t *testing.T) {
    // запрос к сервису
    req, err := http.NewRequest("GET", "/cafe?city=invalid&count=2", nil)
	require.NoError(t, err)
    
    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    // Город, который передаётся в параметре `city`,
    // не поддерживается. 
    // Сервис возвращает код ответа 400 и ошибку `wrong city value` в теле ответа.
    require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}