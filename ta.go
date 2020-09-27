package indicator

import "github.com/markcheno/go-talib"

func Multiply(source []float64, multiplier float64) []float64 {
	ret := make([]float64, len(source))
	for i := 0; i < len(source); i++ {
		ret[i] = source[i] * multiplier
	}
	return ret
}

func Keltner(period int, multiply float64, inHigh, inLow, inClose []float64) (mean, upper, lower []float64) {
	atr := talib.Atr(inHigh, inLow, inClose, period)
	mean = talib.Ema(inClose, 13)
	atr = Multiply(atr, multiply)
	upper = talib.Add(mean, atr)
	lower = talib.Sub(mean, atr)
	return
}
