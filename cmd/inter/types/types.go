package types

import (
	"database/sql"
	"time"
)

type Accounts struct {
	ID            sql.NullInt64 `json:"id"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Created       time.Time     `json:"created"`
	EmailID       string        `json:"email_ID"`
	AccountType   string        `json:"account_type"`
	Password      string        `json:"password"`
	EmailVarified bool          `json:"email_varified"`
}
