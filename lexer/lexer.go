// Page 14
package lexer

type Lexer struct {
	input		 string
	position 	 int   // points to current char
	readPosition int   // after current char
	ch 			 byte  // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
			// reached EOF
			l.ch = 0 // byte code for NUL

		} else {

			l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition ++
}