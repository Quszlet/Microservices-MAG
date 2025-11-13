package utilssql

import (
	"fmt"
	"strings"
	"reflect"
)

func BuildUpdateQuery(table string, data map[string]any, where string) (string, map[string]any) {
	set := make([]string, 0, len(data))
	for k, v := range data {
		if (!reflect.ValueOf(v).IsZero()) {
			set = append(set, k+"=:"+k)
		}
	}
	q := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(set, ", "), where)
	return q, data
}