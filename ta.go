package indicator

import "github.com/markcheno/go-talib"

func Multiply(source []float64, multiplier float64) []float64 {
	ret := make([]float64, len(source))
	for i := 0; i < len(source); i++ {
		ret[i] = source[i] * multiplier
	}
	return ret
}

func Keltner(factor float64, mean, atr []float64) (upper, lower []float64) {
	atr2 := Multiply(atr, factor)
	upper = talib.Add(mean, atr2)
	lower = talib.Sub(mean, atr2)
	return
}
