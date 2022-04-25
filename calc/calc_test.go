package calc

import (
	"math/rand"
	"testing"
	"time"
)

const MaxHistory = 50
const max = 40
const min = 10

func TestCalculateAverage(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	calculator := New(MaxHistory)

	for i := 0; i < MaxHistory*3; i++ {
		calculator.AddNumber(float64(random.Intn(max-min) + min))

		average := calculator.Average()
		trueAverage := averageList(calculator.numbers)
		if average != trueAverage {
			t.Errorf("expected %v but got %v, %v", trueAverage, average, calculator.numbers)
			t.FailNow()
		}

		if len(calculator.numbers) > MaxHistory || len(calculator.numbers) != int(calculator.totalNumbers) {
			t.Errorf("numbers out of sync: len %v, max %v, total %v", len(calculator.numbers), MaxHistory, calculator.totalNumbers)
			t.FailNow()
		}

		t.Logf("average for iteration %d is %v, len %v, total %v", i, calculator.Average(), len(calculator.numbers), calculator.totalNumbers)
	}
}

func averageList(nums []float64) float64 {
	n := 0.0
	for _, num := range nums {
		n += num
	}

	l := len(nums)
	return n / float64(l)
}
