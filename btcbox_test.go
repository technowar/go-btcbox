package btcbox

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTicker(t *testing.T) {

	token := os.Getenv("BTCBOX_KEY")
	secret := os.Getenv("BTCBOX_SECRET")
	api := New(token, secret)

	ticker, err := api.GetTicker()
	log.Printf("err:%v", err)
	require.NoError(t, err, nil)

	log.Printf("err:%v", err)
	require.NoError(t, err, nil)
	log.Printf("ticker.low:%v", ticker.Low)
	log.Printf("ticker.high:%v", ticker.High)
	log.Printf("ticker.last:%v", ticker.Last)
	log.Printf("ticker.buy:%v", ticker.Buy)
	log.Printf("ticker.sell:%v", ticker.Sell)
	return
}
func TestGetBalance(t *testing.T) {
	token := os.Getenv("BTCBOX_KEY")
	secret := os.Getenv("BTCBOX_SECRET")
	api := New(token, secret)

	balance, err := api.GetBalance()
	log.Printf("err:%v", err)
	require.NoError(t, err, nil)
	log.Printf("balance:%v", balance)
	log.Printf("balance.BTCBalance:%v", balance.BTCBalance)
	log.Printf("balance.JPYBalance:%v", balance.JPYBalance)
	log.Printf("balance.ETHBalance:%v", balance.ETHBalance)
	log.Printf("balance.DOGEBalance:%v", balance.DOGEBalance)
	log.Printf("balance.BCHBalance:%v", balance.BCHBalance)
	log.Printf("balance.LTCBalance:%v", balance.LTCBalance)

}
