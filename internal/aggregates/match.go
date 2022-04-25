package aggregates

import (
	"log"

	"github.com/vinitius/smaug/internal/domain"
)

type VWAPAggregate interface {
	Add(value interface{})
	CheckWindowSize()
	VWAP() float64
}

type MatchAggregate struct {
	ProductID         string
	sizeTotal         float64
	priceTotal        float64
	SlidingWindowSize int
	matches           []domain.Match
}

func (a *MatchAggregate) Add(value interface{}) {
	match, ok := value.(domain.Match)
	if !ok {
		return
	}

	a.matches = append(a.matches, match)
	a.sizeTotal += match.ActualSize
	a.priceTotal += match.ActualPrice * match.ActualSize
}

func (a *MatchAggregate) CheckWindowSize() {
	if len(a.matches) == a.SlidingWindowSize {
		log.Printf("==VWAP Reached the window limit of %d occurrences: removing the oldest for %s\n", a.SlidingWindowSize, a.ProductID)
		firstMatch := a.matches[0]
		a.sizeTotal -= firstMatch.ActualSize
		a.priceTotal -= firstMatch.ActualPrice * firstMatch.ActualSize
	}
}

func (a MatchAggregate) VWAP() float64 {
	return a.priceTotal / a.sizeTotal
}
