package database

import (
	"fmt"
	"strings"
)

// ConvertMapToArgsStr converts a map of filters to a string of args for a WHERE clause.
func ConvertMapToArgsStr(filters map[string]any) string {
	args := []string{}
	for key := range filters {
		args = append(args, fmt.Sprintf("%s = :%s", key, key))
	}
	return strings.Join(args, " AND ")
}
