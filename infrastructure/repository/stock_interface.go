package repository

import "github.com/riquellopes/closed-stock-consumer/entity"

type StockInterface interface {
	Save(stock entity.Stock)
}
