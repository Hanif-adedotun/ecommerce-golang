package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		name string
		order Order
		want string
	} {
		{
			name: "Order with valid data",
			order: Order{
				OrderID: 1,
				CustomerID: uuid.New(),
				LineItems: []LineItem{{
					ItemID: uuid.New(),
					Quantity: 1,
					Price: 10,
				}},
				CreatedAt: timePtr(time.Now()),
			},
			want: "Order{OrderID:1,CustomerID:<uuid.UUID>,LineItems:[{ItemID:<uuid.UUID> Quantity:1 Price:10}],CreatedAt:<time.Time>,ShippedAt:<*time.Time>,DeliveredAt:<*time.Time>}",
		},
		{
			name: "Order with empty LineItems",
			order: Order{
				OrderID: 1,
				CustomerID: uuid.New(),
				LineItems: []LineItem{},
				CreatedAt: timePtr(time.Now()),
			},
			want: "Order{OrderID:1,CustomerID:<uuid.UUID>,LineItems:[],CreatedAt:<time.Time>,ShippedAt:<*time.Time>,DeliveredAt:<*time.Time>}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.String(); got != tt.want {
				t.Errorf("Order.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}
