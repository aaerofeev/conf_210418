package main

import "testing"

func TestAdvertTitle(t *testing.T) {
	input := Advert{Room: 1, Square: 45, Address: "Фурманова - Гоголя"}
	expected := "1-комнатная квартира, 45 м², Фурманова - Гоголя"

	title := AdvertTitle(input)

	if title != expected {
		t.Fatalf("Ожидается %s, получено %s", expected, title)
	}
}
