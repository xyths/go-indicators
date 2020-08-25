package indicator

import "github.com/markcheno/go-talib"

// Squeeze Momentum Indicator [LazyBear]
// https://cn.tradingview.com/script/nqQ1DT5a-Squeeze-Momentum-Indicator-LazyBear/
// bbl: BB Length
// kcl: KC Length
// bbf: BB MultFactor
// kcf: KC MulFactor
func SqueezeMomentum(bbl, kcl int, bbf, kcf float64, inHigh, inLow, inClose []float64) ([]float64, []bool, []bool, []bool) {
	basis := talib.Sma(inClose, bbl)
	dev := talib.StdDev(inClose, bbl, bbf)
	upperBB := talib.Add(basis, dev)
	lowerBB := talib.Sub(basis, dev)

	ma := talib.Sma(inClose, kcl)
	// always use true range
	tr := talib.TRange(inHigh, inLow, inClose)
	rangeMa := talib.Sma(tr, kcl)
	for i := 0; i < len(rangeMa); i++ {
		rangeMa[i] = rangeMa[i] * kcf
	}
	upperKC := talib.Add(ma, rangeMa)
	lowerKC := talib.Sub(ma, rangeMa)

	on := make([]bool, len(inHigh))
	off := make([]bool, len(inHigh))
	no := make([]bool, len(inHigh))
	for i := 0; i < len(upperBB); i++ {
		if lowerBB[i] > lowerKC[i] && upperBB[i] < upperKC[i] {
			// bb inside kc
			on[i] = true
		}
		if lowerBB[i] < lowerKC[i] && upperBB[i] > upperKC[i] {
			// kc inside bb
			off[i] = true
		}
		if !on[i] && !off[i] {
			no[i] = true
		}
	}
	maxHigh := talib.Max(inHigh, kcl)
	minLow := talib.Min(inLow, kcl)
	avg1 := talib.MedPrice(maxHigh, minLow)
	avg2 := talib.MedPrice(avg1, basis)
	target := talib.Sub(inClose, avg2)
	val := talib.LinearReg(target, kcl)
	return val, on, off, no
}
