document.addEventListener("DOMContentLoaded", function () {
    const addToCartButtons = document.querySelectorAll(".add-to-cart");

    addToCartButtons.forEach(button => {
        button.addEventListener("click", function () {
            const productId = this.getAttribute("data-id");

            fetch("/cart/add", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                body: "id=" + encodeURIComponent(productId)
            }).then(response => {
                if (response.ok) {
                    alert("Product added to cart!");
                } else {
                    alert("Failed to add product.");
                }
            }).catch(error => console.error("Error:", error));
        });
    });
});
document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".remove-from-cart").forEach(button => {
        button.addEventListener("click", function () {
            let productId = this.getAttribute("data-product-id");

            fetch('/cart/remove', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: 'id=' + encodeURIComponent(productId),
            })
            .then(response => {
                if (response.ok) {
                    location.reload(); // Перезагружаем страницу корзины
                } else {
                    alert('Error removing product from cart');
                }
            })
            .catch(error => console.error('Error:', error));
        });
    });
});
document.addEventListener("DOMContentLoaded", function () {
    const cartKey = "shoppingCart";

    let cartData = JSON.parse(localStorage.getItem(cartKey)) || {};

    document.querySelectorAll(".product-quantity").forEach(quantityElement => {
        const productId = quantityElement.dataset.productId;
        const priceElement = document.getElementById(`price-${productId}`);
        const basePrice = parseFloat(priceElement.parentElement.dataset.price);

        if (cartData[productId]) {
            quantityElement.innerText = cartData[productId].quantity;
            priceElement.innerText = `${cartData[productId].totalPrice || (basePrice * cartData[productId].quantity).toFixed(2)}`;
        }
    });

    document.querySelectorAll(".change-quantity").forEach(button => {
        button.addEventListener("click", function () {
            const productId = this.dataset.productId;
            const action = this.dataset.action;
            const quantityElement = document.querySelector(`.product-quantity[data-product-id="${productId}"]`);
            const priceElement = document.getElementById(`price-${productId}`);
            const basePrice = parseFloat(priceElement.parentElement.dataset.price);
            let quantity = parseInt(quantityElement.innerText);

            if (action === "increase") {
                quantity++;
            } else if (action === "decrease" && quantity > 1) {
                quantity--;
            }

            quantityElement.innerText = quantity;
            priceElement.innerText = (basePrice * quantity).toFixed(2);

            cartData[productId] = { quantity: quantity, price: basePrice };
            localStorage.setItem(cartKey, JSON.stringify(cartData));

            // Отправляем изменения на сервер
            fetch("/cart/update", {
                method: "POST",
                headers: { "Content-Type": "application/x-www-form-urlencoded" },
                body: `id=${encodeURIComponent(productId)}&quantity=${quantity}`
            }).catch(error => console.error("Error updating cart:", error));
        });
    });
});

document.querySelector(".buy").addEventListener("click", function () {
    window.location.href = "/payment";
});

