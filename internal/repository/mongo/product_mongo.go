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

// `productRepository` implements `ProductRepositoryInterface`
var _ interfaces.ProductRepositoryInterface = (*productRepository)(nil)

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) interfaces.ProductRepositoryInterface {
	return &productRepository{collection: collection}
}

// TODO Pagination?
func (p *productRepository) GetProducts() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	cursor, err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []map[string]interface{}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepository) FetchProductByID(id string) (models.Product, error) {
	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}

	err = p.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product)
	return product, err
}

func (p *productRepository) InsertProduct(product models.Product) error {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err := p.collection.InsertOne(context.TODO(), product)
	return err
}

func (p *productRepository) UpdateProduct(id string, product models.Product) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	product.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err = p.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": product},
	)
	return err
}

func (p *productRepository) RemoveProduct(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
