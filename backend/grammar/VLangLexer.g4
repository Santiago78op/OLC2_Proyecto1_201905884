lexer grammar VLangLexer;

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
WHILE_KW    : 'while';
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

NEWLINE : '\n';


// Literales
fragment DIGIT : [0-9];
fragment LETTER : [a-zA-Z];
fragment UNDERSCORE : '_';
fragment CHAR : [!-~];

INT_LITERAL    : DIGIT+;
FLOAT_LITERAL  : DIGIT+ '.' DIGIT+;
CHAR_LITERAL   : CHAR; 
STRING_LITERAL: '"' (~["\r\n\\] | ESC_SEQ)* '"';
BOOL_LITERAL   : 'true' | 'false';
NIL_LITERAL    : 'nil';

// Identificador
ID : (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;

// Secuencia de escape
fragment ESC_SEQ: '\\' [btnfr"'\\]
    ;

// Commentarios
LINE_COMMENT  : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT : '/*' .*? '*/' -> skip;

// Espacios en blanco
WS : [ \t\r\n]+ -> skip;