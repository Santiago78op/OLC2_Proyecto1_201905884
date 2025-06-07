// El nombre de la gramática debe ser el mismo que el archivo
grammar Calculator;

// Gramatical
/*
    El primer simbolo que se determine como el
    simbolo inicial, se convertirá en un mentodo
    al momento de invocar el parser.
 */
prog: statement+ EOF;

statement: expr NEWLINE
         | ID '=' expr NEWLINE
         | NEWLINE;

// Analizador Top-Down -> El elimina la recursión izquierda
expr:  expr ('*'|'/') expr  #MulDiv
     | expr ('+'|'-') expr  #AddSub
     | INT                  #entero
     | ID                   #identificador
     | '(' expr ')'         #Parentesis
     ;                
     
// Lexemas
ID:      [a-zA-Z]+;
INT:     [0-9]+;
NEWLINE: [\r\n]+;
WS:      [ \t]+ -> skip; // Espacios en blanco