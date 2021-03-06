// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"terminator",
		"import",
		"keywordAs",
		"return",
		"var",
		"=",
		":",
		"foreach",
		"(",
		"in",
		")",
		"if",
		"else",
		"raise",
		"{",
		"}",
		"keywordBusinessObject",
		"keywordNode",
		"[",
		",",
		"]",
		"n",
		"message",
		"text",
		"association",
		"to",
		"valuation",
		"using",
		"element",
		"keywordAction",
		"||",
		"&&",
		"+",
		"-",
		"*",
		"/",
		"%",
		".",
		"!",
		"&",
		"==",
		"!=",
		"<=",
		"<",
		">=",
		">",
		"floatLit",
		"intLit",
		"stringLit",
		"boolLit",
		"empty",
		"keywordRaises",
		"identifier",
		"keywordAssociation",
		"keywordTo",
		"keywordValuation",
	},

	idMap: map[string]Type{
		"INVALID":               0,
		"$":                     1,
		"terminator":            2,
		"import":                3,
		"keywordAs":             4,
		"return":                5,
		"var":                   6,
		"=":                     7,
		":":                     8,
		"foreach":               9,
		"(":                     10,
		"in":                    11,
		")":                     12,
		"if":                    13,
		"else":                  14,
		"raise":                 15,
		"{":                     16,
		"}":                     17,
		"keywordBusinessObject": 18,
		"keywordNode":           19,
		"[":                     20,
		",":                     21,
		"]":                     22,
		"n":                     23,
		"message":               24,
		"text":                  25,
		"association":           26,
		"to":                    27,
		"valuation":             28,
		"using":                 29,
		"element":               30,
		"keywordAction":         31,
		"||":                    32,
		"&&":                    33,
		"+":                     34,
		"-":                     35,
		"*":                     36,
		"/":                     37,
		"%":                     38,
		".":                     39,
		"!":                     40,
		"&":                     41,
		"==":                    42,
		"!=":                    43,
		"<=":                    44,
		"<":                     45,
		">=":                    46,
		">":                     47,
		"floatLit":              48,
		"intLit":                49,
		"stringLit":             50,
		"boolLit":               51,
		"empty":                 52,
		"keywordRaises":         53,
		"identifier":            54,
		"keywordAssociation":    55,
		"keywordTo":             56,
		"keywordValuation":      57,
	},
}
