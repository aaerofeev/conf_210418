package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	expected := "<h1>3-комнатная квартира, 25 м², Абая - Жарокова</h1><p>Чисто, есть кондиционер</p>"
	server := httptest.NewServer(mux())

	r, _ := http.Get(server.URL + "/show?id=1")
	body, _ := ioutil.ReadAll(r.Body)

	if expected != string(body) {
		t.Fatalf("Ожидается строка %s, получено %s", expected, string(body))
	}
}
