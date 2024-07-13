package main

import (
	"fmt"
	"net/http"
	"encoding/json"
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
	w.WriteHeader(http.StatusOK)
}
