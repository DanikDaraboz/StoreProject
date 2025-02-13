package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/models"
)

var createdOrder models.Order

func TestOrderRoutes(t *testing.T) {
	t.Run("POST /orders - Create multiple orders", func(t *testing.T) {
		orders := []models.Order{
			{
				UserID: "test-user-1",
				Items: []models.OrderItem{
					{ProductID: "prod1", Quantity: 2, Price: 15.0},
					{ProductID: "prod2", Quantity: 1, Price: 25.0},
				},
				TotalPrice: 55.0,
			},
			{
				UserID: "test-user-2",
				Items: []models.OrderItem{
					{ProductID: "prod3", Quantity: 3, Price: 10.0},
				},
				TotalPrice: 30.0,
			},
			{
				UserID: "test-user-3",
				Items: []models.OrderItem{
					{ProductID: "prod4", Quantity: 1, Price: 100.0},
					{ProductID: "prod5", Quantity: 2, Price: 50.0},
				},
				TotalPrice: 200.0,
			},
		}

		for _, order := range orders {
			payload, _ := json.Marshal(order)
			req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			response := executeRequest(req)
			checkStatus(t, http.StatusCreated, response.Code)

			if err := json.NewDecoder(response.Body).Decode(&createdOrder); err != nil {
				t.Fatalf("Error decoding response: %v", err)
			}

			t.Logf("Order created successfully! Order ID: %s", createdOrder.ID.Hex())
		}
	})

	t.Run("GET /orders{id}", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/orders/"+createdOrder.ID.Hex(), nil)
		response := executeRequest(req)

		checkStatus(t, http.StatusOK, response.Code)

		var fetchedOrder models.Order
		if err := json.NewDecoder(response.Body).Decode(&fetchedOrder); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		if fetchedOrder.ID.Hex() != createdOrder.ID.Hex() {
			t.Errorf("Expected order ID %s, got %s", createdOrder.ID.Hex(), fetchedOrder.ID.Hex())
		}

		t.Logf("UserID of order: %s", fetchedOrder.UserID)
	})

	t.Run("GET /orders", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/orders", nil)
		response := executeRequest(req)

		checkStatus(t, http.StatusOK, response.Code)

		var fetchedOrders []models.Order
		if err := json.NewDecoder(response.Body).Decode(&fetchedOrders); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		for _, order := range fetchedOrders {
			t.Logf("Order: %s", order.ID)
		}
	})

	t.Run("PUT /orders/{id}", func(t *testing.T) {
		update := map[string]interface{}{
			"total_price": 99.99,
		}
		payload, _ := json.Marshal(update)

		t.Log("Current price: 55.0")

		req, _ := http.NewRequest("PUT", "/orders/"+createdOrder.ID.Hex(), bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req)
		checkStatus(t, http.StatusOK, response.Code)

		var updatedOrder models.Order
		if err := json.NewDecoder(response.Body).Decode(&updatedOrder); err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		if updatedOrder.TotalPrice != 99.99 {
			t.Errorf("Expected updated total price to be 99.99, got %v", updatedOrder.TotalPrice)
		}

		t.Logf("New price: %f", updatedOrder.TotalPrice)
	})

	t.Run("DELETE /orders/{id}", func(t *testing.T) {
		t.Logf("Checking if deleted order (%s) still exists...", createdOrder.ID.Hex())

		req, _ := http.NewRequest("DELETE", "/orders/"+createdOrder.ID.Hex(), nil)
		response := executeRequest(req)

		if response.Code == http.StatusNoContent {
			t.Log("Order successfully removed.")
		} else {
			t.Errorf("Order still exists after deletion. Response Code: %d", response.Code)
		}
	})

}
