package indicator

import (
	"github.com/markcheno/go-talib"
	"math"
)

type Sniper struct {
	Length int
}

func (s Sniper) Signal(inHigh, inLow, inClose []float64) []int {
	//out := talib.Sma(inClose, s.Length)
	const slow = 8
	const fast = 5
	vh1 := talib.Ema(talib.Max(talib.MedPrice(inLow, inClose), fast), fast)
	vl1 := talib.Ema(talib.Min(talib.MedPrice(inHigh, inClose), slow), slow)

	e_ema1 := talib.Ema(inClose, 1)
	e_ema2 := talib.Ema(e_ema1, 1)
	e_ema3 := talib.Ema(e_ema2, 1)
	tema := talib.Add(talib.Sub(e_ema1, e_ema2), e_ema3)

	e_e1 := talib.Ema(inClose, slow)
	e_e2 := talib.Ema(e_e1, fast)
	l := len(inClose)
	for i := 0; i < l; i++ {
		e_e1[i] = 2 * e_e1[i]
	}
	dema := talib.Sub(e_e1, e_e2)

	signal := make([]float64, l)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		if tema[i] > dema[i] {
			signal[i] = math.Max(vh1[i], vl1[i])
		} else {
			signal[i] = math.Min(vh1[i], vl1[i])
		}
		if i >= 2 {
			if tema[i] > dema[i] && signal[i] > inLow[i] && signal[i]-signal[i-1] > signal[i-1]-signal[i-2] {
				ret[i] = 1 // buy
			} else if tema[i] < dema[i] && signal[i] < inHigh[i] && signal[i-1]-signal[i] > signal[i-2]-signal[i-1] {
				ret[i] = -1 // sell
			}
		}
	}
	return ret
}
