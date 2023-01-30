# Lox programming language

based on https://craftinginterpreters.com/contents.html

## run
```
go run .
```
### more params
* run without params to run interpreter
* run with a file name to run the interpreter on the file itself


## test
```
go test ./...
```


## Grammar

```
expression     → literal
               | unary
               | binary
               | grouping ;

literal        → NUMBER | STRING | "true" | "false" | "nil" ;
grouping       → "(" expression ")" ;
unary          → ( "-" | "!" ) expression ;
binary         → expression operator expression ;
operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
               | "+"  | "-"  | "*" | "/" ;
```