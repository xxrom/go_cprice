package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/distatus/battery"
	"github.com/getlantern/systray"
	"github.com/robfig/cron"
)

// h2 text-semi-bold details-panel-item--price__value
// https://coinmarketcap.com/currencies/bitcoin/#markets

type state struct {
	Price string
	Cron  *cron.Cron
}

func main() {
	s := &state{}
	systray.Run(s.onReady, s.onExit)
}

func (s *state) onReady() {
	// systray.SetTitle("Hello Natasha - Valerka")
	s.updatePrice()
	s.Cron = cron.New()
	s.Cron.AddFunc("@every 10s", s.updatePrice)
	s.Cron.Start()
	// getBattery()
}

func (s *state) onExit() {
	s.Cron.Stop()
}

func getBattery() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}
	for i, battery := range batteries {
		fmt.Printf("Bat%d: ", i)
		fmt.Printf("state: %f, ", battery.State)
		fmt.Printf("current capacity: %f mWh, ", battery.Current)
		fmt.Printf("last full capacity: %f mWh, ", battery.Full)
		fmt.Printf("design capacity: %f mWh, ", battery.Design)
		fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
		fmt.Printf("voltage: %f V, ", battery.Voltage)
		fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)

		// systray.SetTitle("B - ", battery.State)
	}
}

func (s *state) updatePrice() {
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
	fmt.Printf(price)
	systray.SetTitle("BTC $" + price)
}
