package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main.go/lexicalErrors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type executionResult struct {
	Success  bool   `json:"success"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Error    string `json:"error,omitempty"`
	Duration string `json:"duration"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Receive code from the request body
	var requestData struct {
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 1. Analisis Lexico
	lexicalErrors.FuncionPrueba() // Example function call from the lexicalErrors package
	// 2. Tokens
	// 3. Parser + errores sintácticos
	// New<Nombre de mi gramatica>(Stream)
	// 4. Árbol sintáctico
	// En tu gramatica tienes el axioma, o simbolo inicial
	// Este es el que deberas agregar como parte del parser.
	//tree := parser.Program()

	startTime := time.Now()
	// Simulate code execution (replace with actual execution logic)
	result := executionResult{
		Success:  true,
		Code:     requestData.Code,
		Result:   "Execution result goes here",
		Duration: time.Since(startTime).String(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func main() {
	r := mux.NewRouter()

	// API Routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/status", healthCheck).Methods("GET")
	api.HandleFunc("/execute", executeCode).Methods("POST")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	port := ":8080"
	fmt.Printf("🚀 Servidor Go iniciado en http://localhost%s\n", port)
	fmt.Println("📋 API endpoints disponibles:")
	fmt.Println("  - GET    /api/status")  // Estado del servidor
	fmt.Println("  - POST   /api/execute") // Ejecutar código
	fmt.Println("  - GET    /api/reports") // Obtener reportes

	log.Fatal(http.ListenAndServe(port, handler))
}
