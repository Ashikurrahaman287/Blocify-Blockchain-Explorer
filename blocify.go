package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/block/{blockHash}", blockHandler).Methods("GET")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	log.Println("Blocify Blockchain Explorer running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Render the home page
	http.ServeFile(w, r, "templates/index.html")
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	// Extract block hash from URL
	vars := mux.Vars(r)
	blockHash := vars["blockHash"]

	// Render the block details page
	// For now, just display the block hash
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Block Hash: " + blockHash))
}
