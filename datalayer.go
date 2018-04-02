package main

type Sql struct{}

func (s Sql) Execute(code string, args ...interface{}) Advert {
	return Advert{Room: 3, Square: 25, Address: "Абая - Жарокова", Text: "Чисто, есть кондиционер"}
}

var db Sql

var dataLayer = DataLayer{}
