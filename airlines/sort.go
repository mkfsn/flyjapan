package airlines

type SortByBaseFare []Flight

func (f SortByBaseFare) Len() int {
	return len(f)
}

func (f SortByBaseFare) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f SortByBaseFare) Less(i, j int) bool {
	return f[i].Cheapest() < f[j].Cheapest()
}
