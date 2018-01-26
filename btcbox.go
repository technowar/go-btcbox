package btcbox

import (
	"encoding/json"
	"errors"
	"log"
	"time"
	//simplejson "github.com/bitly/go-simplejson"
)

const (
	// APIURL btcbox API endpoint
	APIURL = "https://www.btcbox.co.jp/api/v1"
)

func init() {
	log.SetFlags(log.Lshortfile)

}

// New returns an instantiated BTCBox struct
func New(apiKey, apiSecret string) *BTCBox {
	client := NewClient(apiKey, apiSecret)
	return &BTCBox{client}
}

// NewWithCustomTimeout returns an instantiated BTCBox struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *BTCBox {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &BTCBox{client}
}

// BTCBox represent a BTCBox client
type BTCBox struct {
	client *Client
}

// GetBalance ..
func (b *BTCBox) GetBalance() (balance Balance, r []byte, err error) {
	payload := map[string]string{}
	//payload["limit"] = fmt.Sprintf("%d", limit)
	r, err = b.client.do("POST", "balance/", payload, true)
	log.Printf("r:%s", string(r))
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &balance); err != nil {
		return
	}

	if len(balance.Code) == 0 {
		balance.Result = true
		balance.Code = "200"
		return
	}

	return balance, nil, errors.New("Authentication failed")
}

// GetTicker ..
func (b *BTCBox) GetTicker() (ticker Ticker, r []byte, err error) {
	r, err = b.client.do("GET", "ticker/", nil, false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &ticker); err != nil {
		return
	}
	return

}
