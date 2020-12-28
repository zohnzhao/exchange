package main
import（
＂fmt"
＂time"
）
var yesterdayPrice,NineteenDaysAgoPrice,yesterdayEMAPrice float32
func main(){
／／不包含当天
yesterday :=DaysAgo(1)
yesterdayDate := yesterday.Format("2006-01-02")
NineteenDaysAgo := DaysAgo(19)
NineteenDaysAgoDate := NineteenDaysAgo.Format("2006-01-02")
fmt.Printf("请输入 v 日的收盘价：＼n",yesterdayDate)
fmt.ScanLn(&yesterdayPrice)
fat.Printf("请输入 %v 日的收盘价：＼n",NineteenDaysAgoDate)
fmt.ScanLn(&NineteenDaysAgoPrice)
fmt.Printf("请输入％v 日的EMA(20)价：＼n",yesterdayDate)
fmt.Scanln(&yesterdayEMAPrice)
buyorSell()
func DaysAgo(day int) time.Time {
now:=time.Nowl)
return now.Add(time.Duration(-day)*time.Hour *24)
func buyorSell(){
switch {
case yesterdayPrice> NineteenDaysAgoPrice && yesterdayPrice > yesterdayEMAPrice:
fmt.Println("继续持仓或者买入＂）
case yesterdayPrice <NineteenDaysAgoPrice && yesterdayPrice <yesterdayEMAPrice:
fmt.Println("卖出＂）
default:
fmt.Println("价格处于中间位置，待确定＂）
｝｝
