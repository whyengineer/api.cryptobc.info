package market

type CoinInfo struct {
	CoinType   string  `json:"coin"`
	BuyAmount  float64 `json:"buyamount"`
	SellAmount float64 `json:"sellamount"`
	Price      float64 `json:"price"`
	Ts         int64   `json:"ts"`
}
