CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(50),
    password VARCHAR(50)
);
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES  users (id),
    title VARCHAR(100),
    content TEXT
);

INSERT INTO users (id, email, password) VALUES (1, 'test@gmail.com', '621');