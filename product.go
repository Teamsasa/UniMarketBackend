package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"strings"
	"strconv"
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
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// GETリクエストのみ許可
	if r.Method != http.MethodGet {
		http.Error(w, "GET method expected", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("getProducts called...")

	query := `
		SELECT 
			products.id, products.user_id, products.name, products.description, products.image_url, products.price,
			categories.name AS category, products.status, products.created_at, products.updated_at
		FROM 
			products
		INNER JOIN 
			categories ON products.category_id = categories.id
	`
	// クエリを実行して、結果を取得
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in query: %s", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 1行ずつ取得して、Product構造体に詰めていく
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.UserID, &product.Name, &product.Description, &product.ImageURL, &product.Price,
			&product.Category, &product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			http.Error(w, fmt.Sprintf("error scanning row: %s", err), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
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

	// リクエストボディの値をDBにINSERT
	query := `
		INSERT INTO products (user_id, name, description, image_url, price, category_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec(query, product.UserID, product.Name, product.Description, product.ImageURL, product.Price, product.Category, product.Status, currentTime, currentTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("error inserting into database: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Product added successfully")
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

	// 商品IDを元にDBから商品を削除
	query := `DELETE FROM products WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting from database: %s", err), http.StatusInternalServerError)
		return
	}

	// 削除された行数を取得
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching affected rows: %s", err), http.StatusInternalServerError)
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
