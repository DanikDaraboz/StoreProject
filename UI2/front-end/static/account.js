document.addEventListener("DOMContentLoaded", function () {
    // ********** Логика регистрации **********
    const signUpForm = document.querySelector("form");
    const signInBtn = document.getElementById("sign-in-btn");
    const userIcon = document.getElementById("user-icon");
    const emailError = document.getElementById("emailError");
    const passwordError = document.getElementById("passwordError");
    const confirmPasswordError = document.getElementById("confirmPasswordError");
    const passwordInput = document.getElementById("password");
    const confirmPasswordInput = document.getElementById("confirm-password");

    if (localStorage.getItem("userLoggedIn") === "true") {
        if (signInBtn) signInBtn.style.display = "none";
        if (userIcon) userIcon.style.display = "inline-block";
    }

    if (signUpForm) {
        signUpForm.addEventListener("submit", function (event) {
            event.preventDefault();
            let valid = true;

            const emailField = document.getElementById("email");
            const email = emailField ? emailField.value.trim() : "";
            if (!email) {
                if (emailError) emailError.textContent = "Email обязателен.";
                valid = false;
            } else {
                if (emailError) emailError.textContent = "";
            }

            const password = passwordInput ? passwordInput.value : "";
            const confirmPassword = confirmPasswordInput ? confirmPasswordInput.value : "";

            if (!password) {
                if (passwordError) passwordError.textContent = "Пароль обязателен.";
                valid = false;
            } else {
                if (passwordError) passwordError.textContent = "";
            }

            if (password !== confirmPassword) {
                if (confirmPasswordError) confirmPasswordError.textContent = "Пароли не совпадают.";
                valid = false;
            } else {
                if (confirmPasswordError) confirmPasswordError.textContent = "";
            }

            if (valid) {
                localStorage.setItem("userLoggedIn", "true");
                localStorage.setItem("userEmail", email);

                if (signInBtn) signInBtn.style.display = "none";
                if (userIcon) userIcon.style.display = "inline-block";

                setTimeout(() => {
                    window.location.href = "/account";
                }, 500);
            }
        });
    }

    // ********** Логика изменения пароля **********
    const changePasswordBtn = document.getElementById("change-password-btn");
    const changePasswordForm = document.getElementById("change-password-form");

    if (changePasswordForm) {
        changePasswordForm.style.display = "none";
    }

    if (changePasswordBtn && changePasswordForm) {
        changePasswordBtn.addEventListener("click", function () {
            changePasswordForm.style.display = (changePasswordForm.style.display === "none" || changePasswordForm.style.display === "") 
                ? "block" 
                : "none";
        });
    }

    if (changePasswordForm) {
        changePasswordForm.addEventListener("submit", function (event) {
            event.preventDefault();

            const newPasswordField = document.getElementById("new-password");
            const confirmNewPasswordField = document.getElementById("confirm-new-password");
            const newPassword = newPasswordField ? newPasswordField.value : "";
            const confirmNewPassword = confirmNewPasswordField ? confirmNewPasswordField.value : "";

            if (newPassword && newPassword === confirmNewPassword) {
                alert("Пароль успешно изменён!");
                newPasswordField.value = "";
                confirmNewPasswordField.value = "";
                changePasswordForm.style.display = "none";
            } else {
                alert("Пароли не совпадают или не введены!");
            }
        });
    }

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

            // Здесь можно сохранить или отправить введённую информацию на сервер
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
