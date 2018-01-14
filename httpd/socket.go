package httpd

import (
	"encoding/json"
	"log"

	"github.com/googollee/go-socket.io"
	"github.com/whyengineer/api.cryptobc.info/caculate"
)

func NewSocketServer() *socketio.Server {
	//create a connect

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		//log.Println("on connection")
		//log.Println(so.Id())
		//join the data channel
		min1c := make(chan caculate.StaInfo)
		CalRes.Min1S.Sub(so.Id(), min1c)
		min5c := make(chan caculate.StaInfo)
		CalRes.Min5S.Sub(so.Id(), min5c)
		min30c := make(chan caculate.StaInfo)
		CalRes.Min30S.Sub(so.Id(), min30c)
		hour1c := make(chan caculate.StaInfo)
		CalRes.Hour1S.Sub(so.Id(), hour1c)
		hour4c := make(chan caculate.StaInfo)
		CalRes.Hour4S.Sub(so.Id(), hour4c)
		secdc := make(chan caculate.StaInfo)
		CalRes.SendS.Sub(so.Id(), secdc)
		done := make(chan struct{})
		dayc := make(chan caculate.StaInfo)
		CalRes.DayS.Sub(so.Id(), dayc)
		go func() {
			for {
				select {
				case secdd := <-secdc:
					as, _ := json.Marshal(&secdd)
					so.Emit("senc"+secdd.CoinType, string(as))
					//log.Println(secdd)
				case min1d := <-min1c:
					//log.Println(so.Id(),data.CoinType,data.Price,data.BuyAmount,data.SellAmount)
					as, _ := json.Marshal(&min1d)
					so.Emit("min1"+min1d.CoinType, string(as))
				case min30d := <-min30c:
					as, _ := json.Marshal(&min30d)
					so.Emit("min30"+min30d.CoinType, string(as))
				case min5d := <-min5c:
					as, _ := json.Marshal(&min5d)
					so.Emit("min5"+min5d.CoinType, string(as))
				case hour1d := <-hour1c:
					as, _ := json.Marshal(&hour1d)
					so.Emit("hour1"+hour1d.CoinType, string(as))
				case hour4d := <-hour4c:
					as, _ := json.Marshal(&hour4d)
					so.Emit("hour4"+hour4d.CoinType, string(as))
				case dayd := <-dayc:
					as, _ := json.Marshal(&dayd)
					so.Emit("day"+dayd.CoinType, string(as))
				case <-done:
					//og.Println("bye bye")
					return
				}
			}
		}()

		// trade,err:=NewTradeApi()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// trade.TradePush("btcusdt","huobi",so)
		// trade.TradePush("ethusdt","huobi",so)
		// so.Join("chat")
		// so.On("chat_message", func(msg string) {
		// 	so.Emit("chat_message", msg)
		// 	log.Println("emit:",msg)
		// 	trade.GetNowInfo("btcusdt","huobi",10)
		// 	so.BroadcastTo("chat", "chat_message", msg)
		// })
		so.On("disconnection", func() {
			close(done)
			//log.Println("on disconnect")
			CalRes.Min1S.Delete(so.Id())
			CalRes.Min5S.Delete(so.Id())
			CalRes.Min30S.Delete(so.Id())
			CalRes.Hour1S.Delete(so.Id())
			CalRes.Hour4S.Delete(so.Id())
			CalRes.DayS.Delete(so.Id())
			CalRes.SendS.Delete(so.Id())
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)

	})

	return server
}
