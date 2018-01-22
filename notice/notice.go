package notice

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/whyengineer/api.cryptobc.info/caculate"
	"github.com/whyengineer/api.cryptobc.info/market"
)

var CalRes *caculate.Cal

type NoticeType struct {
	Coin        string
	StartPrice  float64
	Min5Price   float64
	DayPercent  float64
	Min5Percent float64
	NowPrice    float64
	SecD        chan caculate.StaInfo
}

var Notice *NoticeType

func NewNotic(c *caculate.Cal) {
	CalRes = c
	for _, val := range CalRes.M.Pairs {
		Notice = new(NoticeType)
		Notice.SecD = make(chan caculate.StaInfo)
		Notice.Coin = val
		key := "notice" + val
		c.SendS.Sub(key, Notice.SecD)
		Notice.NoticeCT()
	}
}

func (n *NoticeType) NoticeCT() {
	//get the start price
	nowt := time.Now()
	year := nowt.Year()
	month := int(nowt.Month())
	day := nowt.Day() - 1
	tmp := fmt.Sprintf("%04d%02d%02d", year, month, day)
	key, _ := strconv.Atoi(tmp)
	var d caculate.DayTradeTable
	if CalRes.Db.Where("time_key<=? AND prop=? AND coin_type=?",
		key, "huobi", n.Coin).Order("time_key desc").
		Limit(1).Find(&d).RecordNotFound() {
		log.Println("not find the start price")
	} else {
		n.StartPrice = d.EndPrie
	}
	//get the last hour price,must wait the hot data ready
	time.Sleep(360 * time.Second)
	lastts := nowt.Unix() - 5*60
	for i := lastts; i < 50; i++ {
		key := n.Coin + ":" + strconv.FormatInt(i, 10)
		price, ok := CalRes.HotData.Load(key)
		if ok {
			info := price.(market.CoinInfo)
			n.Min5Price = info.Price
			break
		}
	}

	//start co
	go func() {
		for {
			data := <-n.SecD
			if n.Coin == data.CoinType {
				//log.Println(data.StartPrice)
				n.NowPrice = data.StartPrice
				n.DayPercent = (n.NowPrice - n.StartPrice) / n.StartPrice * 100.0
				if n.Min5Price != 0 {
					n.Min5Percent = (n.NowPrice - n.Min5Price) / n.Min5Price * 100.0
					log.Println(n.Coin, n.Min5Price, n.NowPrice, n.Min5Percent)
				}

			}
		}

	}()

}
