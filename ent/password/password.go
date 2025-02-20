// Code generated by ent, DO NOT EDIT.

package password

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the password type in the database.
	Label = "password"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// Table holds the table name of the password in the database.
	Table = "passwords"
)

// Columns holds all SQL columns for password fields.
var Columns = []string{
	FieldID,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "passwords"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_password",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Password queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}
