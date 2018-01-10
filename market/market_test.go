package market

import (
	"log"
	"testing"
)

func Test_market(t *testing.T) {
	a, err := New([]string{"huobi"}, []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	for {
		b := <-a.DataCh["huobi"]
		log.Println(b)
	}
}
