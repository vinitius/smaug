package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppConfig(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		value    string
		expected interface{}
		call     func() interface{}
	}{
		{
			name:     "product ids",
			env:      "COINBASE_PRODUCT_IDS",
			value:    "BTC-USD",
			expected: []string{"BTC-USD"},
			call:     func() interface{} { return GetProducts() },
		},
		{
			name:     "channels",
			env:      "COINBASE_CHANNELS",
			value:    "matches",
			expected: []string{"matches"},
			call:     func() interface{} { return GetChannels() },
		},
		{
			name:     "service address",
			env:      "COINBASE_SERVICE_ADDRESS",
			value:    "ws-feed.coinbase.com",
			expected: "ws-feed.coinbase.com",
			call:     func() interface{} { return GetCoinbaseServiceAddress() },
		},
		{
			name:     "window size",
			env:      "SLIDING_WINDOW_SIZE",
			value:    "200",
			expected: 200,
			call:     func() interface{} { return GetSlidingWindowSize() },
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv(c.env, c.value)
			assert.Equal(t, c.expected, c.call())
		})
	}
}
