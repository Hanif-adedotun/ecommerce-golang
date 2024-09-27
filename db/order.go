package db

import (
	"context"
	"database/sql"
	"encoding/json"

	// "encoding/json"
	"fmt"

	"github.com/hanif-adedotun/ecommerce-golang/model"
)

type PostgreRepo struct {
	Client *sql.DB
}

func (p *PostgreRepo) Insert(ctx context.Context, order model.Order) (int64, error) {
	lineItems, err := json.Marshal(order.LineItems)
	if err != nil {
		return 0, fmt.Errorf("failed to encode order: %w", err)
	}
	result, err := p.Client.Exec("INSERT INTO orders (order_id, customer_id, created_at, shipped_at, delivered_at) VALUES (?, ?, ?, ?, ?)", order.OrderID, order.CustomerID, order.CreatedAt, order.DeliveredAt, order.ShippedAt)
	if err != nil {
		return 0, fmt.Errorf("CreateOrder: %v", err)
	}

	// Use transactions to add line items to the lineitemsDB
	fmt.Println("Order lineItems:", lineItems)

	// Get the new order's generated ID for the client.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateOrder: %v", err)
	}
	// Return the new album's ID.
	return id, err

}
