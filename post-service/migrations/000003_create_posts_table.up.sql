CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       content TEXT,
                       user_id INT REFERENCES users(id) ON DELETE CASCADE,
                       category_id INT REFERENCES categories(id) ON DELETE CASCADE
);
