package indicator

import (
	"encoding/csv"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestSuperTrend(t *testing.T) {
	filename := os.Getenv("CSV")
	f, err := os.Open(filename)
	defer func() { _ = f.Close() }()
	require.NoError(t, err)
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	require.NoError(t, err)
	timestamp := make([]int64, len(records)-1)
	open := make([]float64, len(records)-1)
	high := make([]float64, len(records)-1)
	low := make([]float64, len(records)-1)
	close1 := make([]float64, len(records)-1)
	for i := 1; i < len(records); i++ {
		timestamp[i-1], _ = strconv.ParseInt(records[i][0], 10, 64)
		open[i-1], _ = strconv.ParseFloat(records[i][1], 0)
		high[i-1], _ = strconv.ParseFloat(records[i][2], 0)
		low[i-1], _ = strconv.ParseFloat(records[i][3], 0)
		close1[i-1], _ = strconv.ParseFloat(records[i][4], 0)
	}

	tsl, trend := SuperTrend(3, 7, high, low, close1)

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	layout := "2006-01-02 15:04:05"
	//return timestamp.Unix(timestamp, 0).In(beijing).Format(layout)
	for i := 0; i < len(tsl); i++ {
		t.Logf("%s %f %f %f %f - %f %v", time.Unix(timestamp[i], 0).In(beijing).Format(layout), open[i], high[i], low[i], close1[i], tsl[i], trend[i])
	}
}
