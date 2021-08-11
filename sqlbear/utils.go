package sqlbear

import "github.com/romberli/go-util/middleware/sql/parser"

// GetSQLInfo Get SQL Information
func GetSQLInfo(sql string) (sqlInfo *parser.Result, err error) {
	p := parser.NewParser(parser.NewVisitorWithDefault())
	sqlInfo, err = p.Parse(sql)
	if err != nil {
		return nil, err
	}

	return sqlInfo, nil
}
