 document.addEventListener("DOMContentLoaded", function () {
        const navLinks = document.querySelectorAll('.nav-item a');
        const catalogText = document.getElementById('catalog-text');
        const categoryLinks = document.querySelectorAll('.category-menu li a');

        // Восстанавливаем активную вкладку из localStorage
        const activeTab = localStorage.getItem("activeTab");
        if (activeTab) {
            navLinks.forEach(link => {
                if (link.getAttribute("href") === activeTab) {
                    link.classList.add("active");
                }
            });
        }

        // Восстанавливаем активную категорию из localStorage
        const activeCategory = localStorage.getItem("activeCategory");
        if (activeCategory) {
            categoryLinks.forEach(link => {
                if (link.getAttribute("href") === activeCategory) {
                    link.classList.add("active");
                    catalogText.classList.add("active"); // Делаем "Catalog" жирным
                }
            });
        }

        // Навигационные вкладки
        navLinks.forEach(link => {
            link.addEventListener("click", function () {
                navLinks.forEach(link => link.classList.remove("active"));
                this.classList.add("active");

                // Сохраняем активную вкладку
                localStorage.setItem("activeTab", this.getAttribute("href"));
                localStorage.removeItem("activeCategory"); // Сбрасываем категорию, если кликнули на вкладку
                catalogText.classList.remove("active"); // Убираем жирный "Catalog"
            });
        });

        // Категории каталога
        categoryLinks.forEach(link => {
            link.addEventListener("click", function () {
                categoryLinks.forEach(link => link.classList.remove("active"));
                this.classList.add("active");

                catalogText.classList.add("active"); // Делаем "Catalog" жирным

                // Сохраняем активную категорию
                localStorage.setItem("activeCategory", this.getAttribute("href"));
                localStorage.removeItem("activeTab"); // Сбрасываем вкладку, если выбрана категория
            });
        });
    });
