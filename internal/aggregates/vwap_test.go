package aggregates

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinitius/smaug/internal/domain"
)

func TestMatchAggregate(t *testing.T) {
	t.Run("get ID successfully", func(t *testing.T) {
		underTest := MatchAggregate{
			ProductID:         "BTC-USD",
			SlidingWindowSize: 10,
		}

		result := underTest.ID()

		assert.Empty(t, underTest.matches)
		assert.Equal(t, "BTC-USD", result)
	})

	t.Run("pass invalid match", func(t *testing.T) {
		underTest := MatchAggregate{
			ProductID:         "BTC-USD",
			SlidingWindowSize: 10,
		}

		underTest.Add("invalid")

		assert.Empty(t, underTest.matches)
	})

	t.Run("calculate VWAP successfully", func(t *testing.T) {
		windowSize := 10
		newMatch := domain.Match{
			Type:        "match",
			TradeID:     10,
			Sequence:    10,
			Time:        "",
			ProductID:   "BTC-USD",
			Side:        "",
			Price:       "6000.00",
			Size:        "1.00",
			ActualSize:  1.00,
			ActualPrice: 6000.00,
		}

		expectedVWAP := 1000.00
		expectedVWAPAfterNewMatch := 1500.00

		underTest := MatchAggregate{
			ProductID:         "BTC-USD",
			SlidingWindowSize: windowSize,
		}

		for i := 0; i < windowSize; i++ {
			match := domain.Match{
				Type:        "match",
				TradeID:     int64(i),
				Sequence:    int64(i),
				Time:        "",
				ProductID:   "BTC-USD",
				Side:        "",
				Price:       "1000.00",
				Size:        "1.00",
				ActualSize:  1.00,
				ActualPrice: 1000.00,
			}
			underTest.Add(match)
			assert.NotEmpty(t, underTest.matches)
			assert.Equal(t, match, underTest.matches[i])
		}

		assert.Equal(t, expectedVWAP, underTest.VWAP())

		oldestMatch := domain.Match{
			Type:        "match",
			TradeID:     0,
			Sequence:    0,
			Time:        "",
			ProductID:   "BTC-USD",
			Side:        "",
			Price:       "1000.00",
			Size:        "1.00",
			ActualSize:  1.00,
			ActualPrice: 1000.00,
		}

		underTest.CheckWindowSize()
		underTest.Add(newMatch)

		assert.NotEmpty(t, underTest.matches)
		assert.Equal(t, windowSize, len(underTest.matches))
		assert.NotContains(t, underTest.matches, oldestMatch, "Expected slice not to contain the oldest match, but it does")
		assert.Equal(t, expectedVWAPAfterNewMatch, underTest.VWAP())
	})
}
