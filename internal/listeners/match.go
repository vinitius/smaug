package listeners

import (
	"encoding/json"
	"log"

	"github.com/vinitius/smaug/internal/aggregates"
	"github.com/vinitius/smaug/internal/domain"
	"github.com/vinitius/smaug/internal/publishers"
	"github.com/vinitius/smaug/pkg/websocket"
)

const (
	MatchType = "match"
)

type MatchListener struct {
	publisher  publishers.VWAPPublisher
	socket     websocket.VWAPWebSocket
	aggregates map[string]aggregates.VWAPAggregate
}

func NewMatchListener(socket websocket.VWAPWebSocket, pub publishers.VWAPPublisher, slidingWindowSize int, productIDs []string) MatchListener {
	index := make(map[string]aggregates.VWAPAggregate, len(productIDs))
	for _, id := range productIDs {
		index[id] = &aggregates.MatchAggregate{ProductID: id, SlidingWindowSize: slidingWindowSize}
	}

	return MatchListener{socket: socket, publisher: pub, aggregates: index}
}

func (l MatchListener) Listen() {
	for {
		var match domain.Match
		_, rawMessage, err := l.socket.Read()
		if err != nil {
			log.Println("error reading message: ", err)
			return
		}

		err = json.Unmarshal(rawMessage, &match)
		if err != nil {
			log.Println("error parsing message: ", err)
			continue
		}

		if match.Type == MatchType {
			l.handle(match)
		}
	}
}

func (l MatchListener) handle(match domain.Match) {
	aggregate, found := l.aggregates[match.ProductID]
	if !found {
		log.Println("unsupported product: moving on: ", match.ProductID)
		return
	}

	err := match.Validate()
	if err != nil {
		log.Println("could not validate incoming match values: ", err)
		return
	}

	aggregate.CheckWindowSize()
	aggregate.Add(match)
	l.publisher.Publish(aggregate)
}
