go-btcbox
==========

go-btcbox is an implementation of the BTCBOX in Golang.

Based off of https://github.com/toorop/go-bittrex/

## Import
```
import "github.com/rakd/go-btcbox"
```

## Usage
~~~ go
package main

import (
	"fmt"
	"github.com/rakd/go-btcbox"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// btcbox client
	client := btcbox.New(btcbox.APIBaseURLForSGD, API_KEY, API_SECRET)

	// Get ticker
	ticker, err := client.GetTicker()
	fmt.Println(err, ticker)
}
~~~


## Stay tuned

- [Follow me on Twitter](https://twitter.com/kaz_lavender)

## Donate

- BTC: 1Ah8sarQ4w9FnsCs8LoG6JuYiFHmrAAy6F
