package sgfparsing

import (
	"errors"
	"fmt"
	"strings"
)

// Node represents an SGF node with properties and child nodes.
type Node struct {
	Properties map[string][]string
	Children   []*Node
}

// charGetter is used to read chars one by one from a string.
type charGetter struct {
	encoded string
	pos     int
}

func (c *charGetter) done() bool {
	return c.pos >= len(c.encoded)
}

// peek returns the next char without incrementing the counter.
func (c *charGetter) peek() string {
	size := 1
	if c.encoded[c.pos:c.pos+size] == "\\" {
		size++
	}
	return c.encoded[c.pos : c.pos+size]
}

// next returns the next char, incrementing the counter.
func (c *charGetter) next() string {
	n := c.peek()
	c.pos += len(n)
	return n
}

// propKey returns a property key, reading until the next "[".
func (c *charGetter) propKey() (string, error) {
	var sb strings.Builder
	for !c.done() && c.peek() != "[" {
		sb.WriteString(c.next())
	}
	if c.done() {
		return "", errors.New("no [ found")
	}
	p := sb.String()
	if p != strings.ToUpper(p) {
		return "", errors.New("property name is not uppercase")
	}
	return p, nil
}

// propKey returns a property value, reading from "[" until "]".
func (c *charGetter) propVal() (string, error) {
	var sb strings.Builder
	if c.done() || c.peek() != "[" {
		return "", errors.New("no [ found")
	}
	c.next()
	for !c.done() && c.peek() != "]" {
		sb.WriteString(c.next())
	}
	if c.done() {
		return "", errors.New("no ] found")
	}
	c.next()
	p := sb.String()
	p = strings.ReplaceAll(p, "\\\t", " ")
	p = strings.ReplaceAll(p, "\\\n", "")
	p = strings.ReplaceAll(p, "\t", " ")
	sb.Reset()
	for i := 0; i < len(p); i++ {
		if p[i] == '\\' {
			i++
		}
		sb.WriteByte(p[i])
	}
	p = sb.String()
	return p, nil
}

// propKey returns the expression between a pair of "()".
func (c *charGetter) bracketed() (string, error) {
	d := 0
	var sb strings.Builder
	for !c.done() {
		n := c.next()
		sb.WriteString(n)
		switch n {
		case "(":
			d++
		case ")":
			d--
			if d == 0 {
				return sb.String(), nil
			}
		}
	}
	return "", errors.New("no ) found")
}

// rest returns the remaining unread string.
func (c *charGetter) rest() string {
	return c.encoded[c.pos:]
}

// parseNode parses a non-bracketed node.
func parseNode(encoded string) (*Node, error) {
	c := &charGetter{encoded, 0}
	if c.next() != ";" {
		return nil, fmt.Errorf("invalid syntax, %q", encoded)
	}
	node := &Node{make(map[string][]string), []*Node{}}
	for !c.done() {
		switch c.peek() {
		// A ";" indicates the rest of the string is a child node.
		case ";":
			if child, err := parseNode(c.rest()); err != nil {
				return nil, err
			} else {
				node.Children = append(node.Children, child)
			}
			return node, nil
		// A "(" indicates a child node starts here; there may be multiple bracketed nodes.
		case "(":
			b, err := c.bracketed()
			if err != nil {
				return nil, err
			}
			child, err := parseBracketed(b)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		// Anything else is a property.
		default:
			k, err := c.propKey()
			if err != nil {
				return nil, err
			}
			// Set one-or-more property values.
			for !c.done() && c.peek() == "[" {
				v, err := c.propVal()
				if err != nil {
					return nil, err
				}
				node.Properties[k] = append(node.Properties[k], v)
			}
		}
	}
	return node, nil
}

// parseNode parses a bracketed node.
func parseBracketed(encoded string) (*Node, error) {
	if encoded == "()" {
		return nil, errors.New("node cannot be empty")
	}
	if encoded[0] != '(' || encoded[len(encoded)-1] != ')' {
		return nil, fmt.Errorf("invalid syntax, %q", encoded)
	}
	return parseNode(encoded[1 : len(encoded)-1])
}

// Parse decodes an SGF string and returns the root node of the tree.
func Parse(encoded string) (*Node, error) {
	if encoded == "" {
		return nil, errors.New("node cannot be empty")
	}
	return parseBracketed(encoded)
}
