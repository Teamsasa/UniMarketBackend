DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INSERT INTO users (id, email, password) VALUES ('1', 'example', 'example');