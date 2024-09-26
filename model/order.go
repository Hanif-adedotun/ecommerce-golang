package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID    uint64 `json:"order_id"`
	CustomerID uuid.UUID `json:"customer_id"`
	LineItems  []LineItem `json:"lineitems"`
	CreatedAt *time.Time `json:"created_at"`
	ShippedAt *time.Time`json:"shipped_at"`
	DeliveredAt *time.Time`json:"delivered_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity int `json:"quantity"`
	Price    int`json:"price"`
}
