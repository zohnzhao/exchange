package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	talib "github.com/markcheno/go-talib"
)

var (
	t   bool
	arr [][]interface{}
)

func init() {
	flag.BoolVar(&t, "t", false, "是否输出连接")
}

func main() {
	flag.Parse()
	if t {
		times()
	} else {
		buyOrSell()
	}

}

func times() {
	t := daysAgo(30)
	t2 := time.Date(t.Year(), t.Month(), t.Day(), 8, 0, 0, 0, time.Local)
	fmt.Printf("https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=1d&startTime=%v", t2.UnixNano()/1e6)
}

func buyOrSell() {
	readData()
	lenData := len(arr)

	if lenData < 20 {
		log.Fatal("参数不够")
	}

	yesterdayPrice := stringToFloat64(arr[lenData-2][4])
	yesterdayEMAPrice := ema(arr[:lenData-1], 18)
	MAPrice := ma(arr[:lenData-1], 18)
	lenMAPrice := len(MAPrice)
	fmt.Println(MAPrice)

	fmt.Printf("昨天 %v 收盘价\t： %v\n", daysAgo(1).Format("2006-01-02"), yesterdayPrice)
	fmt.Printf("昨天 %v EMA价\t： %v\n\n", daysAgo(1).Format("2006-01-02"), yesterdayEMAPrice)
	fmt.Printf("昨天 %v MA价\t： %v\n", daysAgo(1).Format("2006-01-02"), MAPrice[lenMAPrice-1])
	fmt.Printf("前天 %v MA价\t： %v\n", daysAgo(2).Format("2006-01-02"), MAPrice[lenMAPrice-2])

	switch {
	case MAPrice[lenMAPrice-1] > MAPrice[lenMAPrice-2] && yesterdayPrice > yesterdayEMAPrice:
		fmt.Println("继续持仓或者买入")
	case MAPrice[lenMAPrice-1] < MAPrice[lenMAPrice-2] && yesterdayPrice < yesterdayEMAPrice:
		fmt.Println("卖出")
	default:
		fmt.Println("价格处于中间位置，待确定")
	}
}

func readData() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &arr)
	if err != nil {
		log.Fatal(err)
	}
}

func ma(a [][]interface{}, inTimePeriod int) []float64 {
	lenData := len(a)
	s := make([]float64, lenData)

	for i, v := range a {
		s[i] = stringToFloat64(v[4])
	}

	outReal := talib.Ma(s, inTimePeriod, 0)
	return outReal
}

func ema(a [][]interface{}, inTimePeriod int) float64 {
	lenData := len(a)
	s := make([]float64, lenData)

	for i, v := range a {
		s[i] = stringToFloat64(v[4])
	}

	outReal := talib.Ema(s, inTimePeriod)
	return outReal[lenData-1]
}

func stringToFloat64(v interface{}) float64 {
	float, err := strconv.ParseFloat(fmt.Sprint(v), 64)
	if err != nil {
		log.Fatal(err)
	}
	return float
}

func daysAgo(day int) time.Time {
	now := time.Now()
	return now.Add(time.Duration(-day) * time.Hour * 24)
}
