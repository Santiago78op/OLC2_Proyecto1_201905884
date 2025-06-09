grammar VLangParser;

// Reglas Lexicas

// Palabras clave
MUT   : 'mut';
FUNC  : 'func';

// Estructuras
STR         : 'struct';
SLICE_TYPE  : 'slice';

// Control de flujo - wk => keyWord
IF_KW       : 'if';
ELSE_KW     : 'else';
SWITCH_KW   : 'switch';
CASE_KW     : 'case';
DEFAULT_KW  : 'default';
FOR_KW      : 'for';
IN_KW       : 'in';
BREAK_KW    : 'break';
CONTINUE_KW : 'continue';
RETURN_KW   : 'return';


INT_TYPE    : 'int';
FLOAT_TYPE  : 'float64';
STRING_TYPE : 'string';
BOOL_TYPE   : 'bool';
RUNE_TYPE   : 'rune';


// Operadores Aritmeticos
PLUS     : '+';
MINUS    : '-';
MULT     : '*';
DIV      : '/';
MOD      : '%';

// Operadores de Asignacion
ASSIGN      : '=';
PLUS_ASSIGN : '+=';
MINUS_ASSIGN: '-=';

// Operadores de Comparacion
EQ       : '==';
NE       : '!=';
LT       : '<';
LE       : '<=';
GT       : '>';
GE       : '>=';

// Operadores Logicos
AND      : '&&';
OR       : '||';
NOT      : '!';

// Delimitadores
LPAREN   : '(';
RPAREN   : ')';
LBRACE   : '{';
RBRACE   : '}';
LBRACK   : '[';
RBRACK   : ']';
SEMI     : ';';
COLON    : ':';
DOT      : '.';
COMMA    : ',';


// Literales
fragment DIGIT : [0-9];
fragment LETTER : [a-zA-Z];
fragment UNDERSCORE : '_';
fragment CHAR : [!-~];

INT_LITERAL    : DIGIT+;
FLOAT_LITERAL  : DIGIT+ '.' DIGIT+;
CHAR_LITERAL   : CHAR; 
STRING_LITERAL : '"' (~["\\\r\n] | EscapeSequence)* '"';
BOOL_LITERAL   : 'true' | 'false';
NIL_LITERAL    : 'nil';

// Identificador
ID : (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;

// Secuencia de escape
fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' 'n'
    | '\\' 'r'
    | '\\' 't'
    ;

// Commentarios
LINE_COMMENT  : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT : '/*' .*? '*/' -> skip;

// Espacios en blanco
WS : [ \t\r\n]+ -> skip;

// Program entry point
prog: (stmt)* EOF?;

// Demilitador del lenguaje -> indica el final de una sentencia
delim: '\n';

// Sentencias
stmt: 
    stmt_declaration delim
    | stmt_assign delim
    | stmt_transfer delim 
    | stmt_if
    | stmt_switch
    | stmt_while
    | stmt_for
    | func_call delim
    | vect_func delim
    | func_dcl
    | struct_dcl
    ;

// Declaracion de Vector
vect_expr:
    LBRACK ( expression (COMMA expression)* )? RBRACK  # VectorExpr
    ;

vect_item:
    id_pattern (LBRACE expression RBRACE)+   # VectorItem
    ;

vect_prop:
    vect_item DOT id_pattern   # VectorProperty
    ;

vect_func:
    vect_item DOT func_call # VectorFuncCall 
    ;

repeating:
    (var_type | matrix_type) LPAREN ID COLON expression COMMA ID COLON expression RPAREN  # RepeatingDecl
    ;

// tipo de variable
var_type
    : MUT
    ;

type: ID | vector_type | matrix_type;

vector_type: LBRACE ID RBRACK;

matrix_type: aux_matrix_type | LBRACK LBRACK ID RBRACK RBRACK;

aux_matrix_type: LBRACK matrix_type RBRACK;

// Declaracion de variable
stmt_declaration: 
    var_type ID type_annotation (ASSIGN expression)?   # MutableVarDecl
    | ID type_annotation ASSIGN expression             # TypeVarDecl
    ;


// Tipos de datos
type_annotation
    : (INT_TYPE | FLOAT_TYPE | STRING_TYPE | BOOL_TYPE | SLICE_TYPE | ID)
    ;

// Asignacion de Declaracion
assignment_declaration:
    id_pattern ASSIGN expression                  # AssignmentDecl
    | id_pattern op = (
        PLUS_ASSIGN | MINUS_ASSIGN
    ) expression                                  # AugmentedAssignmentDecl
    ;

id_pattern
    : ID (DOT ID)*                                # IdPattern
    ;

// Expresiones primarias
literal
    : INT_LITERAL                                 # IntLiteral
    | FLOAT_LITERAL                               # FloatLiteral
    | STRING_LITERAL                              # StringLiteral
    | BOOL_LITERAL                                # BoolLiteral
    | NIL_LITERAL                                 # NilLiteral
    ;

// Expresiones
expression
    : LPAREN expression RPAREN                    # ParensExpr
    | func_call                                   # FuncCallExpr
    | id_pattern                                  # IdPatternExpr
    | vect_item                                   # VectorItemExpr
    | vect_prop                                   # VectorPropertyExpr
    | vect_func                                   # VectorFuncCallExpr
    | literal                                     # LiteralExpr
    | vect_expr                                   # VectorExpr
    | repeating                                   # RepeatingExpr
    | struct_vect 
    | op = ( NOT | MINUS) expression              # UnaryExpr
    | left = expression op = (
        MULT | DIV | MOD
    ) right = expression                          # BinaryExpr
    | left = expression op = (
        PLUS | MINUS
    ) right = expression                          # BinaryExpr
    | left = expression op = (
        LE | LT | GE | GT 
    ) right = expression                          # BinaryExpr
    | left = expression op = (
        EQ | NE
    ) right = expression                             # BinaryExpr
    | left = expression op = AND right = expression  # BinaryExpr
    | left = expression op = OR right = expression   # BinaryExpr
    ;

if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expression LBRACE stmt* RBRACE # IfChain;
else_stmt: ELSE_KW LBRACE stmt* RBRACE # ElseStmt;

switch_stmt:
	SWITCH_KW expression LBRACE switch_case* default_case? RBRACE # SwitchStmt;

switch_case: CASE_KW expression COLON stmt* # SwitchCase;

default_case: DEFAULT_KW COLON stmt* # DefaultCase;

while_stmt: WHILE_KW expression LBRACE stmt* RBRACE # WhileStmt;

for_stmt:
	FOR_KW ID IN_KW (expression | range) LBRACE stmt* RBRACE # ForStmt;

range: expression DOT DOT DOT expression # NumericRange;

// Lista de argumentos
argument_list
    : expression (COMMA expression)*              # ArgList
    ;

transfer_stmt:
	RETURN_KW expr?	# ReturnStmt
	| BREAK_KW		# BreakStmt
	| CONTINUE_KW	# ContinueStmt;

func_call: id_pattern LPAREN arg_list? RPAREN # FuncCall;

// external names -> num: value, num2: value2
arg_list: func_arg (COMMA func_arg)* # ArgList;
func_arg: (ID COLON)? (ANPERSAND)? (id_pattern | expr) # FuncArg; // 

func_dcl:
	FUNC_KW ID LPAREN param_list? RPAREN (ARROW type)? LBRACE stmt* RBRACE # FuncDecl;

param_list: func_param (COMMA func_param)* # ParamList;
func_param: ID? ID COLON INOUT_KW? type # FuncParam;

// Estructuras de control
strct_dcl: STR ID LBRACE struct_prop* RBRACE     # StructDecl;

struct_prop:
	var_type ID (COLON type)? (ASSIGN expression)?	# StructAttr
    ;

struct_vect: LBRACK ID RBRACK LPAREN RPAREN # StructVector;