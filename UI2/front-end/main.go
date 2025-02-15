package main

import (
	"strconv"
	"encoding/json"
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
	Favorites bool
	Cart     bool
	Quantity int
	Gender string
	Description string
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
		{ID: 1, Name: "EURO 2024 Ball", Category: "Football", Price: 19.99, Image: "/static/images/soccer-ball.png"},
		{ID: 2, Name: "Tennis Racket", Category: "Tennis", Price: 49.99, Image: "/static/images/tennis-racket.png"},
		{ID: 3, Name: "Basketball", Category: "Basketball", Price: 29.99, Image: "/static/images/basketball.png"},
		{ID: 4, Name: "Baseball Glove", Category: "Baseball", Price: 39.99, Image: "/static/images/baseball-glove.png"},
		{ID: 5, Name: "Superstar 2 Shoes", Category: "Shoes", Price: 99.99, Image: "/static/images/shoe1.png"},
		{ID: 6, Name: "Superstar 2 Shoes", Category: "Shoes", Price: 99.99, Image: "/static/images/shoe2.png"},
		{ID: 7, Name: "Country Japan Shoes", Category: "Shoes", Price: 85.99, Image: "/static/images/shoe3.png"},
		{ID: 8, Name: "Samba OG Shoes", Category: "Shoes", Price: 119.99, Image: "/static/images/shoe4.png"},
		{ID: 9, Name: "Superstar 2 shoes", Category: "Shoes", Price: 140.99, Image: "/static/images/shoe5.png"},
		{ID: 10, Name: "SL 72 OG Shoes", Category: "Shoes", Price: 75.99, Image: "/static/images/shoe6.png"},
		{ID: 11, Name: "Samba Messi", Category: "Shoes", Price: 109.99, Image: "/static/images/shoe7.png"},
		{ID: 12, Name: "FIFA 25 Club Ball", Category: "Football", Price: 25.99, Image: "/static/images/sb1.png"},
		{ID: 13, Name: "FIFA 25 Competition Ball", Category: "Football", Price: 25.99, Image: "/static/images/sb2.png"},
		{ID: 14, Name: "MLS 25", Category: "Football", Price: 35.99, Image: "/static/images/sb3.png"},
		{ID: 15, Name: "Tango Glider", Category: "Football", Price: 9.99, Image: "/static/images/sb4.png"},
		{ID: 16, Name: "Argentina Anniversary Ball", Category: "Football", Price: 19.99, Image: "/static/images/sb5.png"},
		{ID: 17, Name: "Messi Club Ball", Category: "Football", Price: 19.99, Image: "/static/images/sb6.png"},
		{ID: 18, Name: "SST Track Jacket", Gender:"Men", Price: 39.99, Image: "/static/images/men1.png"},
		{ID: 19, Name: "Adicolor Teamgeist Tee", Gender:"Women", Price: 29.99, Image: "/static/images/women1.png"},
		{ID: 20, Name: "365 Half-zip", Gender:"Women", Price: 19.99, Image: "/static/images/women2.png"},
		{ID: 21, Name: "Polo Shirt", Gender:"Women", Price: 39.99, Image: "/static/images/women3.png"},
		{ID: 22, Name: "1/4 Zip Tee", Gender:"Men", Price: 59.99, Image: "/static/images/men2.png"},
		{ID: 23, Name: "Puremotion Tee", Gender:"Men", Price: 39.99, Image: "/static/images/men3.png"},
		{ID: 24, Name: "Adidas Originals Hoodie", Gender:"Men", Price: 29.99, Image: "/static/images/men4.png"},
		{ID: 25, Name: "Jacket", Gender:"Men", Price: 89.99, Image: "/static/images/men5.png"},
		{ID: 26, Name: "Yoga Cover-Up", Gender:"Women", Price: 39.99, Image: "/static/images/women4.png"},
		{ID: 27, Name: "3 Stripez Full-Zip Hoodie", Gender:"Women", Price: 39.99, Image: "/static/images/women5.png"},


	}

	// Категории товаров
	categories := []string{"Football", "Tennis", "Basketball", "Baseball","Shoes"}

	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Фильтруем товары по категории
		var filteredProducts []Product
		if len(products) > 4 {
			filteredProducts = products[:4]
		} else {
			filteredProducts = products
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


	http.HandleFunc("/favorites/toggle", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		productID := r.FormValue("id")
	
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == productID {
				products[i].Favorites = !products[i].Favorites // Переключаем значение
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
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
	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Получаем список избранных товаров из запроса
			var favoriteIDs []string
			err := json.NewDecoder(r.Body).Decode(&favoriteIDs)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
	
			// Фильтруем товары
			var favoriteProducts []Product
			for _, product := range products {
				for _, id := range favoriteIDs {
					if fmt.Sprintf("%d", product.ID) == id {
						favoriteProducts = append(favoriteProducts, product)
					}
				}
			}
			// Отправляем страницу с избранными товарами
			data := PageData{
				Title:    "Favorites",
				Products: favoriteProducts,
				Categories: categories,
			}
	
			tmpl, err := template.ParseFiles("templates/favorites.html")
			if err != nil {
				log.Println(err)
				http.Error(w, "Could not load template", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, data)
			return
		}
	
		// GET-запрос просто загружает пустую страницу
		tmpl, err := template.ParseFiles("templates/favorites.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})
	
	
	

	http.HandleFunc("/cart/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		productID := r.FormValue("id")
	
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == productID {
				products[i].Cart = true // Устанавливаем cart = true
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})
	
	http.HandleFunc("/cart/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
	
		r.ParseForm()
		productID := r.FormValue("id")
	
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == productID {
				products[i].Cart = false // Убираем товар из корзины
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})
	
	
	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		var cartProducts []Product
		for _, product := range products {
			if product.Cart { // Фильтруем только те, которые добавлены в корзину
				cartProducts = append(cartProducts, product)
			}
		}
	
		data := PageData{
			Title:      "Shopping Cart",
			Products:   cartProducts,
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/cart.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})


	http.HandleFunc("/cart/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
	
		r.ParseForm()
		productID := r.FormValue("id")
		quantity := r.FormValue("quantity")
	
		for i, product := range products {
			if fmt.Sprintf("%d", product.ID) == productID {
				qty, err := strconv.Atoi(quantity)
				if err != nil || qty < 1 {
					http.Error(w, "Invalid quantity", http.StatusBadRequest)
					return
				}
	
				products[i].Quantity = qty
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	})
	
	
	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:      "Total Payment",
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/payment.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
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
	
	http.HandleFunc("/men", func(w http.ResponseWriter, r *http.Request) {
		var genproduct []Product
		for _, product := range products {
			if product.Gender=="Men" { // Фильтруем только те, которые добавлены в корзину
				genproduct = append(genproduct, product)
			}
		}
	
		data := PageData{
			Title:      "Men's Wear",
			Products:   genproduct,
			Categories: categories,
		}
	
		tmpl, err := template.ParseFiles("templates/men.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/women", func(w http.ResponseWriter, r *http.Request) {
	var genproduct []Product
	for _, product := range products {
		if product.Gender=="Women" { // Фильтруем только те, которые добавлены в корзину
			genproduct = append(genproduct, product)
		}
	}

	data := PageData{
		Title:      "Women's Wear",
		Products:   genproduct,
		Categories: categories,
	}

	tmpl, err := template.ParseFiles("templates/women.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
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
