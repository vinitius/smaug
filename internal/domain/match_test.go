package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMatch(t *testing.T) {
	tests := []struct {
		name        string
		match       Match
		wantedPrice float64
		wantedSize  float64
		err         bool
	}{
		{
			name: "valid match",
			match: Match{
				Price: "1000.00",
				Size:  "1.00",
			},
			wantedSize:  1.00,
			wantedPrice: 1000.00,
			err:         false,
		},
		{
			name: "invalid price",
			match: Match{
				Price: "foo",
				Size:  "1.00",
			},
			wantedSize:  0.00,
			wantedPrice: 0.00,
			err:         true,
		},
		{
			name: "invalid size",
			match: Match{
				Price: "1000.00",
				Size:  "foo",
			},
			wantedSize:  0.00,
			wantedPrice: 1000.00,
			err:         true,
		},
		{
			name: "zeroed price",
			match: Match{
				Price: "0.00",
				Size:  "1.00",
			},
			wantedSize:  1.00,
			wantedPrice: 0.00,
			err:         true,
		},
		{
			name: "zeroed size",
			match: Match{
				Price: "1000.00",
				Size:  "0.00",
			},
			wantedSize:  0.00,
			wantedPrice: 1000.00,
			err:         true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if err := c.match.Validate(); err != nil {
				assert.True(t, c.err)
			} else {
				assert.False(t, c.err)
				assert.Nil(t, err)
			}
			assert.Equal(t, c.wantedPrice, c.match.ActualPrice)
			assert.Equal(t, c.wantedSize, c.match.ActualSize)
		})
	}
}
