package markdown

import (
	"fmt"
	"strings"
)

// Markdown tokens.
const (
	EOF = iota
	Text
	Strong
	Emphasis
	Header
	Newline
	List
	Paragraph
)

// tags maps some markdown tokens to HTML tags.
var tags = map[int]string{
	Strong:    "strong",
	Emphasis:  "em",
	Paragraph: "p",
}

// Lexer reads a markdown string and emits tokens.
type Lexer struct {
	pos  int
	size int
	data string
}

// NewLexer creates a new Lexer.
func NewLexer(markdown string) *Lexer {
	return &Lexer{pos: 0, size: len(markdown), data: markdown}
}

// NextToken emits tokens from markdown.
// The return values are (1) the markdown text, (2) the token type, and (3) if the token is at the start of a line.
func (l *Lexer) NextToken() (text string, token int, firstCol bool) {
	firstCol = l.pos == 0 || l.data[l.pos-1] == '\n'

	if l.pos == l.size {
		return "", EOF, firstCol
	}

	text = l.data[l.pos : l.pos+1]
	switch text {
	// "_" is either Emphasis "_" or Strong "__"; check next char to decide.
	case "_":
		l.pos++
		if l.pos < l.size && l.data[l.pos] == '_' {
			l.pos++
			return "__", Strong, firstCol
		}
		return text, Emphasis, firstCol
	// "*" is a List item.
	case "*":
		if firstCol {
			l.pos++
			return text, List, firstCol
		}
	// `#{1,6}` is a header.
	case "#":
		size := 1
		header := "#"
		for ; l.data[l.pos+size] == '#'; size++ {
			header += "#"
		}
		if firstCol && size < 7 {
			l.pos += size
			return header, Header, firstCol
		}
	case "\n":
		l.pos++
		return text, Newline, firstCol
	}
	l.pos++
	return text, Text, firstCol
}

// EatSpaces discards spaces.
// Spaces after an open tag should be discarded.
func (l *Lexer) EatSpaces() {
	for ; l.pos < l.size && l.data[l.pos] == ' '; l.pos++ {
	}
}

// TokenPrinter turns a stream of markdown tokens into HTML.
type TokenPrinter struct {
	out    strings.Builder
	state  map[int]bool
	header string
}

// NewTokenPrinter returns a new TokenPrinter.
func NewTokenPrinter() *TokenPrinter {
	return &TokenPrinter{
		state: map[int]bool{},
	}
}

// List handles a List token.
func (hw *TokenPrinter) List(text string, token int, firstCol bool) bool {
	if !hw.state[List] {
		hw.Write("<ul>")
	}
	hw.state[List] = true
	hw.state[Paragraph] = true
	hw.Write("<li>")
	return true
}

// Tag handles Strong, Emphasis and Paragraph tokens.
func (hw *TokenPrinter) Tag(text string, token int, firstCol bool) bool {
	hw.state[token] = !hw.state[token]
	if hw.state[token] {
		hw.Text(fmt.Sprintf("<%s>", tags[token]), token, firstCol)
		return true
	}
	hw.Text(fmt.Sprintf("</%s>", tags[token]), token, firstCol)
	return false
}

// Header handles Header tokens of various lengths.
func (hw *TokenPrinter) Header(text string, token int, firstCol bool) bool {
	hw.header = fmt.Sprintf("h%d", len(text))
	hw.state[Paragraph] = true
	hw.Write("<" + hw.header + ">")
	return true
}

// Newline handles Newline tokens, closing tags as needed.
func (hw *TokenPrinter) Newline(text string, token int, firstCol bool) bool {
	if hw.header != "" {
		hw.Write("</" + hw.header + ">")
		hw.header = ""
	}
	if hw.state[List] {
		hw.Write("</li>")
		hw.state[Paragraph] = false
	}
	return true
}

// Text is used to handle Text and generally emit text.
func (hw *TokenPrinter) Text(text string, token int, firstCol bool) bool {
	if hw.state[List] && firstCol {
		hw.Write("</ul>")
		hw.state[List] = false
	}
	if !hw.state[Paragraph] {
		hw.Write("<p>")
	}
	hw.state[Paragraph] = true
	hw.Write(text)
	return false
}

// Write writes output.
func (hw *TokenPrinter) Write(s string) {
	hw.out.WriteString(s)
}

// Close closes out open tags.
// After calling Close(), only String() should be called on this TokenPrinter.
func (hw *TokenPrinter) Close() {
	if hw.header != "" {
		hw.Write("</" + hw.header + ">")
	} else if hw.state[List] {
		hw.Write("</li></ul>")
	} else if hw.state[Paragraph] {
		hw.Write("</p>")
	}
}

// String returns the HTML string.
func (hw *TokenPrinter) String() string {
	return hw.out.String()
}

// Render translates markdown to HTML.
func Render(markdown string) string {
	lexer := NewLexer(markdown)
	printer := NewTokenPrinter()
	ops := map[int]func(string, int, bool) bool{
		Newline:  printer.Newline,
		Text:     printer.Text,
		Strong:   printer.Tag,
		Emphasis: printer.Tag,
		Header:   printer.Header,
		List:     printer.List,
	}

	for {
		// Read markdown tokens.
		text, token, firstCol := lexer.NextToken()
		// On EOF, Close the TokenPrinter and return the string.
		if token == EOF {
			printer.Close()
			return printer.String()
		}
		// Otherwise, handle the token and possible consume spaces.
		if ops[token](text, token, firstCol) {
			lexer.EatSpaces()
		}
	}
}
