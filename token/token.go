package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//Language rules:
	ILLEGAL = "ILLEGAL" //Unkown Characters
	EOF     = "EOF"     //End of File

	//Identifiers + Literals:
	IDENT  = "IDENT" //Represents the identifiers in the language: (variable names)
	INT    = "INT"   // INTEGERS
	STRING = "STRING"

	//Operators:
	ASSIGN   = "="
	ADD      = "+"
	SUBTRACT = "-"
	STAR     = "*"
	FORSLASH = "/"
	EXCLA    = "!"

	//Delimeters:
	COMMA     = ","
	SEMICOLON = ";"

	LPARA = "("
	RPARA = ")"

	LBRAC = "{"
	RBRAC = "}"

	LSIGN = "<"
	RSIGN = ">"

	EQ     = "=="
	NOT_EQ = "!="

	LBOX = "["
	RBOX = "]"

	//KEYWORDS:
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	//RETURN THE TOKEN TYPE IDENT
	return IDENT
}
