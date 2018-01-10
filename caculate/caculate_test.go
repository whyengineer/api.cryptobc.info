package caculate

import (
	"log"
	"testing"
	"time"

	"github.com/whyengineer/api.cryptobc.info/market"
)

func Test_caculate(t *testing.T) {
	a, err := market.New([]string{"huobi"}, []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	_, err = New(a, "huobi")
	if err != nil {
		log.Println(err)
	}
	log.Println("start")
	for {
		time.Sleep(time.Minute)
	}
}
