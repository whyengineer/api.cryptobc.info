package market

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"time"

	"github.com/gorilla/websocket"
)

type HuobiInfo struct {
	CoinType string
	Amount   float64
	Price    float64
	Dir      string
	Ts       int64
}

//huobi json format

type Pong struct {
	Ts float64 `json:"pong"`
}
type Sub struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}
type KLine struct {
	symbol string
	period string
}

type MarketDepth struct {
	symbol string
	depth  string
}

type HuobiDT struct {
	Ch   string     `json:ch`
	Ts   int64      `json:ts`
	Tick HuobiDTone `json:tick`
}
type HuobiDTone struct {
	Id   int64         `json:id`
	Ts   int64         `json:ts`
	Data []HuobiDTone1 `json:data`
}
type HuobiDTone1 struct {
	Id        int64   `json:id`
	Price     float64 `json:price`
	Amount    float64 `json:amount`
	Direction string  `json:"direction"`
	Ts        int64   `json:ts`
}

//channel
var rawdata chan HuobiInfo

type HuobiMarket struct {
	Url  string   //wss://api.huobi.pro/ws
	Pair []string //such as ethusdt btcusdt eosusdt
	Wsc  *websocket.Conn
	Data chan CoinInfo
}

func NewHuobiMarket(url string, pair []string) (chan CoinInfo, error) {
	a := make(chan CoinInfo)

	hm := new(HuobiMarket)
	hm.Url = url
	hm.Pair = append(hm.Pair, pair...)
	hm.Data = a

	rawdata = make(chan HuobiInfo, 10)

	err := hm.Connect()
	if err != nil {
		return a, err
	}
	go hm.calCT()
	return a, nil
}
func (hm *HuobiMarket) SubTopic(topic int, v interface{}) error {
	sub := new(Sub)
	if topic == 1 {
		sub.Id = "id1"
		t := v.(KLine)
		sub.Sub = fmt.Sprintf("market.%s.kline.%s", t.symbol, t.period)
	} else if topic == 2 {
		sub.Id = "id2"
		t := v.(MarketDepth)
		sub.Sub = fmt.Sprintf("market.%s.depth.%s", t.symbol, t.depth)
	} else if topic == 3 {
		sub.Id = "id3"
		sub.Sub = fmt.Sprintf("market.%s.trade.detail", v.(string))
	} else if topic == 4 {
		sub.Id = "id4"
		sub.Sub = fmt.Sprintf("market.%s.detail", v.(string))
	} else {
		return errors.New("invalid topic")
	}
	ret, err := json.Marshal(sub)
	if err != nil {
		log.Println("json parse err:", err)
		return err
	}
	err = hm.Wsc.WriteMessage(websocket.TextMessage, ret)
	return err
}

func (hm *HuobiMarket) Connect() error {
	var err error

	hm.Wsc, _, err = websocket.DefaultDialer.Dial(hm.Url, nil)
	log.Println("start connect huobi websocket")
	if err != nil {
		return err
	}
	go hm.ReadCT()
	for i := range hm.Pair {
		err = hm.SubTopic(3, hm.Pair[i])
		if err != nil {
			return err
		}
	}
	return err
}
func (hm *HuobiMarket) calCT() {
	//realse memory
	//ts := time.Now().Unix()
	pairl := len(hm.Pair)
	num := make([]int, pairl)
	nowts := make([]int64, pairl)
	eachCal := make([]CoinInfo, pairl)
	for {
		data := <-rawdata
		for j, val := range hm.Pair {
			if data.CoinType == val {
				if nowts[j] != data.Ts/1000 {
					if nowts[j] != 0 {
						//write the last calinfo
						eachCal[j].CoinType = val
						eachCal[j].Price /= float64(num[j])
						eachCal[j].Ts = nowts[j]
						//key := val + ":" + strconv.FormatInt(nowts[j], 10)
						//hm.HotData[key] = eachCal[j]
						//log.Println(eachCal[j])
						hm.Data <- eachCal[j]
					}
					nowts[j] = data.Ts / 1000
					num[j] = 0
					eachCal[j].BuyAmount = 0
					eachCal[j].SellAmount = 0
					eachCal[j].Price = 0
				}
				num[j]++
				if data.Dir == "buy" {
					eachCal[j].BuyAmount += data.Amount
				} else {
					eachCal[j].SellAmount += data.Amount
				}
				eachCal[j].Price += data.Price

			}
		}

	}
}
func (hm *HuobiMarket) ReadCT() {
	for {
		_, message, err := hm.Wsc.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			hm.Wsc.Close()
			log.Println("restart connect")
			for {
				err := hm.Connect()
				if err != nil {
					log.Println("open websocket err:", err)
					time.Sleep(time.Second)
				} else {
					break
				}
			}
			return
		}
		reader := bytes.NewReader(message)
		zr, err := gzip.NewReader(reader)
		if err != nil {
			log.Println("zip decompress:", err)
			return
		}
		hm.handlerReceive(zr)
		zr.Close()
	}
}
func (hm *HuobiMarket) handlerReceive(info io.Reader) {
	// io.Copy(os.Stdout, info)
	a, _ := ioutil.ReadAll(info)
	// log.Println(string(a))
	var data map[string]interface{}
	json.Unmarshal(a, &data)
	if _, err := data["ping"]; err {
		ts := data["ping"].(float64)
		ret := &Pong{
			Ts: ts,
		}
		ret1, err := json.Marshal(ret)
		if err != nil {
			log.Println("json encode err:", err)
			return
		}
		err = hm.Wsc.WriteMessage(websocket.TextMessage, ret1)
		if err != nil {
			log.Println("websocket write err:", err)
			return
		}

	}
	if _, err := data["status"]; err {
		log.Println(string(a))
		return
	}
	if _, err := data["ch"]; err {
		trade := HuobiDT{}
		err := json.Unmarshal(a, &trade)
		if err != nil {
			log.Println("json Unmarshal error", err)
		}
		re := regexp.MustCompile(`^market\.(.*?)\.trade\.detail$`)
		a := new(HuobiInfo)
		a.CoinType = re.FindStringSubmatch(trade.Ch)[1]
		for _, v := range trade.Tick.Data {
			a.Ts = v.Ts
			a.Amount = v.Amount
			a.Dir = v.Direction
			a.Price = v.Price

			rawdata <- *a
		}
		return
	}

}
