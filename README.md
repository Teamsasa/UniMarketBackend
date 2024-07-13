## docker 利用方法
Docker Composeを使ってポート8080でサービスを起動
```shell
docker compose up -d --build
```
Docker Composeを使ってサービスを停止し、ボリュームも削除
```shell
docker compose down -v
```
noneイメージを全削除できる有能コマンド
```shell
docker image prune
```

## 各DBテーブル
### users Table

| Column     | Type          | Constraints                    |
|------------|---------------|--------------------------------|
| id         | VARCHAR(255)  | NOT NULL, PRIMARY KEY          |
| username   | VARCHAR(255)  | NOT NULL, UNIQUE               |
| email      | VARCHAR(255)  | NOT NULL                       |
| created_at | TIMESTAMP     | DEFAULT CURRENT_TIMESTAMP      |

### categories Table

| Column     | Type          | Constraints                    |
|------------|---------------|--------------------------------|
| id         | SERIAL        | PRIMARY KEY                    |
| name       | VARCHAR(255)  | NOT NULL, UNIQUE               |

### products Table

| Column      | Type           | Constraints                  |
|-------------|----------------|------------------------------|
| id          | SERIAL         | PRIMARY KEY                  |
| user_id     | VARCHAR(255)   | NOT NULL, FOREIGN KEY        |
| name        | VARCHAR(255)   | NOT NULL                     |
| description | TEXT           |                              |
| image_url   | VARCHAR(255)   | NOT NULL                     |
| price       | DECIMAL(10, 2) | NOT NULL                     |
| category_id | INTEGER        | NOT NULL, FOREIGN KEY        |
| status      | VARCHAR(50)    | NOT NULL DEFAULT 'available' |
| created_at  | TIMESTAMP      | DEFAULT CURRENT_TIMESTAMP    |
| updated_at  | TIMESTAMP      | DEFAULT CURRENT_TIMESTAMP    |

#### 制約
- `FOREIGN KEY (user_id)` は `users(id)` への参照。
- `FOREIGN KEY (category_id)` は `categories(id)` への参照。