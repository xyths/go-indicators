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

// Keltner Channel used by LazyBear
func KCLB(inHigh, inLow, inClose []float64, inTimePeriod int, multiplier float64) (upper, basis, lower []float64) {
	return KC(inHigh, inLow, inClose, inTimePeriod, multiplier, talib.SMA, talib.SMA)
}

// Keltner Channel used by TradingView, use EMA as default
func KCTV(inHigh, inLow, inClose []float64, inTimePeriod int, multiplier float64) (upper, basis, lower []float64) {
	// TODO: TradingView use RMA here
	return KC(inHigh, inLow, inClose, inTimePeriod, multiplier, talib.EMA, talib.EMA)
}

func KC(inHigh, inLow, inClose []float64, inTimePeriod int, multiplier float64, basisMAType, rangeMAType talib.MaType) (upper, basis, lower []float64) {
	basis = talib.Ma(inClose, inTimePeriod, basisMAType)
	tr := talib.TRange(inHigh, inLow, inClose)
	range_ := talib.Ma(tr, inTimePeriod, rangeMAType)
	band := Multiply(range_, multiplier)
	upper = talib.Add(basis, band)
	lower = talib.Sub(basis, band)
	return
}
