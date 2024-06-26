// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DouyinsColumns holds the columns for the "douyins" table.
	DouyinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// DouyinsTable holds the schema information for the "douyins" table.
	DouyinsTable = &schema.Table{
		Name:       "douyins",
		Columns:    DouyinsColumns,
		PrimaryKey: []*schema.Column{DouyinsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DouyinsTable,
	}
)

func init() {
}
