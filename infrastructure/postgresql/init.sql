-- Create the database
CREATE DATABASE simple_ecommerce;

-- Switch to the simple_ecommerce database
\c simple_ecommerce


CREATE SCHEMA simple_ecommerce;

CREATE TABLE IF NOT EXISTS simple_ecommerce.user (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS simple_ecommerce.product (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS simple_ecommerce.cart (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES simple_ecommerce.user(id),
    FOREIGN KEY (product_id) REFERENCES simple_ecommerce.product(id)
);

/* for mock checkout */
CREATE TABLE IF NOT EXISTS simple_ecommerce.payment_method (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    bank_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES simple_ecommerce.user(id)
);

CREATE TABLE IF NOT EXISTS simple_ecommerce.order (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    cart_ids INT[] NOT NULL,
    payment_method_id INT NOT NULL,
    total_price INT NOT NULL,
    is_paid BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES simple_ecommerce.user(id),
    FOREIGN KEY (payment_method_id) REFERENCES simple_ecommerce.payment_method(id)
);
/* Insert sample products */
DO $$ 
BEGIN
    FOR i IN 1..20 LOOP
        INSERT INTO simple_ecommerce.product (title, price, description) 
        VALUES ('Product ' || i, (RANDOM() * 100)::INT, 'Description for product ' || i);
    END LOOP;
END $$;