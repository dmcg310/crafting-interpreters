package main

type _Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) _Scanner {
	return _Scanner{
		source: source,
	}
}

func (s *_Scanner) scanTokens() []Token {
	s.start, s.current = 0, 0
	s.line = 1

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens
}

func (s *_Scanner) scanToken() {
	l := Lox{}
	c := s.advance()
	var token _TokenType

	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			token = BANG_EQUAL
		} else {
			token = BANG
		}
		s.addToken(token, nil)
	case '=':
		if s.match('=') {
			token = EQUAL_EQUAL
		} else {
			token = EQUAL
		}
		s.addToken(token, nil)
	case '<':
		if s.match('=') {
			token = LESS_EQUAL
		} else {
			token = LESS
		}
		s.addToken(token, nil)
	case '>':
		if s.match('=') {
			token = GREATER_EQUAL
		} else {
			token = GREATER
		}
		s.addToken(token, nil)
	case '/':
		if s.match('/') {
			for peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	default:
		l.Error(s.line, "Unexpected character.")
	}
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

func (s *_Scanner) advance() byte {
	return '?'
}

func (s *_Scanner) addToken(tokenType _TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *_Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
