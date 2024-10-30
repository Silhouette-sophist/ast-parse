package main

import (
	"ast-parse/visitor"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// parseGoFile 匹配单一go文件
func parseGoFile(filePath string) error {
	// 1.创建一个新的文件集
	fSet := token.NewFileSet()
	// 2.解析 Go 文件
	node, err := parser.ParseFile(fSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse the file: %v", err)
	}
	// 3.遍历 AST 节点
	simpleVisitor := visitor.SimpleVisitor{}
	ast.Walk(&simpleVisitor, node)
	return nil
}

// parseGoDir 匹配go目录下所有go文件
func parseGoDir(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是 Go 文件
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			return parseGoFile(path)
		}
		return nil
	})
}
