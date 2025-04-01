package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Product represents product data
type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Colors      []Color  `json:"colors"`
	Images      []string `json:"images"`
}

// Color represents a product color option
type Color struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	// Setup file server for static files
	fs := http.FileServer(http.Dir("../frontend/dist/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Setup routes
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/product", handleProductAPI)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Get mock product data
	product := getMockProduct()

	// Convert product data to JSON for hydration
	productJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal product data", http.StatusInternalServerError)
		return
	}

	// Parse template
	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render template with data
	data := map[string]interface{}{
		"Product":    product,
		"ProductJSON": template.JS(productJSON),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleProductAPI(w http.ResponseWriter, r *http.Request) {
	product := getMockProduct()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func getMockProduct() Product {
	return Product{
		ID:          "prod-1",
		Name:        "Premium Leather Jacket",
		Description: "High-quality leather jacket with a modern design. Features multiple pockets and adjustable waist.",
		Price:       299.99,
		Colors: []Color{
			{Name: "Black", Value: "#000000"},
			{Name: "Brown", Value: "#8B4513"},
			{Name: "Tan", Value: "#D2B48C"},
		},
		Images: []string{
			"https://placehold.co/600x400?text=Leather+Jacket+1",
			"https://placehold.co/600x400?text=Leather+Jacket+2",
			"https://placehold.co/600x400?text=Leather+Jacket+3",
		},
	}
}