package main

import (
	"fmt"
)

type DataLayer struct{}

func (d DataLayer) Find(id string) Advert {
	return db.Execute("SELECT * FROM adverts WHERE id = %s", id)
}

type Finder interface {
	Find(id string) Advert
}

func AdvertPresenter(id string, finder Finder) string {
	advert := finder.Find(id)

	return fmt.Sprintf("<h1>%s</h1><p>%s</p>", AdvertTitle(advert), advert.Text)
}
