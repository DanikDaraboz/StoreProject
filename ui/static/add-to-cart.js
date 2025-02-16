document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".add-to-cart").forEach(button => {
        button.addEventListener("click", function() {
            const productID = this.getAttribute("data-product-id");
            const price = parseFloat(this.getAttribute("data-price"));
            const productName = button.getAttribute('data-name');
            const cartItem = {
                product_id: productID,
                quantity: 1,
                price: price,
                name: productName,   // Название товара
            };
            fetch("/cart", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(cartItem)
            })
            .then(response => response.json())
            .then(data => {
                if (data.success) {
                    alert("Product added to cart!");
                } else {
                    alert("Failed to add product to cart.");
                }
            })
            .catch(error => console.error("Error:", error));
        });
    });
});