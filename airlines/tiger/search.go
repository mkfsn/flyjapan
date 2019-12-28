package tiger

import (
	"context"
	"fmt"
	"time"

	"github.com/mkfsn/flyjapan/airlines"
)

func (t *tiger) Search(ctx context.Context, q airlines.Query) (airlines.Result, error) {
	inbound, err := t.getInbound(ctx, q)
	if err != nil {
		return nil, err
	}
	outbound, err := t.getOutbound(ctx, q)
	if err != nil {
		return nil, err
	}
	return airlines.GenerateResult(inbound, outbound), nil
}

func (t *tiger) getFares(ctx context.Context, departure, arrival string, year, month int) (*FareCache, error) {
	path := fmt.Sprintf("%s:%s:%s:%d-%02d", departure, arrival, "TWD", year, month)
	if fares, ok := t.cache[path]; ok {
		return fares, nil
	}

	res, err := NewFareCacheFromClient(ctx, t.client, path)
	if err != nil {
		return nil, err
	}

	t.cache[path] = res
	return res, nil
}

func (t *tiger) getInbound(ctx context.Context, q airlines.Query) (airlines.Flights, error) {
	res, err := t.getFares(ctx, q.DepartureAirportCode, q.ArrivalAirportCode, q.DepartureDate.Year(), int(q.DepartureDate.Month()))
	if err != nil {
		return nil, err
	}
	return t.getFlights(res.Fares[q.DepartureDate.Day()-1], q.DepartureDate, q.DepartureDate, q.DepartureAirportCode, q.ArrivalAirportCode)
}

func (t *tiger) getOutbound(ctx context.Context, q airlines.Query) (airlines.Flights, error) {
	res, err := t.getFares(ctx, q.ArrivalAirportCode, q.DepartureAirportCode, q.ReturnDate.Year(), int(q.ReturnDate.Month()))
	if err != nil {
		return nil, err
	}
	return t.getFlights(res.Fares[q.ReturnDate.Day()-1], q.ReturnDate, q.ReturnDate, q.ArrivalAirportCode, q.DepartureAirportCode)
}

func (t *tiger) getFlights(fare Fare, departure, arrival time.Time, origin, destination string) (airlines.Flights, error) {
	flights := airlines.Flights{
		airlines.Flight{
			DepartureTime: departure,
			ArrivalTime:   arrival,
			FlightID:      "---",
			Origin:        airlines.Airport{Code: origin},
			Destination:   airlines.Airport{Code: destination},
			Fares: []*airlines.Fare{
				{
					Seat:     fare.Available,
					BaseFare: int(fare.FareAmount),
				},
			},
			TaxAdult:  int(fare.TaxesAndFeesAmount),
			TaxChild:  int(fare.TaxesAndFeesAmount),
			TaxInfant: int(fare.TaxesAndFeesAmount),
		},
	}
	return flights, nil
}
