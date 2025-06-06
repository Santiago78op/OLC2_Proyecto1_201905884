package repl

type ReplContext struct {
	// The console is the output of the REPL
	Console *Console
	// The scope is the current scope of the REPL
	ScopeTrace *ScopeTrace
	// Error table is the table of errors
	ErrorTable *ErrorTable
}

func NewReplContext() *ReplContext {
	return &ReplContext{
		Console:    NewConsole(),
		ScopeTrace: NewScopeTrace(),
		ErrorTable: NewErrorTable(),
	}
}
