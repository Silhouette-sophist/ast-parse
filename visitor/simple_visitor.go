package visitor

import (
	"fmt"
	"go/ast"
)

type SimpleVisitor struct {
}

func (receiver *SimpleVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fmt.Println("func decl")
	case *ast.GoStmt:
		fmt.Println("go stat")
	case *ast.BlockStmt:
		fmt.Println("block")
		if len(n.List) > 0 {
			switch n.List[0].(type) {
			case *ast.CaseClause:
				fmt.Println("case clause")
			case *ast.CommClause:
				fmt.Println("common clause")
			}
		}
	}
	return receiver
}
