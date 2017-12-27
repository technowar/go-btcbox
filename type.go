package btcbox

import "github.com/shopspring/decimal"

// Ticker ...
//{"high":1914403,"low":1705274,"buy":1893402,"sell":1902418,"last":1897549,"vol":18490.9318}
type Ticker struct {
	High decimal.Decimal `json:"high"`
	Low  decimal.Decimal `json:"low "`
	Buy  decimal.Decimal `json:"buy"`
	Sell decimal.Decimal `json:"sell"`
	Last decimal.Decimal `json:"last"`
	Vol  decimal.Decimal `json:"vol"`
}

// Balance ...
// {"uid":118029,"nameauth":1,"moflag":1,"btc_balance":0.008,"btc_lock":0,"bch_balance":1.1,"bch_lock":0,"ltc_balance":0.999774,"ltc_lock":0,
// "doge_balance":0,"doge_lock":0,"eth_balance":1.998236,"eth_lock":0,
//"jpy_balance":1218.298,"jpy_lock":0}
type Balance struct {
	UID         int64           `json:"uid"`
	Nameauth    int64           `json:"nameauth"`
	Moflag      int64           `json:"moflag"`
	BTCBalance  decimal.Decimal `json:"btc_balance"`
	BTCLock     decimal.Decimal `json:"btc_lock"`
	BCHBalance  decimal.Decimal `json:"bch_balance"`
	BCHLock     decimal.Decimal `json:"bch_lock"`
	LTCBalance  decimal.Decimal `json:"ltc_balance"`
	LTCLock     decimal.Decimal `json:"ltc_lock"`
	DOGEBalance decimal.Decimal `json:"doge_balance"`
	DOGELock    decimal.Decimal `json:"doge_lock"`
	ETHBalance  decimal.Decimal `json:"eth_balance"`
	ETHLock     decimal.Decimal `json:"eth_lock"`
	JPYBalance  decimal.Decimal `json:"jpy_balance"`
	JPYLock     decimal.Decimal `json:"jpy_lock"`
}
