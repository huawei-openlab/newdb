package orm

import "errors"

// QueryBuilder is the Query builder interface
type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	From(tables ...string) QueryBuilder
	InnerJoin(table string) QueryBuilder
	LeftJoin(table string) QueryBuilder
	RightJoin(table string) QueryBuilder
	On(cond string) QueryBuilder
	Where(cond string) QueryBuilder
	And(cond string) QueryBuilder
	Or(cond string) QueryBuilder
	In(vals ...string) QueryBuilder
	OrderBy(fields ...string) QueryBuilder
	Asc() QueryBuilder
	Desc() QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	GroupBy(fields ...string) QueryBuilder
	Having(cond string) QueryBuilder
	Update(tables ...string) QueryBuilder
	Set(kv ...string) QueryBuilder
	Delete(tables ...string) QueryBuilder
	InsertInto(table string, fields ...string) QueryBuilder
	Values(vals ...string) QueryBuilder
	Subquery(sub string, alias string) string
	String() string
}

// NewQueryBuilder return the QueryBuilder
func NewQueryBuilder(driver string) (qb QueryBuilder, err error) {
	if driver == "mysql" {
		qb = new(MySQLQueryBuilder)
	} else if driver == "tidb" {
		qb = new(TiDBQueryBuilder)
	} else if driver == "postgres" {
		err = errors.New("postgres query builder is not supported yet")
	} else if driver == "sqlite" {
		err = errors.New("sqlite query builder is not supported yet")
	} else {
		err = errors.New("unknown driver for query builder")
	}
	return
}
