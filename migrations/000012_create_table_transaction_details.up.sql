CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    seat VARCHAR(10) NOT NULL,
    transaction_id INT REFERENCES transactions (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
