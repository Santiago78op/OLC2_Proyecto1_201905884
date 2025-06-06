grammar VLangGrammar;

// Reglas Lexicas

// Palabras clave
MUT      : 'mut';
FUNC     : 'func';
STRUCT   : 'struct';
IF       : 'if';
ELSE     : 'else';
FOR      : 'for';
SWITCH   : 'switch';
CASE     : 'case';
DEFAULT  : 'default';
BREAK    : 'break';
CONTINUE : 'continue';
RETURN   : 'return';
IN       : 'in';
NIL      : 'nil';

// Tipos Primitivos
INT_TYPE    : 'int';
FLOAT_TYPE  : 'float64';
STRING_TYPE : 'string';
BOOL_TYPE   : 'bool';

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
LBRACE   : '{';
RBRACE   : '}';
LBRACKET : '[';
RBRACKET : ']';
SEMICOLON: ';';
COMMA    : ',';
DOT      : '.';
COLON    : ':';
AMPERSAND: '&';

// Identificadores y Literales
fragment DIGIT : [0-9];
fragment LETTER : [a-zA-Z];
fragment UNDERSCORE : '_';

INT      : DIGIT+;
FLOAT    : DIGIT+ '.' DIGIT+;
STRING   : '"' (~["\\\r\n] | EscapeSequence)* '"';
RUNE     : '\'' (~['\\\r\n] | EscapeSequence) '\'';
ID       : (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;

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

// Reglas de la gramatica

// Program entry point
program: (declaration | function_declaration | struct_declaration)* function_declaration* EOF;

// Declarations
declaration
    : variable_declaration
    | assignment_statement
    ;

variable_declaration
    : MUT IDENTIFIER type_annotation? (ASSIGN expression)?     # MutableVarDecl
    | IDENTIFIER type_annotation ASSIGN expression            # TypedVarDecl
    | IDENTIFIER DECLARE_ASSIGN expression                    # InferredVarDecl
    ;

type_annotation: INT_TYPE | FLOAT_TYPE | STRING_TYPE | BOOL_TYPE | RUNE_TYPE | slice_type | IDENTIFIER;

slice_type: LBRACKET RBRACKET type_annotation;

// Function declarations
function_declaration
    : FUNC IDENTIFIER LPAREN parameter_list? RPAREN type_annotation? block_statement
    | FUNC LPAREN receiver RPAREN IDENTIFIER LPAREN parameter_list? RPAREN type_annotation? block_statement
    ;

receiver: IDENTIFIER type_annotation;

parameter_list: parameter (COMMA parameter)*;
parameter: IDENTIFIER type_annotation;

// Struct declarations
struct_declaration: STRUCT IDENTIFIER LBRACE struct_field+ RBRACE;
struct_field: type_annotation IDENTIFIER SEMICOLON?;

// Statements
statement
    : declaration
    | assignment_statement
    | if_statement
    | for_statement
    | switch_statement
    | break_statement
    | continue_statement
    | return_statement
    | expression_statement
    | block_statement
    ;

assignment_statement
    : IDENTIFIER ASSIGN expression                           # SimpleAssignment
    | IDENTIFIER PLUS_ASSIGN expression                      # PlusAssignment
    | IDENTIFIER MINUS_ASSIGN expression                     # MinusAssignment
    | IDENTIFIER LBRACKET expression RBRACKET ASSIGN expression  # IndexAssignment
    | IDENTIFIER DOT IDENTIFIER ASSIGN expression            # FieldAssignment
    ;

if_statement: IF expression block_statement (ELSE if_statement | ELSE block_statement)?;

for_statement
    : FOR expression block_statement                                    # ForCondition
    | FOR IDENTIFIER DECLARE_ASSIGN expression SEMICOLON expression SEMICOLON assignment_statement block_statement  # ForLoop
    | FOR IDENTIFIER (COMMA IDENTIFIER)? IN expression block_statement  # ForIn
    ;

switch_statement: SWITCH expression LBRACE case_clause* default_clause? RBRACE;
case_clause: CASE expression COLON statement*;
default_clause: DEFAULT COLON statement*;

break_statement: BREAK;
continue_statement: CONTINUE;
return_statement: RETURN expression?;
expression_statement: expression;
block_statement: LBRACE statement* RBRACE;

// Expressions
expression
    : primary_expression                                      # Primary
    | expression DOT IDENTIFIER                              # FieldAccess
    | expression LBRACKET expression RBRACKET                # IndexAccess
    | expression LPAREN argument_list? RPAREN                # FunctionCall
    | MINUS expression                                        # UnaryMinus
    | NOT expression                                          # UnaryNot
    | expression op=(MULTIPLY | DIVIDE | MODULO) expression  # MulDivMod
    | expression op=(PLUS | MINUS) expression                # AddSub
    | expression op=(LESS_THAN | LESS_EQUAL | GREATER_THAN | GREATER_EQUAL) expression  # Relational
    | expression op=(EQUAL | NOT_EQUAL) expression           # Equality
    | expression AND expression                               # LogicalAnd
    | expression OR expression                                # LogicalOr
    ;

primary_expression
    : IDENTIFIER                                              # IdentifierExpr
    | INT_LITERAL                                             # IntLiteral
    | FLOAT_LITERAL                                           # FloatLiteral
    | STRING_LITERAL                                          # StringLiteral
    | RUNE_LITERAL                                            # RuneLiteral
    | TRUE                                                    # TrueLiteral
    | FALSE                                                   # FalseLiteral
    | NIL                                                     # NilLiteral
    | slice_literal                                           # SliceLiteralExpr
    | struct_literal                                          # StructLiteralExpr
    | builtin_function_call                                   # BuiltinCall
    | LPAREN expression RPAREN                                # ParenExpr
    ;

slice_literal: LBRACKET RBRACKET type_annotation LBRACE expression_list? RBRACE;
struct_literal: IDENTIFIER LBRACE field_assignment_list? RBRACE;
field_assignment_list: field_assignment (COMMA field_assignment)*;
field_assignment: IDENTIFIER COLON expression;

builtin_function_call
    : PRINT LPAREN argument_list? RPAREN                     # PrintCall
    | ATOI LPAREN expression RPAREN                          # AtoiCall
    | PARSE_FLOAT LPAREN expression RPAREN                   # ParseFloatCall
    | TYPE_OF LPAREN expression RPAREN                       # TypeOfCall
    | INDEX_OF LPAREN expression COMMA expression RPAREN     # IndexOfCall
    | JOIN LPAREN expression COMMA expression RPAREN         # JoinCall
    | LEN LPAREN expression RPAREN                           # LenCall
    | APPEND LPAREN expression COMMA expression RPAREN       # AppendCall
    ;

argument_list: expression (COMMA expression)*;
expression_list: expression (COMMA expression)*;