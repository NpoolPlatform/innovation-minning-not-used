// Code generated by entc, DO NOT EDIT.

package team

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTeamName holds the string denoting the team_name field in the database.
	FieldTeamName = "team_name"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldLeaderID holds the string denoting the leader_id field in the database.
	FieldLeaderID = "leader_id"
	// FieldMemberIds holds the string denoting the member_ids field in the database.
	FieldMemberIds = "member_ids"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the team in the database.
	Table = "teams"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldTeamName,
	FieldLogo,
	FieldLeaderID,
	FieldMemberIds,
	FieldIntroduction,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
