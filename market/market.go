package market

// const(
// 	EXCHANGE_NAME=["huobi"]
// 	HOT_DATA="huobi"
// )
type Market struct {
	DataCh map[string]chan CoinInfo
	Pairs  []string
	ExP    []string
}

// the exl is a array of exchange platform,now support huobi
// the hot is one of the exl,storage the 1 hour data in the memory
func New(exl []string, pairs []string) (*Market, error) {
	a := new(Market)
	a.DataCh = make(map[string]chan CoinInfo)
	a.Pairs = append(a.Pairs, pairs...)
	a.ExP = append(a.ExP, exl...)
	for _, val := range exl {
		if val == "huobi" {
			huobi, err := NewHuobiMarket("wss://api.huobi.pro/ws", pairs)
			if err != nil {
				return nil, err
			}
			a.DataCh[val] = huobi
		}
	}
	return a, nil
}
