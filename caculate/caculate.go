package caculate

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/whyengineer/api.cryptobc.info/market"
)

var dbChan chan map[string]StaInfo

type Cal struct {
	HotEx   string
	HotData map[string]market.CoinInfo //second hot data
	Db      *gorm.DB
	M       *market.Market

	eachPC map[string]chan market.CoinInfo
	eachDC map[string]chan market.CoinInfo

	min1  map[string]StaInfo
	min5  map[string]StaInfo
	min30 map[string]StaInfo
	hour1 map[string]StaInfo
	hour4 map[string]StaInfo
	day   map[string]StaInfo
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
	a.eachPC = make(map[string]chan market.CoinInfo)
	//each every chan
	for _, val := range m.Piars {
		a.eachDC[val] = make(chan market.CoinInfo, 5)
	}
	//each every platform
	for _, val := range m.ExP {
		a.eachPC[val] = make(chan market.CoinInfo, 5)
	}
	// a.min1 = new(map[string]StaInfo)
	// a.min5 = new(map[string]StaInfo)
	// a.min30 = new(map[string]StaInfo)
	// a.hour1 = new(map[string]StaInfo)
	// a.hour4 = new(map[string]StaInfo)
	// a.day = new(map[string]StaInfo)

	//db channel
	dbChan := make(chan map[string]StaInfo, 10)
	//conect Db
	a.Db, err = gorm.Open("mysql", "test:12345678@tcp(123.56.216.29:3306)/coins?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	//migrate the table
	a.Migrate()
	//
}
func (m *Cal) Migrate() {
	m.Db.AutoMigrate(&Min1TradeTable{}, &Min5TradeTable{}, &Min30TradeTable{}, &Hour1TradeTable{}, &Hour4TradeTable{}, &DayTradeTable{})
}
func (m *Cal) CalStaSend(sta StaInfo, num int) {
	if num != 0 {

	}
}
func (c *Cal) Distribute() {
	c.DistributePlat()
	for _, plat := range c.M.ExP {
		c.DistributeCoin(plat)
	}

}
func (c *Cal) Calculate() {

}
func (c *Cal) DistributeCoin(flat string) {
	go func() {
		for {
			//get the data from the platform channle
			coin := <-m.eachPC[plat]
			//save hot data
			m.eachDC[coin.CoinType] <- coin
		}

	}()
}
func (c *Cal) DistributePlat() {
	for plat, datac := range m.M.DataCh {
		go func() {
			for {
				//the exchang data
				each := <-datac
				//save hot data
				if m.HotEx == plat {
					key := each.CoinType + ":" + strconv.FormatInt(each.Ts, 10)
					m.HotData[key] = each
				}
				//to each plat
				m.eachPC[plat] <- each
			}

		}()
	}
}
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
					log.Println(data)
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
