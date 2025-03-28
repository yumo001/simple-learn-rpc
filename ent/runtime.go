// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/yumo001/simple-learn-rpc/ent/schema"
	"github.com/yumo001/simple-learn-rpc/ent/xaddress"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	xaddressMixin := schema.XAddress{}.Mixin()
	xaddressMixinFields0 := xaddressMixin[0].Fields()
	_ = xaddressMixinFields0
	xaddressFields := schema.XAddress{}.Fields()
	_ = xaddressFields
	// xaddressDescCreatedAt is the schema descriptor for created_at field.
	xaddressDescCreatedAt := xaddressMixinFields0[1].Descriptor()
	// xaddress.DefaultCreatedAt holds the default value on creation for the created_at field.
	xaddress.DefaultCreatedAt = xaddressDescCreatedAt.Default.(func() time.Time)
	// xaddressDescUpdatedAt is the schema descriptor for updated_at field.
	xaddressDescUpdatedAt := xaddressMixinFields0[2].Descriptor()
	// xaddress.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	xaddress.DefaultUpdatedAt = xaddressDescUpdatedAt.Default.(func() time.Time)
	// xaddress.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	xaddress.UpdateDefaultUpdatedAt = xaddressDescUpdatedAt.UpdateDefault.(func() time.Time)
}
