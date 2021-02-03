package indicator

import "github.com/markcheno/go-talib"

func Efi(inTimePeriod int, inReal, inVol []float64) []float64 {
	l := len(inReal)
	outReal := make([]float64, l)
	for i := 1; i < l; i++ {
		outReal[i] = inVol[i] * (inReal[i] - inReal[i-1])
	}
	return talib.Ema(outReal, inTimePeriod)
}
