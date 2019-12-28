package tiger

import (
	"net/http"
	"net/http/cookiejar"

	"github.com/mkfsn/flyjapan/airlines"
)

type tiger struct {
	client *http.Client
	cache  map[string]*FareCache
}

func New() (airlines.Airline, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	tiger := &tiger{
		client: &http.Client{Jar: jar},
		cache:  make(map[string]*FareCache),
	}
	return tiger, nil
}
