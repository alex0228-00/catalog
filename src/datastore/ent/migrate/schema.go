// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SystemsColumns holds the columns for the "systems" table.
	SystemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "host", Type: field.TypeString},
		{Name: "unique_identifier", Type: field.TypeString},
		{Name: "credentials", Type: field.TypeString, Size: 2147483647},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_by", Type: field.TypeString, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// SystemsTable holds the schema information for the "systems" table.
	SystemsTable = &schema.Table{
		Name:       "systems",
		Columns:    SystemsColumns,
		PrimaryKey: []*schema.Column{SystemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SystemsTable,
	}
)

func init() {
}
