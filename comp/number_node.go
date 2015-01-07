package comp

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LukeMauldin/calc/ast"
	"github.com/LukeMauldin/calc/token"
)

type numberNode struct {
	kind token.Token
	i    int
	f    float64
}

func newNumberNode(lit *ast.BasicLit) *numberNode {
	ret := &numberNode{kind: lit.Kind}
	if lit.Kind == token.INTEGER {
		i, err := strconv.Atoi(lit.Lit)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ret.i = i
	} else if lit.Kind == token.FLOAT {
		f, err := strconv.ParseFloat(lit.Lit, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ret.f = f
	} else {
		fmt.Println("Unsupported token: " + lit.Kind.String())
	}
	return ret
}

func (f *numberNode) add(other *numberNode) *numberNode {
	if f.kind == token.INTEGER && other.kind == token.INTEGER {
		//Perform operation directly on integers
		i := f.i + other.i
		return &numberNode{kind: token.INTEGER, i: i}
	} else {
		//Ensure both are floats and perform operation
		f := f.asFloat() + other.asFloat()
		return &numberNode{kind: token.FLOAT, f: f}
	}
}

func (f *numberNode) sub(other *numberNode) *numberNode {
	if f.kind == token.INTEGER && other.kind == token.INTEGER {
		//Perform operation directly on integers
		i := f.i - other.i
		return &numberNode{kind: token.INTEGER, i: i}
	} else {
		//Ensure both are floats and perform operation
		f := f.asFloat() - other.asFloat()
		return &numberNode{kind: token.FLOAT, f: f}
	}
}

func (f *numberNode) mul(other *numberNode) *numberNode {
	if f.kind == token.INTEGER && other.kind == token.INTEGER {
		//Perform operation directly on integers
		i := f.i * other.i
		return &numberNode{kind: token.INTEGER, i: i}
	} else {
		//Ensure both are floats and perform operation
		f := f.asFloat() * other.asFloat()
		return &numberNode{kind: token.FLOAT, f: f}
	}
}

func (f *numberNode) quo(other *numberNode) *numberNode {
	if f.kind == token.INTEGER && other.kind == token.INTEGER {
		//Perform operation directly on integers
		i := f.i / other.i
		return &numberNode{kind: token.INTEGER, i: i}
	} else {
		//Ensure both are floats and perform operation
		f := f.asFloat() / other.asFloat()
		return &numberNode{kind: token.FLOAT, f: f}
	}
}

func (f *numberNode) rem(other *numberNode) *numberNode {
	if f.kind == token.INTEGER && other.kind == token.INTEGER {
		//Perform operation directly on integers
		i := f.i % other.i
		return &numberNode{kind: token.INTEGER, i: i}
	} else {
		//Ensure both are floats and perform operation
		panic("Operator '%' undefined for float")
	}
}

func (f *numberNode) asFloat() float64 {
	if f.kind == token.FLOAT {
		return f.f
	} else {
		return float64(f.i)
	}
}

func (f *numberNode) String() string {
	if f.kind == token.INTEGER {
		return strconv.Itoa(f.i)
	} else {
		return strconv.FormatFloat(f.f, 'f', 2, 64)
	}
}
