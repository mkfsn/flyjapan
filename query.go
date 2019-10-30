package flyjapan

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/mkfsn/flyjapan/airlines"
)

type querier interface {
	Query(context.Context) (Result, error)
}

type flightsQuery struct {
	options
}

// Query(From(date), To(date), Airport(), Airport(), SortBy(fn), Airline(), Airline(), RepeatWeeks(n), FilterBy(fn))
func Query(ctx context.Context, setters ...Option) (<-chan Result, error) {
	q := &flightsQuery{}
	for _, setter := range setters {
		setter(&q.options)
	}
	if err := q.options.Validate(); err != nil {
		return nil, err
	}
	return q.Query(ctx)
}

func (q *flightsQuery) Query(ctx context.Context) (<-chan Result, error) {
	var wg sync.WaitGroup
	ch := make(chan Result, len(q.Airlines))
	for _, airline := range q.Airlines {
		airline := airline
		wg.Add(1)
		go func() {
			defer wg.Done()
			q.query(airline, ch)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch, nil
}

func (q *flightsQuery) query(airline airlines.Searcher, ch chan Result) {
	for _, source := range q.SourceAirports {
		for _, destination := range q.DestinationAirports {
			inbound, outbound, err := q.queryRaw(airline, source, destination)
			if err != nil {
				log.Println("error:", err)
				continue
			}
			inbound = inbound.FilterBy(airlines.AvailableFlight)
			outbound = outbound.FilterBy(airlines.AvailableFlight)
			ch <- Result{Inbound: inbound, Outbound: outbound}
		}
	}
}

func (q *flightsQuery) queryRaw(airline airlines.Searcher, source, destination string) (airlines.Flights, airlines.Flights, error) {
	var inbound, outbound airlines.Flights
	for _, dates := range q.ToAndFrom {
		res, err := fetch(airline, source, destination, dates.from, dates.to)
		if err != nil {
			return nil, nil, err
		}
		inbound, outbound = append(inbound, res.InBound()...), append(outbound, res.OutBound()...)
	}
	return inbound, outbound, nil
}

func fetch(airline airlines.Searcher, from, to string, begin, end time.Time) (airlines.Result, error) {
	return airline.Search(context.Background(), airlines.Query{
		DepartureDate:        begin,
		ReturnDate:           end,
		DepartureAirportCode: from,
		ArrivalAirportCode:   to,
		IsReturn:             true,
		AdultCount:           1,
	})
}
