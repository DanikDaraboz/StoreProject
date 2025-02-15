document.addEventListener("DOMContentLoaded", function () {
    const favoriteIcons = document.querySelectorAll(".favorites-icon");

    // Загружаем избранные товары из localStorage
    let favorites = JSON.parse(localStorage.getItem("favorites")) || {};

    // Применяем состояние из localStorage
    favoriteIcons.forEach(icon => {
        const productId = icon.dataset.id;
        if (favorites[productId]) {
            icon.classList.add("favorited");
        }

        // Обработчик клика по иконке
        icon.addEventListener("click", function () {
            const isFavorited = !favorites[productId]; // Переключаем состояние
            favorites[productId] = isFavorited;

            // Обновляем LocalStorage
            localStorage.setItem("favorites", JSON.stringify(favorites));

            // Обновляем стиль иконки
            if (isFavorited) {
                icon.classList.add("favorited");
            } else {
                icon.classList.remove("favorited");
            }
        });
    });

    // Передача избранных товаров на сервер (для страницы избранного)
    if (window.location.pathname === "/favorites") {
        fetch("/favorites", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(Object.keys(favorites).filter(id => favorites[id]))
        }).then(response => response.text())
          .then(html => document.body.innerHTML = html);
    }
});
