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
`cool` - classroom object oriented language. Compile `cool` to MIPS assembly. Run by spim exec. Language:
* abstraction
* static typing
* inheritance
* memory management


this will return 1, explicit returns
```
class Main {
  main():Int { 1 };
};

```
class Main {
  i: IO <- new IO;
  main():IO {
    i.out_string("Hello, world!");
  };
};
```