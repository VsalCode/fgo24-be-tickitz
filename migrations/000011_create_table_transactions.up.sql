CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer_fullname VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255) NOT NULL,
    customer_phone VARCHAR(20),
    amount DECIMAL(10, 2) NOT NULL CHECK (amount > 0),
    cinema VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    show_time TIME NOT NULL,
    show_date DATE NOT NULL,
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    movie_id INT REFERENCES movies (id) ON DELETE CASCADE,
    payment_method_id INT REFERENCES payment_method (id) ON DELETE RESTRICT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
