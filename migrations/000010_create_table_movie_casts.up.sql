CREATE TABLE movie_casts (
    movie_id INT,
    cast_id INT,
    PRIMARY KEY (movie_id, cast_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE,
    FOREIGN KEY (cast_id) REFERENCES casts (id) ON DELETE CASCADE
);
