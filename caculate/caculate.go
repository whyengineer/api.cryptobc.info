package caculate

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whyengineer/api.cryptobc.info/market"
)

type Cal struct {
}

type TradeDayRes struct {
	gorm.Model
	Year          int
	Month         int
	Day           int
	CoinType      string
	LowPrice      float64
	LowPriceTime  time.Time
	HighPrice     float64
	HighPriceTime time.Time
	FinalPrice    float64
}
type TradeInfo struct {
	BuyAmount  float64
	SellAmount float64
	Price      float64
	CoinType   string
	Ts         int64 //for 1min ts/60*60
}

var dayDone chan struct{}
var minDataChan chan TradeInfo
var dayRes TradeDayRes

func writeDb() error {
	year, mon, day := time.Now().Date()
	dayRes.Year = year
	dayRes.Month = int(mon)
	dayRes.Day = day
	//create the connect
	db, err := gorm.Open("mysql", "test:12345678@tcp(123.56.216.29:3306)/coins?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db.AutoMigrate(&TradeInfo{}, &TradeDayRes{})
	go func() {
		defer db.Close()
		for {
			select {
			case mindata := <-minDataChan:
				db.Create(&mindata)
				log.Println(mindata)
				log.Println(dayRes)
			case <-dayDone:
				db.Create(&dayRes)
				log.Println(dayRes)
				year, mon, day := time.Now().Date()
				dayRes.Year = year
				dayRes.Month = int(mon)
				dayRes.Day = day
				dayRes.LowPrice = 0
				dayRes.HighPrice = 0
			}
		}
	}()

	return err

}
func EachCal(cointype string, m interface{}) {
	huobi, ok := m.(*market.HuobiMarket)
	if ok {
		secondTick := time.Tick(time.Second)
		minTick := time.Tick(time.Minute)
		for {
			select {
			case secTime := <-secondTick:
				ts := int64(secTime.Unix()) - 1
				key := cointype + ":" + strconv.FormatInt(ts, 10)
				data, ok := huobi.HotData[key]
				if ok {
					if dayRes.HighPrice < data.Price || dayRes.HighPrice == 0 {
						dayRes.HighPrice = data.Price
						dayRes.HighPriceTime = secTime
					}
					if dayRes.LowPrice > data.Price || dayRes.LowPrice == 0 {
						dayRes.LowPrice = data.Price
						dayRes.LowPriceTime = secTime
					}
				}
			case minTime := <-minTick:
				ts := int64(minTime.Unix())
				var fdata market.CalInfo
				var j int
				for i := ts; i > ts-60; i-- {
					key := cointype + ":" + strconv.FormatInt(i, 10)
					data, ok := huobi.HotData[key]
					if ok {
						fdata.BuyAmount += data.BuyAmount
						fdata.SellAmount += data.SellAmount
						fdata.Price += data.Price
						j++
					}
				}
				if j != 0 {
					fdata.Price /= float64(j)
				}
				var xx TradeInfo
				xx.Ts = ts / 60 * 60 //1mindata
				xx.CoinType = cointype
				xx.BuyAmount = fdata.BuyAmount
				xx.SellAmount = fdata.SellAmount
				xx.Price = fdata.Price
				minDataChan <- xx
			}

		}
	}

}
func Start(a interface{}) {
	minDataChan = make(chan TradeInfo)
	dayDone = make(chan struct{})
	//
	go func() {
		for {
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			dayDone <- struct{}{}
		}

	}()
	//start db
	err := writeDb()
	if err != nil {
		log.Println(err)
		return
	}
	//huobi
	huobi, ok := a.(*market.HuobiMarket)
	if ok {
		go func() {
			for _, val := range huobi.Pair {
				//caculate each coin
				go EachCal(val, a)
			}
		}()
	} else {
		log.Println("format error")
	}
}
