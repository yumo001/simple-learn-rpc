// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yumo001/simple-learn-rpc/ent/xaddress"
)

// XAddress is the model entity for the XAddress schema.
type XAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Default holds the value of the "default" field.
	Default int64 `json:"default,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// CountryID holds the value of the "country_id" field.
	CountryID int64 `json:"country_id,omitempty"`
	// Street holds the value of the "street" field.
	Street string `json:"street,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone        string `json:"phone,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*XAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case xaddress.FieldID, xaddress.FieldUserID, xaddress.FieldDefault, xaddress.FieldCountryID:
			values[i] = new(sql.NullInt64)
		case xaddress.FieldFirstName, xaddress.FieldLastName, xaddress.FieldStreet, xaddress.FieldProvince, xaddress.FieldCity, xaddress.FieldPostalCode, xaddress.FieldPhone:
			values[i] = new(sql.NullString)
		case xaddress.FieldCreatedAt, xaddress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the XAddress fields.
func (x *XAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case xaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			x.ID = uint64(value.Int64)
		case xaddress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				x.CreatedAt = value.Time
			}
		case xaddress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				x.UpdatedAt = value.Time
			}
		case xaddress.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				x.UserID = value.Int64
			}
		case xaddress.FieldDefault:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field default", values[i])
			} else if value.Valid {
				x.Default = value.Int64
			}
		case xaddress.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				x.FirstName = value.String
			}
		case xaddress.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				x.LastName = value.String
			}
		case xaddress.FieldCountryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field country_id", values[i])
			} else if value.Valid {
				x.CountryID = value.Int64
			}
		case xaddress.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street", values[i])
			} else if value.Valid {
				x.Street = value.String
			}
		case xaddress.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				x.Province = value.String
			}
		case xaddress.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				x.City = value.String
			}
		case xaddress.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				x.PostalCode = value.String
			}
		case xaddress.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				x.Phone = value.String
			}
		default:
			x.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the XAddress.
// This includes values selected through modifiers, order, etc.
func (x *XAddress) Value(name string) (ent.Value, error) {
	return x.selectValues.Get(name)
}

// Update returns a builder for updating this XAddress.
// Note that you need to call XAddress.Unwrap() before calling this method if this XAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (x *XAddress) Update() *XAddressUpdateOne {
	return NewXAddressClient(x.config).UpdateOne(x)
}

// Unwrap unwraps the XAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (x *XAddress) Unwrap() *XAddress {
	_tx, ok := x.config.driver.(*txDriver)
	if !ok {
		panic("ent: XAddress is not a transactional entity")
	}
	x.config.driver = _tx.drv
	return x
}

// String implements the fmt.Stringer.
func (x *XAddress) String() string {
	var builder strings.Builder
	builder.WriteString("XAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", x.ID))
	builder.WriteString("created_at=")
	builder.WriteString(x.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(x.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", x.UserID))
	builder.WriteString(", ")
	builder.WriteString("default=")
	builder.WriteString(fmt.Sprintf("%v", x.Default))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(x.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(x.LastName)
	builder.WriteString(", ")
	builder.WriteString("country_id=")
	builder.WriteString(fmt.Sprintf("%v", x.CountryID))
	builder.WriteString(", ")
	builder.WriteString("street=")
	builder.WriteString(x.Street)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(x.Province)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(x.City)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(x.PostalCode)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(x.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// XAddresses is a parsable slice of XAddress.
type XAddresses []*XAddress
