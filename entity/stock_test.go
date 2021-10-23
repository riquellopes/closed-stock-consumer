package entity_test

import (
	"testing"

	"github.com/riquellopes/closed-stock-consumer/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateStock(t *testing.T) {
	s, err := entity.NewStock("KLBN11")
	assert.Nil(t, err)
	assert.Equal(t, s.Code, "KLBN11")
}
