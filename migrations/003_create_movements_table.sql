CREATE TABLE movements (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id),
    type VARCHAR(50) NOT NULL CHECK (type IN ('Entrada', 'Saída')),
    quantity INT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);