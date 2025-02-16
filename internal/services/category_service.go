package services

import (
	"github.com/DanikDaraboz/StoreProject/internal/models"
	repoInterface "github.com/DanikDaraboz/StoreProject/internal/repository/interfaces"
	"github.com/DanikDaraboz/StoreProject/internal/services/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ interfaces.CategoryServicesInterface = (*categoryService)(nil)

type categoryService struct {
	repo repoInterface.CategoryRepositoryInterface
}

func NewCategoryService(repo repoInterface.CategoryRepositoryInterface) interfaces.CategoryServicesInterface {
	return &categoryService{repo: repo}
}

func (c categoryService) CreateCategory(category *models.Category) (primitive.ObjectID, error) {
	return c.repo.CreateCategory(category)
}

func (c categoryService) GetAllCategories() ([]models.Category, error) {
	return c.repo.GetAllCategories()
}

func (c categoryService) GetCategoryByID(id string) (*models.Category, error) {
	return c.repo.GetCategoryByID(id)
}

func (c categoryService) UpdateCategory(id string, category *models.Category) error {
	return c.repo.UpdateCategory(id, category)
}

func (c categoryService) DeleteCategory(id string) error {
	return c.repo.DeleteCategory(id)
}
