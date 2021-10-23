package entity

import (
	"time"
)

type LastPrice struct {
	Price string
	Font  string
	Year  int
}

type Stock struct {
	Code      string
	LastPrice []LastPrice
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewStock(code string) (*Stock, error) {
	s := &Stock{
		Code:      code,
		CreatedAt: time.Now(),
	}

	return s, nil
}

func (s *Stock) addStock(price, font string) error {
	s.LastPrice = append(s.LastPrice, LastPrice{
		Price: price,
		Font:  font,
		Year:  time.Now().Year(),
	})
	return nil
}
