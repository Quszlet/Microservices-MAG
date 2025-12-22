package utilssql

import (
	"fmt"
	"reflect"
	"strings"
)

func BuildUpdateQuery(table string, data map[string]any, where string) (string, map[string]any) {
	set := make([]string, 0, len(data))
	for k, v := range data {
		if !reflect.ValueOf(v).IsZero() {
			set = append(set, k+"=:"+k)
		}
	}
	q := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(set, ", "), where)
	return q, data
}

func BuildGetQuery(table string, fields []string, where string) (string, map[string]any) {
	q := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, ", "), table, where)
	return q, nil
}
