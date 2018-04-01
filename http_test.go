package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
)

func TestHTTP(t *testing.T) {
	expected := "<h1>3-комнатная квартира, 35 м², Абая - Жарокова</h1><p>Чисто, есть кондиционер</p>"
	server := httptest.NewServer(mux())
	defer server.Close()
	
	r, err := http.Get(server.URL + "/show?id=1")
	
	if err != nil {
		t.Fatalf("Сервер вернул ошибку: %s", err)
	}
	
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	
	if expected != string(body) {
		t.Fatalf("Ожидиется строка %s, получено %s", expected, string(body))
	}
}
