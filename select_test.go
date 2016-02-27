package qbx

import "testing"

func TestSimpleSelectStmt(t *testing.T) {
	b := &SelectStmt{}
	b.Select("id", "name", "email").From("account").Limit(10)
	sql, err := b.ToSQL()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sql)
}
