package orm

import (
	"fmt"
)

// mysql dbBaser implementation.
type dbBaseTidb struct {
	dbBase
}

var _ dbBaser = new(dbBaseTidb)

// get mysql operator.
func (d *dbBaseTidb) OperatorSQL(operator string) string {
	return mysqlOperators[operator]
}

// get mysql table field types.
func (d *dbBaseTidb) DbTypes() map[string]string {
	return mysqlTypes
}

// show table sql for mysql.
func (d *dbBaseTidb) ShowTablesQuery() string {
	return "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema = DATABASE()"
}

// show columns sql of table for mysql.
func (d *dbBaseTidb) ShowColumnsQuery(table string) string {
	return fmt.Sprintf("SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE FROM information_schema.columns "+
		"WHERE table_schema = DATABASE() AND table_name = '%s'", table)
}

// execute sql to check index exist.
func (d *dbBaseTidb) IndexExists(db dbQuerier, table string, name string) bool {
	row := db.QueryRow("SELECT count(*) FROM information_schema.statistics "+
		"WHERE table_schema = DATABASE() AND table_name = ? AND index_name = ?", table, name)
	var cnt int
	row.Scan(&cnt)
	return cnt > 0
}

// create new mysql dbBaser.
func newdbBaseTidb() dbBaser {
	b := new(dbBaseTidb)
	b.ins = b
	return b
}
