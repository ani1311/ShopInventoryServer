package mapUtils

import (
	"math"
	"strconv"
)

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1str, lon1str, lat2str, lon2str string) float64 {
	lat1, _ := strconv.ParseFloat(lat1str, 64)
	lon1, _ := strconv.ParseFloat(lon1str, 64)
	lat2, _ := strconv.ParseFloat(lat2str, 64)
	lon2, _ := strconv.ParseFloat(lon2str, 64)

	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
