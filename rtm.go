package indicator

import "github.com/markcheno/go-talib"

func Rtm(atrPeriod, emaPeriod int, factor float64, inHigh, inLow, inClose []float64) (mean, upper, lower, atr []float64) {
	atr = talib.Atr(inHigh, inLow, inClose, atrPeriod)
	mean = talib.Ema(inClose, emaPeriod)
	upper, lower = Keltner(factor, mean, atr)
	return mean, upper, lower, atr
}
