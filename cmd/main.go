package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/vinitius/smaug/internal/listeners"
	"github.com/vinitius/smaug/internal/publishers"
	"github.com/vinitius/smaug/pkg/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Dependencies & Configs
	socket := websocket.NewCoinbaseWebSocket()
	products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
	channels := []string{"matches"}
	u := url.URL{Scheme: "wss", Host: "ws-feed.exchange.coinbase.com"}

	// Connect
	cleanup, err := socket.Connect(u.String())
	if err != nil {
		log.Panicf("could not connect to socket: %s", err.Error())
	}
	defer cleanup()

	// Subscribe
	err = socket.Subscribe(channels, products)
	if err != nil {
		log.Panicf("could not subscribe to channels: %s", err.Error())
	}

	// Listen for `match` events
	done := make(chan bool)
	go func() {
		defer close(done)
		matchListener := listeners.NewMatchListener(publishers.NewLogPublisher(), 200, products)
		matchListener.Listen(&socket)
	}()

	// Control Panel
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("gracefully shutting down")
			err := socket.Close()
			if err != nil {
				log.Println("close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
