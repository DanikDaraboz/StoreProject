package handlers

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
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

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "templates/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		ts, err := template.ParseFS(ui.Files, page)
		if err != nil {
			return nil, err
		}
		cache[filepath.Base(page)] = ts
	}

	return cache, nil
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}) error {
	cache, err := NewTemplateCache()
	if err != nil {
		return err
	}

	tmpl, ok := cache[tmplName]
	if !ok {
		return fmt.Errorf("template not found: %s", tmplName)
	}

	return tmpl.Execute(w, data)
}
