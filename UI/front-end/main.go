package main

import (
	"fmt"
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
		idParam := r.URL.Query().Get("id") // Получаем id товара из URL
		if idParam == "" {
			http.Error(w, "Product ID is required", http.StatusBadRequest)
			return
		}
	
		var selectedProduct Product
		found := false
	
		for _, product := range products {
			if fmt.Sprintf("%d", product.ID) == idParam { // Сравниваем ID товара
				selectedProduct = product
				found = true
				break
			}
		}
	
		if !found {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
	
		// Создаем структуру данных для страницы, включающую продукт и категории
		data := struct {
			Product   Product
			Categories []string
		}{
			Product:   selectedProduct,
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/product.html")
		if err != nil {
			log.Println("Error loading template:", err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)  // Передаем данные в шаблон
	})
	

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:      "All Products",
			Products:   products, // Показываем все товары
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/all.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})



	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		category := strings.TrimPrefix(r.URL.Path, "/category/") // Получаем название категории из URL
	
		var filteredProducts []Product
		for _, product := range products {
			if strings.ToLower(product.Category) == strings.ToLower(category) {
				filteredProducts = append(filteredProducts, product)
			}
		}
	
		if len(filteredProducts) == 0 {
			http.Error(w, "No products found in this category", http.StatusNotFound)
			return
		}
	
		data := PageData{
			Title:      category + " Products",
			Products:   filteredProducts,
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/category.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
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
	http.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем наличие куки "authenticated"

	
		// Загружаем страницу аккаунта
		tmpl, err := template.ParseFiles("templates/account.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load account page", http.StatusInternalServerError)
			return
		}
	
		// Передаем информацию в шаблон
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
