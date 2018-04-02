package main

import (
	"testing"
	"time"
)

func TestUpFreeAdvert(t *testing.T) {
	recentlyTime := time.Now().Add(-time.Hour)
	futureTime := time.Now().Add(time.Hour)

	cases := map[string]struct {
		advert      Advert
		expectAddAt time.Time
		expectError error
	}{
		"success":          {Advert{AddAt: time.Now().Add(-time.Hour * 25)}, time.Now(), nil},
		"already_provided": {Advert{AddAt: recentlyTime}, recentlyTime, ErrAlreadyProvided},
		"from_future":      {Advert{AddAt: futureTime}, futureTime, ErrAlreadyProvided},
	}

	for name, d := range cases {
		t.Run(name, func(t *testing.T) {
			err := UpFreeAdvert(&d.advert)

			if d.expectAddAt.Format(time.RFC3339) != d.advert.AddAt.Format(time.RFC3339) {
				t.Fatalf("Ожидается дата %s, получено %s", d.expectAddAt, d.advert.AddAt)
			}

			if d.expectError != err {
				t.Fatalf(`Ожидается ошибка "%s", получено "%s"`, d.expectError, err)
			}
		})
	}
}
