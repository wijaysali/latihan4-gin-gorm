CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    customer_id INT UNIQUE,
    address TEXT,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10, 2)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_price DECIMAL(10, 2),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT,
    price DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE product_categories (
    product_id INT,
    category_id INT,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    PRIMARY KEY (product_id, category_id)
);

INSERT INTO customers (name, email) VALUES 
('John Doe', 'john.doe@example.com'),
('Jane Smith', 'jane.smith@example.com');

INSERT INTO addresses (customer_id, address) VALUES 
(1, '123 Main St, Anytown'),
(2, '456 Oak St, Othertown');

INSERT INTO products (name, price) VALUES 
('Product A', 10.00),
('Product B', 20.00),
('Product C', 30.00);

INSERT INTO orders (customer_id, total_price) VALUES 
(1, 50.00),
(2, 70.00);

INSERT INTO order_details (order_id, product_id, quantity, price) VALUES 
(1, 1, 2, 20.00),
(1, 2, 1, 20.00),
(2, 2, 2, 40.00),
(2, 3, 1, 30.00);

INSERT INTO categories (name) VALUES 
('Electronics'),
('Books'),
('Clothing');

INSERT INTO product_categories (product_id, category_id) VALUES 
(1, 1),
(1, 2),
(2, 3),
(3, 1);


