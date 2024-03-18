package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
    "moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
    countStr := req.URL.Query().Get("count")
    if countStr == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("count missing"))
        return
    }

    count, err := strconv.Atoi(countStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong count value"))
        return
    }

    city := req.URL.Query().Get("city")

    cafe, ok := cafeList[city]
    if !ok {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("wrong city value"))
        return
    }

    if count > len(cafe) {
        count = len(cafe)
    }

    answer := strings.Join(cafe[:count], ",")

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(answer))
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler:= http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, responseRecorder.Code, http.StatusOK)
    assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityNotSupport(t *testing.T) {
    bodyResponse := `wrong city value`

    req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=Tula", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, responseRecorder.Body, bodyResponse)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4
    
    req := httptest.NewRequest(http.MethodGet, "/cafe?count=9&city=moscow", nil)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.NotEqual(t, responseRecorder.Code, http.StatusOK)
    assert.Len(t, responseRecorder.Body, totalCount)
}
