package main

import (
	"time"
	"errors"
)

var ErrAlreadyProvided = errors.New("время бесплатного продления не наступило")

type Advert struct {
	Room int
	Address string
	Square int
	AddAt time.Time
	Text string
}

func UpFreeAdvert(advert *Advert) error {
	if time.Since(advert.AddAt) > time.Hour * 24 {
		advert.AddAt = time.Now()
		return nil
	}
	
	return ErrAlreadyProvided
}