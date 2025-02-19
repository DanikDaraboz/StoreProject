document.getElementById("payButton").addEventListener("click", async () => {
    try {
        const response = await fetch("/api/cart"); // Получаем корзину пользователя
        const cart = await response.json();
        
        if (!cart.items || cart.items.length === 0) {
            alert("Your cart is empty");
            return;
        }

        const orderData = {
            user_id: cart.user_id, // Получаем ID пользователя
            items: cart.items.map(item => ({
                product_id: item.product_id,
                quantity: item.quantity,
                price: item.price
            })),
            total_price: cart.items.reduce((sum, item) => sum + item.price * item.quantity, 0)
        };

        const orderResponse = await fetch("/api/orders", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(orderData)
        });

        if (!orderResponse.ok) {
            throw new Error("Failed to create order");
        }

        const order = await orderResponse.json();
        alert("Order created successfully! Order ID: " + order.id);
    } catch (error) {
        console.error("Error creating order:", error);
        alert("Error creating order");
    }
});
