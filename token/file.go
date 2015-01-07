// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package token

type File struct {
	base  Pos
	name  string
	src   string
	lines []Pos
}

func NewFile(name, src string) *File {
	return &File{
		base:  1,
		name:  name,
		src:   src,
		lines: make([]Pos, 0, 16),
	}
}

func (f *File) AddLine(offset Pos) {
	if offset >= f.base-1 && offset < f.base+Pos(len(f.src)) {
		f.lines = append(f.lines, offset)
	}
}

func (f *File) Base() Pos {
	return f.base
}

func (f *File) Pos(offset Pos) Pos {
	if offset >= Pos(len(f.src)) {
		panic("illegal file offset")
	}
	return Pos(f.base + offset)
}

func (f *File) Position(p Pos) Position {
	col, row := p, Pos(1)

	for i, nl := range f.lines {
		if p < f.Pos(nl) && i == 0 {
			//Handle scenario where p is one the first row
			break
		} else if p == f.Pos(nl) && i == 0 {
			//Handle scenario where p is a newline character
			col, row = 0, Pos(i+2)
		} else if p > f.Pos(nl) {
			col, row = p-f.Pos(nl), Pos(i+2)
		}
	}

	return Position{Filename: f.name, Col: col, Row: row}
}

func (f *File) Size() int {
	return len(f.src)
}
