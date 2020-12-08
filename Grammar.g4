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
    : R_AND? R_DOT (ID (R_DOT ID)*)? R_SPREAD?
    ;
// spread
//     : R_DOT R_DOT R_DOT
//     ;

R_AND                   : '&';
R_BRACE_L               : '{';
R_BRACE_R               : '}';
R_PIPE                  : '|';
R_SPREAD                : '>>>';
R_DOT                   : '.';
R_END                   : 'end';
R_STOP                  : 'stop';
R_DOLAR                 : '$';

fragment    DIGITO      : [0-9] ;
fragment    CHARACTER   : [a-zA-Z_] ;
fragment    CHARACTER_UP: [A-Z];

ID                      : R_DOLAR? CHARACTER ( CHARACTER | DIGITO )* ;

R_LINE_COMMENT          : '//' .*? '\r'? '\n' -> skip ;  // Match "#" stuff '\n'
R_WS                    : (' '|'\t'|'\r'|'\n')+ -> skip ;
R_COMMENT               : '/*' .*? '*/' -> skip ; // Match "/*" stuff "*/"