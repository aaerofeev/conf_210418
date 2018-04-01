package main

import "testing"

type MockLayer struct {}

func (m MockLayer) Find(id string) Advert {
	return Advert{Room: 3, Square: 85, Address: "Абая - Ауэзова", Text: "Метро в шаговой доступности"}
}

func TestAdvertPresenter(t *testing.T) {
	mock := MockLayer{}
	html := AdvertPresenter("1", mock)
	expected := "<h1>3-комнатная квартира, 85 м², Абая - Ауэзова</h1><p>Метро в шаговой доступности</p>"
	
	if html != expected {
		t.Fatalf(`Ожидается "%s", получен "%s"`, expected, html)
	}
}
