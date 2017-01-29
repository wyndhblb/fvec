/*
   Base Vector Interface Object methods

   The base objects structs are in vepr.proto and generated
*/

package vepr

import (
	"strings"
)

// VectorType interface
type VectorType interface {

	// CassandraCreateType string for the create type (if nessesary)
	// the string will be blank if no create is needed
	CassandraCreateType(keyspace string) string

	// CassandraType the matching types in cassandra for the tuple
	CassandraType() string
}

// Vector interface
type Vector interface {
	// IsVector
	IsVector() bool

	// CassandraCreateType create a cassandra type query/string
	CassandraCreateType(keyspace string) string

	// CassandraType the type in cassandra
	CassandraType() string

	// Key key name
	Key() string

	// GetName get the name object
	GetName() *VName

	// UniqueId unique id of the vector
	UniqueId() uint64

	// UniqueIdString unique id as a base36 string of the vector
	UniqueIdString() string

	// TypeName the type name of the vector
	TypeName() string
}

type VectorMap interface {
	Vector
	IsMap() bool
}

type VectorList interface {
	Vector
	IsList() bool
}

type VectorSet interface {
	Vector
	IsSet() bool
}

// list objects sorts by the key name
type VectorSlice []Vector

func (p VectorSlice) Len() int { return len(p) }
func (p VectorSlice) Less(i, j int) bool {
	return strings.Compare(p[i].GetName().Key, p[j].GetName().Key) < 0
}
func (p VectorSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// GetVector
// based on a short hand notation l,s,m (list, set, map) + combos of s,i,d (string, int, double)
// grab the appropriate vector object
// for instance basetype == l and subtype == ss  --> VLStrStr
func GetVector(basetype byte, subtypes string) Vector {
	switch basetype {
	case 'l':
		switch subtypes {
		case "s":
			return new(VLStr)
		case "i":
			return new(VLInt)
		case "d":
			return new(VLDbl)
		case "ss":
			return new(VLStrStr)

		case "si":
			return new(VLStrInt)
		case "sd":
			return new(VLStrDbl)
		case "is":
			return new(VLIntStr)
		case "id":
			return new(VLIntDbl)
		case "ii":
			return new(VLIntInt)
		case "ds":
			return new(VLDblStr)
		case "di":
			return new(VLDblInt)
		case "dd":
			return new(VLDblDbl)
		default:
			return nil
		}
	case 's':
		switch subtypes {
		case "s":
			return new(VSStr)
		case "i":
			return new(VSInt)
		case "d":
			return new(VSDbl)
		case "ss":
			return new(VSStrStr)

		case "si":
			return new(VSStrInt)
		case "sd":
			return new(VSStrDbl)
		case "is":
			return new(VSIntStr)
		case "id":
			return new(VSIntDbl)
		case "ii":
			return new(VSIntInt)
		case "ds":
			return new(VSDblStr)
		case "di":
			return new(VSDblInt)
		case "dd":
			return new(VSDblDbl)
		default:
			return nil
		}
	case 'm':
		switch subtypes {
		case "ss":
			return new(VMStrStr)
		case "si":
			return new(VMStrInt)
		case "sd":
			return new(VMStrDbl)
		case "is":
			return new(VMIntStr)
		case "id":
			return new(VMIntDbl)
		case "ii":
			return new(VMIntInt)

		case "sss":
			return new(VMStrTPStrStr)
		case "ssi":
			return new(VMStrTPStrInt)
		case "ssd":
			return new(VMStrTPStrDbl)
		case "sis":
			return new(VMStrTPIntStr)
		case "sii":
			return new(VMStrTPIntInt)
		case "sid":
			return new(VMStrTPIntDbl)
		case "sds":
			return new(VMStrTPDblStr)
		case "sdi":
			return new(VMStrTPDblInt)
		case "sdd":
			return new(VMStrTPDblDbl)

		case "iss":
			return new(VMIntTPStrStr)
		case "isi":
			return new(VMIntTPStrInt)
		case "isd":
			return new(VMIntTPStrDbl)
		case "iis":
			return new(VMIntTPIntStr)
		case "iii":
			return new(VMIntTPIntInt)
		case "iid":
			return new(VMIntTPIntDbl)
		case "ids":
			return new(VMIntTPDblStr)
		case "idi":
			return new(VMIntTPDblInt)
		case "idd":
			return new(VMIntTPDblDbl)

		default:
			return nil
		}
	}
	return nil
}

// GetVectorFromString returns a vector from a string
// the string can be either the short form (i.e. msss or object names VMStrTPStrStr
// will return nil if nothing matches
func GetVectorFromString(vec string) Vector {
	return stringToVector(vec)
}

// GetVectorFromCassType
// based on the cassandra "column" type get the vector object
func GetVectorFromCassType(casstype string) Vector {
	// remove all spaces
	casstype = strings.Replace(casstype, " ", "", -1)
	if gots, ok := CASSANDRA_TYPE_MAP[casstype]; ok {
		return stringToVector(gots)
	}
	return nil
}
