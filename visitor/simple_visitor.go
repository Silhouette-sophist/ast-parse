package visitor

import (
	"go/ast"
	"log"
)

type SimpleVisitor struct {
}

func (receiver *SimpleVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		log.Printf("file %v", n.Name)
	case *ast.FuncDecl:
		log.Printf("func decl %v", n)
	case *ast.GoStmt:
		log.Printf("go stat %v", n)
	case *ast.BlockStmt:
		log.Printf("block, %v", n)
		if len(n.List) > 0 {
			switch n.List[0].(type) {
			case *ast.CaseClause:
				log.Printf("case clause")
			case *ast.CommClause:
				log.Printf("common clause")
			}
		}
	case *ast.IfStmt:
		log.Printf("ifStmt, %v", n)
	}
	return receiver
}
