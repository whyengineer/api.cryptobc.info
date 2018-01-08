package market


type CoinInfo struct {
	CoinType    string
	Amount      float64
	Price       float64
	Dir         string
	Ts          int64
	Prop        string
}



type CalInfo struct{
	CoinType string
	BuyAmount float64
	SellAmount float64
	Price float64
}