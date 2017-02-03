/*
   Base Scalar objects to match
*/

package fvec

// Scalar (the basic elements)
type Scalar interface {
	// IsScalar
	IsScalar() bool
	// CassandraCreateType string for the create type (if nessesary)
	// the string will be blank if no create is needed
	CassandraCreateType(keyspace string) string

	// CassandraType the matching types in cassandra for the tuple
	CassandraType() string

	// JavaType the string of the "type" in Java
	JavaType() string

	// GoType the string name of the golang type
	GoType() string
}

type DoubleType float64
type FloatType float32
type IntType int64
type CounterType int64
type StringType string
type ByteType []byte

func (d *DoubleType) CassandraCreateType(keyspace string) string {
	return ""
}
func (d *FloatType) CassandraCreateType(keyspace string) string {
	return ""
}
func (d *IntType) CassandraCreateType(keyspace string) string {
	return ""
}
func (d *StringType) CassandraCreateType(keyspace string) string {
	return ""
}
func (d *CounterType) CassandraCreateType(keyspace string) string {
	return ""
}
func (d *ByteType) CassandraCreateType(keyspace string) string {
	return ""
}

func (d *DoubleType) CassandraType() string {
	return "double"
}
func (d *FloatType) CassandraType() string {
	return "float"
}
func (d *IntType) CassandraType() string {
	return "varint"
}
func (d *StringType) CassandraType() string {
	return "varchar"
}
func (d *CounterType) CassandraType() string {
	return "counter"
}
func (d *ByteType) CassandraType() string {
	return "blob"
}

func (d *DoubleType) JavaType() string {
	return "Double"
}

func (d *FloatType) JavaType() string {
	return "Float"
}
func (d *IntType) JavaType() string {
	return "Integer"
}
func (d *CounterType) JavaType() string {
	return "Integer"
}
func (d *StringType) JavaType() string {
	return "String"
}
func (d *ByteType) JavaType() string {
	return "byte[]"
}

func (d *DoubleType) GoType() string {
	return "DoubleType"
}
func (d *FloatType) GoType() string {
	return "FloatType"
}
func (d *IntType) GoType() string {
	return "IntType"
}
func (d *CounterType) GoType() string {
	return "CounterType"
}
func (d *StringType) GoType() string {
	return "StringType"
}
func (d *ByteType) GoType() string {
	return "ByteType"
}

func (d *DoubleType) IsScalar() bool {
	return true
}
func (d *FloatType) IsScalar() bool {
	return true
}
func (d *IntType) IsScalar() bool {
	return true
}
func (d *CounterType) IsScalar() bool {
	return true
}
func (d *StringType) IsScalar() bool {
	return true
}
func (d *ByteType) IsScalar() bool {
	return true
}

var SCALAR_SHORT_NAME_MAP = map[string]string{
	"i": "IntType",
	"d": "DoubleType",
	"f": "FloatType",
	"b": "ByteType",
	"s": "StringType",
	"c": "CounterType",
}

// IsCounter counters are a special kind of thing in cassandra
// which will require it's own table separate from the main vector table
func IsCounter(vec string) bool {
	switch vec {
	case "ct", "c", "counter", "CounterType":
		return true
	default:
		return false
	}
}

// GetScalarFromString returns a vector from a string
// the string can be either the short form (i.e. msss or object names VMStrTPStrStr
// will return nil if nothing matches
func GetScalarFromString(vec string) Scalar {
	switch vec {
	case "ct", "c", "counter", "CounterType":
		return new(CounterType)
	case "d", "double", "dbl", "DoubleType":
		return new(DoubleType)
	case "f", "float", "flt", "FloatType":
		return new(FloatType)
	case "b", "bytes", "byte", "ByteType":
		return new(ByteType)
	case "s", "str", "string", "varchar", "chars", "StringType":
		return new(StringType)
	case "i", "int", "int64", "integer", "IntType":
		return new(IntType)
	}
	return nil
}

// GetScalarFromCassType
// based on the cassandra "column" type get the vector object
func GetScalarFromCassType(casstype string) Scalar {
	switch casstype {
	case "counter":
		return new(CounterType)
	case "double", "decimal":
		return new(DoubleType)
	case "float":
		return new(FloatType)
	case "blob":
		return new(ByteType)
	case "varchar", "text", "ascii":
		return new(StringType)
	case "int", "bigint", "varint":
		return new(IntType)
	}
	return nil
}
