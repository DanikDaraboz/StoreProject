package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.CartServicesInterface = (*cartServices)(nil)

type cartServices struct {
	cartRepo repoInterface.CartRepositoryInterface
}

func NewCartServices(cartRepo repoInterface.CartRepositoryInterface) interfaces.CartServicesInterface {
	return &cartServices{cartRepo: cartRepo}
}

func (c *cartServices) AddItemToCart(userID string, item *models.CartItem) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	cart, err := c.cartRepo.FindCartByUserID(objectID)
	if err != nil {
		// if no cart exist, create one
		if err == mongo.ErrNoDocuments {
			newCart := models.Cart{
				UserID:     objectID,
				Items:      []models.CartItem{*item},
				TotalPrice: item.Price * float64(item.Quantity),
			}
			return c.cartRepo.InsertCart(&newCart)
		}
		return err
	}

	// Check if item already exists in the cart
	index, exists := itemExists(cart, item.ProductID)
	if exists {
		cart.Items[index].Quantity += item.Quantity
	} else {
		cart.Items = append(cart.Items, *item)
	}

	cart.TotalPrice = calculateTotal(cart.Items)

	return c.cartRepo.UpdateCart(cart)

}

func (c *cartServices) GetCartItems(userID string) ([]models.CartItem, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return []models.CartItem{}, err
	}

	var cart *models.Cart
	cart, err = c.cartRepo.FindCartByUserID(objectID)
	if err != nil {
		return []models.CartItem{}, err
	}

	return cart.Items, nil
}

func (c *cartServices) RemoveItemFromCart(userID string, itemID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	objectItemID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return err
	}

	var cart *models.Cart
	cart, err = c.cartRepo.FindCartByUserID(objectID)
	if err != nil {
		// if no cart exist, create one
		if err == mongo.ErrNoDocuments {
			return errors.New("no cart found to delete item")
		}
		return err
	}

	// Remove the item from the cart items slice
	var updatedItems []models.CartItem
	for _, item := range cart.Items {
		if item.ProductID != objectItemID {
			updatedItems = append(updatedItems, item)
		}
	}

	// Update the cart model
	cart.Items = updatedItems
	cart.TotalPrice = calculateTotal(updatedItems)

	return c.cartRepo.UpdateCart(cart)
}

func (c *cartServices) ClearCart(userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	var cart *models.Cart
	cart, err = c.cartRepo.FindCartByUserID(objectID)
	if err != nil {
		return err
	}

	// Clear items and reset total price
	cart.Items = []models.CartItem{}
	cart.TotalPrice = 0.0

	return c.cartRepo.UpdateCart(cart)
}

func (c *cartServices) UpdateCart(cart *models.Cart) error {
	return c.cartRepo.UpdateCart(cart)
}

func itemExists(cart *models.Cart, productID primitive.ObjectID) (int, bool) {
	for index, item := range cart.Items {
		if item.ProductID == productID {
			return index, true
		}
	}
	return -1, false
}

func calculateTotal(items []models.CartItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
