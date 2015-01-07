// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package scan

import (
	"strings"
	"unicode"

	"github.com/LukeMauldin/calc/token"
)

type Scanner struct {
	ch      rune
	offset  token.Pos
	roffset token.Pos
	src     string
	file    *token.File
}

func (s *Scanner) Init(file *token.File, src string) {
	s.file = file
	s.offset, s.roffset = 0, 0
	s.src = src
	s.file.AddLine(s.offset)

	s.next()
}

func (s *Scanner) Scan() (lit string, tok token.Token, pos token.Pos) {
	s.skipWhitespace()

	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}

	lit, pos = string(s.ch), s.file.Pos(s.offset)
	switch s.ch {
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '+':
		tok = token.ADD
	case '-':
		tok = token.SUB
	case '*':
		tok = token.MUL
	case '/':
		tok = token.QUO
	case '%':
		tok = token.REM
	case ';':
		s.skipComment()
		s.next()
		return s.Scan()
	default:
		if s.offset >= token.Pos(len(s.src)-1) {
			tok = token.EOF
		} else {
			tok = token.ILLEGAL
		}
	}

	s.next()

	return
}

func (s *Scanner) next() {
	s.ch = rune(0)
	if s.roffset < token.Pos(len(s.src)) {
		s.offset = s.roffset
		s.ch = rune(s.src[s.offset])
		if s.ch == '\n' {
			s.file.AddLine(s.offset)
		}
		s.roffset++
	}
}

func (s *Scanner) scanNumber() (string, token.Token, token.Pos) {
	start := s.offset

	for unicode.IsDigit(s.ch) || s.ch == '.' {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}

	//Detect if str is a integer or float
	str := s.src[start:offset]
	decimalCount := strings.Count(str, ".")
	var tokenType token.Token
	switch decimalCount {
	case 0:
		tokenType = token.INTEGER
	case 1:
		tokenType = token.FLOAT
	default:
		tokenType = token.ILLEGAL
	}
	return str, tokenType, s.file.Pos(start)
}

func (s *Scanner) skipComment() {
	for s.ch != '\n' && s.offset < token.Pos(len(s.src)-1) {
		s.next()
	}
}

func (s *Scanner) skipWhitespace() {
	for unicode.IsSpace(s.ch) {
		s.next()
	}
}
