package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mkfsn/flyjapan"
	"github.com/mkfsn/flyjapan/peach"
)

func main() {
	peach, err := peach.New()
	if err != nil {
		log.Fatalln("error:", err)
	}

	ctx := context.Background()

	week := time.Hour * 24 * 7
	places := []string{"HND", "KIX", "OKA"}
	for _, place := range places {
		for i := 0; i < 23; i++ {

			friday := time.Date(2019, 3, 22, 0, 0, 0, 0, time.Local).Add(week * time.Duration(i))
			saturday := time.Date(2019, 3, 23, 0, 0, 0, 0, time.Local).Add(week * time.Duration(i))
			sunday := time.Date(2019, 3, 24, 0, 0, 0, 0, time.Local).Add(week * time.Duration(i))
			monday := time.Date(2019, 3, 25, 0, 0, 0, 0, time.Local).Add(week * time.Duration(i))

			x, err := fetch(peach, ctx, "TPE", place, friday, sunday)
			if err != nil {
				break
			}
			y, err := fetch(peach, ctx, "TPE", place, saturday, monday)
			if err != nil {
				break
			}

			var i, j flyjapan.Flight

			i1 := x.InBound().Cheapest()
			i2 := y.InBound().Cheapest()
			if a, b := i1.CheapestFare(), i2.CheapestFare(); a.BaseFare > b.BaseFare {
				i = i1
			} else {
				i = i2
			}

			j1 := x.OutBound().Cheapest()
			j2 := y.OutBound().Cheapest()
			if a, b := j1.CheapestFare(), j2.CheapestFare(); a.BaseFare > b.BaseFare {
				j = j1
			} else {
				j = j2
			}

			f1, f2 := i.CheapestFare().BaseFare, j.CheapestFare().BaseFare
			if f1 != 0 && f2 != 0 && f1+f2 < 5000 {
				fmt.Printf("[%s(%v)->%s(%v)]%5v --> [%s(%v)->%s(%v)]%5v = %5v\n",
					i.OriginCode,
					i.DepartureTime,
					i.DestinationCode,
					i.ArrivalTime,
					f1,
					j.OriginCode,
					j.DepartureTime,
					j.DestinationCode,
					j.ArrivalTime,
					f2,
					f1+f2,
				)
			}
		}
	}
}

func fetch(peach flyjapan.Searcher, ctx context.Context, from, to string, begin, end time.Time) (flyjapan.Result, error) {
	return peach.Search(ctx, flyjapan.Query{
		DepartureDate:        begin,
		ReturnDate:           end,
		DepartureAirportCode: from,
		ArrivalAirportCode:   to,
		IsReturn:             true,
		AdultCount:           1,
	})
}
