package flyjapan

import (
	"time"

	"github.com/mkfsn/flyjapan/airlines"
	"github.com/mkfsn/flyjapan/airlines/peach"
)

type Option func(*options)

type datePair struct {
	from time.Time
	to   time.Time
}

// Query(ToAndFrom(from, to), SourceAirport(), DestinationAirport(), SortBy(fn), Airline(), Airline(), RepeatWeeks(n), FilterBy(fn))
type options struct {
	ToAndFrom           []datePair
	SourceAirports      []string
	DestinationAirports []string
	Airlines            []airlines.Searcher
	err                 error
}

func ToAndFrom(from, to time.Time) Option {
	return func(o *options) {
		o.ToAndFrom = append(o.ToAndFrom, datePair{from: from, to: to})
	}
}

func SourceAirport(airports ...string) Option {
	return func(o *options) {
		o.SourceAirports = airports
	}
}

func DestinationAirport(airports ...string) Option {
	return func(o *options) {
		o.DestinationAirports = airports
	}
}

func Airline(airline airlines.Airline) Option {
	return func(o *options) {
		var searcher airlines.Searcher
		var err error
		switch airline {
		case airlines.AirlinePeach:
			searcher, err = peach.New()
		default:
			err = ErrUnsupportedAirline
		}
		if err != nil {
			o.err = err
		}
		o.Airlines = append(o.Airlines, searcher)
	}
}

func (o *options) Validate() error {
	return o.err
}
