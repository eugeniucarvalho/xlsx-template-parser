wget http://www.antlr.org/download/antlr-4.7-complete.jar
alias antlr='java -jar $PWD/antlr-4.7-complete.jar'

antlr -Dlanguage=Go -o parser Calc.g4 
