package domain

import (
	"fmt"
	"strconv"
)

type Match struct {
	Type        string  `json:"type"`
	TradeID     int64   `json:"trade_id"`
	Sequence    int64   `json:"sequence"`
	Time        string  `json:"time"`
	ProductID   string  `json:"product_id"`
	Side        string  `json:"side"`
	Price       string  `json:"price"`
	Size        string  `json:"size"`
	ActualPrice float64 `json:"-"`
	ActualSize  float64 `json:"-"`
}

func (m *Match) Validate() error {
	var err error
	m.ActualPrice, err = strconv.ParseFloat(m.Price, 64)
	if err != nil {
		return err
	}

	m.ActualSize, err = strconv.ParseFloat(m.Size, 64)
	if err != nil {
		return err
	}

	if m.ActualSize <= 0 || m.ActualPrice <= 0 {
		return fmt.Errorf("invalid Match values for price and/or size: %v", m)
	}

	return nil
}
