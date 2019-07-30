package sema

import (
	"fmt"
	"github.com/dapperlabs/bamboo-node/pkg/language/runtime/ast"
	"github.com/dapperlabs/bamboo-node/pkg/language/runtime/common"
)

// astTypeConversionError

type astTypeConversionError struct {
	invalidASTType ast.Type
}

func (e *astTypeConversionError) Error() string {
	return fmt.Sprintf("cannot convert unsupported AST type: %#+v", e.invalidASTType)
}

// unsupportedAssignmentTargetExpression

type unsupportedAssignmentTargetExpression struct {
	target ast.Expression
}

func (e *unsupportedAssignmentTargetExpression) Error() string {
	return fmt.Sprintf("cannot assign to unsupported target expression: %#+v", e.target)
}

// RedeclarationError

// TODO: show previous declaration

type RedeclarationError struct {
	Kind        common.DeclarationKind
	Name        string
	Pos         *ast.Position
	PreviousPos *ast.Position
}

func (e *RedeclarationError) Error() string {
	return fmt.Sprintf("cannot redeclare already declared %s: %s", e.Kind.Name(), e.Name)
}

func (e *RedeclarationError) StartPosition() *ast.Position {
	return e.Pos
}

func (e *RedeclarationError) EndPosition() *ast.Position {
	return e.Pos
}

// NotDeclaredError

type NotDeclaredError struct {
	ExpectedKind common.DeclarationKind
	Name         string
	StartPos     *ast.Position
	EndPos       *ast.Position
}

func (e *NotDeclaredError) Error() string {
	return fmt.Sprintf("cannot find %s in this scope: %s", e.ExpectedKind.Name(), e.Name)
}

func (e *NotDeclaredError) SecondaryError() string {
	return "not found in this scope"
}

func (e *NotDeclaredError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *NotDeclaredError) EndPosition() *ast.Position {
	return e.EndPos
}

// AssignmentToConstantError

type AssignmentToConstantError struct {
	Name     string
	StartPos *ast.Position
	EndPos   *ast.Position
}

func (e *AssignmentToConstantError) Error() string {
	return fmt.Sprintf("cannot assign to constant: %s", e.Name)
}

func (e *AssignmentToConstantError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *AssignmentToConstantError) EndPosition() *ast.Position {
	return e.EndPos
}

// TypeMismatchError

type TypeMismatchError struct {
	ExpectedType Type
	ActualType   Type
	StartPos     *ast.Position
	EndPos       *ast.Position
}

func (e *TypeMismatchError) Error() string {
	return "mismatched types"
}

func (e *TypeMismatchError) SecondaryError() string {
	return fmt.Sprintf(
		"expected `%s`, found `%s`",
		e.ExpectedType.String(),
		e.ActualType.String(),
	)
}

func (e *TypeMismatchError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *TypeMismatchError) EndPosition() *ast.Position {
	return e.EndPos
}

// NotIndexableTypeError

type NotIndexableTypeError struct {
	Type     Type
	StartPos *ast.Position
	EndPos   *ast.Position
}

func (e *NotIndexableTypeError) Error() string {
	return fmt.Sprintf("cannot index into value which has type: %s", e.Type.String())
}

func (e *NotIndexableTypeError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *NotIndexableTypeError) EndPosition() *ast.Position {
	return e.EndPos
}

// NotIndexingTypeError

type NotIndexingTypeError struct {
	Type     Type
	StartPos *ast.Position
	EndPos   *ast.Position
}

func (e *NotIndexingTypeError) Error() string {
	return fmt.Sprintf("cannot index with value which has type: %s", e.Type.String())
}

func (e *NotIndexingTypeError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *NotIndexingTypeError) EndPosition() *ast.Position {
	return e.EndPos
}

// NotCallableError

type NotCallableError struct {
	Type     Type
	StartPos *ast.Position
	EndPos   *ast.Position
}

func (e *NotCallableError) Error() string {
	return fmt.Sprintf("cannot call type: %s", e.Type.String())
}

func (e *NotCallableError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *NotCallableError) EndPosition() *ast.Position {
	return e.EndPos
}

// ArgumentCountError

type ArgumentCountError struct {
	ParameterCount int
	ArgumentCount  int
	StartPos       *ast.Position
	EndPos         *ast.Position
}

func (e *ArgumentCountError) Error() string {
	return fmt.Sprintf(
		"incorrect number of arguments: got %d, need %d",
		e.ArgumentCount,
		e.ParameterCount,
	)
}

func (e *ArgumentCountError) StartPosition() *ast.Position {
	return e.StartPos
}

func (e *ArgumentCountError) EndPosition() *ast.Position {
	return e.EndPos
}
