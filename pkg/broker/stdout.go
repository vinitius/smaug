package broker

import (
	"log"

	"github.com/vinitius/smaug/internal/aggregates"
)

type VWAPMessageBrokerClient interface {
	Send(aggregate aggregates.VWAPAggregate)
}

type StdoutBrokerClient struct{}

func NewStdoutBrokerClient() StdoutBrokerClient {
	return StdoutBrokerClient{}
}

func (s StdoutBrokerClient) Send(aggregate aggregates.VWAPAggregate) {
	log.Printf("===[Publishing] Current VWAP for [%s] is: %0.2f\n", aggregate.ID(), aggregate.VWAP())
}
