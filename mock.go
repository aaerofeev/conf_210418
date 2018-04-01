package main

import (
	"fmt"
)

type Finder interface {
	Find(id string) Advert
}

// Слой базы данных
type DataLayer struct {}

// Получает объявление из базы данных
func (d DataLayer) Find(id string) Advert {
	return db.Execute("SELECT * FROM adverts WHERE id = %s", id)
}

// Получает HTML представление объявления
func AdvertPresenter(id string, finder Finder) string {
	advert := finder.Find(id)
	
	return fmt.Sprintf("<h1>%s</h1><p>%s</p>", AdvertTitle(advert), advert.Text)
}
