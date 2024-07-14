package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID        string
	Username  string
	Email     string
	CreatedAt string
}

var cognitoRegion string
var clientId string
var jwksURL string

func main() {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		log.Fatal("データベースの環境変数が設定されていません")
	}

	cognitoRegion = os.Getenv("COGNITO_REGION")
	clientId = os.Getenv("COGNITO_CLIENT_ID")
	jwksURL = os.Getenv("TOKEN_KEY_URL")

	if cognitoRegion == "" || clientId == "" || jwksURL == "" {
		log.Fatal("Cognitoの環境変数が設定されていません")
	}

	var err error

	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/checkemail", checkEmail)
	http.HandleFunc("/resnedemail", resendEmail)
	http.HandleFunc("/getProducts/", getProducts)
	http.HandleFunc("/addProduct", addProduct)
	http.HandleFunc("/editProduct/", editProduct)
	http.HandleFunc("/deleteProduct/", deleteProduct)
	http.HandleFunc("/getImages/", getImages)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler called...")

	var user User
	err := db.QueryRow("SELECT id, username, email, created_at  FROM users WHERE id = $1", "1").Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		fmt.Printf("error in query: %s", err)
		return
	}

	fmt.Fprintf(w, "id: %s, username: %s, email: %s, created_at: %s", user.ID, user.Username, user.Email, user.CreatedAt)
}
