package repository

import (
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/repository/mongo"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	ProductRepo interfaces.ProductRepositoryInterface
	OrderRepo   interfaces.OrderRepositoryInterface
	CartRepo    interfaces.CartRepositoryInterface
	UserRepo    interfaces.UserRepositoryInterface
	SessionRepo interfaces.SessionRepositoryInterface
}

func NewRepositories(db *mongodriver.Database) *Repositories {
	return &Repositories{
		ProductRepo: mongo.NewProductRepository(db.Collection("products")),
		OrderRepo:   mongo.NewOrderRepository(db.Collection("orders")),
		CartRepo:    mongo.NewCartRepository(db.Collection("carts")),
		UserRepo:    mongo.NewUserRepository(db.Collection("users")),
		SessionRepo: mongo.NewSessionRepository(db.Collection("sessions")),
	}
}
