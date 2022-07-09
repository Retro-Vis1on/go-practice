package main

import (
	"github.com/Retro-Vis1on/go-practice/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
