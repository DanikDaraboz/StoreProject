package main

import (
	"html/template"
	"net/http"
	"log"
	"strings"
)

// Структуры для продукта и данных страницы
type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Image    string
}

type PageData struct {
	Title    string
	Products []Product
	Categories []string
}

func main() {
	// Статические файлы (CSS, изображения и т.д.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Пример товаров
	products := []Product{
		{ID: 1, Name: "Soccer Ball", Category: "Football", Price: 19.99, Image: "/static/images/soccer-ball.jpg"},
		{ID: 2, Name: "Tennis Racket", Category: "Tennis", Price: 49.99, Image: "/static/images/tennis-racket.jpg"},
		{ID: 3, Name: "Basketball", Category: "Basketball", Price: 29.99, Image: "/static/images/basketball.jpg"},
		{ID: 4, Name: "Baseball Glove", Category: "Baseball", Price: 39.99, Image: "/static/images/baseball-glove.jpg"},
	}

	// Категории товаров
	categories := []string{"Football", "Tennis", "Basketball", "Baseball"}

	// Обработка главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Получаем категорию из запроса (если есть)
		category := r.URL.Query().Get("category")

		// Фильтруем товары по категории, если категория указана
		var filteredProducts []Product
		for _, product := range products {
			if category == "" || strings.ToLower(product.Category) == strings.ToLower(category) || category == "all" {
				filteredProducts = append(filteredProducts, product)
			}
		}

		// Данные страницы
		data := PageData{
			Title:     "Sports Goods Store",
			Products:  filteredProducts,
			Categories: categories,
		}

		// Загрузка шаблона и вывод
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	// Обработка страницы товара
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		// Пример товара
		product := Product{ID: 1, Name: "Soccer Ball", Price: 19.99, Image: "/static/soccer-ball.jpg"}
		tmpl, err := template.ParseFiles("templates/product.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, product)
	})

	// Запуск сервера
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
