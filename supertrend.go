package indicator

import "github.com/markcheno/go-talib"

// Supertrend V1.0 - Buy or Sell Signal
// https://cn.tradingview.com/chart/5wBFaZWw/
// trend true = green, false = red
// red->green = false->true = buy
// green->red = true->false = sell
func SuperTrend(factor float64, period int, inHigh, inLow, inClose []float64) ([]float64, []bool) {
	l := len(inHigh)
	hl2 := talib.MedPrice(inHigh, inLow)
	atr := talib.Atr(inHigh, inLow, inClose, period)

	up := make([]float64, l)
	down := make([]float64, l)
	trendUp := make([]float64, l)
	trendDown := make([]float64, l)
	trend := make([]bool, l)
	tsl := make([]float64, l)
	color := make([]bool, l)
	for i := 0; i < l; i++ {
		up[i] = hl2[i] - atr[i]*factor
		down[i] = hl2[i] + atr[i]*factor
		if i == 0 {
			trend[i] = true
			color[i] = true
			continue
		}
		if inClose[i-1] > trendUp[i-1] {
			if up[i] >= trendUp[i-1] {
				trendUp[i] = up[i]
			} else {
				trendUp[i] = trendUp[i-1]
			}
		} else {
			trendUp[i] = up[i]
		}
		if inClose[i-1] < trendDown[i-1] {
			if down[i] <= trendDown[i-1] {
				trendDown[i] = down[i]
			} else {
				trendDown[i] = trendDown[i-1]
			}
		} else {
			trendDown[i] = down[i]
		}

		if inClose[i] > trendDown[i-1] {
			trend[i] = true
		} else if inClose[i] < trendUp[i-1] {
			trend[i] = false
		} else {
			trend[i] = trend[i-1]
		}

		if trend[i] {
			tsl[i] = trendUp[i]
			color[i] = true
		} else {
			tsl[i] = trendDown[i]
			color[i] = false
		}
	}

	return tsl, trend
}
