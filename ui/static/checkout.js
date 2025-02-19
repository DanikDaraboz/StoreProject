document.getElementById('checkoutBtn')?.addEventListener('click', function () {
    // Скрываем контейнер с корзиной (с анимацией исчезновения)
    const cartContainer = document.querySelector('.cart');
    cartContainer.classList.add('cart-container'); // Добавляем класс для анимации исчезновения
  
    // Показываем форму для оформления заказа (с анимацией появления)
    const checkoutFormContainer = document.getElementById('checkout-form-container');
    checkoutFormContainer.style.display = 'block'; // Показываем форму
    checkoutFormContainer.classList.add('checkout-form'); // Добавляем класс для анимации появления
  });
  