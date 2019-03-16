package flyjapan

type FilterFunc func(f Flight) bool

func AvailableFlight(f Flight) bool {
	if len(f.Fares) == 0 {
		return false
	}
	var seats int
	for _, fare := range f.Fares {
		seats += fare.Seat
	}
	return seats != 0
}
