package flyjapan

import (
	"context"
	"time"

	"github.com/mkfsn/flyjapan/airlines"
)

type Query struct {
	queryArguments
}

func NewQuery(setters ...QueryArgument) *Query {
	var job Query
	for _, setter := range setters {
		setter(&job.queryArguments)
	}
	return &job
}

func (q *Query) query(ctx context.Context) (*Result, error) {
	inbound, outbound, err := q.queryFlights(ctx)
	if err != nil {
		return nil, err
	}
	inbound = inbound.FilterBy(airlines.AvailableFlight)
	outbound = outbound.FilterBy(airlines.AvailableFlight)
	return &Result{Inbound: inbound, Outbound: outbound}, nil
}

func (q *Query) queryFlights(ctx context.Context) (airlines.Flights, airlines.Flights, error) {
	var inbound, outbound airlines.Flights
	for _, dates := range q.DateFromTo {
		res, err := fetch(ctx, q.Airline, q.SourceAirport, q.DestinationAirport, dates.from, dates.to)
		if err != nil {
			return nil, nil, err
		}
		inbound, outbound = append(inbound, res.InBound()...), append(outbound, res.OutBound()...)
	}
	return inbound, outbound, nil
}

func fetch(ctx context.Context, airline airlines.Searcher, from, to string, begin, end time.Time) (airlines.Result, error) {
	return airline.Search(ctx, airlines.Query{
		DepartureDate:        begin,
		ReturnDate:           end,
		DepartureAirportCode: from,
		ArrivalAirportCode:   to,
		IsReturn:             true,
		AdultCount:           1,
	})
}
