parser grammar VLangGrammar;

// Importar las reglas lexicas
options { tokenVocab = VLangLexer; }

// Program entry point
program: (stmt)* EOF?
    ;

// Sentencias
stmt: 
    decl_stmt 
    | assign_stmt 
    | block_ind 
    | transfer_stmt  
    | if_stmt
	| switch_stmt
    | while_stmt
	| for_stmt
    | func_call 
    | vect_func 
    | func_dcl
    | strct_dcl
    ;

// Inicia Declaracion de variable
// Ejemplo: Mut variable_1 int = 10
// Ejemplo: Mut variable_2 int 
decl_stmt: 
    var_type ID type ASSIGN expression  # MutVarDecl
    | var_type ID ASSIGN expression     # ValueDecl 
    | ID type ASSIGN expression         # VarAssDecl  
    | ID ASSIGN type expression         # VarVectDecl  // vector []int = {1,2,3,4}
    | var_type ID vector_type           # MutSliceDecl // mut slice []int
    ;

// Inicia Declaracion de Vector
// {1,2,3,4}
vect_expr:
    LBRACE ( expression (COMMA expression)* )? RBRACE  # VectorItemLis
    ;

// vector_1[0]
vect_item:
    id_pattern (LBRACK expression RBRACK)+   # VectorItem
    ;

// llamada a un vector por medio de una propiedad
// ejemplo: vector_1.id
vect_prop:
    vect_item DOT id_pattern   # VectorProperty
    ;

vect_func:
    vect_item DOT func_call   # VectorFuncCall 
    ;

repeating:
    (vector_type | matrix_type) LPAREN ID COLON expression COMMA ID COLON expression RPAREN  # RepeatingDecl
    ;
// Termina Declaracion Vectores

// tipo de variable
var_type
    : MUT
    ;

// Inicia Declaracion de Vectores
// [] int, [] float, [] String, [] bool
vector_type: LBRACK RBRACK ID
    ;

type: 
    ID 
    | vector_type 
    | matrix_type
    ;

matrix_type: 
    aux_matrix_type 
    | LBRACK LBRACK ID RBRACK RBRACK
    ;

aux_matrix_type: LBRACK matrix_type RBRACK
    ;

// Termina Declaracion de Variables

// Inicia Asignacion de Variables
// varSring = "Hola"
assign_stmt:
    id_pattern ASSIGN expression                  # AssignmentDecl
    | id_pattern op = (
        PLUS_ASSIGN | MINUS_ASSIGN
    ) expression                                  # ArgAddAssigDecl
    | vect_item op = ( 
        PLUS_ASSIGN 
        | MINUS_ASSIGN 
        | ASSIGN) expression	                  # VectorAssign
    ;

// variable ASSIGN expression
// num2 
id_pattern // (a.a)
    : ID (DOT ID)*                                # IdPattern
    ;
// Termina Asignacion de Variables

// Literales
// Ejemplo 3   = INT_LITERAL
// Ejemplo 4.5 = FLOAT_LITERAL
literal
    : INT_LITERAL                                 # IntLiteral
    | FLOAT_LITERAL                               # FloatLiteral
    | STRING_LITERAL                              # StringLiteral
    | BOOL_LITERAL                                # BoolLiteral
    | NIL_LITERAL                                 # NilLiteral
    ;

// Incremento y decremento
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
    ;

// Inicio Expresiones
expression
    : LPAREN expression RPAREN                       # ParensExpr 
    | func_call                                      # FuncCallExpr 
    | id_pattern                                     # IdPatternExpr
    | vect_item                                      # VectorItemExpr
    | vect_prop                                      # VectorPropertyExpr
    | vect_func                                      # VectorFuncCallExpr
    | literal                                        # LiteralExpr
    | vect_expr                                      # VectorExpr
    | repeating                                      # RepeatingExpr
    | incredecre                                     # incredecr
    | op = ( NOT | MINUS) expression                 # UnaryExpr
    | left = expression op = (
        MULT | DIV | MOD
    ) right = expression                             # BinaryExpr
    | left = expression op = (
        PLUS | MINUS
    ) right = expression                             # BinaryExpr
    | left = expression op = (
        LE | LT | GE | GT 
    ) right = expression                             # BinaryExpr
    | left = expression op = (
        EQ | NE
    ) right = expression                             # BinaryExpr
    | left = expression op = AND right = expression  # BinaryExpr
    | left = expression op = OR right = expression   # BinaryExpr
    ;
// Terminan Expresiones

// Inicia Sentencias de Control If
if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expression LBRACE stmt* RBRACE   # IfChain;

else_stmt: ELSE_KW LBRACE stmt* RBRACE           # ElseStmt;
// Termina Sentencias de Control If

// Inicia Sentencias de Control Switch
switch_stmt:
	SWITCH_KW expression LBRACE switch_case* default_case? RBRACE # SwitchStmt;

switch_case: CASE_KW expression COLON stmt* # SwitchCase;

default_case: DEFAULT_KW COLON stmt* # DefaultCase;
// Termina Sentencias de Control Switch

// Inicia Sentencias de Control While
while_stmt: WHILE_KW expression LBRACE stmt* RBRACE # WhileStmt;
// Termina Sentencia de Control While

// Inicia Sentencias de Iteracion For
for_stmt:
    FOR_KW expression LBRACE stmt* RBRACE                                        # ForStmtCond
    | FOR_KW assign_stmt SEMI expression SEMI expression LBRACE stmt* RBRACE     # ForAssCond

	| FOR_KW ID COMMA expression IN_KW (expression | range) LBRACE stmt* RBRACE  # ForStmt ;

range: expression DOT DOT DOT expression # NumericRange;
// Termina Sentencias de Iteracion For

// Inicia Sentencias de Transferencia
transfer_stmt:
	RETURN_KW expression?	# ReturnStmt
	| BREAK_KW		        # BreakStmt
	| CONTINUE_KW	        # ContinueStmt;
// Termina Sentencias de Transferencia

// Inicia Llamadas a funcion 
// Ejemplo Println(5)  
func_call: id_pattern LPAREN arg_list? RPAREN # FuncCall;
// Termina Llamadas a funcion

// Inicia Bloques de codigo
block_ind: LBRACE stmt* RBRACE # BlockInd;
// Termina Bloques de codigo

// external names -> variable int, num2 float, 5
arg_list: func_arg (COMMA func_arg)* # ArgList;

// 5
func_arg: (ID)? (id_pattern | expression) # FuncArg; // 

func_dcl:
	FUNC ID LPAREN param_list? RPAREN (type)? LBRACE stmt* RBRACE # FuncDecl;

param_list: func_param (COMMA func_param)* # ParamList;
func_param: ID type                        # FuncParam;

// Inicia Estructuras de control
strct_dcl: STR ID LBRACE struct_prop* RBRACE # StructDecl;

struct_prop:
	var_type ID (COLON type)? (ASSIGN expression)?	# StructAttr
    ;

struct_vector: LBRACK ID RBRACK LPAREN RPAREN       # StructVector
    ;
// Termina Estructuras de control