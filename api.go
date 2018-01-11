package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/whyengineer/api.cryptobc.info/caculate"
	"github.com/whyengineer/api.cryptobc.info/httpd"
	"github.com/whyengineer/api.cryptobc.info/market"
)

func main() {
	httpd.HttpdCT()
	log.Println("start httpd")
	m, err := market.New([]string{"huobi"}, []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	_, err = caculate.New(m, "huobi")
	if err != nil {
		log.Println(err)
	}
	log.Println("start")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
}
