package factory

import (
	"github.com/mkfsn/flyjapan/airlines"
	"github.com/mkfsn/flyjapan/airlines/peach"
)

func CreateAirline(airline airlines.AirlineName) (airlines.Airline, error) {
	switch airline {
	case airlines.AirlinePeach:
		return peach.New()
	default:
	}
	return nil, ErrUnsupportedAirline
}
