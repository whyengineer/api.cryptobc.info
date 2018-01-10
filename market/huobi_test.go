package market

import (
	"log"
	"testing"
)

func Test_huobi(t *testing.T) {
	a, err := NewHuobiMarket("wss://api.huobi.pro/ws", []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	for {
		b := <-a
		log.Println(b)
	}
}
