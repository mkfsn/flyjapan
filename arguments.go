package flyjapan

import (
	"time"

	"github.com/mkfsn/flyjapan/airlines"
)

type QueryArgument func(*queryArguments)

type dateFromTo struct {
	from time.Time
	to   time.Time
}

// Query(DateFromTo(from, to), SourceAirport(), DestinationAirport(), SortBy(fn), Airline(), Airline(), RepeatWeeks(n), FilterBy(fn))
type queryArguments struct {
	DateFromTo         []dateFromTo
	SourceAirport      string
	DestinationAirport string
	Airline            airlines.Airline
}

func DateFromTo(from, to time.Time) QueryArgument {
	return func(o *queryArguments) {
		o.DateFromTo = append(o.DateFromTo, dateFromTo{from: from, to: to})
	}
}

func SourceAirport(airport string) QueryArgument {
	return func(o *queryArguments) {
		o.SourceAirport = airport
	}
}

func DestinationAirport(airport string) QueryArgument {
	return func(o *queryArguments) {
		o.DestinationAirport = airport
	}
}

func Airline(airline airlines.Airline) QueryArgument {
	return func(o *queryArguments) {
		o.Airline = airline
	}
}
