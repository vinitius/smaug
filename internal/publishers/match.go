package publishers

import (
	"log"
)

type VWAPPublisher interface {
	Publish(productID string, vwap float64)
}

type LogPublisher struct{}

func NewLogPublisher() LogPublisher {
	return LogPublisher{}
}

func (l LogPublisher) Publish(productID string, vwap float64) {
	log.Printf("===[Publishing] Current VWAP for [%s] is: %0.2f\n", productID, vwap)
}
