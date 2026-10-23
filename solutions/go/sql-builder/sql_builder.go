package sqlbuilder

import (
	"fmt"
	"strings"
)

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	if len(columns) == 0 {
		columns = append(columns, "*")
	}
	return fmt.Sprintf("SELECT %s FROM %s;", strings.Join(columns, ", "), table)
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	var out []string
	var sb strings.Builder
	for _, row := range rows {
		sb.WriteString("INSERT INTO ")
		sb.WriteString(table)
		sb.WriteString(" (")
		for i, col := range row {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(col.Column)
		}
		sb.WriteString(") VALUES ('")
		for i, col := range row {
			if i > 0 {
				sb.WriteString("', '")
			}
			sb.WriteString(col.Value)
		}
		sb.WriteString("');")
		out = append(out, sb.String())
		sb.Reset()
	}
	return out
}
