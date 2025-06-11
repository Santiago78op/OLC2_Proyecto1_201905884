package main

import (
	"encoding/json"
	"fmt"
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
	// Receive code from the request body
	w.Header().Set("Content-Type", "application/json")

	var requestData struct {
		Code *string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Desreferenciar el puntero para obtener el string
	if requestData.Code == nil {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}

	// Desreferenciar el puntero para obtener el string
	codeString := *requestData.Code
	fmt.Println("Received code for execution:")
	fmt.Println(codeString)

	//Pasar a string requestData.Code

	// 1. Analisis Lexico
	lexicalErrorListener := errors.NewLexicalErrorListener()
	lexer := compiler.NewVLangLexer(antlr.NewInputStream("print(\"hola\")"))

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)
	// 2. Tokens
	// New<Nombre de mi gramatica>(Stream) < id, valor>  token
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 3. Analisis Sintactico Parser + errores sintácticos
	parser := compiler.NewVLangGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	// 4. Árbol sintáctico
	// En tu gramatica tienes el axioma, o simbolo inicial
	// Este es el que deberas agregar como parte del parser.
	tree := parser.Program() // Aquí se debe llamar al método adecuado según tu gramática

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

	intepretationEndTime := time.Now()

	startTime := time.Now()

	reportEndTime := time.Now()

	fmt.Println("Interpretation finished")

	fmt.Println("Interpretation time:", intepretationEndTime.Sub(startTime))
	fmt.Println("Total (with report) time:", reportEndTime.Sub(intepretationEndTime))
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
