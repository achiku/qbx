package qbx

import (
	"bytes"
	"errors"
	"fmt"
)

// SelectStmt builds select statement
type SelectStmt struct {
	Comment    string
	IsDistinct bool

	Columns []string
	Set     interface{}

	LimitCount  int64
	OffsetCount int64
}

// ToSQL creates sql statement
func (b *SelectStmt) ToSQL() (string, error) {
	if len(b.Columns) == 0 {
		return "", errors.New("at least one column must be specified")
	}
	buf := new(bytes.Buffer)
	buf.WriteString("SELECT ")
	if b.IsDistinct {
		buf.WriteString("DISTINCT ")
	}
	for i, col := range b.Columns {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(col)
	}

	if b.Set != nil {
		buf.WriteString(" FROM ")
		// subquery has to be able to be used here
		buf.WriteString(b.Set.(string))
	}

	if b.LimitCount >= 0 {
		buf.WriteString(" LIMIT ")
		buf.WriteString(fmt.Sprint(b.LimitCount))
	}
	return buf.String(), nil
}

// Select creates select part of SQL
func (b *SelectStmt) Select(columns ...string) *SelectStmt {
	b.Columns = columns
	return b
}

// From creates from part of SQL
func (b *SelectStmt) From(set interface{}) *SelectStmt {
	b.Set = set
	return b
}

// Distinct creates distinct part of SQL
func (b *SelectStmt) Distinct() *SelectStmt {
	b.IsDistinct = true
	return b
}

// Limit create limit part of SQL
func (b *SelectStmt) Limit(n uint64) *SelectStmt {
	b.LimitCount = int64(n)
	return b
}
