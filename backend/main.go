package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	// Importa el paquete de pruebas que contiene la lógica de ejecución

	"main.go/errors"
	compiler "main.go/grammar"
	"main.go/repl"
)

type executionResult struct {
	Success bool         `json:"success"`
	Errors  []repl.Error `json:"errors"`
	Output  string       `json:"output"`
	CSTSvg  string       `json:"cstSvg"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 🔍 DEBUG: Verificar Content-Type
	contentType := r.Header.Get("Content-Type")
	fmt.Printf("🔹 Content-Type recibido: %s\n", contentType)

	// 🔍 DEBUG: Leer el body completo para debugging
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("❌ Error leyendo body: %v\n", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("🔹 Body raw recibido: %s\n", string(bodyBytes))
	fmt.Printf("🔹 Longitud del body: %d bytes\n", len(bodyBytes))

	// Recrear el reader para el JSON decoder
	// (necesario porque ya leímos el body)
	if len(bodyBytes) == 0 {
		fmt.Println("❌ Body está vacío")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	// Cambiar la estructura para usar string en lugar de *string
	var requestData struct {
		Code string `json:"code"`
	}

	// Decodificar desde los bytes leídos
	if err := json.Unmarshal(bodyBytes, &requestData); err != nil {
		fmt.Printf("❌ Error decodificando JSON: %v\n", err)
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Verificar si el código está presente
	if requestData.Code == "" {
		fmt.Println("❌ Campo 'code' está vacío")
		http.Error(w, "Code field is required and cannot be empty", http.StatusBadRequest)
		return
	}

	// Convertir el código a string
	codeString := string(requestData.Code)

	fmt.Printf("✅ Código recibido exitosamente:\n%s\n", codeString)
	fmt.Printf("🔹 Longitud del código: %d caracteres\n", len(codeString))

	// 1. Análisis Léxico
	lexicalErrorListener := errors.NewLexicalErrorListener()
	lexer := compiler.NewVLangLexer(antlr.NewInputStream("print(10"))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// 2. Tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 3. Análisis Sintáctico Parser + errores sintácticos
	parser := compiler.NewVLangGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	fmt.Printf("Errores léxicos: %v\n", lexicalErrorListener.ErrorTable.Errors)
	fmt.Printf("Errores sintácticos: %v\n", syntaxErrorListener.ErrorTable.Errors)

	// 4. Árbol sintáctico
	tree := parser.Program()

	// ✅ VERIFICACIONES CRÍTICAS
	fmt.Printf("🔹 Tree creado: %T\n", tree)
	fmt.Printf("🔹 Tree es nil: %v\n", tree == nil)
	fmt.Printf("🔹 Tree text: %s\n", tree.GetText())
	fmt.Printf("🔹 Errores en ErrorTable: %d\n", len(syntaxErrorListener.ErrorTable.Errors))

	dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
	dclVisitor.Visit(tree)

	replVisitor := repl.NewVisitor(dclVisitor)
	replVisitor.Visit(tree)

	cstReport := ""

	interpretationEndTime := time.Now()
	startTime := time.Now()
	reportEndTime := time.Now()

	fmt.Println("Interpretation finished")
	fmt.Println("Interpretation time:", interpretationEndTime.Sub(startTime))
	fmt.Println("Total (with report) time:", reportEndTime.Sub(interpretationEndTime))
	fmt.Println("")

	// Imprimir Output
	fmt.Println("Output:", replVisitor.Console.GetOutput())

	executionResult := executionResult{
		Success: true,
		Output:  replVisitor.Console.GetOutput(),
		CSTSvg:  cstReport,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(executionResult)
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
