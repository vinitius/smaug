package aggregates

import (
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
		firstMatch := a.matches[0]
		a.sizeTotal -= firstMatch.ActualSize
		a.priceTotal -= firstMatch.ActualPrice * firstMatch.ActualSize
		a.matches = append(a.matches[:0], a.matches[1:]...)
	}
}

func (a MatchAggregate) VWAP() float64 {
	return a.priceTotal / a.sizeTotal
}
