package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
  "strconv"
)

type stocksArray []singleStock

type singleStock struct {
	Symbol   string `json:"symbol"`
}

var (
	err      error
	stocks   stocksArray
	response *http.Response
	body     []byte
)

func stockTicker01() { //this one is to get past an initial level of json so if it looks like this [{stuff:stuff}] and has the hard brackets on it this is the best way
  //you have to have stocksArray and singleStock for it to work. also the varialbes you see above

  // Use http://finance.google.com/finance/info?client=ig&q=NASDAQ:GOOG to get a JSON response
	response, err = http.Get("https://api.coinmarketcap.com/v1/ticker/")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// Read the data into a byte slice
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Remove whitespace from response
	data := bytes.TrimSpace(body)

	// Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &stocks)
	if err != nil {
		fmt.Println(err)
    //no idea what most of this stuff does but its important
	}

	fmt.Println(stocks[0].Symbol)
  fmt.Println(stocks[4].Symbol)
  fmt.Println("")
}


//Ticker is a go struct of Bitfinex's json ticker
type Ticker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	//Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

func stockTicker02() {
	client := &http.Client{Timeout: time.Second * 5}
	resp, _ := client.Get("https://api.bitfinex.com/v1/pubticker/btcusd")

	json.NewDecoder(resp.Body).Decode(&ticker)

	lastprice, _ := strconv.ParseFloat(ticker.LastPrice, 16)
	fmt.Println("Last Price: ", lastprice)
	fmt.Println("High: ", ticker.High)
}

var ticker Ticker
func main() {
  for i := 0; i < 11; i++ {
	   stockTicker01() //the one for more complicated json
     stockTicker02()//the one for more simiple json
  }
}
