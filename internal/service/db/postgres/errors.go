package postgres

import "strings"

func IsDuplicateErr(err error) bool {
	return strings.Contains(err.Error(), "[service.db.postgres] Duplicate entry")
}
