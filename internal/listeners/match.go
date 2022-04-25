package listeners

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/vinitius/smaug/internal/aggregates"
	"github.com/vinitius/smaug/internal/domain"
	internalerrors "github.com/vinitius/smaug/internal/errors"
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
	forever    bool
}

func NewMatchListener(socket websocket.VWAPWebSocket, pub publishers.VWAPPublisher, slidingWindowSize int, productIDs []string, forever bool) MatchListener {
	index := make(map[string]aggregates.VWAPAggregate, len(productIDs))
	for _, id := range productIDs {
		index[id] = &aggregates.MatchAggregate{ProductID: id, SlidingWindowSize: slidingWindowSize}
	}

	return MatchListener{socket: socket, publisher: pub, aggregates: index, forever: forever}
}

func (l MatchListener) Listen() error {
	var errParseMatch internalerrors.ErrParseMatch
	for l.forever {
		if err := l.receiveIncomingMatches(); err != nil {
			if errors.Is(err, errParseMatch) {
				continue
			}
			return err
		}
	}

	return l.receiveIncomingMatches()
}

func (l MatchListener) receiveIncomingMatches() error {
	var match domain.Match
	_, rawMessage, err := l.socket.Read()
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage, &match)
	if err != nil {
		log.Println("parse error:", err)
		return internalerrors.ErrParseMatch{}
	}

	if match.Type == MatchType {
		l.handle(match)
	}

	return nil
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
