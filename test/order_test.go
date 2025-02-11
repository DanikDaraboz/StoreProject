package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/DanikDaraboz/StoreProject/internal/models"
)

var createdOrder models.Order

func TestCreateOrder(t *testing.T) {
	order := models.Order{
		UserID: "test-user",
		Items: []models.OrderItem{
			{ProductID: "prod1", Quantity: 2, Price: 15.0},
			{ProductID: "prod2", Quantity: 1, Price: 25.0},
		},
		TotalPrice: 55.0,
	}

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

func TestGetOrder(t *testing.T) {
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
}

func TestUpdateOrder(t *testing.T) {
	update := map[string]interface{}{
		"total_price": 99.99,
	}
	payload, _ := json.Marshal(update)

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
}


func TestGetDeletedOrder(t *testing.T) {
	t.Logf("Checking if deleted order (%s) still exists...", createdOrder.ID.Hex())

	req, _ := http.NewRequest("DELETE", "/orders/"+createdOrder.ID.Hex(), nil)
	response := executeRequest(req)

	if response.Code == http.StatusNoContent {
		t.Log("Order successfully removed.")
	} else {
		t.Errorf("Order still exists after deletion. Response Code: %d", response.Code)
	}
}
