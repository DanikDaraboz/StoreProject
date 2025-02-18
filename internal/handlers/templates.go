package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/DanikDaraboz/StoreProject/internal/models"
	"github.com/DanikDaraboz/StoreProject/ui"
)

type TemplateData struct {
	Title      string
	User       *models.User
	Users      *models.User
	Product    *models.Product
	Products   *[]models.Product
	Cart       *models.Cart
	CartItem   *models.CartItem
	Order      *models.Order
	Orders     *[]models.Order
	Category   *models.Category
	Categories *[]models.Category
}

func Mul(a float64, b int) float64 {
	return a * float64(b)
}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Define a function map and include custom functions like Mul
	funcMap := template.FuncMap{
		"mul": Mul,
	}

	// Find all HTML templates in the "templates" folder
	pages, err := fs.Glob(ui.Files, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("error globbing templates: %w", err)
	}

	// Parse each template with the custom functions
	for _, page := range pages {
		ts, err := template.New(filepath.Base(page)).Funcs(funcMap).ParseFS(ui.Files, page)
		if err != nil {
			return nil, fmt.Errorf("error parsing template %s: %w", page, err)
		}
		cache[filepath.Base(page)] = ts
	}

	return cache, nil
}
