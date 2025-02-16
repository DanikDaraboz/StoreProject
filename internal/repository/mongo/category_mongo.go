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

var _ interfaces.CategoryRepositoryInterface = (*categoryRepo)(nil)

type categoryRepo struct {
	collection *mongo.Collection
}

func NewCategoryRepository(collection *mongo.Collection) interfaces.CategoryRepositoryInterface {
	return &categoryRepo{collection: collection}
}

func (c categoryRepo) CreateCategory(category models.Category) (primitive.ObjectID, error) {
	category.ID = primitive.NewObjectID()
	category.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	category.UpdatedAt = category.CreatedAt

	res, err := c.collection.InsertOne(context.TODO(), category)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (c categoryRepo) GetAllCategories() ([]models.Category, error) {
	cursor, err := c.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var categories []models.Category
	if err := cursor.All(context.TODO(), &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepo) GetCategoryByID(id string) (models.Category, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Category{}, err
	}

	var category models.Category
	err = c.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (c categoryRepo) UpdateCategory(id string, category models.Category) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{
		"name":        category.Name,
		"description": category.Description,
		"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
	}}

	_, err = c.collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

func (r categoryRepo) DeleteCategory(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
