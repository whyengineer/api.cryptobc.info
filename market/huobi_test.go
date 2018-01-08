package market

import(
	"testing"
	"log"
	"time"
)
func Test_huobi(t *testing.T){
	a:=NewHuobiMarket("wss://api.huobi.pro/ws",[]string{"btcusdt","ethusdt"})
	err:=a.Connect()
	if err != nil {
		log.Println(err)
	}
	for{
		time.Sleep(time.Second*10)
	}
}