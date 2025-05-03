CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255),
    category VARCHAR(255),
    price NUMERIC(10, 2) NOT NULL,
    quantity INT NOT NULL,
    expiry_date DATE,
    sku VARCHAR(255) UNIQUE NOT NULL
);