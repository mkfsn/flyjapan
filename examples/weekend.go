package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/mkfsn/flyjapan"
	"github.com/mkfsn/flyjapan/airlines"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"FUK", "HND", "KIX", "OKA"}
	}
	for _, city := range args {
		fetchByDestination(city)
	}
}

func fetchByDestination(city string) {
	// Query(From(date), To(date), Airport(), Airport(), SortBy(fn), Airline(), Airline(), RepeatWeeks(n), FilterBy(fn))
	today := time.Now().Truncate(time.Hour * 24)
	friday := today.Add(time.Duration(12-int(today.Weekday())%7) * 24 * time.Hour)
	for i := 0; i < 28; i, friday = i+1, friday.Add(time.Hour*24*7) {
		saturday, sunday, monday := friday.Add(time.Hour*24), friday.Add(time.Hour*24*2), friday.Add(time.Hour*24*3)
		ch, err := flyjapan.Query(
			context.Background(),
			flyjapan.Airline(airlines.AirlinePeach),
			flyjapan.DestinationAirport(city),
			flyjapan.SourceAirport("TPE"),
			flyjapan.ToAndFrom(friday, sunday),
			flyjapan.ToAndFrom(friday, monday),
			flyjapan.ToAndFrom(saturday, monday),
		)
		if err != nil {
			log.Fatalln("error:", err)
		}
		for res := range ch {
			handleResult(res)
		}
	}
}

func handleResult(res flyjapan.Result) {
	inBound := res.Inbound.FilterBy(airlines.AvailableFlight)
	outBound := res.Outbound.FilterBy(airlines.AvailableFlight)
	sort.Sort(airlines.SortByBaseFare(inBound))
	sort.Sort(airlines.SortByBaseFare(outBound))
	for i := 0; i < len(inBound) && i < 3; i++ {
		for j := 0; j < len(outBound) && j < 3; j++ {
			x, y := inBound[i], outBound[j]
			a, b := x.Cheapest(), y.Cheapest()
			if a+b > 4500 {
				continue
			}
			dateFormat := "2006/01/02 15:04:05"
			fmt.Printf("%s%v%s%s%v%5v",
				BgBlue(Black(fmt.Sprintf("[%s]", x.Origin.Code))),
				BgMagenta(Black(x.DepartureTime.Format(dateFormat))),
				BgGray(Black("->")),
				BgBlue(Black(fmt.Sprintf("[%s]", x.Destination.Code))),
				BgMagenta(Black(x.ArrivalTime.Format(dateFormat))),
				BgBrown(Black(a)),
			)
			fmt.Printf(" + ")
			fmt.Printf("%s%v%s%s%v%5v",
				BgBlue(Black(fmt.Sprintf("[%s]", y.Origin.Code))),
				BgMagenta(Black(y.DepartureTime.Format(dateFormat))),
				BgGray(Black("->")),
				BgBlue(Black(fmt.Sprintf("[%s]", y.Destination.Code))),
				BgMagenta(Black(y.ArrivalTime.Format(dateFormat))),
				BgBrown(Black(b)),
			)
			if a+b < 3000 {
				fmt.Printf(" = %5v\n", BgGreen(Red(a+b)))
			} else {
				fmt.Printf(" = %5v\n", BgGreen(Black(a+b)))
			}
		}
	}
}
