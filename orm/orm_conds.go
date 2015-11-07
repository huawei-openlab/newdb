package orm

import (
	"fmt"
	"strings"
)

// ExprSep define the expression seperation
const (
	ExprSep = "__"
)

type condValue struct {
	exprs  []string
	args   []interface{}
	cond   *Condition
	isOr   bool
	isNot  bool
	isCond bool
}

// Condition struct.
// work for WHERE conditions.
type Condition struct {
	params []condValue
}

// NewCondition return new condition struct
func NewCondition() *Condition {
	c := &Condition{}
	return c
}

// And add expression to condition
func (c Condition) And(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.And> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: strings.Split(expr, ExprSep), args: args})
	return &c
}

// AndNot add NOT expression to condition
func (c Condition) AndNot(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.AndNot> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: strings.Split(expr, ExprSep), args: args, isNot: true})
	return &c
}

// AndCond combine a condition to current condition
func (c *Condition) AndCond(cond *Condition) *Condition {
	c = c.clone()
	if c == cond {
		panic(fmt.Errorf("<Condition.AndCond> cannot use self as sub cond"))
	}
	if cond != nil {
		c.params = append(c.params, condValue{cond: cond, isCond: true})
	}
	return c
}

// Or add OR expression to condition
func (c Condition) Or(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.Or> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: strings.Split(expr, ExprSep), args: args, isOr: true})
	return &c
}

// OrNot add OR NOT expression to condition
func (c Condition) OrNot(expr string, args ...interface{}) *Condition {
	if expr == "" || len(args) == 0 {
		panic(fmt.Errorf("<Condition.OrNot> args cannot empty"))
	}
	c.params = append(c.params, condValue{exprs: strings.Split(expr, ExprSep), args: args, isNot: true, isOr: true})
	return &c
}

// OrCond combine a OR condition to current condition
func (c *Condition) OrCond(cond *Condition) *Condition {
	c = c.clone()
	if c == cond {
		panic(fmt.Errorf("<Condition.OrCond> cannot use self as sub cond"))
	}
	if cond != nil {
		c.params = append(c.params, condValue{cond: cond, isCond: true, isOr: true})
	}
	return c
}

// IsEmpty check the condition arguments are empty or not.
func (c *Condition) IsEmpty() bool {
	return len(c.params) == 0
}

// clone clone a condition
func (c Condition) clone() *Condition {
	return &c
}
