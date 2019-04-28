package main

import (
	"fmt"
	"go-rest-api-basic/app"
	"go-rest-api-basic/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/profile/get", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/api/profile/create", controllers.CreateProfile).Methods("POST")
	router.HandleFunc("/api/profile/update", controllers.UpdateProfile).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
