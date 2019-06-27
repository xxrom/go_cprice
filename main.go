package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Hello Natasha - Valerka")
	getPrice()
}

// h2 text-semi-bold details-panel-item--price__value
// https://coinmarketcap.com/currencies/bitcoin/#markets
func onExit() {

}

func getPrice() {
	url := "https://coinmarketcap.com/currencies/bitcoin/#markets"
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	price := doc.Find(".details-panel-item--price__value").Text()

	systray.SetTitle("BTC $" + price)
}
