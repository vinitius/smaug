package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/vinitius/smaug/pkg/broker"

	"github.com/vinitius/smaug/pkg/config"

	"github.com/vinitius/smaug/internal/listeners"
	"github.com/vinitius/smaug/internal/publishers"

	"github.com/vinitius/smaug/pkg/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Configs & Dependencies
	u := url.URL{Scheme: "wss", Host: config.GetCoinbaseServiceAddress()}
	windowSize := config.GetSlidingWindowSize()
	products := config.GetProducts()
	channels := config.GetChannels()
	socket := websocket.NewCoinbaseWebSocket()
	messageClient := broker.NewStdoutBrokerClient()
	matchListener := listeners.NewMatchListener(&socket, publishers.NewLocalPublisher(messageClient), windowSize, products)

	// Connect
	cleanup, err := socket.Connect(u.String())
	if err != nil {
		log.Fatalf("could not connect to socket: %s", err.Error())
	}
	defer cleanup()

	// Subscribe
	err = socket.Subscribe(channels, products)
	if err != nil {
		log.Fatalf("could not subscribe to channels: %s", err.Error())
	}

	// Listen for `match` events
	done := make(chan bool)
	go func() {
		defer close(done)
		matchListener.Listen()
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
				log.Println("close error:", err)
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
