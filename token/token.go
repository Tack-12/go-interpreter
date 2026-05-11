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
	IDENT = "IDENT" //Represents the identifiers in the language: (variable names)
	INT   = "INT"   // INTEGERS

	//Operators:
	ASSIGN = "="
	ADD    = "+"

	//Delimeters:
	COMMA     = ","
	SEMICOLON = ";"

	LPARA = "("
	RPARA = ")"

	LBRAC = "{"
	RBRAC = "}"

	//KEYWORDS:
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	//RETURN THE TOKEN TYPE IDENT
	return IDENT
}
