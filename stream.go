package bitcoinCentral

import (
	"encoding/json"
	//"fmt"
	"github.com/toorop/go-socket.io"
	"log"
)

// Ticker
type Ticker struct {
	/*
	   high: 389.9,
	        low: 367,
	        volume: 106.89926211,
	        bid: 371,
	        ask: 374.9,
	        midpoint: 372.95,
	        vwap: 377.73505884,
	        at: 1396075634,
	        price: 374.9,
	        variation: -0.6098,
	        currency: 'EUR'
	*/

	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	Midpoint  float64 `json:"midpoint"`
	Vwap      float64 `json:"vwap"`
	At        int64   `json:"at"`
	Price     float64 `json:"price"`
	Variation float64 `json:"variation"`
	Currency  string  `json:"currency"`
}

// Bid
type Bid struct {
	Timestamp int64   `json:"timestamp"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Category  string  `json:"category"`
}

// Ask
type Ask struct {
	Timestamp int64   `json:"timestamp"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Category  string  `json:"category"`
}

/*
trades:
   [ { price: 298.00001,
       traded_btc: 0.26845636,
       timestamp: 1397409616000,
       currency: 'EUR' } ] }
*/

// Trade
type Trade struct {
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
	Currency  string  `json:"currency"`
	TradedBtc float64 `json:"traded_btc"`
}

func GetStream(tickersChan *chan Ticker, bidsChan *chan Bid, asksChan *chan Ask, tradesChan *chan Trade) {
	client, err := socketio.Dial(WS_ENDPOINT, WS_RESSOURCE)
	if err != nil {
		log.Fatalln(err)
	}
	client.On("connect", func(ns *socketio.NameSpace) {
		log.Println("connected")
	})
	err = client.On("stream", func(ns *socketio.NameSpace, stream map[string]*json.RawMessage) {
		//log.Println("On a un stream")
		//log.Println(ns)
		//log.Println(stream)
		if stream["ticker"] != nil {
			var ticker Ticker
			err := json.Unmarshal(*stream["ticker"], &ticker)
			if err != nil {
				log.Println(err)
			}
			//log.Println(ticker)
			*tickersChan <- ticker
			//log.Println(fmt.Sprintf("High: %.6f Low: %.6f Bid: %.6f Ask: %.6f Price: %.6f Currency: %s", ticker.High, ticker.Low, ticker.Bid, ticker.Ask, ticker.Price, ticker.Currency))
		}
		if stream["bids"] != nil {
			var bids []Bid
			err := json.Unmarshal(*stream["bids"], &bids)
			if err != nil {
				log.Println(err)
			}
			for _, bid := range bids {
				*bidsChan <- bid
			}
		}
		if stream["asks"] != nil {
			var asks []Ask
			err := json.Unmarshal(*stream["asks"], &asks)
			if err != nil {
				log.Println(err)
			}
			for _, ask := range asks {
				*asksChan <- ask
			}
		}
		if stream["trades"] != nil {
			var trades []Trade
			err := json.Unmarshal(*stream["trades"], &trades)
			if err != nil {
				log.Println(err)
			}
			for _, trade := range trades {
				*tradesChan <- trade
			}
		}

	})

	if err != nil {
		log.Println("ERROR")
		log.Fatalln(err)
	}
	client.Run()
}
