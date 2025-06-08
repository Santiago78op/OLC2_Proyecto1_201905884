lexer grammar VLangLexer;

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