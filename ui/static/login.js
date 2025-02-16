document.addEventListener("DOMContentLoaded", function () {
    // Проверяем, есть ли пользователь в localStorage (или можно сделать запрос к серверу)
    if (localStorage.getItem("user")) {
        showUserIcon();
    }

    document.getElementById("loginForm").addEventListener("submit", async function (event) {
        event.preventDefault(); // Остановить стандартную отправку формы

        // Получаем данные формы
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        // Формируем JSON
        const requestData = { email, password };

        try {
            const response = await fetch("/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(requestData)
            });

            const responseData = await response.json();
            console.log("Server Response:", responseData);

            if (response.ok) {
                alert("Login successful!");

                // Сохраняем пользователя в localStorage (или использовать JWT в cookies)
                localStorage.setItem("user", JSON.stringify(responseData.user));

                // Обновляем UI
                showUserIcon();

                // Перенаправляем на страницу аккаунта
                window.location.href = "/account";
            } else {
                alert(responseData.message || "Login failed!");
            }
        } catch (error) {
            console.error("Error:", error);
        }
    });

    function showUserIcon() {
        const userIcon = document.getElementById("user-icon");
        const signInBtn = document.getElementById("sign-in-btn");

        if (userIcon && signInBtn) {
            userIcon.style.display = "inline-block"; // Показываем иконку пользователя
            signInBtn.style.display = "none"; // Скрываем кнопку входа

            // Можно добавить первую букву имени пользователя, если есть
            const user = JSON.parse(localStorage.getItem("user"));
            if (user && user.name) {
                userIcon.textContent = user.name[0].toUpperCase(); // Первая буква имени
            }
        }
    }
});
