package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/whyengineer/api.cryptobc.info/caculate"
	"github.com/whyengineer/api.cryptobc.info/market"
)

func main() {

	m, err := market.New([]string{"huobi"}, []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	_, err = caculate.New(m, "huobi")
	if err != nil {
		log.Println(err)
	}
	//notice.NewNotic(cal)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
}
