package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
	db    *gorm.DB
)

func main() {
	fmt.Println("Starting server Connection.....")
	InitDB()
	router := mux.NewRouter()
	router.HandleFunc("/api/register", register).Methods("POST")
	router.HandleFunc("/api/login", login).Methods("POST")
	router.HandleFunc("/api/activeUsers", activeUsers).Methods("GET")
	router.HandleFunc("/api/logout", logout).Methods("PUT")
	http.ListenAndServe(":8000", router)
}
