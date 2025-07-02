CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    vote_average INT CHECK (
        vote_average >= 0
        AND vote_average <= 10
    ),
    poster_path VARCHAR(500),
    backdrop_path VARCHAR(500),
    release_date DATE,
    runtime INT CHECK (runtime > 0),
    popularity INT DEFAULT 0,
    admin_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
