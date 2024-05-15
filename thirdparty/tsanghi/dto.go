package tsanghi

type Ticker struct {
	Msg  string       `json:"msg"`
	Code int          `json:"code"`
	Data []TickerData `json:"data"`
}
type TickerData struct {
	Ticker string  `json:"ticker"`
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}
