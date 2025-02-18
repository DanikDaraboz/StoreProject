package services

import (
	"errors"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"github.com/DanikDaraboz/StoreProject/pkg/logger"
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

func (c *cartServices) GetCart(userID primitive.ObjectID) (*models.Cart, error) {
	cart, err := c.cartRepo.FindCartByUserID(userID)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return nil, err
	}
	return cart, nil
}

func (c *cartServices) UpdateCartItem(userID primitive.ObjectID, productID primitive.ObjectID, quantity int) error {
    cart, err := c.cartRepo.FindCartByUserID(userID)
    if err != nil {
        logger.ErrorLogger.Println("error", err)
        return err
    }

    // Update the cart item quantity if exists
   	updated := false
    for i, item := range cart.Items {
        if item.ProductID == productID {
            if quantity <= 0 {
                // Remove the item if quantity is 0 or negative
                cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
            } else {
                cart.Items[i].Quantity = quantity
            }
            updated = true
            break
        }
    }
    if !updated {
        return errors.New("cart item not found")
    }

    // Recalculate the total price
    cart.TotalPrice = calculateTotal(cart.Items)

    // Persist the updated cart
    if err := c.cartRepo.UpdateCart(cart); err != nil {
        logger.ErrorLogger.Println("error", err)
        return err
    }

    return nil
}


func (c *cartServices) AddItemToCart(userID primitive.ObjectID, item *models.CartItem) error {
	cart, err := c.cartRepo.FindCartByUserID(userID)
	logger.ErrorLogger.Println("additemtocart service")
	if err != nil {
		// if no cart exists, create one
		if err == mongo.ErrNoDocuments {
			newCart := models.Cart{
				UserID:     userID,
				Items:      []models.CartItem{*item},
				TotalPrice: item.Price * float64(item.Quantity),
			}
			logger.ErrorLogger.Println("error", err)
			return c.cartRepo.InsertCart(&newCart)
		}
		logger.ErrorLogger.Println("error", err)
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

	err = c.cartRepo.UpdateCart(cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
	}
	return err
}

func (c *cartServices) GetCartItems(userID primitive.ObjectID) ([]models.CartItem, error) {
	var cart *models.Cart
	cart, err := c.cartRepo.FindCartByUserID(userID)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return []models.CartItem{}, err
	}

	return cart.Items, nil
}

func (c *cartServices) RemoveItemFromCart(userID primitive.ObjectID, itemID string) error {
	objectItemID, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return err
	}

	var cart *models.Cart
	cart, err = c.cartRepo.FindCartByUserID(userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			err = errors.New("no cart found to delete item")
			logger.ErrorLogger.Println("error", err)
			return err
		}
		logger.ErrorLogger.Println("error", err)
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

	err = c.cartRepo.UpdateCart(cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
	}
	return err
}

func (c *cartServices) ClearCart(userID primitive.ObjectID) error {
	var cart *models.Cart
	cart, err := c.cartRepo.FindCartByUserID(userID)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
		return err
	}

	// Clear items and reset total price
	cart.Items = []models.CartItem{}
	cart.TotalPrice = 0.0

	err = c.cartRepo.UpdateCart(cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
	}
	return err
}

func (c *cartServices) UpdateCart(cart *models.Cart) error {
	err := c.cartRepo.UpdateCart(cart)
	if err != nil {
		logger.ErrorLogger.Println("error", err)
	}
	return err
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
