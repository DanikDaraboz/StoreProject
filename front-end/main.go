package main

import (
	"html/template"
	"log"
	"net/http"
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
	Title      string
	Products   []Product
	Categories []string
}

func main() {
	// Статические файлы (CSS, изображения и т.д.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Пример товаров
	products := []Product{
		{ID: 1, Name: "Soccer Ball", Category: "Football", Price: 19.99, Image: "/static/images/soccer-ball.png"},
		{ID: 2, Name: "Tennis Racket", Category: "Tennis", Price: 49.99, Image: "/static/images/tennis-racket.png"},
		{ID: 3, Name: "Basketball", Category: "Basketball", Price: 29.99, Image: "/static/images/basketball.png"},
		{ID: 4, Name: "Baseball Glove", Category: "Baseball", Price: 39.99, Image: "/static/images/baseball-glove.png"},
	}

	// Категории товаров
	categories := []string{"Football", "Tennis", "Basketball", "Baseball"}

	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		category := r.URL.Query().Get("category")

		// Фильтруем товары по категории
		var filteredProducts []Product
		for _, product := range products {
			if category == "" || strings.ToLower(product.Category) == strings.ToLower(category) || category == "all" {
				filteredProducts = append(filteredProducts, product)
			}
		}

		// Данные страницы
		data := PageData{
			Title:      "Sports Goods Store",
			Products:   filteredProducts,
			Categories: categories,
		}

		// Загружаем шаблон
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	// Страница товара
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		product := Product{ID: 1, Name: "Soccer Ball", Price: 19.99, Image: "/static/images/soccer-ball.png"}
		tmpl, err := template.ParseFiles("templates/product.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, product)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	// Страница авторизации (Login)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/authorization.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	// Страница регистрации (Sign Up)
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/signup.html")
		if err != nil {
			log.Println("Error loading register.html:", err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	// Обработчик 404 ошибки
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/404.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		tmpl.Execute(w, nil)
	})

	// Запуск сервера
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
