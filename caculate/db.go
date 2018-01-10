package caculate

type Min1TradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Min5TradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Min30TradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	Min        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Hour1TradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type Hour4TradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	Hour       int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
type DayTradeTable struct {
	Prop       string
	CoinType   string
	Year       int
	Month      int
	Day        int
	BuyAmount  float64
	SellAmount float64
	HighPrice  float64
	HighTs     int64
	LowPrice   float64
	LowTs      int64
	StartPrice float64
	EndPrie    float64
}
