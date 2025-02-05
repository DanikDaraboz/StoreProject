db = db.getSiblingDB("ecommerce");

db.products.insertMany([
    { name: "Tshirt", price: 1000 },
    { name: "Hat", price: 500 }
]);

db.users.insertMany([
    { name: "John Doe", email: "john@example.com" }
]);

print("Test data inserted");
