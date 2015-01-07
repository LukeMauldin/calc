// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package comp

import (
	"fmt"
	"os"

	"github.com/LukeMauldin/calc/ast"
	"github.com/LukeMauldin/calc/parse"
	"github.com/LukeMauldin/calc/token"
)

type compiler struct {
	fp *os.File
}

func CompileFile(fname, src string) {

	var c compiler
	fp, err := os.Create(fname + ".c")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	f := parse.ParseFile(fname, src)
	if f == nil {
		os.Exit(1)
	}
	c.fp = fp
	c.compFile(f)
}

func (c *compiler) compNode(node ast.Node) *numberNode {
	switch n := node.(type) {
	case *ast.BasicLit:
		return newNumberNode(n)
	case *ast.BinaryExpr:
		return c.compBinaryExpr(n)
	default:
		panic("unreachable")
	}
}

func (c *compiler) compBinaryExpr(b *ast.BinaryExpr) *numberNode {
	var tmp *numberNode

	tmp = c.compNode(b.List[0])

	for _, node := range b.List[1:] {
		switch b.Op {
		case token.ADD:
			tmp = tmp.add(c.compNode(node))
		case token.SUB:
			tmp = tmp.sub(c.compNode(node))
		case token.MUL:
			tmp = tmp.mul(c.compNode(node))
		case token.QUO:
			tmp = tmp.quo(c.compNode(node))
		case token.REM:
			tmp = tmp.rem(c.compNode(node))
		}
	}

	return tmp
}

func (c *compiler) compFile(f *ast.File) {
	fmt.Fprintln(c.fp, "#include <stdio.h>")
	fmt.Fprintln(c.fp, "int main(void) {")
	fmt.Fprintf(c.fp, "printf(\"%%s\", \"%s\");\n", c.compNode(f.Root))
	fmt.Fprintln(c.fp, "return 0;")
	fmt.Fprintln(c.fp, "}")
}
