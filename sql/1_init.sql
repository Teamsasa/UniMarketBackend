DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS products;

CREATE TABLE users (
  id VARCHAR(255) NOT NULL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  image_url VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  category_id INTEGER NOT NULL,
  status VARCHAR(50) NOT NULL DEFAULT 'available',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO users (id, username, email, created_at) VALUES
  ('1', 'example', 'example', '2024-07-10 10:00:00');

INSERT INTO categories (name) VALUES
('数学'),
('物理学'),
('化学'),
('生物学'),
('コンピュータサイエンス'),
('工学'),
('経済学'),
('経営学'),
('文学'),
('歴史'),
('心理学'),
('社会学'),
('哲学'),
('政治学'),
('法学'),
('医学'),
('看護学'),
('環境科学'),
('教育学'),
('言語学'),
('その他');

-- テスト商品データ
INSERT INTO products (user_id, name, description, image_url, price, category_id, status, created_at, updated_at) VALUES
  -- 数学
  ('1', '数学の本1', '数学に関する本です。', 'https://example.com/image1.jpg', 1000, '1', 'available', '2024-07-10 10:00:00', '2024-07-10 10:00:00'),
  ('1', '数学の本2', '数学に関する本です。', 'https://example.com/image2.jpg', 1000, '1', 'available', '2024-07-11 11:00:00', '2024-07-11 11:00:00'),
  ('1', '数学の本3', '数学に関する本です。', 'https://example.com/image3.jpg', 1000, '1', 'available', '2024-07-12 12:00:00', '2024-07-12 12:00:00'),
  ('1', '数学の本4', '数学に関する本です。', 'https://example.com/image4.jpg', 1000, '1', 'available', '2024-07-13 13:00:00', '2024-07-13 13:00:00'),
  ('1', '数学の本5', '数学に関する本です。', 'https://example.com/image5.jpg', 1000, '1', 'available', '2024-07-14 14:00:00', '2024-07-14 14:00:00'),
  ('1', '数学の本6', '数学に関する本です。', 'https://example.com/image6.jpg', 1000, '1', 'available', '2024-07-15 15:00:00', '2024-07-15 15:00:00'),
  ('1', '数学の本7', '数学に関する本です。', 'https://example.com/image7.jpg', 1000, '1', 'available', '2024-07-16 16:00:00', '2024-07-16 16:00:00'),
  ('1', '数学の本8', '数学に関する本です。', 'https://example.com/image8.jpg', 1000, '1', 'available', '2024-07-17 17:00:00', '2024-07-17 17:00:00'),
  ('1', '数学の本9', '数学に関する本です。', 'https://example.com/image9.jpg', 1000, '1', 'available', '2024-07-18 18:00:00', '2024-07-18 18:00:00'),
  ('1', '数学の本10', '数学に関する本です。', 'https://example.com/image10.jpg', 1000, '1', 'available', '2024-07-19 19:00:00', '2024-07-19 19:00:00'),

  -- 物理学
  ('1', '物理学の本1', '物理学に関する本です。', 'https://example.com/image1.jpg', 1500, '2', 'available', '2024-07-20 10:00:00', '2024-07-20 10:00:00'),
  ('1', '物理学の本2', '物理学に関する本です。', 'https://example.com/image2.jpg', 1500, '2', 'available', '2024-07-21 11:00:00', '2024-07-21 11:00:00'),
  ('1', '物理学の本3', '物理学に関する本です。', 'https://example.com/image3.jpg', 1500, '2', 'available', '2024-07-22 12:00:00', '2024-07-22 12:00:00'),
  ('1', '物理学の本4', '物理学に関する本です。', 'https://example.com/image4.jpg', 1500, '2', 'available', '2024-07-23 13:00:00', '2024-07-23 13:00:00'),
  ('1', '物理学の本5', '物理学に関する本です。', 'https://example.com/image5.jpg', 1500, '2', 'available', '2024-07-24 14:00:00', '2024-07-24 14:00:00'),
  ('1', '物理学の本6', '物理学に関する本です。', 'https://example.com/image6.jpg', 1500, '2', 'available', '2024-07-25 15:00:00', '2024-07-25 15:00:00'),
  ('1', '物理学の本7', '物理学に関する本です。', 'https://example.com/image7.jpg', 1500, '2', 'available', '2024-07-26 16:00:00', '2024-07-26 16:00:00'),
  ('1', '物理学の本8', '物理学に関する本です。', 'https://example.com/image8.jpg', 1500, '2', 'available', '2024-07-27 17:00:00', '2024-07-27 17:00:00'),
  ('1', '物理学の本9', '物理学に関する本です。', 'https://example.com/image9.jpg', 1500, '2', 'available', '2024-07-28 18:00:00', '2024-07-28 18:00:00'),
  ('1', '物理学の本10', '物理学に関する本です。', 'https://example.com/image10.jpg', 1500, '2', 'available', '2024-07-29 19:00:00', '2024-07-29 19:00:00'),

  -- 化学
  ('1', '化学の本1', '化学に関する本です。', 'https://example.com/image1.jpg', 1200, '3', 'available', '2024-07-30 10:00:00', '2024-07-30 10:00:00'),
  ('1', '化学の本2', '化学に関する本です。', 'https://example.com/image2.jpg', 1200, '3', 'available', '2024-07-31 11:00:00', '2024-07-31 11:00:00'),
  ('1', '化学の本3', '化学に関する本です。', 'https://example.com/image3.jpg', 1200, '3', 'available', '2024-08-01 12:00:00', '2024-08-01 12:00:00'),
  ('1', '化学の本4', '化学に関する本です。', 'https://example.com/image4.jpg', 1200, '3', 'available', '2024-08-02 13:00:00', '2024-08-02 13:00:00'),
  ('1', '化学の本5', '化学に関する本です。', 'https://example.com/image5.jpg', 1200, '3', 'available', '2024-08-03 14:00:00', '2024-08-03 14:00:00'),
  ('1', '化学の本6', '化学に関する本です。', 'https://example.com/image6.jpg', 1200, '3', 'available', '2024-08-04 15:00:00', '2024-08-04 15:00:00'),
  ('1', '化学の本7', '化学に関する本です。', 'https://example.com/image7.jpg', 1200, '3', 'available', '2024-08-05 16:00:00', '2024-08-05 16:00:00'),
  ('1', '化学の本8', '化学に関する本です。', 'https://example.com/image8.jpg', 1200, '3', 'available', '2024-08-06 17:00:00', '2024-08-06 17:00:00'),
  ('1', '化学の本9', '化学に関する本です。', 'https://example.com/image9.jpg', 1200, '3', 'available', '2024-08-07 18:00:00', '2024-08-07 18:00:00'),
  ('1', '化学の本10', '化学に関する本です。', 'https://example.com/image10.jpg', 1200, '3', 'available', '2024-08-08 19:00:00', '2024-08-08 19:00:00'),

  -- 生物学
  ('1', '生物学の本1', '生物学に関する本です。', 'https://example.com/image1.jpg', 1300, '4', 'available', '2024-08-09 10:00:00', '2024-08-09 10:00:00'),
  ('1', '生物学の本2', '生物学に関する本です。', 'https://example.com/image2.jpg', 1300, '4', 'available', '2024-08-10 11:00:00', '2024-08-10 11:00:00'),
  ('1', '生物学の本3', '生物学に関する本です。', 'https://example.com/image3.jpg', 1300, '4', 'available', '2024-08-11 12:00:00', '2024-08-11 12:00:00'),
  ('1', '生物学の本4', '生物学に関する本です。', 'https://example.com/image4.jpg', 1300, '4', 'available', '2024-08-12 13:00:00', '2024-08-12 13:00:00'),
  ('1', '生物学の本5', '生物学に関する本です。', 'https://example.com/image5.jpg', 1300, '4', 'available', '2024-08-13 14:00:00', '2024-08-13 14:00:00'),
  ('1', '生物学の本6', '生物学に関する本です。', 'https://example.com/image6.jpg', 1300, '4', 'available', '2024-08-14 15:00:00', '2024-08-14 15:00:00'),
  ('1', '生物学の本7', '生物学に関する本です。', 'https://example.com/image7.jpg', 1300, '4', 'available', '2024-08-15 16:00:00', '2024-08-15 16:00:00'),
  ('1', '生物学の本8', '生物学に関する本です。', 'https://example.com/image8.jpg', 1300, '4', 'available', '2024-08-16 17:00:00', '2024-08-16 17:00:00'),
  ('1', '生物学の本9', '生物学に関する本です。', 'https://example.com/image9.jpg', 1300, '4', 'available', '2024-08-17 18:00:00', '2024-08-17 18:00:00'),
  ('1', '生物学の本10', '生物学に関する本です。', 'https://example.com/image10.jpg', 1300, '4', 'available', '2024-08-18 19:00:00', '2024-08-18 19:00:00'),

  -- コンピュータサイエンス
  ('1', 'コンピュータサイエンスの本1', 'コンピュータサイエンスに関する本です。', 'https://example.com/image1.jpg', 2000, '5', 'available', '2024-08-19 10:00:00', '2024-08-19 10:00:00'),
  ('1', 'コンピュータサイエンスの本2', 'コンピュータサイエンスに関する本です。', 'https://example.com/image2.jpg', 2000, '5', 'available', '2024-08-20 11:00:00', '2024-08-20 11:00:00'),
  ('1', 'コンピュータサイエンスの本3', 'コンピュータサイエンスに関する本です。', 'https://example.com/image3.jpg', 2000, '5', 'available', '2024-08-21 12:00:00', '2024-08-21 12:00:00'),
  ('1', 'コンピュータサイエンスの本4', 'コンピュータサイエンスに関する本です。', 'https://example.com/image4.jpg', 2000, '5', 'available', '2024-08-22 13:00:00', '2024-08-22 13:00:00'),
  ('1', 'コンピュータサイエンスの本5', 'コンピュータサイエンスに関する本です。', 'https://example.com/image5.jpg', 2000, '5', 'available', '2024-08-23 14:00:00', '2024-08-23 14:00:00'),
  ('1', 'コンピュータサイエンスの本6', 'コンピュータサイエンスに関する本です。', 'https://example.com/image6.jpg', 2000, '5', 'available', '2024-08-24 15:00:00', '2024-08-24 15:00:00'),
  ('1', 'コンピュータサイエンスの本7', 'コンピュータサイエンスに関する本です。', 'https://example.com/image7.jpg', 2000, '5', 'available', '2024-08-25 16:00:00', '2024-08-25 16:00:00'),
  ('1', 'コンピュータサイエンスの本8', 'コンピュータサイエンスに関する本です。', 'https://example.com/image8.jpg', 2000, '5', 'available', '2024-08-26 17:00:00', '2024-08-26 17:00:00'),
  ('1', 'コンピュータサイエンスの本9', 'コンピュータサイエンスに関する本です。', 'https://example.com/image9.jpg', 2000, '5', 'available', '2024-08-27 18:00:00', '2024-08-27 18:00:00'),
  ('1', 'コンピュータサイエンスの本10', 'コンピュータサイエンスに関する本です。', 'https://example.com/image10.jpg', 2000, '5', 'available', '2024-08-28 19:00:00', '2024-08-28 19:00:00'),

  -- 工学
  ('1', '工学の本1', '工学に関する本です。', 'https://example.com/image1.jpg', 1800, '6', 'available', '2024-08-29 10:00:00', '2024-08-29 10:00:00'),
  ('1', '工学の本2', '工学に関する本です。', 'https://example.com/image2.jpg', 1800, '6', 'available', '2024-08-30 11:00:00', '2024-08-30 11:00:00'),
  ('1', '工学の本3', '工学に関する本です。', 'https://example.com/image3.jpg', 1800, '6', 'available', '2024-08-31 12:00:00', '2024-08-31 12:00:00'),
  ('1', '工学の本4', '工学に関する本です。', 'https://example.com/image4.jpg', 1800, '6', 'available', '2024-09-01 13:00:00', '2024-09-01 13:00:00'),
  ('1', '工学の本5', '工学に関する本です。', 'https://example.com/image5.jpg', 1800, '6', 'available', '2024-09-02 14:00:00', '2024-09-02 14:00:00'),
  ('1', '工学の本6', '工学に関する本です。', 'https://example.com/image6.jpg', 1800, '6', 'available', '2024-09-03 15:00:00', '2024-09-03 15:00:00'),
  ('1', '工学の本7', '工学に関する本です。', 'https://example.com/image7.jpg', 1800, '6', 'available', '2024-09-04 16:00:00', '2024-09-04 16:00:00'),
  ('1', '工学の本8', '工学に関する本です。', 'https://example.com/image8.jpg', 1800, '6', 'available', '2024-09-05 17:00:00', '2024-09-05 17:00:00'),
  ('1', '工学の本9', '工学に関する本です。', 'https://example.com/image9.jpg', 1800, '6', 'available', '2024-09-06 18:00:00', '2024-09-06 18:00:00'),
  ('1', '工学の本10', '工学に関する本です。', 'https://example.com/image10.jpg', 1800, '6', 'available', '2024-09-07 19:00:00', '2024-09-07 19:00:00'),

  -- その他
  ('1', 'その他の本1', 'その他の本です。', 'https://example.com/image1.jpg', 1500, '21', 'available', '2024-09-08 10:00:00', '2024-09-08 10:00:00'),
  ('1', 'その他の本2', 'その他の本です。', 'https://example.com/image2.jpg', 1500, '21', 'available', '2024-09-09 11:00:00', '2024-09-09 11:00:00'),
  ('1', 'その他の本3', 'その他の本です。', 'https://example.com/image3.jpg', 1500, '21', 'available', '2024-09-10 12:00:00', '2024-09-10 12:00:00'),
  ('1', 'その他の本4', 'その他の本です。', 'https://example.com/image4.jpg', 1500, '21', 'available', '2024-09-11 13:00:00', '2024-09-11 13:00:00'),
  ('1', 'その他の本5', 'その他の本です。', 'https://example.com/image5.jpg', 1500, '21', 'available', '2024-09-12 14:00:00', '2024-09-12 14:00:00'),
  ('1', 'その他の本6', 'その他の本です。', 'https://example.com/image6.jpg', 1500, '21', 'available', '2024-09-13 15:00:00', '2024-09-13 15:00:00'),
  ('1', 'その他の本7', 'その他の本です。', 'https://example.com/image7.jpg', 1500, '21', 'available', '2024-09-14 16:00:00', '2024-09-14 16:00:00'),
  ('1', 'その他の本8', 'その他の本です。', 'https://example.com/image8.jpg', 1500, '21', 'available', '2024-09-15 17:00:00', '2024-09-15 17:00:00'),
  ('1', 'その他の本9', 'その他の本です。', 'https://example.com/image9.jpg', 1500, '21', 'available', '2024-09-16 18:00:00', '2024-09-16 18:00:00'),
  ('1', 'その他の本10', 'その他の本です。', 'https://example.com/image10.jpg', 1500, '21', 'available', '2024-09-17 19:00:00', '2024-09-17 19:00:00');
