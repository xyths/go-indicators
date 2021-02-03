package indicator

// 1: green, -1: red, 0: blue
func Impulse(ema, macdHist []float64) []int {
	if len(ema) != len(macdHist) {
		return nil
	}
	l := len(ema)
	out := make([]int, l)
	for i := 1; i < l; i++ {
		flag1 := ema[i] > ema[i-1]
		flag2 := macdHist[i] > macdHist[i-1]
		if flag1 && flag2 {
			out[i] = 1
		} else if !flag1 && !flag2 {
			out[i] = -1
		}
	}
	return out
}
