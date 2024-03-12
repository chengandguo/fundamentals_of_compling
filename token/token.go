package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // 标识符
	INT   = "INT"

	ASSIGN = "="
	EQUAL  = "=="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	// control characters
	SPACE = " "
	LF    = "\n"

	RETURN = "RETURN"
)

// token/token.go

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

/*
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
*/

type Tokenizer string

func (t Tokenizer) Tokenizer() {
	var index int = 0 // current index
	var total int = len(t)
	// var lineNumber int = 0 // current line number
	var result []Token

	getErrorMsg := func(expected string, got string) string {
		msg := fmt.Sprintf("expected %s, but got %s", expected, got)
		return msg
	}

	next := func(targetChar string) {
		if string(t[index]) != targetChar {
			msg := getErrorMsg(targetChar, string(t[index]))
			panic(msg)
		}
		index++
	}

	handleLParenthesis := func() {
		next(LPAREN)
		token := Token{Type: LPAREN, Literal: "("}
		result = append(result, token)
	}

	handleRParenthesis := func() {
		next(RPAREN)
		token := Token{Type: RPAREN, Literal: ")"}
		result = append(result, token)
	}

	handleAssign := func() {
		next(ASSIGN)
		token := Token{Type: ASSIGN, Literal: "="}
		result = append(result, token)
	}

	isLetter := func(r rune) bool {
		return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r == '_'
	}

	handleVariableName := func() {
		variableName := ""
		for isLetter(rune(t[index])) {
			variableName += string(t[index])
			index++
		}

		if variableName != "" {
			ident := LookupIdent(variableName)
			token := Token{Type: ident, Literal: variableName}
			result = append(result, token)
		} else {
			index++
		}
	}

	handleSemicolon := func() {
		next(";")
		token := Token{Type: SEMICOLON, Literal: ";"}
		result = append(result, token)
	}

	handleLBrace := func() {
		next("{")
		token := Token{Type: LBRACE, Literal: "{"}
		result = append(result, token)
	}

	handleRBrace := func() {
		next("}")
		token := Token{Type: RBRACE, Literal: "}"}
		result = append(result, token)
	}

	handleComma := func() {
		next(",")
		token := Token{Type: COMMA, Literal: ","}
		result = append(result, token)
	}

	handlePlus := func() {
		next("+")
		token := Token{Type: PLUS, Literal: "+"}
		result = append(result, token)
	}

	loop := func() {
		for index < total {
			current := string(t[index])
			switch current {
			case " ":
				index++
			case "\n":
				index++
			case "=":
				handleAssign()
			case ";":
				handleSemicolon()
			case "{":
				handleLBrace()
			case "}":
				handleRBrace()
			case ",":
				handleComma()
			case "(":
				handleLParenthesis()
			case ")":
				handleRParenthesis()
			case "+":
				handlePlus()
			default:
				handleVariableName()
			}
		}
	}

	loop()
	// fmt.Printf("%#v", result)
	fmt.Println(result)
}
