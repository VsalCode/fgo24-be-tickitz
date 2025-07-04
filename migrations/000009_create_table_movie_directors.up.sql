CREATE TABLE movie_directors (
    movie_id INT,
    director_id INT,
    PRIMARY KEY (movie_id, director_id),
    FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE,
    FOREIGN KEY (director_id) REFERENCES directors (id) ON DELETE CASCADE
);
