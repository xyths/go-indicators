package indicator

import "testing"

func TestImpulse(t *testing.T) {
	tests := []struct {
		EMA    []float64
		MACD   []float64
		Expect []int
	}{
		{
			EMA:    []float64{0, 0, 0},
			MACD:   []float64{0, 0, 0},
			Expect: []int{0, -1, -1},
		},
		{
			EMA:    []float64{1, 2, 3},
			MACD:   []float64{2, 3, 4},
			Expect: []int{0, 1, 1},
		},
		{
			EMA:    []float64{3, 2, 1},
			MACD:   []float64{4, 3, 2},
			Expect: []int{0, -1, -1},
		},
		{
			EMA:    []float64{1, 2, 3, 3, 2, 1},
			MACD:   []float64{2, 3, 4, 4, 3, 2},
			Expect: []int{0, 1, 1, -1, -1, -1},
		},
	}
	for i, tt := range tests {
		actual := Impulse(tt.EMA, tt.MACD)
		if len(actual) != len(tt.EMA) {
			t.Errorf("bad len: expect %d, actual %d", len(tt.EMA), len(actual))
			continue
		}
		for j := 0; j < len(tt.Expect); j++ {
			if actual[j] != tt.Expect[j] {
				t.Errorf("[%d][%d] expect %d, actual %d", i, j, tt.Expect[j], actual[j])
			}
		}
	}
}
