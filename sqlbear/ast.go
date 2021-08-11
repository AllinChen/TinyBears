package sqlbear

import (
	"fmt"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"

	"vitess.io/vitess/go/vt/sqlparser"
)

type Audit struct {
	Query  string              // 查询语句
	Stmt   sqlparser.Statement // 通过Vitess解析出的抽象语法树
	TiStmt []ast.StmtNode      // 通过TiDB解析出的抽象语法树
}

// NewQuery4Audit return a struct for Audit
func NewAudit(sql, charset, collation string) (*Audit, error) {
	q := &Audit{Query: sql}
	// vitess 语法解析不上报，以 tidb parser 为主
	q.Stmt, _ = sqlparser.Parse(sql)
	// tdib parser 语法解析
	var warns []error
	var err error
	sqlInfo, err := GetSQLInfo(sql)
	if err == nil {
		for columnName, columnType := range sqlInfo.ColumnTypes {
			fmt.Println(columnName, columnType)
		}
	}
	q.TiStmt, warns, err = parser.New().Parse(sql, charset, collation)
	if len(warns) > 0 {
		return q, fmt.Errorf(fmt.Sprint(warns))
	}

	return q, err
}
