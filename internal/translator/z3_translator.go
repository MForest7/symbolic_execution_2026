// Package translator содержит реализацию транслятора в Z3
package translator

import (
	"fmt"

	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

// Z3Translator транслирует символьные выражения в Z3 формулы
type Z3Translator struct {
	ctx    *z3.Context
	config *z3.Config
	vars   map[string]z3.Value // Кэш переменных
}

// NewZ3Translator создаёт новый экземпляр Z3 транслятора
func NewZ3Translator() *Z3Translator {
	config := &z3.Config{}
	ctx := z3.NewContext(config)

	return &Z3Translator{
		ctx:    ctx,
		config: config,
		vars:   make(map[string]z3.Value),
	}
}

// GetContext возвращает Z3 контекст
func (zt *Z3Translator) GetContext() interface{} {
	return zt.ctx
}

// Reset сбрасывает состояние транслятора
func (zt *Z3Translator) Reset() {
	zt.vars = make(map[string]z3.Value)
}

// Close освобождает ресурсы
func (zt *Z3Translator) Close() {
	// Z3 контекст закрывается автоматически
}

// TranslateExpression транслирует символьное выражение в Z3
func (zt *Z3Translator) TranslateExpression(expr symbolic.SymbolicExpression) (interface{}, error) {
	return expr.Accept(zt), nil
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// VisitVariable транслирует символьную переменную в Z3
func (zt *Z3Translator) VisitVariable(expr *symbolic.SymbolicVariable) interface{} {
	// Проверить, есть ли переменная в кэше
	// Если нет - создать новую Z3 переменную соответствующего типа
	// Добавить в кэш и вернуть

	if v, ok := zt.vars[expr.Name]; ok {
		return v
	} else {
		switch expr.Type() {
		case symbolic.IntType:
			{
				zt.vars[expr.Name] = zt.ctx.IntConst(expr.Name)
			}
		case symbolic.BoolType:
			{
				zt.vars[expr.Name] = zt.ctx.BoolConst(expr.Name)
			}
		default:
			{
				panic("Неизвестный тип")
			}
		}

		return zt.vars[expr.Name]
	}
}

// VisitIntConstant транслирует целочисленную константу в Z3
func (zt *Z3Translator) VisitIntConstant(expr *symbolic.IntConstant) interface{} {
	return zt.ctx.FromInt(expr.Value, zt.ctx.IntSort())
}

// VisitBoolConstant транслирует булеву константу в Z3
func (zt *Z3Translator) VisitBoolConstant(expr *symbolic.BoolConstant) interface{} {
	return zt.ctx.FromBool(expr.Value)
}

// VisitBinaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitBinaryOperation(expr *symbolic.BinaryOperation) interface{} {
	le, _ := zt.TranslateExpression(expr.Left)
	ri, _ := zt.TranslateExpression(expr.Right)

	left := le.(z3.Int)
	right := ri.(z3.Int)

	switch expr.Operator {
	case symbolic.ADD:
		return left.Add(right)
	case symbolic.SUB:
		return left.Sub(right)
	case symbolic.MUL:
		return left.Mul(right)
	case symbolic.DIV:
		return left.Div(right)
	case symbolic.MOD:
		return left.Mod(right)
	case symbolic.EQ:
		return left.Eq(right)
	case symbolic.NE:
		return left.NE(right)
	case symbolic.LT:
		return left.LT(right)
	case symbolic.LE:
		return left.LE(right)
	case symbolic.GT:
		return left.GT(right)
	case symbolic.GE:
		return left.GE(right)
	}

	panic("")
}

// VisitLogicalOperation транслирует логическую операцию в Z3
func (zt *Z3Translator) VisitLogicalOperation(expr *symbolic.LogicalOperation) interface{} {
	var operands []z3.Bool
	for _, subexpr := range expr.Operands {
		operand, _ := zt.TranslateExpression(subexpr)
		operands = append(operands, operand.(z3.Bool))
	}

	switch expr.Operator {
	case symbolic.AND:
		return operands[0].And(operands[1:]...)
	case symbolic.OR:
		return operands[0].Or(operands[1:]...)
	case symbolic.NOT:
		return operands[0].Not()
	case symbolic.IMPLIES:
		return operands[0].Implies(operands[1])
	}

	panic(fmt.Sprintf("неподдерживаемый логический оператор: %s", expr.Operator))
}

// Вспомогательные методы

// createZ3Variable создаёт Z3 переменную соответствующего типа
func (zt *Z3Translator) createZ3Variable(name string, exprType symbolic.ExpressionType) z3.Value {
	// TODO: Реализовать (вспомогательный метод)
	// Создать Z3 переменную на основе типа
	panic("не реализовано")
}

// castToZ3Type приводит значение к нужному Z3 типу
func (zt *Z3Translator) castToZ3Type(value interface{}, targetType symbolic.ExpressionType) (z3.Value, error) {
	// TODO: Реализовать (вспомогательный метод)
	// Безопасно привести interface{} к конкретному Z3 типу
	panic("не реализовано")
}
