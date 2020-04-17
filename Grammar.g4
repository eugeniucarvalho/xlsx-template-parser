/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
 /**
 TODO:
 SwitchCASE incompleto
 */

grammar Grammar; 

expression
    : R_BRACE_L R_BRACE_L actionTypes R_BRACE_R R_BRACE_R  EOF
    ;

actionTypes
    : command? path 
    | R_END 
    | R_STOP
    ;

command 
    : 'range'
    | 'if'
    | 'with'
    | 'block'
    ;

path
    : R_AND? R_DOT (ID (R_DOT ID)*)?
    ;

R_AND                   : '&';
R_BRACE_L               : '{';
R_BRACE_R               : '}';
R_PIPE                  : '|';
R_DOT                   : '.';
R_END                   : 'end';
R_STOP                  : 'stop';

fragment    DIGITO      : [0-9] ;
fragment    CHARACTER   : [a-zA-Z_] ;
fragment    CHARACTER_UP: [A-Z];

ID                      : CHARACTER ( CHARACTER | DIGITO )* ;
T_INTEIRO               : DIGITO+ ;

R_LINE_COMMENT          : '//' .*? '\r'? '\n' -> skip ;  // Match "#" stuff '\n'
R_WS                    : (' '|'\t'|'\r'|'\n')+ -> skip ;
R_COMMENT               : '/*' .*? '*/' -> skip ; // Match "/*" stuff "*/"