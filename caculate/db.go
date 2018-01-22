package caculate

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Min1TradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Min5TradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Min30TradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Hour1TradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Hour4TradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type DayTradeTable struct {
	Id         int `gorm:"AUTO_INCREMENT;primary_key" json:"-"`
	TimeKey    int `gorm:"index;type:bigint" json:"-"`
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}

func WriteMin1(t time.Time, d StaInfo, db *gorm.DB) {
	var tb Min1TradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tb.Hour = t.Hour()
	tb.Min = t.Minute()
	tmp := fmt.Sprintf("%04d%02d%02d%02d%02d", tb.Year, tb.Month, tb.Day, tb.Hour, tb.Min)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
func WriteMin5(t time.Time, d StaInfo, db *gorm.DB) {
	var tb Min5TradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tb.Hour = t.Hour()
	tb.Min = t.Minute()
	tmp := fmt.Sprintf("%04d%02d%02d%02d%02d", tb.Year, tb.Month, tb.Day, tb.Hour, tb.Min)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
func WriteMin30(t time.Time, d StaInfo, db *gorm.DB) {
	var tb Min30TradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tb.Hour = t.Hour()
	tb.Min = t.Minute()
	tmp := fmt.Sprintf("%04d%02d%02d%02d%02d", tb.Year, tb.Month, tb.Day, tb.Hour, tb.Min)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
func WriteHour1(t time.Time, d StaInfo, db *gorm.DB) {
	var tb Hour1TradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tb.Hour = t.Hour()
	tmp := fmt.Sprintf("%04d%02d%02d%02d", tb.Year, tb.Month, tb.Day, tb.Hour)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
func WriteHour4(t time.Time, d StaInfo, db *gorm.DB) {
	var tb Hour4TradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tb.Hour = t.Hour()
	tmp := fmt.Sprintf("%04d%02d%02d%02d", tb.Year, tb.Month, tb.Day, tb.Hour)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
func WriteDay(t time.Time, d StaInfo, db *gorm.DB) {
	var tb DayTradeTable
	//time
	tb.Day = t.Day()
	tb.Month = int(t.Month())
	tb.Year = t.Year()
	tmp := fmt.Sprintf("%04d%02d%02d", tb.Year, tb.Month, tb.Day)
	tb.TimeKey, _ = strconv.Atoi(tmp)
	//data
	tb.BuyAmount = d.BuyAmount
	tb.SellAmount = d.SellAmount
	tb.StartPrice = d.StartPrice
	tb.EndPrie = d.EndPrie
	tb.CoinType = d.CoinType
	tb.Prop = d.Prop
	tb.HighPrice = d.HighPrice
	tb.HighTs = d.HighTs
	tb.LowPrice = d.LowPrice
	tb.LowTs = d.LowTs
	//db.Create(&tb)
}
