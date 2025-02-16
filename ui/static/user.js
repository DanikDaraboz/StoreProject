document.addEventListener("DOMContentLoaded", async function () {
    try {
        const response = await fetch("/user", { credentials: "include" });
        if (!response.ok) throw new Error("Ошибка загрузки данных");

        const user = await response.json();

        document.getElementById("user-name").textContent = user.username || "Не указано";
        document.getElementById("user-email").textContent = user.email || "Не указано";
        document.getElementById("user-age").textContent = user.age || "Не указано";
        document.getElementById("user-phone").textContent = user.phone || "Не указано";
        document.getElementById("user-address").textContent = user.address || "Не указано";
    } catch (error) {
        console.error("Ошибка загрузки данных аккаунта:", error);
        document.getElementById("user-info").innerHTML = "<p>Не удалось загрузить информацию.</p>";
    }
});


document.getElementById("logout").addEventListener("click", async function () {
    try {
        const response = await fetch("/logout", {
            method: "POST",
            credentials: "include",
        });

        if (response.ok) {
            localStorage.clear(); // Очистка локального хранилища
            window.location.href = "/login"; // Перенаправление на страницу входа
        } else {
            throw new Error("Ошибка при выходе");
        }
    } catch (error) {
        console.error("Ошибка выхода:", error);
        alert("Не удалось выйти из аккаунта");
    }
});



document.getElementById("change-password-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    const newPassword = document.getElementById("new-password").value;
    const confirmNewPassword = document.getElementById("confirm-new-password").value;

    if (newPassword !== confirmNewPassword) {
        alert("Passwords do not match!");
        return;
    }

    try {
        const response = await fetch("/change-password", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ password: newPassword }),
            credentials: "include",
        });

        if (response.ok) {
            alert("Password updated successfully!");
            document.getElementById("change-password-form").reset();
        } else {
            throw new Error("Failed to update password");
        }
    } catch (error) {
        console.error("Password change error:", error);
        alert("Error changing password");
    }
});


document.getElementById("add-info-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    const userData = {
        name: document.getElementById("name").value,
        age: document.getElementById("age").value,
        phone: document.getElementById("phone").value,
        address: document.getElementById("address").value
    };

    try {
        const response = await fetch("/update-user", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(userData)
        });

        if (!response.ok) throw new Error("Ошибка обновления данных");

        alert("Информация успешно обновлена!");

        location.reload(); // Перезагружаем страницу, чтобы загрузить актуальные данные
    } catch (error) {
        console.error("Ошибка:", error);
        alert("Не удалось обновить информацию");
    }
});

