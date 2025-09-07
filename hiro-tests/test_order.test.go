package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestOrder(t *testing.T) {
	order := Order{
		OrderID:    1,
		CustomerID: uuid.New(),
		LineItems: []LineItem{{
			ItemID:   uuid.New(),
			Quantity: 1,
			Price:    10,
		}},
		CreatedAt: time.Now(),
	}

	// Test OrderID
	if order.OrderID != 1 {
		t.Errorf("OrderID is not 1, got %d", order.OrderID)
	}

	// Test CustomerID
	if order.CustomerID == uuid.Nil {
		t.Errorf("CustomerID is not set, got %v", order.CustomerID)
	}

	// Test LineItems
	if len(order.LineItems) != 1 {
		t.Errorf("LineItems length is not 1, got %d", len(order.LineItems))
	}

	// Test CreatedAt
	if order.CreatedAt == nil {
		t.Errorf("CreatedAt is not set, got %v", order.CreatedAt)
	}
}

func TestLineItem(t *testing.T) {
	lineItem := LineItem{
		ItemID:   uuid.New(),
		Quantity: 1,
		Price:    10,
	}

	// Test ItemID
	if lineItem.ItemID == uuid.Nil {
		t.Errorf("ItemID is not set, got %v", lineItem.ItemID)
	}

	// Test Quantity
	if lineItem.Quantity != 1 {
		t.Errorf("Quantity is not 1, got %d", lineItem.Quantity)
	}

	// Test Price
	if lineItem.Price != 10 {
		t.Errorf("Price is not 10, got %d", lineItem.Price)
	}
}
