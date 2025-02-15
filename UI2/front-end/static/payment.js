function updateTotalAmount() {
    let cartData = JSON.parse(localStorage.getItem("shoppingCart")) || {};
    let totalAmount = 0;

    Object.keys(cartData).forEach(productId => {
        let item = cartData[productId];

        if (item && item.quantity && item.price) {
            totalAmount += item.price * item.quantity;
        } else {
            console.warn(`Ошибка в товаре ${productId}:`, item);
        }
    });

    let totalAmountElement = document.getElementById("total-amount");
    if (totalAmountElement) {
        totalAmountElement.textContent = `Total: ${totalAmount.toFixed(2)} USD`;
    } else {
        console.error("Элемент #total-amount не найден в payment.html");
    }
}

document.addEventListener("DOMContentLoaded", updateTotalAmount);

document.querySelector(".payment-button").addEventListener("click", function () {
    window.location.href = "/h";
});

