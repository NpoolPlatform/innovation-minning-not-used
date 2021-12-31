// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CapitalsColumns holds the columns for the "capitals" table.
	CapitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "introduction", Type: field.TypeString},
		{Name: "logo", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CapitalsTable holds the schema information for the "capitals" table.
	CapitalsTable = &schema.Table{
		Name:       "capitals",
		Columns:    CapitalsColumns,
		PrimaryKey: []*schema.Column{CapitalsColumns[0]},
	}
	// LaunchTimesColumns holds the columns for the "launch_times" table.
	LaunchTimesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "project_id", Type: field.TypeUUID, Unique: true},
		{Name: "network_name", Type: field.TypeString},
		{Name: "network_version", Type: field.TypeString},
		{Name: "introduction", Type: field.TypeString},
		{Name: "launch_time", Type: field.TypeUint32},
		{Name: "incentive", Type: field.TypeBool},
		{Name: "incentive_total", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// LaunchTimesTable holds the schema information for the "launch_times" table.
	LaunchTimesTable = &schema.Table{
		Name:       "launch_times",
		Columns:    LaunchTimesColumns,
		PrimaryKey: []*schema.Column{LaunchTimesColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "team_id", Type: field.TypeUUID},
		{Name: "capital_ids", Type: field.TypeJSON},
		{Name: "introduction", Type: field.TypeString},
		{Name: "logo", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "team_name", Type: field.TypeString, Unique: true},
		{Name: "logo", Type: field.TypeString},
		{Name: "leader_id", Type: field.TypeUUID},
		{Name: "member_ids", Type: field.TypeJSON},
		{Name: "introduction", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// TechniqueAnalysesColumns holds the columns for the "technique_analyses" table.
	TechniqueAnalysesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "author_id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "project_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// TechniqueAnalysesTable holds the schema information for the "technique_analyses" table.
	TechniqueAnalysesTable = &schema.Table{
		Name:       "technique_analyses",
		Columns:    TechniqueAnalysesColumns,
		PrimaryKey: []*schema.Column{TechniqueAnalysesColumns[0]},
	}
	// TrendAnalysesColumns holds the columns for the "trend_analyses" table.
	TrendAnalysesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "project_id", Type: field.TypeUUID},
		{Name: "author_id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// TrendAnalysesTable holds the schema information for the "trend_analyses" table.
	TrendAnalysesTable = &schema.Table{
		Name:       "trend_analyses",
		Columns:    TrendAnalysesColumns,
		PrimaryKey: []*schema.Column{TrendAnalysesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "introduction", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_first_name_last_name",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CapitalsTable,
		LaunchTimesTable,
		ProjectsTable,
		TeamsTable,
		TechniqueAnalysesTable,
		TrendAnalysesTable,
		UsersTable,
	}
)

func init() {
}
