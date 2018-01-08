package caculate

import (
	"strconv"
	"time"

	"github.com/whyengineer/api.cryptobc.info/market"
)

func StartCaculate(a interface{}) {
	//huobi
	huobi, ok := a.(market.HuobiMarket)
	if ok {
		pairl := len(huobi.Pair)
		keyl := make([]string, pairl)
		go func() {
			secondTick := time.Tick(time.Second)
			for {
				nowTime := <-secondTick
				ts := int64(nowTime.Unix())
				for i, val := range huobi.Pair {
					keyl[i] = val + ":" + strconv.FormatInt(ts, 10)
				}
			}
		}()
	}
}
