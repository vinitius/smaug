package listeners

import (
	"encoding/json"
	"errors"
	"testing"

	internalerrors "github.com/vinitius/smaug/internal/errors"

	"github.com/stretchr/testify/assert"
	"github.com/vinitius/smaug/internal/aggregates"
	"github.com/vinitius/smaug/internal/domain"

	"github.com/vinitius/smaug/test/mocks"
)

const (
	PublishFunc = "Publish"
	ReadFunc    = "Read"
)

func TestMatchListener(t *testing.T) {
	t.Run("listen successfully", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		matchJson := []byte(`{"type": "match", "price": "1000.00", "size": "1.00", "product_id": "BTC-USD"}`)
		var expectedMatch domain.Match
		err := json.Unmarshal(matchJson, &expectedMatch)
		if err != nil {
			t.Fatalf("could not parse expected match: %v", err)
		}
		expectedAggregate := aggregates.MatchAggregate{
			ProductID:         "BTC-USD",
			SlidingWindowSize: windowSize,
		}
		err = expectedMatch.Validate()
		if err != nil {
			t.Fatalf("could not validate expected match: %v", err)
		}
		expectedAggregate.Add(expectedMatch)

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, matchJson, nil)
		pubMock.On(PublishFunc, &expectedAggregate).Return()

		err = underTest.Listen()

		assert.Nil(t, err)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertCalled(t, PublishFunc, &expectedAggregate)
	})

	t.Run("listen with read error", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		expectedError := errors.New("read error")

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, nil, expectedError)

		err := underTest.Listen()

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertNotCalled(t, PublishFunc)
	})

	t.Run("listen to invalid message", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		invalidJson := []byte(`{"type'''''': /match", "price": "1000.00", "size": "1.00", "product_id": "BTC-USD"}`)
		var expectedError internalerrors.ErrParseMatch

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, invalidJson, nil)

		err := underTest.Listen()

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, expectedError)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertNotCalled(t, PublishFunc)
	})

	t.Run("listen to `type` other than `match`", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		matchJson := []byte(`{"type": "any_other_type", "price": "1000.00", "size": "1.00", "product_id": "BTC-USD"}`)

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, matchJson, nil)

		err := underTest.Listen()

		assert.Nil(t, err)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertNotCalled(t, PublishFunc)
	})

	t.Run("listen to unsupported product", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		matchJson := []byte(`{"type": "match", "price": "1000.00", "size": "1.00", "product_id": "UNKNOWN"}`)

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, matchJson, nil)

		err := underTest.Listen()

		assert.Nil(t, err)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertNotCalled(t, PublishFunc)
	})

	t.Run("listen to invalid prices/sizes", func(t *testing.T) {
		windowSize := 10
		pubMock := new(mocks.VWAPPublisher)
		socketMock := new(mocks.VWAPWebSocket)
		products := []string{"BTC-USD", "ETH-USD", "ETH-BTC"}
		matchJson := []byte(`{"type": "match", "price": "foo", "size": "bah", "product_id": "BTC-USD"}`)

		underTest := NewMatchListener(socketMock, pubMock, windowSize, products, false)
		socketMock.On(ReadFunc).Return(0, matchJson, nil)

		err := underTest.Listen()

		assert.Nil(t, err)
		socketMock.AssertCalled(t, ReadFunc)
		pubMock.AssertNotCalled(t, PublishFunc)
	})
}
