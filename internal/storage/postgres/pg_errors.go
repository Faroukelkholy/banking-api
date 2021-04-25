package postgres

import "strings"

var (
	NoCustomerID = `no customer available with this id`
	NoAccountID  = `no account available with this id`
)

func IsViolateFK(strErr string) bool {
	return strings.Contains(strErr, "#23503")
}
