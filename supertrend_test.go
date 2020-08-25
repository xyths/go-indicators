package indicator

import (
	"encoding/csv"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func TestSuperTrend(t *testing.T) {
	filename := os.Getenv("CSV")
	f, err := os.Open(filename)
	defer f.Close()
	require.NoError(t, err)
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	require.NoError(t, err)
	time := make([]string, len(records)-1)
	open := make([]float64, len(records)-1)
	high := make([]float64, len(records)-1)
	low := make([]float64, len(records)-1)
	close := make([]float64, len(records)-1)
	for i := 1; i < len(records); i++ {
		time[i-1] = records[i][0]
		open[i-1], _ = strconv.ParseFloat(records[i][1], 0)
		high[i-1], _ = strconv.ParseFloat(records[i][2], 0)
		low[i-1], _ = strconv.ParseFloat(records[i][3], 0)
		close[i-1], _ = strconv.ParseFloat(records[i][4], 0)
	}

	tsl, trend := SuperTrend(3, 7, high, low, close)
	for i := 0; i < len(tsl); i++ {
		t.Logf("%s %f %f %f %f - %f %v", time[i], open[i], high[i], low[i], close[i], tsl[i], trend[i])
	}
}
