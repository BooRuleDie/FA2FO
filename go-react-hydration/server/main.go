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
	fs := http.FileServer(http.Dir("../client/dist/assets/"))
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

	// Create template function map
	funcMap := template.FuncMap{
		"css": func(s string) template.CSS {
			return template.CSS(s)
		},
	}

	// Parse template with function map
	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render template with data
	data := map[string]any{
		"Product":     product,
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
			{Name: "Black", Value: "rgb(0,0,0)"},
			{Name: "Brown", Value: "rgb(139,69,19)"},
			{Name: "Tan", Value: "rgb(210,180,140)"},
		},
		Images: []string{
			"https://placehold.co/600x400?text=Leather+Jacket+1",
			"https://placehold.co/600x400?text=Leather+Jacket+2",
			"https://placehold.co/600x400?text=Leather+Jacket+3",
		},
	}
}
