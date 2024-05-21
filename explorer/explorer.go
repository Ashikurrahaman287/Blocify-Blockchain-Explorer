package explorer

import (
	"blocify/blockchain"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Explorer represents the blockchain explorer
type Explorer struct {
	Blockchain *blockchain.Blockchain
}

// NewExplorer creates a new blockchain explorer
func NewExplorer(bc *blockchain.Blockchain) *Explorer {
	return &Explorer{Blockchain: bc}
}

// StartServer starts the web server for the blockchain explorer
func (exp *Explorer) StartServer() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", exp.homeHandler).Methods("GET")
	router.HandleFunc("/block/{blockHash}", exp.blockHandler).Methods("GET")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	log.Println("Blocify Blockchain Explorer running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// homeHandler handles requests to the home page
func (exp *Explorer) homeHandler(w http.ResponseWriter, r *http.Request) {
	// Render the home page
	http.ServeFile(w, r, "templates/index.html")
}

// blockHandler handles requests to view block details
func (exp *Explorer) blockHandler(w http.ResponseWriter, r *http.Request) {
	// Extract block hash from URL
	vars := mux.Vars(r)
	blockHash := vars["blockHash"]

	// Retrieve block details from the blockchain
	block := exp.Blockchain.GetBlockByHash(blockHash)
	if block == nil {
		http.Error(w, "Block not found", http.StatusNotFound)
		return
	}

	// Render the block details page
	// For now, just display the block hash
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Block Hash: %s\n", block.Hash)
}
