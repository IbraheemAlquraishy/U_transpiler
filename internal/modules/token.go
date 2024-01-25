package token

type Tokentype string

const (
	Illegal = "illegal"
	EOF     = "EOF"

	Ident = "ident"
	//ident type consider a keyword
	Intt   = "int"
	Strt   = "string"
	Boolt  = "bool"
	Floatt = "float"
	//values
	Int   = "int"
	Str   = "string"
	Bool  = "bool"
	Float = "float"
	//one char operator
	Assign  = "="
	Plus    = "+"
	Sub     = "-"
	Multi   = "*"
	Div     = "/"
	Greater = ">"
	Lower   = "<"
	Not     = "!"
	Power   = "^"
	//two char operators
	Isequal        = "=="
	Notequal       = "!="
	Greaterorequal = ">="
	Lowerorequal   = "<="
	Inc            = "++"
	Dec            = "--"
	Plusequal      = "+="
	Subequal       = "-="
	Multiequal     = "*="
	Divequal       = "/="
	Powerequal     = "^="
	//other symbols
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//other keywords
	FUNCTION = "FUNCTION"
	TRUE     = "true"
	False    = "false"
	If       = "if"
	Else     = "else"
	Return   = "return"
)

type Token struct {
	Type Tokentype
	Lit  string
}

var keywords = map[string]Tokentype{
	"func":   FUNCTION,
	"int":    Intt,
	"string": Strt,
	"bool":   Boolt,
	"float":  Floatt,
	"true":   TRUE,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func Lookupident(ident string) Tokentype {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
