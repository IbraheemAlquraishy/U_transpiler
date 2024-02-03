# U_transpiler
A transpiler for U (a language i made) to c++
## about this project
I was bored when i decided I want to make my own transpiler just to waste some time and a challenge to test my skills
## How this works
well, we get the input from the file as string and we input it in the lexer after the lexer it goes to the parser to get ast after that the type checker check for errors in the code and in the end the coder convert it to c++ and writes the new code in a file
### lexer
the lexer is simple we set the tokens of the lanuage in the token module "can be found in ./internal/modules/token" it checks ever single char in the input string to match it with the tokens defined 
the lexer will output an array of tokens
### parser
parser convert the array of tokens to a tree "ast" this tree make it easier to convert and to check for errors "the tree is defined in ./internal/modules/ast"
the tree consist of 2 types statment and expression 
### type checker
after the parser finishes it will check for errors in the code and its also responsible of datatype evalation ":= operator" and functions overloading 
### coder
the coder will have an objects of the lexer parser and 2 from the checkers the 1st checker will check the functions and the other the rest this is just to have a privte vars inside the functions 
then it will simply write the tree and c++ code and output it
## syntax of U
### declaring var
```
x int;
```
or
```
x:=5;
```
### if condition
```
if (condition){
  body
}
```
### for loop
```
for x int=0;x<5;x++{

}
```
#### yeah i know its weird that for dont need () but if do
### function declartion
```
func FUNCTIONNAME(PARAMS)RETURNTYPE{
BODY
}
```
#### return types are the same as data types plus void
### keywords
```
print x;
println "5";
input y;
```
## data types in U
int 
string
bool
float
## roadmap
[]args for the transpiler
[]the use of goroutins to get things faster
[]arrays and len keyword
