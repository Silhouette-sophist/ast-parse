在 Go 的 go/ast 包中，AST（抽象语法树）节点被组织成一个层级结构。每种节点类型都实现了 ast.Node 接口。以下是一些主要的 AST 节点类型及其分类：

1. 根节点
   *ast.File: 表示一个 Go 源文件，包括包声明、导入声明、常量、变量、类型和函数等。
2. 声明节点
   *ast.GenDecl: 表示一般声明（如常量、变量、类型声明）。
   子节点:
   *ast.ValueSpec: 表示常量或变量的值说明。
   *ast.TypeSpec: 表示类型声明。
   *ast.FuncDecl: 表示函数声明。
3. 表达式节点
   *ast.BasicLit: 表示基本字面量，如字符串、整数等。
   *ast.Ident: 表示标识符（如变量、函数名等）。
   *ast.SelectorExpr: 表示选择器表达式（如 pkg.Name）。
   *ast.CallExpr: 表示函数调用表达式，如 f(x)。
   *ast.BinaryExpr: 表示二元操作符表达式（如 a + b）。
   *ast.UnaryExpr: 表示一元操作符表达式（如 -a）。
4. 语句节点
   *ast.BlockStmt: 表示语句块。
   *ast.IfStmt: 表示 if 语句。
   *ast.ForStmt: 表示 for 循环。
   *ast.SwitchStmt: 表示 switch 语句。
   *ast.ReturnStmt: 表示返回语句。
   *ast.ExprStmt: 表示表达式语句。
5. 类型节点
   *ast.ArrayType: 表示数组类型。
   *ast.StructType: 表示结构体类型。
   *ast.InterfaceType: 表示接口类型。
   *ast.MapType: 表示映射类型。
   *ast.StarExpr: 表示指针类型。
6. 其他节点
   *ast.Package: 表示包，通常由 go/parser 解析后返回。
   *ast.Comment: 表示单行评论。
   *ast.CommentGroup: 表示一组评论。