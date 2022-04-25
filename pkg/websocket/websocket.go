package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vinitius/smaug/internal/domain"
)

type VWAPWebSocket interface {
	Connect(url string) (func(), error)
	Read() (int, []byte, error)
	Subscribe(channels, productIDs []string) error
	Close() error
}

type CoinbaseWebSocket struct {
	conn *websocket.Conn
}

func NewCoinbaseWebSocket() CoinbaseWebSocket {
	return CoinbaseWebSocket{}
}

func (c *CoinbaseWebSocket) Connect(url string) (func(), error) {
	log.Printf("connecting to %s", url)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	return func() {
		_ = conn.Close()
	}, nil
}

func (c CoinbaseWebSocket) Read() (int, []byte, error) {
	return c.conn.ReadMessage()
}

func (c CoinbaseWebSocket) Subscribe(channels, productIDs []string) error {
	log.Printf("subscribing to %s for the following products: %s", channels, productIDs)
	sub, err := json.Marshal(domain.Subscription{
		Type:       "subscribe",
		ProductIDs: productIDs,
		Channels:   channels,
	})
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, sub)
	if err != nil {
		return err
	}

	return nil
}

func (c CoinbaseWebSocket) Close() error {
	err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	return nil
}
