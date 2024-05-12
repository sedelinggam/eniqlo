package lumen

import "strings"

func CheckErrorSQLUnique(err error) bool {
	return strings.Contains(err.Error(), "violates unique constraint")
}

func CheckErrorSQLNotFound(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}

func CheckRelationNotExist(err error) bool {
	return strings.Contains(err.Error(), "does not exist (SQLSTATE 42P01)}")
}
