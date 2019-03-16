package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/mkfsn/flyjapan"
	"github.com/mkfsn/flyjapan/peach"
)

func main() {
	peach, err := peach.New()
	if err != nil {
		log.Fatalln("error:", err)
	}
	for _, place := range []string{"HND", "KIX", "OKA"} {
		fetchFlightToCity(peach, place)
	}
}

func fetchFlightToCity(searcher flyjapan.Searcher, city string) {
	today := time.Now().Truncate(time.Hour * 24)
	nextFriday := today.Add((time.Duration(12-int(today.Weekday())%7) * 24 * time.Hour))
	for i := 0; i < 23; i, nextFriday = i+1, nextFriday.Add(time.Hour*24*7) {
		inBound, outBound := fetchWeekendFlightToCity(searcher, city, nextFriday)
		inBound = inBound.FilterBy(flyjapan.AvailableFlight)
		outBound = outBound.FilterBy(flyjapan.AvailableFlight)
		sort.Sort(flyjapan.SortByBaseFare(inBound))
		sort.Sort(flyjapan.SortByBaseFare(outBound))
		for i := 0; i < len(inBound) && i < 2; i++ {
			for j := 0; j < len(outBound) && j < 2; j++ {
				x, y := inBound[i], outBound[j]
				a, b := x.Cheapest(), y.Cheapest()
				if a+b > 3500 {
					continue
				}
				dateFormat := "2006/01/02 15:04:05"
				fmt.Printf("[%s(%v)->%s(%v)]%5v --> [%s(%v)->%s(%v)]%5v = %5v\n",
					x.Origin.Code, x.DepartureTime.Format(dateFormat), x.Destination.Code, x.ArrivalTime.Format(dateFormat), a,
					y.Origin.Code, y.DepartureTime.Format(dateFormat), y.Destination.Code, y.ArrivalTime.Format(dateFormat), b,
					a+b,
				)
			}
		}
	}
}

func fetchWeekendFlightToCity(searcher flyjapan.Searcher, city string, friday time.Time) (flyjapan.Flights, flyjapan.Flights) {
	saturday := friday.Add(time.Hour * 24 * 1)
	sunday := friday.Add(time.Hour * 24 * 2)
	monday := friday.Add(time.Hour * 24 * 3)
	x, err := fetch(searcher, "TPE", city, friday, sunday)
	if err != nil {
		return nil, nil
	}
	y, err := fetch(searcher, "TPE", city, saturday, monday)
	if err != nil {
		return nil, nil
	}
	inBound := append(x.InBound(), y.InBound()...)
	outBound := append(x.OutBound(), y.OutBound()...)
	return inBound, outBound
}

func fetch(peach flyjapan.Searcher, from, to string, begin, end time.Time) (flyjapan.Result, error) {
	return peach.Search(context.Background(), flyjapan.Query{
		DepartureDate:        begin,
		ReturnDate:           end,
		DepartureAirportCode: from,
		ArrivalAirportCode:   to,
		IsReturn:             true,
		AdultCount:           1,
	})
}