package factory

import (
	"strings"

	"github.com/mkfsn/flyjapan/airlines"
	"github.com/mkfsn/flyjapan/airlines/peach"
	"github.com/mkfsn/flyjapan/airlines/tiger"
)

func CreateAirline(name airlines.AirlineName) (airlines.Airline, error) {
	airline := airlines.AirlineName(strings.ToLower(string(name)))
	switch airline {
	case airlines.AirlinePeach:
		return peach.New()
	case airlines.AirlineTiger:
		return tiger.New()
	default:
	}
	return nil, ErrUnsupportedAirline
}
