package peach

type Fare struct {
	BaseFare     int    `json:"baseFare"`
	BookingClass string `json:"bookingClass"`
	BookingType  string `json:"bookingType"`
	Discounted   bool   `json:"discounted"`
	Fare         int    `json:"fare"`
	FareCode     string `json:"fareCode"`
	FareId       string `json:"fareId"`
	IsMin        bool   `json:"isMin"`
	IsSale       bool   `json:"isSale"`
	IsStaff      bool   `json:"isStaff"`
	Seat         int    `json:"seat"`
}

type Flight struct {
	ArrivalTime       string          `json:"arrivalTime"`
	ArrivalTimezone   string          `json:"arrivalTimezone"`
	DepartureTime     string          `json:"departureTime"`
	DepartureTimezone int             `json:"departureTimezone"`
	Destination       string          `json:"destination"`
	DestinationCode   string          `json:"destinationCode"`
	Fares             map[string]Fare `json:"fares"`
	FlightDuration    int             `json:"flightDuration"`
	FlightId          string          `json:"flightId"`
	FlightNumber      string          `json:"flightNumber"`
	Origin            string          `json:"origin"`
	OriginCode        string          `json:"originCode"`
	TaxAdult          string          `json:"taxAdult"`
	TaxChild          int             `json:"taxChild"`
	TaxInfant         int             `json:"taxInfant"`
}
