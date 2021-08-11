# TinyBears
go_utils

## htmlbear
    - GetHTML(url string) string
    - ReGet(html, compile string) (res [][]string)

## sqlbear
    - NewAudit(sql, charset, collation string) (*Audit, error)
    - GetSQLInfo(sql string) (sqlInfo *parser.Result, err error)