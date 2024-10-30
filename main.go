package main

import (
	"ast-parse/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func parseGoFile(filePath string) error {
	// 1.创建一个新的文件集
	fset := token.NewFileSet()
	// 2.解析 Go 文件
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse the file: %v", err)
	}
	// 3.遍历 AST 节点
	simpleVisitor := visitor.SimpleVisitor{}
	ast.Walk(&simpleVisitor, node)
	return nil
}

func main() {
	filePath := "/Users/silhouette/codeworks/test-pool/main.go" // 替换为你的 Go 文件路径
	if err := parseGoFile(filePath); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
