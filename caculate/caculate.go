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
	HotEx   string
	HotData map[string]market.CoinInfo //second hot data
	Db      *gorm.DB
	M       *market.Market

	eachDC map[string]chan market.CoinInfo
}

type StaInfo struct {
	Prop       string
	CoinType   string
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}

//new market data
func New(m *market.Market, hot string) (*Cal, error) {
	var err error
	a := new(Cal)
	a.HotData = make(map[string]market.CoinInfo)
	a.HotEx = hot
	a.M = m
	a.eachDC = make(map[string]chan market.CoinInfo)
	//each every chan
	for _, plat := range m.ExP {
		for _, coin := range m.Pairs {
			key := plat + coin
			a.eachDC[key] = make(chan market.CoinInfo, 5)
			//log.Println("used pairs:", plat, coin)
			a.Calculate(a.eachDC[key], plat, coin)
		}
	}

	//conect Db
	a.Db, err = gorm.Open("mysql", "test:12345678@tcp(123.56.216.29:3306)/coins?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	//migrate the table
	a.Migrate()
	//
	a.DistributeCoin()

	return a, err
}
func (c *Cal) Migrate() {
	c.Db.AutoMigrate(&Min1TradeTable{}, &Min5TradeTable{}, &Min30TradeTable{}, &Hour1TradeTable{}, &Hour4TradeTable{}, &DayTradeTable{})
}

func (c *Cal) Calculate(data chan market.CoinInfo, plat string, coin string) {

	go func() {
		// cointype:=coin
		// theplat:=plat
		// var min1, min5, min30, hour1, hour4, day StaInfo
		stal := map[string]*StaInfo{
			"min1":  &StaInfo{},
			"min5":  &StaInfo{},
			"min30": &StaInfo{},
			"hour1": &StaInfo{},
			"hour4": &StaInfo{},
			"day":   &StaInfo{}}

		for _, val := range stal {
			val.CoinType = coin
			val.Prop = plat
		}
		//start on a zero minute point
		now := time.Now()
		next := now.Add(time.Minute)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
		start := time.NewTimer(next.Sub(now))
		var coin market.CoinInfo
		for {
			select {
			case <-start.C:

				//render the start point
				min1Tick := time.Tick(time.Minute)
				log.Println("zero point")
				for _, val := range stal {
					val.StartPrice = coin.Price
				}
				go func() {
					for {
						select {
						case nowt := <-min1Tick:
							//handle min1
							t1 := stal["min1"]
							t1.EndPrie = coin.Price
							//save
							WriteMin1(nowt, *t1, c.Db)
							//refresh
							t1.StartPrice = coin.Price
							t1.BuyAmount = 0
							t1.SellAmount = 0
							t1.HighPrice = 0
							t1.LowPrice = 0
							//handle min5
							if nowt.Minute()%5 == 0 {
								t5 := stal["min5"]
								t5.EndPrie = coin.Price
								//save
								WriteMin5(nowt, *t5, c.Db)
								//refresh
								t5.StartPrice = coin.Price
								t5.BuyAmount = 0
								t5.SellAmount = 0
								t5.HighPrice = 0
								t5.LowPrice = 0
							}
							if nowt.Minute()%30 == 0 {
								t30 := stal["min30"]
								t30.EndPrie = coin.Price
								//save
								WriteMin30(nowt, *t30, c.Db)
								//refresh
								t30.StartPrice = coin.Price
								t30.BuyAmount = 0
								t30.SellAmount = 0
								t30.HighPrice = 0
								t30.LowPrice = 0
							}
							if nowt.Minute() == 0 {
								th1 := stal["hour1"]
								th1.EndPrie = coin.Price
								//save
								WriteHour1(nowt, *th1, c.Db)
								//refresh
								th1.StartPrice = coin.Price
								th1.BuyAmount = 0
								th1.SellAmount = 0
								th1.HighPrice = 0
								th1.LowPrice = 0
							}
							if nowt.Hour()%4 == 0 {
								th4 := stal["hour4"]
								th4.EndPrie = coin.Price
								//save
								WriteHour4(nowt, *th4, c.Db)
								//refresh
								th4.StartPrice = coin.Price
								th4.BuyAmount = 0
								th4.SellAmount = 0
								th4.HighPrice = 0
								th4.LowPrice = 0
							}
							if nowt.Hour() == 0 {
								td := stal["day"]
								td.EndPrie = coin.Price
								//save
								WriteDay(nowt, *td, c.Db)
								//refresh
								td.StartPrice = coin.Price
								td.BuyAmount = 0
								td.SellAmount = 0
								td.HighPrice = 0
								td.LowPrice = 0
							}
						}
					}
				}()

			case coin = <-data:
				//log.Println(coin)
				for _, val := range stal {
					if coin.Price > val.HighPrice || val.HighPrice == 0 {
						val.HighPrice = coin.Price
						val.HighTs = coin.Ts
					}
					if coin.Price < val.LowPrice || val.LowPrice == 0 {
						val.LowPrice = coin.Price
						val.LowTs = coin.Ts
					}
					val.BuyAmount += coin.BuyAmount
					val.SellAmount += coin.SellAmount
					//log.Println(*val)
				}
			}
			//stadata.BuyAmount+=
		}
	}()
}
func (c *Cal) DistributeCoin() {
	//start garbage delete
	go func() {
		time.Sleep(time.Hour)
		secondTick := time.Tick(time.Second)
		for {
			nowt := <-secondTick
			ts := nowt.Unix()
			for _, coin := range c.M.Pairs {
				key := coin + ":" + strconv.FormatInt(ts, 10)
				delete(c.HotData, key)
			}

		}
	}()
	for plat, datac := range c.M.DataCh {
		go func() {
			for {
				//the exchang data
				each := <-datac
				//save hot data
				if c.HotEx == plat {
					key := each.CoinType + ":" + strconv.FormatInt(each.Ts, 10)
					c.HotData[key] = each
				}
				//to each coin
				c.eachDC[plat+each.CoinType] <- each
			}

		}()
	}
}
