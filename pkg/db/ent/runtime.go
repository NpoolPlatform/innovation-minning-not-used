// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/capital"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/schema"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/team"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	capitalFields := schema.Capital{}.Fields()
	_ = capitalFields
	// capitalDescCreateAt is the schema descriptor for create_at field.
	capitalDescCreateAt := capitalFields[4].Descriptor()
	// capital.DefaultCreateAt holds the default value on creation for the create_at field.
	capital.DefaultCreateAt = capitalDescCreateAt.Default.(func() uint32)
	// capitalDescUpdateAt is the schema descriptor for update_at field.
	capitalDescUpdateAt := capitalFields[5].Descriptor()
	// capital.DefaultUpdateAt holds the default value on creation for the update_at field.
	capital.DefaultUpdateAt = capitalDescUpdateAt.Default.(func() uint32)
	// capital.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	capital.UpdateDefaultUpdateAt = capitalDescUpdateAt.UpdateDefault.(func() uint32)
	// capitalDescDeleteAt is the schema descriptor for delete_at field.
	capitalDescDeleteAt := capitalFields[6].Descriptor()
	// capital.DefaultDeleteAt holds the default value on creation for the delete_at field.
	capital.DefaultDeleteAt = capitalDescDeleteAt.Default.(func() uint32)
	// capitalDescID is the schema descriptor for id field.
	capitalDescID := capitalFields[0].Descriptor()
	// capital.DefaultID holds the default value on creation for the id field.
	capital.DefaultID = capitalDescID.Default.(func() uuid.UUID)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreateAt is the schema descriptor for create_at field.
	teamDescCreateAt := teamFields[7].Descriptor()
	// team.DefaultCreateAt holds the default value on creation for the create_at field.
	team.DefaultCreateAt = teamDescCreateAt.Default.(func() uint32)
	// teamDescUpdateAt is the schema descriptor for update_at field.
	teamDescUpdateAt := teamFields[8].Descriptor()
	// team.DefaultUpdateAt holds the default value on creation for the update_at field.
	team.DefaultUpdateAt = teamDescUpdateAt.Default.(func() uint32)
	// team.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	team.UpdateDefaultUpdateAt = teamDescUpdateAt.UpdateDefault.(func() uint32)
	// teamDescDeleteAt is the schema descriptor for delete_at field.
	teamDescDeleteAt := teamFields[9].Descriptor()
	// team.DefaultDeleteAt holds the default value on creation for the delete_at field.
	team.DefaultDeleteAt = teamDescDeleteAt.Default.(func() uint32)
	// teamDescID is the schema descriptor for id field.
	teamDescID := teamFields[0].Descriptor()
	// team.DefaultID holds the default value on creation for the id field.
	team.DefaultID = teamDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateAt is the schema descriptor for create_at field.
	userDescCreateAt := userFields[4].Descriptor()
	// user.DefaultCreateAt holds the default value on creation for the create_at field.
	user.DefaultCreateAt = userDescCreateAt.Default.(func() uint32)
	// userDescUpdateAt is the schema descriptor for update_at field.
	userDescUpdateAt := userFields[5].Descriptor()
	// user.DefaultUpdateAt holds the default value on creation for the update_at field.
	user.DefaultUpdateAt = userDescUpdateAt.Default.(func() uint32)
	// user.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	user.UpdateDefaultUpdateAt = userDescUpdateAt.UpdateDefault.(func() uint32)
	// userDescDeleteAt is the schema descriptor for delete_at field.
	userDescDeleteAt := userFields[6].Descriptor()
	// user.DefaultDeleteAt holds the default value on creation for the delete_at field.
	user.DefaultDeleteAt = userDescDeleteAt.Default.(func() uint32)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
