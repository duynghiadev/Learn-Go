package engine

import (
	"time"
	"trading-engine/pkg/uuid"
)

type Fill struct {
	MakerOrderId uuid.UUID
	TakerOrderId uuid.UUID
	Price        float64
	Size         float64
	Timestamp    int64
}

func NewFill(takerOrderId uuid.UUID, makerOrderId uuid.UUID, price, size float64) *Fill {
	return &Fill{
		MakerOrderId: makerOrderId,
		TakerOrderId: takerOrderId,
		Price:        price,
		Size:         size,
		Timestamp:    time.Now().Unix(),
	}
}
