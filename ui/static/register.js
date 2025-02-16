document.getElementById("signupForm").addEventListener("submit", async function(event) {
    event.preventDefault(); // Остановить стандартную отправку формы

    // Получаем данные формы
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const confirmPassword = document.getElementById("confirm-password").value;

    // Очистка сообщений об ошибках
    document.getElementById("emailError").textContent = "";
    document.getElementById("passwordError").textContent = "";
    document.getElementById("confirmPasswordError").textContent = "";

    // Простая валидация
    if (password.length < 6) {
        document.getElementById("passwordError").textContent = "Password must be at least 6 characters.";
        return;
    }
    if (password !== confirmPassword) {
        document.getElementById("confirmPasswordError").textContent = "Passwords do not match.";
        return;
    }

    // Формируем JSON
    const requestData = { email, password };

    try {
        const response = await fetch("/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(requestData)
        });

        const responseData = await response.json();
        console.log("Server Response:", responseData);

        if (response.ok) {
            alert("Registration successful!");
            window.location.href = "/login";
        } else {
            alert(responseData.message || "Registration failed!");
        }
    } catch (error) {
        console.error("Error:", error);
    }
});