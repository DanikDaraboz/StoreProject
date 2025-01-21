import React from "react";

const ProductList = () => {
    const products = [
        { id: 1, name: "Basketball", price: "$20" },
        { id: 2, name: "Running Shoes", price: "$120" },
    ];

    return (
        <div className="product-list">
            {products.map((product) => (
                <div key={product.id} className="product-card">
                    <h2>{product.name}</h2>
                    <p>{product.price}</p>
                </div>
            ))}
        </div>
    );
};

export default ProductList;
