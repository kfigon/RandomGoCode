Compiler course on edx from Stanford

# languages
* compilers - takes input (program) - produces executable. This is another program, that takes data and produces output. Offline
* interpreters - takes input (program) and data and produces direct output. Online

# compilation process

### lexical analysis
recognize words - syntax of language. Divide program text into words - `tokens`

`if x==y then z=1; else z=2;`
tokens:
* if
* then
* else
* ;
* z=1, z=2
* x==y
* spaces

### parsing
understand words. Identify role of token in the text. Group in higher level constructs. Building a tree of `if-then-else` statement:
* x==y - predicate
* == - relation
* = - assignment

### semantic analysis
understanding meaning of high level structure. **This is hard!** Compilers are only catching inconsistencies, we are very limited. We don't know what the program is meant to do. e.g.

* strict rules to avoid ambiguity - variable shadowing
* type checking

### optimization
auto modification of program so that they run faster/use less memory.
e.g. for integers we can optimize:

```
x=y*0
x=0
```
### code generation
producing an exe (c->machine code) or a program in other language (js->ts, c->asm)

# Our example language:
`cool` - classroom object oriented language. Compile `cool` to MIPS assembly. Run by spim exec (mips simulator). Language:
* abstraction
* static typing
* inheritance
* memory management

* https://nathanfriend.io/cooltojs/

this will return 1, explicit returns
```
class Main {
  main():Int { 1 };
};
```

```
class Main {
  i: IO <- new IO;
  main():IO {
    i.out_string("Hello, world!");
  };
};
```

```
class Main {
  i: IO <- new IO;
  atoi: A2I <- new A2I;

  -- dont care about return type
  main():Object { 
    let num: Int <- 0,
        result: Int <- 0
     in {

      i.out_string("provide input\n");
      num <- atoi.a2i(i.in_string());
      i.out_string("provided ".concat(atoi.i2a(num).concat("\n")));

      result <- fact(num);
      i.out_string("Result\n");
      i.out_string(atoi.itoa(result));
      i.out_string(atoi.itoa(factIter(num)));
    }
  };

  fact(i: Int): Int {
    if (i = 0) then 1 else i * fact(i-1) fi
  };

  factIter(i: Int): Int {
    let res: Int <- 1 in {
      while (not (i = 0)) loop 
        {
          res <- res * i;
          i <- i - 1;
        }
      pool; -- end of loop
      res; -- return statement to let
    }

  };
};
```

# Lexical Analysis (LA)
tokenize text input (classify program substring) and communicate tokens to parser.
list of tokens - pairs with `class` and `string corresponding`(lexeme)

* recognise substrings corresponding to tokens (lexemes)
* identify token class of each lexeme

classes: identifiers, keywords, `(`, `)`, numbers, strings, whitespace (grouped, 3 spaces are 1 token - whitespace) etc.

example of token - LA output:
* <ID, "foo"> - identifier
* <OP, "="> - assignment

```
if (i==j) z = 0;
else z = 1;
```
find whitespace, keywords, identifiers, numbers, equal operator, (, ), ;, =

