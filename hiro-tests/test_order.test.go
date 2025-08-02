package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestOrder(t *testing.T) {
	// Arrange
	orderID := uint64(1)
	customerID := uuid.New()
	createdAt := time.Now()
	shippedAt := time.Now()
	deliveredAt := time.Now()
	lineItem := LineItem{
		ItemID:   uuid.New(),
		Quantity: 2,
		Price:    10,
	}
	order := Order{
		OrderID:    orderID,
		CustomerID: customerID,
		LineItems:  []LineItem{lineItem},
		CreatedAt: &createdAt,
		ShippedAt: &shippedAt,
		DeliveredAt: &deliveredAt,
	}

	// Act
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		t.Errorf("failed to marshal order: %w", err)
	}

	var unmarshaledOrder Order
	err = json.Unmarshal(jsonOrder, &unmarshaledOrder)
	if err != nil {
		t.Errorf("failed to unmarshal order: %w", err)
	}

	// Assert
	if unmarshaledOrder.OrderID != orderID {
		t.Errorf("expected order ID %d, but got %d", orderID, unmarshaledOrder.OrderID)
	}
	if unmarshaledOrder.CustomerID != customerID {
		t.Errorf("expected customer ID %s, but got %s", customerID, unmarshaledOrder.CustomerID)
	}
	if len(unmarshaledOrder.LineItems) != 1 {
		t.Errorf("expected 1 line item, but got %d", len(unmarshaledOrder.LineItems))
	}
	if unmarshaledOrder.LineItems[0].ItemID != lineItem.ItemID {
		t.Errorf("expected item ID %s, but got %s", lineItem.ItemID, unmarshaledOrder.LineItems[0].ItemID)
	}
	if unmarshaledOrder.LineItems[0].Quantity != lineItem.Quantity {
		t.Errorf("expected quantity %d, but got %d", lineItem.Quantity, unmarshaledOrder.LineItems[0].Quantity)
	}
	if unmarshaledOrder.LineItems[0].Price != lineItem.Price {
		t.Errorf("expected price %d, but got %d", lineItem.Price, unmarshaledOrder.LineItems[0].Price)
	}
	if unmarshaledOrder.CreatedAt == nil || unmarshaledOrder.CreatedAt.Format(time.RFC3339) != createdAt.Format(time.RFC3339) {
		t.Errorf("expected created at %s, but got %s", createdAt, unmarshaledOrder.CreatedAt)
	}
	if unmarshaledOrder.ShippedAt == nil || unmarshaledOrder.ShippedAt.Format(time.RFC3339) != shippedAt.Format(time.RFC3339) {
		t.Errorf("expected shipped at %s, but got %s", shippedAt, unmarshaledOrder.ShippedAt)
	}
	if unmarshaledOrder.DeliveredAt == nil || unmarshaledOrder.DeliveredAt.Format(time.RFC3339) != deliveredAt.Format(time.RFC3339) {
		t.Errorf("expected delivered at %s, but got %s", deliveredAt, unmarshaledOrder.DeliveredAt)
	}
}
