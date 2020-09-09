package delivery

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// inisiasi router package
var router = mux.NewRouter()

// Router init
func Router() {
	// route untuk home
	router.HandleFunc("/", homeHandler)

	router.HandleFunc("/blogs", blogsHandler).Methods("POST")

	// base route
	http.Handle("/", router)
	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
