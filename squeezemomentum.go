package indicator

import "github.com/markcheno/go-talib"

const (
	Black int = 1
	Gray      = -1
	Blue      = 0
)

type Result struct {
	Trend int // 1 squeeze, 2 up, -2 down, 0 stop
	Last  int
}

type Detail struct {
	Value      []float64
	SqueezeOn  []bool
	SqueezeOff []bool
	NoSqueeze  []bool
	KCUpper    []float64
	KCMiddle   []float64
	KCLower    []float64
	BBUpper    []float64
	BBMiddle   []float64
	BBLower    []float64
}

// Squeeze Momentum Indicator [LazyBear]
// https://cn.tradingview.com/script/nqQ1DT5a-Squeeze-Momentum-Indicator-LazyBear/
// bbl: BB Length
// kcl: KC Length
// bbf: BB MultFactor
// kcf: KC MulFactor
func Squeeze(bbl, kcl int, bbf, kcf float64, inHigh, inLow, inClose []float64) (r Result, d Detail) {
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

	squeezeOn := make([]bool, len(inHigh))
	squeezeOff := make([]bool, len(inHigh))
	noSqueeze := make([]bool, len(inHigh))
	for i := 0; i < len(upperBB); i++ {
		// bb inside kc
		squeezeOn[i] = lowerBB[i] > lowerKC[i] && upperBB[i] < upperKC[i]
		// kc inside bb
		squeezeOff[i] = lowerBB[i] < lowerKC[i] && upperBB[i] > upperKC[i]

		noSqueeze[i] = !squeezeOn[i] && !squeezeOff[i]
	}
	maxHigh := talib.Max(inHigh, kcl)
	minLow := talib.Min(inLow, kcl)
	avg1 := talib.MedPrice(maxHigh, minLow)
	avg2 := talib.MedPrice(avg1, ma)
	target := talib.Sub(inClose, avg2)
	val := talib.LinearReg(target, kcl)
	d.Value = val
	d.SqueezeOn = squeezeOn
	d.SqueezeOff = squeezeOff
	d.NoSqueeze = noSqueeze
	d.KCUpper = upperKC
	d.KCMiddle = ma
	d.KCLower = lowerKC
	d.BBUpper = upperBB
	d.BBMiddle = basis
	d.BBLower = lowerBB
	r = Summary(d)
	return
}

func Summary(d Detail) (r Result) {
	l := len(d.Value)
	current := l - 2
	if l < 3 {
		return
	}
	if d.NoSqueeze[current] {
		return
	}

	if d.SqueezeOn[current] {
		// count the dark cross
		last := 0
		for i := current; d.SqueezeOn[i] && i >= 0; i-- {
			last++
		}
		r.Trend = 1 // squeeze
		r.Last = last

	} else {
		// find first gray cross
		firstGrayIndex := 0
		for i := current; !d.SqueezeOn[i] && i >= 0; i-- {
			firstGrayIndex = i
		}
		isUptrend := d.Value[firstGrayIndex] > 0
		trendStopped := false
		stopIndex := 0
		for i := firstGrayIndex; i <= current; i++ {
			if isUptrend && d.Value[i] <= d.Value[i-1] || !isUptrend && d.Value[i] >= d.Value[i-1] {
				trendStopped = true
				stopIndex = i
				break
			}
		}
		if !trendStopped {
			if isUptrend {
				r.Trend = 2
			} else {
				r.Trend = -2
			}
			r.Last = current - firstGrayIndex + 1
		} else {
			r.Trend = 0
			r.Last = current - stopIndex + 1
		}
	}
	return
}
