// Code generated by entc, DO NOT EDIT.

package launchtime

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the launchtime type in the database.
	Label = "launch_time"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldNetworkName holds the string denoting the network_name field in the database.
	FieldNetworkName = "network_name"
	// FieldNetworkVersion holds the string denoting the network_version field in the database.
	FieldNetworkVersion = "network_version"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// FieldLaunchTime holds the string denoting the launch_time field in the database.
	FieldLaunchTime = "launch_time"
	// FieldIncentive holds the string denoting the incentive field in the database.
	FieldIncentive = "incentive"
	// FieldIncentiveTotal holds the string denoting the incentive_total field in the database.
	FieldIncentiveTotal = "incentive_total"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the launchtime in the database.
	Table = "launch_times"
)

// Columns holds all SQL columns for launchtime fields.
var Columns = []string{
	FieldID,
	FieldProjectID,
	FieldNetworkName,
	FieldNetworkVersion,
	FieldIntroduction,
	FieldLaunchTime,
	FieldIncentive,
	FieldIncentiveTotal,
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
