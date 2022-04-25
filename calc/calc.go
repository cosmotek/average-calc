package calc

type AverageCalculator struct {
	MaxHistory uint64

	totalNumbers uint64
	numbers      []float64
	sum          float64
}

func New(maxHistory uint64) AverageCalculator {
	return AverageCalculator{
		MaxHistory: maxHistory,
		numbers:    []float64{},
	}
}

func (a *AverageCalculator) AddNumber(number float64) {
	a.numbers = append(a.numbers, number)
	a.sum += number

	if a.totalNumbers == a.MaxHistory {
		a.sum -= a.numbers[0]
		a.numbers = a.numbers[1:]
	} else {
		a.totalNumbers += 1
	}
}

func (a AverageCalculator) Average() float64 {
	return a.sum / float64(a.totalNumbers)
}

func (a *AverageCalculator) Reset() {
	a.totalNumbers = 0
	a.sum = 0
	a.numbers = []float64{}
}
