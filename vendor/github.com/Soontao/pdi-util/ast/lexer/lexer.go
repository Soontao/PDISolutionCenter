// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/Soontao/pdi-util/ast/token"
)

const (
	NoState    = -1
	NumStates  = 165
	NumSymbols = 223
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: ';'
1: '\n'
2: ';'
3: 'a'
4: 's'
5: 'b'
6: 'u'
7: 's'
8: 'i'
9: 'n'
10: 'e'
11: 's'
12: 's'
13: 'o'
14: 'b'
15: 'j'
16: 'e'
17: 'c'
18: 't'
19: 'n'
20: 'o'
21: 'd'
22: 'e'
23: 'r'
24: 'a'
25: 'i'
26: 's'
27: 'e'
28: 's'
29: 'a'
30: 'c'
31: 't'
32: 'i'
33: 'o'
34: 'n'
35: '.'
36: '.'
37: 't'
38: 'r'
39: 'u'
40: 'e'
41: 'f'
42: 'a'
43: 'l'
44: 's'
45: 'e'
46: 'i'
47: 'm'
48: 'p'
49: 'o'
50: 'r'
51: 't'
52: 'r'
53: 'e'
54: 't'
55: 'u'
56: 'r'
57: 'n'
58: 'v'
59: 'a'
60: 'r'
61: '='
62: ':'
63: 'f'
64: 'o'
65: 'r'
66: 'e'
67: 'a'
68: 'c'
69: 'h'
70: '('
71: 'i'
72: 'n'
73: ')'
74: 'i'
75: 'f'
76: 'e'
77: 'l'
78: 's'
79: 'e'
80: 'r'
81: 'a'
82: 'i'
83: 's'
84: 'e'
85: '{'
86: '}'
87: 'm'
88: 'e'
89: 's'
90: 's'
91: 'a'
92: 'g'
93: 'e'
94: 't'
95: 'e'
96: 'x'
97: 't'
98: 'a'
99: 's'
100: 's'
101: 'o'
102: 'c'
103: 'i'
104: 'a'
105: 't'
106: 'i'
107: 'o'
108: 'n'
109: 't'
110: 'o'
111: 'v'
112: 'a'
113: 'l'
114: 'u'
115: 'a'
116: 't'
117: 'i'
118: 'o'
119: 'n'
120: 'u'
121: 's'
122: 'i'
123: 'n'
124: 'g'
125: 'e'
126: 'l'
127: 'e'
128: 'm'
129: 'e'
130: 'n'
131: 't'
132: ','
133: '|'
134: '|'
135: '&'
136: '&'
137: '+'
138: '-'
139: '*'
140: '/'
141: '%'
142: '!'
143: '&'
144: '='
145: '='
146: '!'
147: '='
148: '<'
149: '='
150: '<'
151: '>'
152: '='
153: '>'
154: '.'
155: '['
156: ']'
157: 'n'
158: '/'
159: '/'
160: '\n'
161: '/'
162: '*'
163: '*'
164: '*'
165: '/'
166: '_'
167: '0'
168: '0'
169: 'x'
170: 'X'
171: 'e'
172: 'E'
173: '+'
174: '-'
175: '`'
176: '`'
177: '"'
178: '\'
179: '"'
180: '"'
181: '\'
182: 'n'
183: '\'
184: 'r'
185: '\'
186: 't'
187: '='
188: '='
189: '!'
190: '='
191: '<'
192: '<'
193: '='
194: '>'
195: '>'
196: '='
197: '+'
198: '-'
199: '|'
200: '^'
201: '*'
202: '/'
203: '%'
204: '<'
205: '<'
206: '>'
207: '>'
208: '&'
209: '&'
210: '^'
211: ' '
212: '\t'
213: '\r'
214: '\n'
215: 'a'-'z'
216: 'A'-'Z'
217: '0'-'9'
218: '0'-'7'
219: 'a'-'f'
220: 'A'-'F'
221: '1'-'9'
222: .
*/