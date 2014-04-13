package main

import (
	"fmt"
	"github.com/toorop/go-bitcoin-central"
	"log"
)

func main() {
	tickersChan := make(chan bitcoinCentral.Ticker)
	bidsChan := make(chan bitcoinCentral.Bid)
	asksChan := make(chan bitcoinCentral.Ask)
	tradesChan := make(chan bitcoinCentral.Trade)
	go func() {
		for {
			select {
			case ticker := <-tickersChan:
				log.Println(fmt.Sprintf("TICKER - Timestamp: %d High: %.6f Low: %.6f Bid: %.6f Ask: %.6f Price: %.6f Currency: %s Volume: %.8f, Midpoint: %.6f, Variation: %.6f, Volume-weighted average price: %.6f", ticker.At, ticker.High, ticker.Low, ticker.Bid, ticker.Ask, ticker.Price, ticker.Currency, ticker.Volume, ticker.Midpoint, ticker.Variation, ticker.Vwap))
				break
			case bid := <-bidsChan:
				log.Println(fmt.Sprintf("BID - Timestamp: %d Amount: %.6f Price: %.6f Currency: %s Category: %s", bid.Timestamp, bid.Amount, bid.Price, bid.Currency, bid.Category))
				break
			case ask := <-asksChan:
				log.Println(fmt.Sprintf("ASK - Timestamp: %d Amount: %.6f Price: %.6f Currency: %s Category: %s", ask.Timestamp, ask.Amount, ask.Price, ask.Currency, ask.Category))
				break
			case trade := <-tradesChan:
				log.Println(fmt.Sprintf("ASK - Timestamp: %d Traded BTC: %.6f Price: %.6f Currency: %s", trade.Timestamp, trade.TradedBtc, trade.Price, trade.Currency))
				break
			}
		}
	}()
	bitcoinCentral.GetStream(&tickersChan, &bidsChan, &asksChan, &tradesChan)

}
