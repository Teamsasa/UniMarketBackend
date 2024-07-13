DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

INSERT INTO users (id, email, password) VALUES
  ('1', 'example', 'example');