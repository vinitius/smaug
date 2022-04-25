package domain

import (
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

func (m *Match) ParseValues() error {
	var err error
	m.ActualPrice, err = strconv.ParseFloat(m.Price, 64)
	if err != nil {
		return err
	}

	m.ActualSize, err = strconv.ParseFloat(m.Size, 64)
	if err != nil {
		return err
	}

	return nil
}
