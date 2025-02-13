document.addEventListener("DOMContentLoaded", function () {
    // ... другой код (регистрация, изменение пароля и т.д.)

    // ********** Логика добавления информации **********
    const addInfoBtn = document.getElementById("add-info-btn");
    const addInfoForm = document.getElementById("add-info-form");

    // Изначально скрываем форму добавления информации
    if (addInfoForm) {
        addInfoForm.style.display = "none";
    }

    if (addInfoBtn && addInfoForm) {
        addInfoBtn.addEventListener("click", function () {
            addInfoForm.style.display = (addInfoForm.style.display === "none" || addInfoForm.style.display === "")
                ? "block"
                : "none";
        });
    }

    if (addInfoForm) {
        addInfoForm.addEventListener("submit", function (event) {
            event.preventDefault();

            // Получаем значения полей формы
            const nameField = document.getElementById("name");
            const ageField = document.getElementById("age");
            const phoneField = document.getElementById("phone");
            const addressField = document.getElementById("address");

            const name = nameField ? nameField.value.trim() : "";
            const age = ageField ? ageField.value.trim() : "";
            const phone = phoneField ? phoneField.value.trim() : "";
            const address = addressField ? addressField.value.trim() : "";

            // Простая валидация (можно расширить по необходимости)
            if (!name || !age || !phone || !address) {
                alert("Пожалуйста, заполните все поля!");
                return;
            }

            // Сохраняем информацию в localStorage
            localStorage.setItem("userName", name);
            localStorage.setItem("userAge", age);
            localStorage.setItem("userPhone", phone);
            localStorage.setItem("userAddress", address);

            alert("Информация успешно добавлена!");

            // Очистка полей и скрытие формы после отправки
            nameField.value = "";
            ageField.value = "";
            phoneField.value = "";
            addressField.value = "";
            addInfoForm.style.display = "none";
        });
    }
});