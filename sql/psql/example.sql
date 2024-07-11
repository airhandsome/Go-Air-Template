-- example.sql

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- A simple query to retrieve a user by ID
SELECT * FROM users WHERE id = $1;

-- A query to insert a new user
INSERT INTO users (username, email) VALUES ($1,$2) RETURNING id;

-- A query to update a user's email
UPDATE users SET email = $1 WHERE id =$2;

-- A query to delete a user by ID
DELETE FROM users WHERE id = $1;
