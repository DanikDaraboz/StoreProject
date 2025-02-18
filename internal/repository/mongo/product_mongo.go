package mongo

import (
	"context"
	"time"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ interfaces.ProductRepositoryInterface = (*productRepository)(nil)

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) interfaces.ProductRepositoryInterface {
	return &productRepository{collection: collection}
}

func (r *productRepository) GetProducts(categoryID string) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Construct query: filter by category if provided
	filter := bson.M{}
	if categoryID != "" {
		oid, err := primitive.ObjectIDFromHex(categoryID)
		if err != nil {
			return nil, err
		}
		filter["category_id"] = oid
	}

	// Execute query
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode results into products slice
	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepository) FetchProductByID(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &product, err
	}

	err = p.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	return &product, err
}

func (p *productRepository) InsertProduct(product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	product.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := p.collection.InsertOne(ctx, product)
	return err
}

func (p *productRepository) UpdateProduct(id string, product *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := p.collection.UpdateOne(
		ctx,
		bson.M{"_id": product.ID},
		bson.M{"$set": product},
	)
	return err
}

func (p *productRepository) RemoveProduct(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
