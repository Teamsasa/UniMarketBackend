package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	UserID      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Status      string  `json:"status"`
	University  string  `json:"university"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// プリフライトリクエスト
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// GETリクエストのみ許可
	if r.Method != http.MethodGet {
		http.Error(w, "GET method expected", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("getProducts called...")

	// URLパスの一部を取得
	path := strings.TrimPrefix(r.URL.Path, "/getProducts/")

	// クエリのベース
	query := `
		SELECT 
			products.id, products.user_id, products.name, products.description, products.image_url, products.price,
			categories.name AS category, products.status, university, products.created_at, products.updated_at
		FROM 
			products
		INNER JOIN 
			categories ON products.category_id = categories.id
	`

	var rows *sql.Rows
	var err error

	// パスパラメーターが空でない場合、LIKE検索を追加
	if path != "" {
		query += " WHERE products.name LIKE $1 OR products.description LIKE $1"
		rows, err = db.Query(query, "%"+path+"%")
	} else {
		rows, err = db.Query(query)
	}

	// クエリを実行して、結果を取得
	if err != nil {
		http.Error(w, fmt.Sprintf("error in query: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// cookieから大学を取得
	university, err := getCookie(r, "university")
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting university from cookie: %s", err), http.StatusBadRequest)
		return
	}

	// 1行ずつ取得して、Product構造体に詰めていく
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.UserID, &product.Name, &product.Description, &product.ImageURL, &product.Price,
			&product.Category, &product.Status, &product.University, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("error scanning row: %s", err), http.StatusInternalServerError)
			return
		}
		if product.University == university {
			products = append(products, product)
		}
	}
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("error in rows iteration: %s", err), http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %s", err), http.StatusInternalServerError)
		return
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	// POSTリクエストのみ許可
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("addProduct called...")

	// リクエストボディをデコード
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %s", err), http.StatusBadRequest)
		return
	}

	// cookieから大学を取得
	university, err := getCookie(r, "university")
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting university from cookie: %s", err), http.StatusBadRequest)
		return
	}

	product.ImageURL = "./images/default.jpg"

	// リクエストボディの値をDBにINSERT
	query := `
		INSERT INTO products (user_id, name, description, image_url, price, category_id, status, university, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec(query, product.UserID, product.Name, product.Description, product.ImageURL, product.Price, product.Category, product.Status, university, currentTime, currentTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("error inserting into database: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Product added successfully")
}

func editProduct(w http.ResponseWriter, r *http.Request) {

	// PUTリクエストのみ許可
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// URLからIDを抽出
	path := r.URL.Path
	segments := strings.Split(path, "/")
	if len(segments) < 3 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	id := segments[2] // /editProduct/{id} の {id} 部分を取得

	fmt.Println("editProduct called...")

	// cookieから大学を取得
	university, err := getCookie(r, "university")
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting university from cookie: %s", err), http.StatusBadRequest)
		return
	}

	// リクエストボディをデコード
	var product Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request body: %s", err), http.StatusBadRequest)
		return
	}

	// 更新するフィールドと値を保持するマップ
	fieldsToUpdate := map[string]interface{}{}
	if product.Name != "" {
		fieldsToUpdate["name"] = product.Name
	}
	if product.Description != "" {
		fieldsToUpdate["description"] = product.Description
	}
	if product.ImageURL != "" {
		fieldsToUpdate["image_url"] = product.ImageURL
	}
	if product.Price != 0 {
		fieldsToUpdate["price"] = product.Price
	}
	if product.Category != "" {
		fieldsToUpdate["category_id"] = product.Category
	}
	if product.Status != "" {
		fieldsToUpdate["status"] = product.Status
	}

	// SQLクエリとパラメータを動的に構築
	query := "UPDATE products SET "
	params := []interface{}{}
	i := 1
	for field, value := range fieldsToUpdate {
		if i > 1 {
			query += ", "
		}
		query += fmt.Sprintf("%s = $%d", field, i)
		params = append(params, value)
		i++
	}
	query += fmt.Sprintf(", updated_at = $%d WHERE id = $%d AND university = $%d", i, i+1, i+2)
	currentTime := time.Now()
	params = append(params, currentTime)
	params = append(params, id)
	params = append(params, university)

	// データベースを更新
	result, err := db.Exec(query, params...)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating database: %s", err), http.StatusInternalServerError)
		return
	}

	// 影響を受けた行の数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("error checking affected rows: %s", err), http.StatusBadRequest)
		return
	}

	// 更新された行がない場合はエラーを返す
	if rowsAffected == 0 {
		http.Error(w, "no rows updated", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// DELETEリクエストのみ許可
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("deleteProduct called...")

	// URLパスの一部を取得
	path := strings.TrimPrefix(r.URL.Path, "/deleteProduct/")

	// パスパラメーターが空の場合はエラーを返す
	if path == "" {
		http.Error(w, "Product ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// cookieから大学を取得
	university, err := getCookie(r, "university")
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting university from cookie: %s", err), http.StatusBadRequest)
		return
	}

	// 商品IDを元にDBから商品を削除
	query := `DELETE FROM products WHERE id = $1 AND university = $2`
	result, err := db.Exec(query, id, university)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting from database: %s", err), http.StatusInternalServerError)
		return
	}

	// 削除された行数を取得
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching affected rows: %s", err), http.StatusBadRequest)
		return
	}

	// 削除された行数が0の場合はエラーを返す
	if rowsAffected == 0 {
		http.Error(w, "No product found with the given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Product deleted successfully")
}

func getImages(w http.ResponseWriter, r *http.Request) {
	// GETリクエストのみ許可
	if r.Method != http.MethodGet {
		http.Error(w, "GET method expected", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("getImages called...")

	// IDをURLから取得
	idStr := strings.TrimPrefix(r.URL.Path, "/getImages/")
	if idStr == "" {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	fmt.Println("ID:", id)
	// クエリのベース
	query := `SELECT image_url FROM products WHERE id = $1`

	// クエリを実行して、結果を取得
	var imageUrl string
	err = db.QueryRow(query, id).Scan(&imageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No image found for the given ID", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// 画像を返す
	http.ServeFile(w, r, imageUrl)
}
