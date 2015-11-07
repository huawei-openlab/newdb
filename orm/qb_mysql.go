package orm

import (
	"fmt"
	"strconv"
	"strings"
)

// CommaSpace is the seperation
const CommaSpace = ", "

// MySQLQueryBuilder is the SQL build
type MySQLQueryBuilder struct {
	Tokens []string
}

// Select will join the fields
func (qb *MySQLQueryBuilder) Select(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "SELECT", strings.Join(fields, CommaSpace))
	return qb
}

// From join the tables
func (qb *MySQLQueryBuilder) From(tables ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "FROM", strings.Join(tables, CommaSpace))
	return qb
}

// InnerJoin INNER JOIN the table
func (qb *MySQLQueryBuilder) InnerJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "INNER JOIN", table)
	return qb
}

// LeftJoin LEFT JOIN the table
func (qb *MySQLQueryBuilder) LeftJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "LEFT JOIN", table)
	return qb
}

// RightJoin RIGHT JOIN the table
func (qb *MySQLQueryBuilder) RightJoin(table string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "RIGHT JOIN", table)
	return qb
}

// On join with on cond
func (qb *MySQLQueryBuilder) On(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "ON", cond)
	return qb
}

// Where join the Where cond
func (qb *MySQLQueryBuilder) Where(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "WHERE", cond)
	return qb
}

// And join the and cond
func (qb *MySQLQueryBuilder) And(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "AND", cond)
	return qb
}

// Or join the or cond
func (qb *MySQLQueryBuilder) Or(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "OR", cond)
	return qb
}

// In join the IN (vals)
func (qb *MySQLQueryBuilder) In(vals ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "IN", "(", strings.Join(vals, CommaSpace), ")")
	return qb
}

// OrderBy join the Order by fields
func (qb *MySQLQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "ORDER BY", strings.Join(fields, CommaSpace))
	return qb
}

// Asc join the asc
func (qb *MySQLQueryBuilder) Asc() QueryBuilder {
	qb.Tokens = append(qb.Tokens, "ASC")
	return qb
}

// Desc join the desc
func (qb *MySQLQueryBuilder) Desc() QueryBuilder {
	qb.Tokens = append(qb.Tokens, "DESC")
	return qb
}

// Limit join the limit num
func (qb *MySQLQueryBuilder) Limit(limit int) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "LIMIT", strconv.Itoa(limit))
	return qb
}

// Offset join the offset num
func (qb *MySQLQueryBuilder) Offset(offset int) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "OFFSET", strconv.Itoa(offset))
	return qb
}

// GroupBy join the Group by fields
func (qb *MySQLQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "GROUP BY", strings.Join(fields, CommaSpace))
	return qb
}

// Having join the Having cond
func (qb *MySQLQueryBuilder) Having(cond string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "HAVING", cond)
	return qb
}

// Update join the update table
func (qb *MySQLQueryBuilder) Update(tables ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "UPDATE", strings.Join(tables, CommaSpace))
	return qb
}

// Set join the set kv
func (qb *MySQLQueryBuilder) Set(kv ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "SET", strings.Join(kv, CommaSpace))
	return qb
}

// Delete join the Delete tables
func (qb *MySQLQueryBuilder) Delete(tables ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "DELETE")
	if len(tables) != 0 {
		qb.Tokens = append(qb.Tokens, strings.Join(tables, CommaSpace))
	}
	return qb
}

// InsertInto join the insert SQL
func (qb *MySQLQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	qb.Tokens = append(qb.Tokens, "INSERT INTO", table)
	if len(fields) != 0 {
		fieldsStr := strings.Join(fields, CommaSpace)
		qb.Tokens = append(qb.Tokens, "(", fieldsStr, ")")
	}
	return qb
}

// Values join the Values(vals)
func (qb *MySQLQueryBuilder) Values(vals ...string) QueryBuilder {
	valsStr := strings.Join(vals, CommaSpace)
	qb.Tokens = append(qb.Tokens, "VALUES", "(", valsStr, ")")
	return qb
}

// Subquery join the sub as alias
func (qb *MySQLQueryBuilder) Subquery(sub string, alias string) string {
	return fmt.Sprintf("(%s) AS %s", sub, alias)
}

// String join all Tokens
func (qb *MySQLQueryBuilder) String() string {
	return strings.Join(qb.Tokens, " ")
}
