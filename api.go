package main

import(
	//"github.com/whyengineer/api.cryptobc.info/notice"
	"github.com/whyengineer/api.cryptobc.info/httpd"
	"log"
	"os"
	"os/signal"
	"github.com/whyengineer/api.cryptobc.info/market"
)

func main(){
	httpd.HttpdCT()
	log.Println("start httpd")
	huobi:=market.NewHuobiMarket("wss://api.huobi.pro/ws",[]string{"btcusdt","ethusdt","eosusdt"})
	err:=huobi.Connect()
	if err!=nil{
		log.Println(err)
		return
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,os.Interrupt)
	<-sigs
}