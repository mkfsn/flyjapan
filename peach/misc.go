package peach

import (
	"fmt"
	"io"
	"strings"

	"github.com/mkfsn/flyjapan"
)

func parseQuery(q flyjapan.Query) io.Reader {
	data := map[string]interface{}{
		"flight_search_parameter[0][departure_date]":         q.DepartureDate.Format("2006/01/02"),
		"flight_search_parameter[0][return_date]":            q.ReturnDate.Format("2006/01/02"),
		"flight_search_parameter[0][departure_airport_code]": q.DepartureAirportCode,
		"flight_search_parameter[0][arrival_airport_code]":   q.ArrivalAirportCode,
		"flight_search_parameter[0][is_return]":              q.IsReturn,
		"adult_count":                                        q.AdultCount,
		"child_count":                                        q.ChildCount,
		"infant_count":                                       q.InfantCount,
		"r":                                                  "static_search",
	}
	var result []string
	for k, v := range data {
		result = append(result, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.NewReader(strings.Join(result, "&"))
}
