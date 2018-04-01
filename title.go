package main

import "fmt"

func AdvertTitle(advert Advert) string {
	return fmt.Sprintf("%d-комнатная квартира, %d м², %s", advert.Room, advert.Square, advert.Address)
}
