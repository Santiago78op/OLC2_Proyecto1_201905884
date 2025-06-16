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

	"main.go/ast"
	"main.go/cst"
	"main.go/errors"
	compiler "main.go/grammar"
	"main.go/repl"
)

type executionResult struct {
	Success       bool                `json:"success"`
	Errors        []repl.Error        `json:"errors"`
	Output        string              `json:"output"`
	CSTSvg        string              `json:"cstSvg"`
	AST           string              `json:"ast"`
	Symbols       []repl.ReportSymbol `json:"symbols"`
	ScopeTrace    repl.ReportTable    `json:"scopeTrace"`
	ErrorSummary  map[string]int      `json:"errorSummary"`
	ExecutionTime int64               `json:"executionTime"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Leer y procesar el body (código existente)
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("❌ Error leyendo body: %v\n", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("🔹 Body raw recibido: %s\n", string(bodyBytes))

	if len(bodyBytes) == 0 {
		fmt.Println("❌ Body está vacío")
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var requestData struct {
		Code string `json:"code"`
	}

	if err := json.Unmarshal(bodyBytes, &requestData); err != nil {
		fmt.Printf("❌ Error decodificando JSON: %v\n", err)
		http.Error(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	if requestData.Code == "" {
		fmt.Println("❌ Campo 'code' está vacío")
		http.Error(w, "Code field is required and cannot be empty", http.StatusBadRequest)
		return
	}

	codeString := requestData.Code
	for len(codeString) > 0 && (codeString[0] == '\n' || codeString[0] == '\r') {
		codeString = codeString[1:]
	}

	fmt.Printf("✅ Código recibido exitosamente:\n%s\n", codeString)

	// =========== ANÁLISIS Y EJECUCIÓN ===========
	startTime := time.Now()

	// 1. Generar CST Report en paralelo
	resultChannel := make(chan string, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error generando CST Report:", r)
				resultChannel <- ""
			}
		}()
		resultChannel <- cst.CstReport(codeString)
	}()

	// 2. Análisis Léxico
	lexicalErrorListener := errors.NewLexicalErrorListener()
	lexer := compiler.NewVLangLexer(antlr.NewInputStream(codeString))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// 3. Análisis Sintáctico
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := compiler.NewVLangGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	// 4. Generar AST
	tree := parser.Program()

	// Verificar si hubo errores críticos que impiden continuar
	hasCompilationErrors := len(syntaxErrorListener.ErrorTable.Errors) > 0

	fmt.Printf("🔹 Errores de compilación: %d\n", len(syntaxErrorListener.ErrorTable.Errors))
	if hasCompilationErrors {
		for _, err := range syntaxErrorListener.ErrorTable.Errors {
			fmt.Printf("   - %s (Línea %d, Col %d): %s\n", err.GetDisplayName(), err.Line, err.Column, err.Msg)
		}
	}

	var replVisitor *repl.ReplVisitor
	var output string = ""

	// 5. Solo continuar con análisis semántico si no hay errores críticos de sintaxis
	if !hasCompilationErrors {
		// Análisis Semántico y Ejecución
		dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
		dclVisitor.Visit(tree)

		replVisitor = repl.NewVisitor(dclVisitor)
		replVisitor.Visit(tree)
		output = replVisitor.Console.GetOutput()
	} else {
		// Si hay errores de compilación, no ejecutar pero crear visitor básico para reportes
		dclVisitor := repl.NewDclVisitor(syntaxErrorListener.ErrorTable)
		replVisitor = repl.NewVisitor(dclVisitor)
		output = "" // No hay salida si no se pudo compilar
	}

	interpretationEndTime := time.Now()

	// 6. Obtener CST Report
	cstReport := <-resultChannel

	// Si no hay CST report, intentar generar uno básico
	if cstReport == "" {
		cstReport = generateBasicAST(tree)
	}

	reportEndTime := time.Now()

	// =========== GENERAR REPORTES ===========

	// Determinar si la ejecución fue exitosa
	success := !hasCompilationErrors && len(syntaxErrorListener.ErrorTable.Errors) == 0

	// Generar tabla de símbolos
	scopeReport := replVisitor.ScopeTrace.Report()
	symbols := extractSymbolsFromScope(scopeReport)

	// Crear resumen de errores
	errorSummary := syntaxErrorListener.ErrorTable.GetErrorsSummary()

	fmt.Printf("🔹 Resumen de errores: %+v\n", errorSummary)
	fmt.Printf("🔹 Tiempo de interpretación: %v\n", interpretationEndTime.Sub(startTime))
	fmt.Printf("🔹 Tiempo total: %v\n", reportEndTime.Sub(startTime))
	fmt.Printf("🔹 Salida: %s\n", output)

	// Generar AST nativo
	var finalAST string
	if tree != nil {
		fmt.Println("🌳 Generando AST nativo...")
		astNode := ast.GenerateNativeAST(tree)
		finalAST = ast.GenerateASTSVG(astNode)
		fmt.Println("✅ AST nativo generado exitosamente")
	} else {
		fmt.Println("❌ No se pudo generar el árbol de análisis")
		finalAST = `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="200">
			<rect width="400" height="200" fill="#1e1e1e"/>
			<text x="200" y="100" text-anchor="middle" fill="#ffffff">Error en análisis sintáctico</text>
		</svg>`
	}

	// Crear resultado con información detallada
	result := executionResult{
		Success:       success,
		Errors:        syntaxErrorListener.ErrorTable.Errors,
		Output:        output,
		CSTSvg:        finalAST,
		AST:           finalAST,
		Symbols:       symbols,
		ScopeTrace:    scopeReport,
		ErrorSummary:  errorSummary,                                        // Agregar resumen de errores
		ExecutionTime: interpretationEndTime.Sub(startTime).Milliseconds(), // Tiempo en ms
	}

	// Enviar respuesta
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("❌ Error encoding response: %v\n", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("✅ Respuesta enviada exitosamente\n")
}

// Función auxiliar para extraer símbolos del scope report
func extractSymbolsFromScope(scopeReport repl.ReportTable) []repl.ReportSymbol {
	var allSymbols []repl.ReportSymbol

	// Función recursiva para extraer símbolos de todos los scopes
	var extractFromScope func(scope repl.ReportScope, scopeName string)
	extractFromScope = func(scope repl.ReportScope, scopeName string) {
		// Agregar variables
		for _, symbol := range scope.Vars {
			symbol.Scope = scopeName // Asegurar que tenga el scope correcto
			allSymbols = append(allSymbols, symbol)
		}

		// Agregar funciones
		for _, symbol := range scope.Funcs {
			symbol.Scope = scopeName
			allSymbols = append(allSymbols, symbol)
		}

		// Agregar estructuras
		for _, symbol := range scope.Structs {
			symbol.Scope = scopeName
			allSymbols = append(allSymbols, symbol)
		}

		// Procesar scopes hijos recursivamente
		for _, childScope := range scope.ChildScopes {
			childScopeName := scopeName + "." + childScope.Name
			extractFromScope(childScope, childScopeName)
		}
	}

	extractFromScope(scopeReport.GlobalScope, "global")
	return allSymbols
}

// Función auxiliar para generar AST básico si falla el CST report
func generateBasicAST(tree antlr.ParseTree) string {
	if tree == nil {
		return `<svg xmlns="http://www.w3.org/2000/svg" width="400" height="200">
            <rect width="400" height="200" fill="#1e1e1e"/>
            <text x="200" y="100" text-anchor="middle" fill="#ffffff" font-family="Arial" font-size="16">
                No se pudo generar el AST
            </text>
        </svg>`
	}

	// Generar un SVG básico con información del árbol
	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="600" height="400">
        <rect width="600" height="400" fill="#1e1e1e"/>
        <circle cx="300" cy="100" r="40" fill="#007acc" stroke="#ffffff" stroke-width="2"/>
        <text x="300" y="105" text-anchor="middle" fill="#ffffff" font-family="Arial" font-size="12">Program</text>
        <text x="300" y="200" text-anchor="middle" fill="#cccccc" font-family="Arial" font-size="14">
            AST generado exitosamente
        </text>
        <text x="300" y="220" text-anchor="middle" fill="#cccccc" font-family="Arial" font-size="12">
            Texto del árbol: %s
        </text>
    </svg>`, tree.GetText()[:min(50, len(tree.GetText()))])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
