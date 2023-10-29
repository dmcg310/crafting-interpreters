package scanner

import (
	"github.com/dmcg310/glox/src/report"
	"github.com/dmcg310/glox/src/token"
	"strconv"
)

type _Scanner struct {
	source   string
	tokens   []token.Token
	start    int
	current  int
	line     int
	reporter report.Reporter
	keywords map[string]token.TTokentype
}

func NewScanner(source string, reporter report.Reporter) _Scanner {
	s := _Scanner{
		source:   source,
		reporter: reporter,
	}
	s.InitKeywords()

	return s
}

func (s *_Scanner) ScanTokens() []token.Token {
	s.start, s.current = 0, 0
	s.line = 1

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens
}

func (s *_Scanner) scanToken() {
	c := s.advance()
	var _token token.TTokentype

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
	case '!':
		if s.match('=') {
			_token = token.BANG_EQUAL
		} else {
			_token = token.BANG
		}
		s.addToken(_token, nil)
	case '=':
		if s.match('=') {
			_token = token.EQUAL_EQUAL
		} else {
			_token = token.EQUAL
		}
		s.addToken(_token, nil)
	case '<':
		if s.match('=') {
			_token = token.LESS_EQUAL
		} else {
			_token = token.LESS
		}
		s.addToken(_token, nil)
	case '>':
		if s.match('=') {
			_token = token.GREATER_EQUAL
		} else {
			_token = token.GREATER
		}
		s.addToken(_token, nil)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			s.reporter.Error(s.line, "Unexpected character.")
		}
	}
}

func (s *_Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType, exists := s.keywords[text]
	if !exists {
		tokenType = token.IDENTIFIER
	}

	s.addToken(tokenType, nil)
}

func (s *_Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// look for a fractional part
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// consume the '.'
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	substr := s.source[s.start:s.current]
	val, _ := strconv.ParseFloat(substr, 64)
	s.addToken(token.NUMBER, val)
}

func (s *_Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		s.reporter.Error(s.line, "Unterminated string.")
		return
	}

	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addToken(token.STRING, value)
}

func (s *_Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *_Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *_Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *_Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *_Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *_Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *_Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *_Scanner) addToken(tokenType token.TTokentype, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.NewToken(tokenType, text, literal, s.line))
}

func (s *_Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *_Scanner) InitKeywords() {
	s.keywords = make(map[string]token.TTokentype)
	s.keywords["and"] = token.AND
	s.keywords["class"] = token.CLASS
	s.keywords["else"] = token.ELSE
	s.keywords["false"] = token.FALSE
	s.keywords["for"] = token.FOR
	s.keywords["fun"] = token.FUN
	s.keywords["if"] = token.IF
	s.keywords["nil"] = token.NIL
	s.keywords["or"] = token.OR
	s.keywords["print"] = token.PRINT
	s.keywords["return"] = token.RETURN
	s.keywords["super"] = token.SUPER
	s.keywords["this"] = token.THIS
	s.keywords["true"] = token.TRUE
	s.keywords["var"] = token.VAR
	s.keywords["while"] = token.WHILE
}
