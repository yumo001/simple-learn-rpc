package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type XAddress struct {
	ent.Schema
}

func (XAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("default"),
		field.String("first_name"),
		field.String("last_name"),
		field.Int64("country_id"),
		field.String("street"),
		field.String("province"),
		field.String("city"),
		field.String("postal_code"),
		field.String("phone"),
	}
}

func (XAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

func (XAddress) Edges() []ent.Edge {
	return nil
}
