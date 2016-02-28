# qbx

Go query builder for [pgx](https://github.com/jackc/pgx). Still under active development.


## Why created

No database agnostic, just purely tuned for PostgreSQL query building + loading data to struct.


## Usage

```go
b := &qbx.SelectStmt{}
sql, err := b.Select("id", "name", "email").From("account").Limit(10).ToSQL()
if err != nil {
	t.Fatal(err)
}
log.Println(sql) // SELECT id, name, email FROM account LIMIT 10	
```	
   	
## Inspired by

Huge thanks to both of the libraries.

- [dbr](https://github.com/gocraft/dbr)
- [squirrel](https://github.com/Masterminds/squirrel)
