package evaluator

import (
	"go-interpreter/ast"
	"go-interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObj(node.Value)
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {

	var result object.Object

	for _, statements := range stmts {
		result = Eval(statements)
	}

	return result
}

func nativeBoolToBooleanObj(input bool) *object.Boolean {

	if input {
		return TRUE
	}

	return FALSE
}
