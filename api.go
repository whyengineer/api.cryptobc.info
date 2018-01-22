package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/whyengineer/api.cryptobc.info/caculate"
	"github.com/whyengineer/api.cryptobc.info/httpd"
	"github.com/whyengineer/api.cryptobc.info/market"
	"github.com/whyengineer/api.cryptobc.info/notice"
)

func main() {

	m, err := market.New([]string{"huobi"}, []string{"btcusdt", "ethusdt", "eosusdt"})
	if err != nil {
		log.Println(err)
	}
	cal, err := caculate.New(m, "huobi")
	if err != nil {
		log.Println(err)
	}
	httpd.HttpdCT(cal)
	log.Println("start httpd")
	notice.NewNotic(cal)
	log.Println("start notice")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
}
