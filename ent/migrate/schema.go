// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "price", Type: field.TypeFloat32},
		{Name: "stock", Type: field.TypeInt},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// PasswordsColumns holds the columns for the "passwords" table.
	PasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "password", Type: field.TypeString},
		{Name: "user_password", Type: field.TypeInt, Nullable: true},
	}
	// PasswordsTable holds the schema information for the "passwords" table.
	PasswordsTable = &schema.Table{
		Name:       "passwords",
		Columns:    PasswordsColumns,
		PrimaryKey: []*schema.Column{PasswordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "passwords_users_password",
				Columns:    []*schema.Column{PasswordsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CartsTable,
		ItemsTable,
		PasswordsTable,
		UsersTable,
	}
)

func init() {
	PasswordsTable.ForeignKeys[0].RefTable = UsersTable
}
