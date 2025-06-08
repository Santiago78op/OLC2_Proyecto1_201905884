grammar VLangParser;

// Reglas Lexicas

// Palabras clave
MUT      : 'mut';
PRINT    : 'print';
PRINTLN  : 'println';
NIL      : 'nil';

// Tipos Primitivos
INT_TYPE    : 'int';
FLOAT_TYPE  : 'float64';
STRING_TYPE : 'string';
BOOL_TYPE   : 'bool';
SLICE_TYPE  : 'slice';

// Literales
TRUE     : 'true';
FALSE    : 'false';

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
COMMA    : ',';


// Identificadores y Literales
fragment DIGIT : [0-9];
fragment LETTER : [a-zA-Z];
fragment UNDERSCORE : '_';

INT_LITERAL    : DIGIT+;
FLOAT_LITERAL  : DIGIT+ '.' DIGIT+;
STRING_LITERAL : '"' (~["\\\r\n] | EscapeSequence)* '"';
ID             : (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;

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
stmt: declaration delim
    ;

// Declaraciones
declaration: 
    variable_declaration
    | assignment_declaration
    ;

// Declaracion de variable
variable_declaration
    : MUT ID type_annotation (ASSIGN expression)?     # MutableVarDecl
    | ID type_annotation ASSIGN expression            # TypedVarDecl
    ;

// Tipos de datos
type_annotation
    : (INT_TYPE | FLOAT_TYPE | STRING_TYPE | BOOL_TYPE | SLICE_TYPE | ID)
    ;

// Asignacion de Declaracion
assignment_declaration
    : ID ASSIGN expression                       # AssignmentDecl
    | ID PLUS_ASSIGN expression                  # PlusAssignmentDecl
    | ID MINUS_ASSIGN expression                 # MinusAssignmentDecl 
    ;

// Expresiones
expression
    : primary_expression                          # PrimaryExpr
    | expression LPAREN expression RPAREN         # ParenthesizedExpr
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

// Expresiones primarias
primary_expression
    : ID                                          # IdentifierExpr
    | INT_LITERAL                                 # IntLiteral
    | FLOAT_LITERAL                               # FloatLiteral
    | STRING_LITERAL                              # StringLiteral
    | TRUE                                        # TrueLiteral
    | FALSE                                       # FalseLiteral
    | NIL                                         # NilLiteral
    | builtin_function_call                       # BuiltinCall
    | LPAREN expression RPAREN                    # ParenExpr
    ;

// Llamadas a funciones integradas
builtin_function_call
    : PRINT LPAREN   argument_list? RPAREN          # PrintCall
    | PRINTLN LPAREN argument_list? RPAREN        # PrintlnCall
    ;

// Lista de argumentos
argument_list
    : expression (COMMA expression)*              # ArgList
    ;