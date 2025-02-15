document.addEventListener("DOMContentLoaded", function () {
    // Получаем элементы страницы по id
    const userIconContainer = document.getElementById("user-icon-container"); // контейнер иконки пользователя
    const signInBtn = document.getElementById("sign-in-btn"); // кнопка входа
    const userEmailSpan = document.getElementById("user-email"); // элемент для отображения email
    const userInitialSpan = document.getElementById("user-initial"); // элемент для отображения первой буквы email
    const logoutBtn = document.getElementById("logout"); // кнопка выхода
    const changePasswordForm = document.getElementById("change-password-form"); // форма смены пароля

    // Проверка, авторизован ли пользователь
    if (localStorage.getItem("userLoggedIn") === "true") {
        const userEmail = localStorage.getItem("userEmail");

        // Если email сохранён, отображаем его и первую букву
        if (userEmail) {
            userEmailSpan.textContent = userEmail;
            userInitialSpan.textContent = userEmail.charAt(0).toUpperCase();
        }

        // Отображаем иконку пользователя и скрываем кнопку входа
        if (userIconContainer) userIconContainer.style.display = "inline-block";
        if (signInBtn) signInBtn.style.display = "none";
    } else {
        // Если пользователь не авторизован, скрываем иконку и показываем кнопку входа
        if (userIconContainer) userIconContainer.style.display = "none";
        if (signInBtn) signInBtn.style.display = "inline-block";
    }

    // Логика выхода из аккаунта
    if (logoutBtn) {
        logoutBtn.addEventListener("click", function () {
            // Удаляем данные о пользователе из localStorage
            localStorage.removeItem("userLoggedIn");
            localStorage.removeItem("userEmail");
            // Дополнительно можно удалять другие данные, если они сохраняются

            // Перенаправляем на главную страницу (или любую другую)
            window.location.href = "/";
        });
    }

    // Логика изменения пароля
    if (changePasswordForm) {
        changePasswordForm.addEventListener("submit", function (event) {
            event.preventDefault(); // Отменяем стандартное поведение формы

            // Получаем значения из полей формы
            const newPassword = document.getElementById("new-password").value;
            const confirmNewPassword = document.getElementById("confirm-new-password").value;

            if (newPassword && newPassword === confirmNewPassword) {
                // Здесь можно отправить запрос на сервер для обновления пароля
                alert("Пароль успешно изменён!");
                // Дополнительные действия: очистить поля, обновить интерфейс и т.д.
            } else {
                alert("Пароли не совпадают или не введены!");
            }
        });
    }
});
