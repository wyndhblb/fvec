// GENERATED with python3 genproto.py
// Edit at your own risk

package vepr

import "reflect"
import "fmt"

//**************** Tuple: double double VTDblDbl **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTDblDbl) CassandraType() string {
	return "VTDblDbl"
}

// RedisInsertCmd returns the redis add command
func (t *VTDblDbl) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTDblDbl) RedisValue() string {
	return fmt.Sprintf("%v %v", t.Key, t.Value)
}

//**************** Tuple: double int64 VTDblInt **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTDblInt) CassandraType() string {
	return "VTDblInt"
}

// RedisInsertCmd returns the redis add command
func (t *VTDblInt) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTDblInt) RedisValue() string {
	return fmt.Sprintf("%v %d", t.Key, t.Value)
}

//**************** Tuple: double string VTDblStr **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTDblStr) CassandraType() string {
	return "VTDblStr"
}

// RedisInsertCmd returns the redis add command
func (t *VTDblStr) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTDblStr) RedisValue() string {
	return fmt.Sprintf("%v %s", t.Key, t.Value)
}

//**************** Tuple: int64 double VTIntDbl **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTIntDbl) CassandraType() string {
	return "VTIntDbl"
}

// RedisInsertCmd returns the redis add command
func (t *VTIntDbl) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTIntDbl) RedisValue() string {
	return fmt.Sprintf("%d %v", t.Key, t.Value)
}

//**************** Tuple: int64 int64 VTIntInt **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTIntInt) CassandraType() string {
	return "VTIntInt"
}

// RedisInsertCmd returns the redis add command
func (t *VTIntInt) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTIntInt) RedisValue() string {
	return fmt.Sprintf("%d %d", t.Key, t.Value)
}

//**************** Tuple: int64 string VTIntStr **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTIntStr) CassandraType() string {
	return "VTIntStr"
}

// RedisInsertCmd returns the redis add command
func (t *VTIntStr) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTIntStr) RedisValue() string {
	return fmt.Sprintf("%d %s", t.Key, t.Value)
}

//**************** Tuple: string double VTStrDbl **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTStrDbl) CassandraType() string {
	return "VTStrDbl"
}

// RedisInsertCmd returns the redis add command
func (t *VTStrDbl) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTStrDbl) RedisValue() string {
	return fmt.Sprintf("%s %v", t.Key, t.Value)
}

//**************** Tuple: string int64 VTStrInt **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTStrInt) CassandraType() string {
	return "VTStrInt"
}

// RedisInsertCmd returns the redis add command
func (t *VTStrInt) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTStrInt) RedisValue() string {
	return fmt.Sprintf("%s %d", t.Key, t.Value)
}

//**************** Tuple: string string VTStrStr **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VTStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the tuple
func (t *VTStrStr) CassandraType() string {
	return "VTStrStr"
}

// RedisInsertCmd returns the redis add command
func (t *VTStrStr) RedisInsertCmd(key string) string {
	return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VTStrStr) RedisValue() string {
	return fmt.Sprintf("%s %s", t.Key, t.Value)
}

//**************** Set: double  VSDbl **********************/

// VecType the type of vector (set)
func (t *VSDbl) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSDbl) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDbl) TypeName() string {
	return "VSDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VSDbl) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the set
func (t *VSDbl) CassandraType() string {
	return "set<double>"
}

//**************** Set: int64  VSInt **********************/

// VecType the type of vector (set)
func (t *VSInt) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSInt) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSInt) TypeName() string {
	return "VSInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VSInt) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the set
func (t *VSInt) CassandraType() string {
	return "set<bigint>"
}

//**************** Set: string  VSStr **********************/

// VecType the type of vector (set)
func (t *VSStr) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSStr) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStr) TypeName() string {
	return "VSStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VSStr) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the set
func (t *VSStr) CassandraType() string {
	return "set<varchar>"
}

//**************** Set: double double VSDblDbl **********************/

// VecType the type of vector (set)
func (t *VSDblDbl) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSDblDbl) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSDblDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSDblDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSDblDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSDblDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSDblDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblDbl) TypeName() string {
	return "VSDblDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSDblDbl) CassandraType() string {
	return "set<frozen<VTDblDbl>>"
}

//**************** Set: double int64 VSDblInt **********************/

// VecType the type of vector (set)
func (t *VSDblInt) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSDblInt) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSDblInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSDblInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSDblInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSDblInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSDblInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblInt) TypeName() string {
	return "VSDblInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSDblInt) CassandraType() string {
	return "set<frozen<VTDblInt>>"
}

//**************** Set: double string VSDblStr **********************/

// VecType the type of vector (set)
func (t *VSDblStr) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSDblStr) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSDblStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSDblStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSDblStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSDblStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSDblStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblStr) TypeName() string {
	return "VSDblStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSDblStr) CassandraType() string {
	return "set<frozen<VTDblStr>>"
}

//**************** Set: int64 double VSIntDbl **********************/

// VecType the type of vector (set)
func (t *VSIntDbl) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSIntDbl) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSIntDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSIntDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSIntDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSIntDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSIntDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntDbl) TypeName() string {
	return "VSIntDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSIntDbl) CassandraType() string {
	return "set<frozen<VTIntDbl>>"
}

//**************** Set: int64 int64 VSIntInt **********************/

// VecType the type of vector (set)
func (t *VSIntInt) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSIntInt) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSIntInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSIntInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSIntInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSIntInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSIntInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntInt) TypeName() string {
	return "VSIntInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSIntInt) CassandraType() string {
	return "set<frozen<VTIntInt>>"
}

//**************** Set: int64 string VSIntStr **********************/

// VecType the type of vector (set)
func (t *VSIntStr) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSIntStr) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSIntStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSIntStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSIntStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSIntStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSIntStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntStr) TypeName() string {
	return "VSIntStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSIntStr) CassandraType() string {
	return "set<frozen<VTIntStr>>"
}

//**************** Set: string double VSStrDbl **********************/

// VecType the type of vector (set)
func (t *VSStrDbl) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSStrDbl) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSStrDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSStrDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSStrDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSStrDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSStrDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrDbl) TypeName() string {
	return "VSStrDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSStrDbl) CassandraType() string {
	return "set<frozen<VTStrDbl>>"
}

//**************** Set: string int64 VSStrInt **********************/

// VecType the type of vector (set)
func (t *VSStrInt) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSStrInt) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSStrInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSStrInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSStrInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSStrInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSStrInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrInt) TypeName() string {
	return "VSStrInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSStrInt) CassandraType() string {
	return "set<frozen<VTStrInt>>"
}

//**************** Set: string string VSStrStr **********************/

// VecType the type of vector (set)
func (t *VSStrStr) VecType() string {
	return "set"
}

// IsSet is a set type
func (t *VSStrStr) IsSet() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VSStrStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VSStrStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VSStrStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VSStrStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VSStrStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VSStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrStr) TypeName() string {
	return "VSStrStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VSStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the set
func (t *VSStrStr) CassandraType() string {
	return "set<frozen<VTStrStr>>"
}

//**************** List: double  VLDbl **********************/

// VecType the type of vector (list)
func (t *VLDbl) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLDbl) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDbl) TypeName() string {
	return "VLDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VLDbl) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLDbl) CassandraType() string {
	return "list<double>"
}

//**************** List: int64  VLInt **********************/

// VecType the type of vector (list)
func (t *VLInt) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLInt) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLInt) TypeName() string {
	return "VLInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VLInt) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLInt) CassandraType() string {
	return "list<bigint>"
}

//**************** List: string  VLStr **********************/

// VecType the type of vector (list)
func (t *VLStr) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLStr) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStr) TypeName() string {
	return "VLStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *VLStr) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLStr) CassandraType() string {
	return "list<varchar>"
}

//**************** List: double double VLDblDbl **********************/

// VecType the type of vector (list)
func (t *VLDblDbl) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLDblDbl) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLDblDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLDblDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLDblDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLDblDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLDblDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblDbl) TypeName() string {
	return "VLDblDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblDbl) CassandraType() string {
	return "list<frozen<VTDblDbl>>"
}

//**************** List: double int64 VLDblInt **********************/

// VecType the type of vector (list)
func (t *VLDblInt) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLDblInt) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLDblInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLDblInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLDblInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLDblInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLDblInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblInt) TypeName() string {
	return "VLDblInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblInt) CassandraType() string {
	return "list<frozen<VTDblInt>>"
}

//**************** List: double string VLDblStr **********************/

// VecType the type of vector (list)
func (t *VLDblStr) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLDblStr) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLDblStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLDblStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLDblStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLDblStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLDblStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblStr) TypeName() string {
	return "VLDblStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblStr) CassandraType() string {
	return "list<frozen<VTDblStr>>"
}

//**************** List: int64 double VLIntDbl **********************/

// VecType the type of vector (list)
func (t *VLIntDbl) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLIntDbl) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLIntDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLIntDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLIntDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLIntDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLIntDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntDbl) TypeName() string {
	return "VLIntDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntDbl) CassandraType() string {
	return "list<frozen<VTIntDbl>>"
}

//**************** List: int64 int64 VLIntInt **********************/

// VecType the type of vector (list)
func (t *VLIntInt) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLIntInt) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLIntInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLIntInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLIntInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLIntInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLIntInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntInt) TypeName() string {
	return "VLIntInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntInt) CassandraType() string {
	return "list<frozen<VTIntInt>>"
}

//**************** List: int64 string VLIntStr **********************/

// VecType the type of vector (list)
func (t *VLIntStr) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLIntStr) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLIntStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLIntStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLIntStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLIntStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLIntStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntStr) TypeName() string {
	return "VLIntStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntStr) CassandraType() string {
	return "list<frozen<VTIntStr>>"
}

//**************** List: string double VLStrDbl **********************/

// VecType the type of vector (list)
func (t *VLStrDbl) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLStrDbl) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLStrDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLStrDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLStrDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLStrDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLStrDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrDbl) TypeName() string {
	return "VLStrDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrDbl) CassandraType() string {
	return "list<frozen<VTStrDbl>>"
}

//**************** List: string int64 VLStrInt **********************/

// VecType the type of vector (list)
func (t *VLStrInt) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLStrInt) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLStrInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLStrInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLStrInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLStrInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLStrInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrInt) TypeName() string {
	return "VLStrInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrInt) CassandraType() string {
	return "list<frozen<VTStrInt>>"
}

//**************** List: string string VLStrStr **********************/

// VecType the type of vector (list)
func (t *VLStrStr) VecType() string {
	return "list"
}

// IsList this is a list type
func (t *VLStrStr) IsList() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VLStrStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VLStrStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VLStrStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VLStrStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VLStrStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VLStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrStr) TypeName() string {
	return "VLStrStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrStr) CassandraType() string {
	return "list<frozen<VTStrStr>>"
}

//**************** Map: int64 double VMIntDbl **********************/

// VecType the type of vector (map)
func (t *VMIntDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntDbl) TypeName() string {
	return "VMIntDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntDbl) CassandraType() string {
	return "map<bigint, frozen<VTIntDbl>>"
}

//**************** Map: int64 int64 VMIntInt **********************/

// VecType the type of vector (map)
func (t *VMIntInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntInt) TypeName() string {
	return "VMIntInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntInt) CassandraType() string {
	return "map<bigint, frozen<VTIntInt>>"
}

//**************** Map: int64 string VMIntStr **********************/

// VecType the type of vector (map)
func (t *VMIntStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntStr) TypeName() string {
	return "VMIntStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntStr) CassandraType() string {
	return "map<bigint, frozen<VTIntStr>>"
}

//**************** Map: string double VMStrDbl **********************/

// VecType the type of vector (map)
func (t *VMStrDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrDbl) TypeName() string {
	return "VMStrDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrDbl) CassandraType() string {
	return "map<varchar, frozen<VTStrDbl>>"
}

//**************** Map: string int64 VMStrInt **********************/

// VecType the type of vector (map)
func (t *VMStrInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrInt) TypeName() string {
	return "VMStrInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrInt) CassandraType() string {
	return "map<varchar, frozen<VTStrInt>>"
}

//**************** Map: string string VMStrStr **********************/

// VecType the type of vector (map)
func (t *VMStrStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrStr) TypeName() string {
	return "VMStrStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrStr) CassandraType() string {
	return "map<varchar, frozen<VTStrStr>>"
}

//**************** Map: double double VMIntTPDblDbl **********************/

// VecType the type of vector (map)
func (t *VMIntTPDblDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPDblDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPDblDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPDblDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPDblDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPDblDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPDblDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblDbl) TypeName() string {
	return "VMIntTPDblDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblDbl) CassandraType() string {
	return "map<double, frozen<VTDblDbl>>"
}

//**************** Map: double int64 VMIntTPDblInt **********************/

// VecType the type of vector (map)
func (t *VMIntTPDblInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPDblInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPDblInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPDblInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPDblInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPDblInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPDblInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblInt) TypeName() string {
	return "VMIntTPDblInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblInt) CassandraType() string {
	return "map<double, frozen<VTDblInt>>"
}

//**************** Map: double string VMIntTPDblStr **********************/

// VecType the type of vector (map)
func (t *VMIntTPDblStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPDblStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPDblStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPDblStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPDblStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPDblStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPDblStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblStr) TypeName() string {
	return "VMIntTPDblStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblStr) CassandraType() string {
	return "map<double, frozen<VTDblStr>>"
}

//**************** Map: int64 double VMIntTPIntDbl **********************/

// VecType the type of vector (map)
func (t *VMIntTPIntDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPIntDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPIntDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPIntDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPIntDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPIntDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPIntDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntDbl) TypeName() string {
	return "VMIntTPIntDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntDbl) CassandraType() string {
	return "map<bigint, frozen<VTIntDbl>>"
}

//**************** Map: int64 int64 VMIntTPIntInt **********************/

// VecType the type of vector (map)
func (t *VMIntTPIntInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPIntInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPIntInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPIntInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPIntInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPIntInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPIntInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntInt) TypeName() string {
	return "VMIntTPIntInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntInt) CassandraType() string {
	return "map<bigint, frozen<VTIntInt>>"
}

//**************** Map: int64 string VMIntTPIntStr **********************/

// VecType the type of vector (map)
func (t *VMIntTPIntStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPIntStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPIntStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPIntStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPIntStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPIntStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPIntStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntStr) TypeName() string {
	return "VMIntTPIntStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntStr) CassandraType() string {
	return "map<bigint, frozen<VTIntStr>>"
}

//**************** Map: string double VMIntTPStrDbl **********************/

// VecType the type of vector (map)
func (t *VMIntTPStrDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPStrDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPStrDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPStrDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPStrDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPStrDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPStrDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrDbl) TypeName() string {
	return "VMIntTPStrDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrDbl) CassandraType() string {
	return "map<varchar, frozen<VTStrDbl>>"
}

//**************** Map: string int64 VMIntTPStrInt **********************/

// VecType the type of vector (map)
func (t *VMIntTPStrInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPStrInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPStrInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPStrInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPStrInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPStrInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPStrInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrInt) TypeName() string {
	return "VMIntTPStrInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrInt) CassandraType() string {
	return "map<varchar, frozen<VTStrInt>>"
}

//**************** Map: string string VMIntTPStrStr **********************/

// VecType the type of vector (map)
func (t *VMIntTPStrStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMIntTPStrStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMIntTPStrStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMIntTPStrStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMIntTPStrStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMIntTPStrStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMIntTPStrStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMIntTPStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrStr) TypeName() string {
	return "VMIntTPStrStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrStr) CassandraType() string {
	return "map<varchar, frozen<VTStrStr>>"
}

//**************** Map: double double VMStrTPDblDbl **********************/

// VecType the type of vector (map)
func (t *VMStrTPDblDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPDblDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPDblDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPDblDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPDblDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPDblDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPDblDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblDbl) TypeName() string {
	return "VMStrTPDblDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblDbl) CassandraType() string {
	return "map<double, frozen<VTDblDbl>>"
}

//**************** Map: double int64 VMStrTPDblInt **********************/

// VecType the type of vector (map)
func (t *VMStrTPDblInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPDblInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPDblInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPDblInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPDblInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPDblInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPDblInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblInt) TypeName() string {
	return "VMStrTPDblInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblInt) CassandraType() string {
	return "map<double, frozen<VTDblInt>>"
}

//**************** Map: double string VMStrTPDblStr **********************/

// VecType the type of vector (map)
func (t *VMStrTPDblStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPDblStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPDblStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPDblStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPDblStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPDblStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPDblStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblStr) TypeName() string {
	return "VMStrTPDblStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblStr) CassandraType() string {
	return "map<double, frozen<VTDblStr>>"
}

//**************** Map: int64 double VMStrTPIntDbl **********************/

// VecType the type of vector (map)
func (t *VMStrTPIntDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPIntDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPIntDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPIntDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPIntDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPIntDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPIntDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntDbl) TypeName() string {
	return "VMStrTPIntDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntDbl) CassandraType() string {
	return "map<bigint, frozen<VTIntDbl>>"
}

//**************** Map: int64 int64 VMStrTPIntInt **********************/

// VecType the type of vector (map)
func (t *VMStrTPIntInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPIntInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPIntInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPIntInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPIntInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPIntInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPIntInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntInt) TypeName() string {
	return "VMStrTPIntInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntInt) CassandraType() string {
	return "map<bigint, frozen<VTIntInt>>"
}

//**************** Map: int64 string VMStrTPIntStr **********************/

// VecType the type of vector (map)
func (t *VMStrTPIntStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPIntStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPIntStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPIntStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPIntStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPIntStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPIntStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntStr) TypeName() string {
	return "VMStrTPIntStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntStr) CassandraType() string {
	return "map<bigint, frozen<VTIntStr>>"
}

//**************** Map: string double VMStrTPStrDbl **********************/

// VecType the type of vector (map)
func (t *VMStrTPStrDbl) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPStrDbl) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPStrDbl) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPStrDbl) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPStrDbl) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPStrDbl) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPStrDbl) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrDbl) TypeName() string {
	return "VMStrTPStrDbl"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrDbl) CassandraType() string {
	return "map<varchar, frozen<VTStrDbl>>"
}

//**************** Map: string int64 VMStrTPStrInt **********************/

// VecType the type of vector (map)
func (t *VMStrTPStrInt) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPStrInt) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPStrInt) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPStrInt) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPStrInt) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPStrInt) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPStrInt) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrInt) TypeName() string {
	return "VMStrTPStrInt"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrInt) CassandraType() string {
	return "map<varchar, frozen<VTStrInt>>"
}

//**************** Map: string string VMStrTPStrStr **********************/

// VecType the type of vector (map)
func (t *VMStrTPStrStr) VecType() string {
	return "map"
}

// IsMap is a set type
func (t *VMStrTPStrStr) IsMap() bool {
	return true
}

// GetName returns the Name of the vector
func (t *VMStrTPStrStr) GetName() *VName {
	return t.Name
}

// Key returns the key of the vector
func (t *VMStrTPStrStr) Key() string {
	return t.Name.Key
}

// Tags returns the tags of the vector
func (t *VMStrTPStrStr) Tags() Tags {
	return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *VMStrTPStrStr) UniqueId() uint64 {
	return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *VMStrTPStrStr) UniqueIdString() string {
	return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *VMStrTPStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrStr) TypeName() string {
	return "VMStrTPStrStr"
}

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrStr) CassandraType() string {
	return "map<varchar, frozen<VTStrStr>>"
}

var TupleNames = [...]string{"VTDblDbl", "VTDblInt", "VTDblStr", "VTIntDbl", "VTIntInt", "VTIntStr", "VTStrDbl", "VTStrInt", "VTStrStr"}
var ListNames = [...]string{"VLDbl", "VLInt", "VLStr", "VLDblDbl", "VLDblInt", "VLDblStr", "VLIntDbl", "VLIntInt", "VLIntStr", "VLStrDbl", "VLStrInt", "VLStrStr"}
var SetNames = [...]string{"VSDbl", "VSInt", "VSStr", "VSDblDbl", "VSDblInt", "VSDblStr", "VSIntDbl", "VSIntInt", "VSIntStr", "VSStrDbl", "VSStrInt", "VSStrStr"}
var MapNames = [...]string{"VMIntDbl", "VMIntInt", "VMIntStr", "VMStrDbl", "VMStrInt", "VMStrStr", "VMIntTPDblDbl", "VMIntTPDblInt", "VMIntTPDblStr", "VMIntTPIntDbl", "VMIntTPIntInt", "VMIntTPIntStr", "VMIntTPStrDbl", "VMIntTPStrInt", "VMIntTPStrStr", "VMStrTPDblDbl", "VMStrTPDblInt", "VMStrTPDblStr", "VMStrTPIntDbl", "VMStrTPIntInt", "VMStrTPIntStr", "VMStrTPStrDbl", "VMStrTPStrInt", "VMStrTPStrStr"}
var SHORT_NAME_MAP = map[string]string{
	"ld":   "VLDbl",
	"ldd":  "VLDblDbl",
	"ldi":  "VLDblInt",
	"lds":  "VLDblStr",
	"li":   "VLInt",
	"lid":  "VLIntDbl",
	"lii":  "VLIntInt",
	"lis":  "VLIntStr",
	"ls":   "VLStr",
	"lsd":  "VLStrDbl",
	"lsi":  "VLStrInt",
	"lss":  "VLStrStr",
	"mid":  "VMIntDbl",
	"midd": "VMIntTPDblDbl",
	"midi": "VMIntTPDblInt",
	"mids": "VMIntTPDblStr",
	"mii":  "VMIntInt",
	"miid": "VMIntTPIntDbl",
	"miii": "VMIntTPIntInt",
	"miis": "VMIntTPIntStr",
	"mis":  "VMIntStr",
	"misd": "VMIntTPStrDbl",
	"misi": "VMIntTPStrInt",
	"miss": "VMIntTPStrStr",
	"msd":  "VMStrDbl",
	"msdd": "VMStrTPDblDbl",
	"msdi": "VMStrTPDblInt",
	"msds": "VMStrTPDblStr",
	"msi":  "VMStrInt",
	"msid": "VMStrTPIntDbl",
	"msii": "VMStrTPIntInt",
	"msis": "VMStrTPIntStr",
	"mss":  "VMStrStr",
	"mssd": "VMStrTPStrDbl",
	"mssi": "VMStrTPStrInt",
	"msss": "VMStrTPStrStr",
	"sd":   "VSDbl",
	"sdd":  "VSDblDbl",
	"sdi":  "VSDblInt",
	"sds":  "VSDblStr",
	"si":   "VSInt",
	"sid":  "VSIntDbl",
	"sii":  "VSIntInt",
	"sis":  "VSIntStr",
	"ss":   "VSStr",
	"ssd":  "VSStrDbl",
	"ssi":  "VSStrInt",
	"sss":  "VSStrStr",
}
var CASSANDRA_TYPE_MAP = map[string]string{
	"list<bigint>":                  "VLInt",
	"list<double>":                  "VLDbl",
	"list<frozen<VTDblDbl>>":        "VLDblDbl",
	"list<frozen<VTDblInt>>":        "VLDblInt",
	"list<frozen<VTDblStr>>":        "VLDblStr",
	"list<frozen<VTIntDbl>>":        "VLIntDbl",
	"list<frozen<VTIntInt>>":        "VLIntInt",
	"list<frozen<VTIntStr>>":        "VLIntStr",
	"list<frozen<VTStrDbl>>":        "VLStrDbl",
	"list<frozen<VTStrInt>>":        "VLStrInt",
	"list<frozen<VTStrStr>>":        "VLStrStr",
	"list<varchar>":                 "VLStr",
	"map<bigint,frozen<VTIntDbl>>":  "VMStrTPIntDbl",
	"map<bigint,frozen<VTIntInt>>":  "VMStrTPIntInt",
	"map<bigint,frozen<VTIntStr>>":  "VMStrTPIntStr",
	"map<double,frozen<VTDblDbl>>":  "VMStrTPDblDbl",
	"map<double,frozen<VTDblInt>>":  "VMStrTPDblInt",
	"map<double,frozen<VTDblStr>>":  "VMStrTPDblStr",
	"map<varchar,frozen<VTStrDbl>>": "VMStrTPStrDbl",
	"map<varchar,frozen<VTStrInt>>": "VMStrTPStrInt",
	"map<varchar,frozen<VTStrStr>>": "VMStrTPStrStr",
	"set<bigint>":                   "VSInt",
	"set<double>":                   "VSDbl",
	"set<frozen<VTDblDbl>>":         "VSDblDbl",
	"set<frozen<VTDblInt>>":         "VSDblInt",
	"set<frozen<VTDblStr>>":         "VSDblStr",
	"set<frozen<VTIntDbl>>":         "VSIntDbl",
	"set<frozen<VTIntInt>>":         "VSIntInt",
	"set<frozen<VTIntStr>>":         "VSIntStr",
	"set<frozen<VTStrDbl>>":         "VSStrDbl",
	"set<frozen<VTStrInt>>":         "VSStrInt",
	"set<frozen<VTStrStr>>":         "VSStrStr",
	"set<varchar>":                  "VSStr",
}

func stringToVector(vec string) Vector {
	switch vec {
	case "ld":
		return new(VLDbl)
	case "VLDbl":
		return new(VLDbl)
	case "ldd":
		return new(VLDblDbl)
	case "VLDblDbl":
		return new(VLDblDbl)
	case "ldi":
		return new(VLDblInt)
	case "VLDblInt":
		return new(VLDblInt)
	case "lds":
		return new(VLDblStr)
	case "VLDblStr":
		return new(VLDblStr)
	case "li":
		return new(VLInt)
	case "VLInt":
		return new(VLInt)
	case "lid":
		return new(VLIntDbl)
	case "VLIntDbl":
		return new(VLIntDbl)
	case "lii":
		return new(VLIntInt)
	case "VLIntInt":
		return new(VLIntInt)
	case "lis":
		return new(VLIntStr)
	case "VLIntStr":
		return new(VLIntStr)
	case "ls":
		return new(VLStr)
	case "VLStr":
		return new(VLStr)
	case "lsd":
		return new(VLStrDbl)
	case "VLStrDbl":
		return new(VLStrDbl)
	case "lsi":
		return new(VLStrInt)
	case "VLStrInt":
		return new(VLStrInt)
	case "lss":
		return new(VLStrStr)
	case "VLStrStr":
		return new(VLStrStr)
	case "mid":
		return new(VMIntDbl)
	case "VMIntDbl":
		return new(VMIntDbl)
	case "midd":
		return new(VMIntTPDblDbl)
	case "VMIntTPDblDbl":
		return new(VMIntTPDblDbl)
	case "midi":
		return new(VMIntTPDblInt)
	case "VMIntTPDblInt":
		return new(VMIntTPDblInt)
	case "mids":
		return new(VMIntTPDblStr)
	case "VMIntTPDblStr":
		return new(VMIntTPDblStr)
	case "mii":
		return new(VMIntInt)
	case "VMIntInt":
		return new(VMIntInt)
	case "miid":
		return new(VMIntTPIntDbl)
	case "VMIntTPIntDbl":
		return new(VMIntTPIntDbl)
	case "miii":
		return new(VMIntTPIntInt)
	case "VMIntTPIntInt":
		return new(VMIntTPIntInt)
	case "miis":
		return new(VMIntTPIntStr)
	case "VMIntTPIntStr":
		return new(VMIntTPIntStr)
	case "mis":
		return new(VMIntStr)
	case "VMIntStr":
		return new(VMIntStr)
	case "misd":
		return new(VMIntTPStrDbl)
	case "VMIntTPStrDbl":
		return new(VMIntTPStrDbl)
	case "misi":
		return new(VMIntTPStrInt)
	case "VMIntTPStrInt":
		return new(VMIntTPStrInt)
	case "miss":
		return new(VMIntTPStrStr)
	case "VMIntTPStrStr":
		return new(VMIntTPStrStr)
	case "msd":
		return new(VMStrDbl)
	case "VMStrDbl":
		return new(VMStrDbl)
	case "msdd":
		return new(VMStrTPDblDbl)
	case "VMStrTPDblDbl":
		return new(VMStrTPDblDbl)
	case "msdi":
		return new(VMStrTPDblInt)
	case "VMStrTPDblInt":
		return new(VMStrTPDblInt)
	case "msds":
		return new(VMStrTPDblStr)
	case "VMStrTPDblStr":
		return new(VMStrTPDblStr)
	case "msi":
		return new(VMStrInt)
	case "VMStrInt":
		return new(VMStrInt)
	case "msid":
		return new(VMStrTPIntDbl)
	case "VMStrTPIntDbl":
		return new(VMStrTPIntDbl)
	case "msii":
		return new(VMStrTPIntInt)
	case "VMStrTPIntInt":
		return new(VMStrTPIntInt)
	case "msis":
		return new(VMStrTPIntStr)
	case "VMStrTPIntStr":
		return new(VMStrTPIntStr)
	case "mss":
		return new(VMStrStr)
	case "VMStrStr":
		return new(VMStrStr)
	case "mssd":
		return new(VMStrTPStrDbl)
	case "VMStrTPStrDbl":
		return new(VMStrTPStrDbl)
	case "mssi":
		return new(VMStrTPStrInt)
	case "VMStrTPStrInt":
		return new(VMStrTPStrInt)
	case "msss":
		return new(VMStrTPStrStr)
	case "VMStrTPStrStr":
		return new(VMStrTPStrStr)
	case "sd":
		return new(VSDbl)
	case "VSDbl":
		return new(VSDbl)
	case "sdd":
		return new(VSDblDbl)
	case "VSDblDbl":
		return new(VSDblDbl)
	case "sdi":
		return new(VSDblInt)
	case "VSDblInt":
		return new(VSDblInt)
	case "sds":
		return new(VSDblStr)
	case "VSDblStr":
		return new(VSDblStr)
	case "si":
		return new(VSInt)
	case "VSInt":
		return new(VSInt)
	case "sid":
		return new(VSIntDbl)
	case "VSIntDbl":
		return new(VSIntDbl)
	case "sii":
		return new(VSIntInt)
	case "VSIntInt":
		return new(VSIntInt)
	case "sis":
		return new(VSIntStr)
	case "VSIntStr":
		return new(VSIntStr)
	case "ss":
		return new(VSStr)
	case "VSStr":
		return new(VSStr)
	case "ssd":
		return new(VSStrDbl)
	case "VSStrDbl":
		return new(VSStrDbl)
	case "ssi":
		return new(VSStrInt)
	case "VSStrInt":
		return new(VSStrInt)
	case "sss":
		return new(VSStrStr)
	case "VSStrStr":
		return new(VSStrStr)
	}
	return nil
}

var typeRegistry = make(map[string]reflect.Type)

func init() {
	typeRegistry["ld"] = reflect.TypeOf(VLDbl{})
	typeRegistry["VLDbl"] = reflect.TypeOf(VLDbl{})
	typeRegistry["ldd"] = reflect.TypeOf(VLDblDbl{})
	typeRegistry["VLDblDbl"] = reflect.TypeOf(VLDblDbl{})
	typeRegistry["ldi"] = reflect.TypeOf(VLDblInt{})
	typeRegistry["VLDblInt"] = reflect.TypeOf(VLDblInt{})
	typeRegistry["lds"] = reflect.TypeOf(VLDblStr{})
	typeRegistry["VLDblStr"] = reflect.TypeOf(VLDblStr{})
	typeRegistry["li"] = reflect.TypeOf(VLInt{})
	typeRegistry["VLInt"] = reflect.TypeOf(VLInt{})
	typeRegistry["lid"] = reflect.TypeOf(VLIntDbl{})
	typeRegistry["VLIntDbl"] = reflect.TypeOf(VLIntDbl{})
	typeRegistry["lii"] = reflect.TypeOf(VLIntInt{})
	typeRegistry["VLIntInt"] = reflect.TypeOf(VLIntInt{})
	typeRegistry["lis"] = reflect.TypeOf(VLIntStr{})
	typeRegistry["VLIntStr"] = reflect.TypeOf(VLIntStr{})
	typeRegistry["ls"] = reflect.TypeOf(VLStr{})
	typeRegistry["VLStr"] = reflect.TypeOf(VLStr{})
	typeRegistry["lsd"] = reflect.TypeOf(VLStrDbl{})
	typeRegistry["VLStrDbl"] = reflect.TypeOf(VLStrDbl{})
	typeRegistry["lsi"] = reflect.TypeOf(VLStrInt{})
	typeRegistry["VLStrInt"] = reflect.TypeOf(VLStrInt{})
	typeRegistry["lss"] = reflect.TypeOf(VLStrStr{})
	typeRegistry["VLStrStr"] = reflect.TypeOf(VLStrStr{})
	typeRegistry["mid"] = reflect.TypeOf(VMIntDbl{})
	typeRegistry["VMIntDbl"] = reflect.TypeOf(VMIntDbl{})
	typeRegistry["midd"] = reflect.TypeOf(VMIntTPDblDbl{})
	typeRegistry["VMIntTPDblDbl"] = reflect.TypeOf(VMIntTPDblDbl{})
	typeRegistry["midi"] = reflect.TypeOf(VMIntTPDblInt{})
	typeRegistry["VMIntTPDblInt"] = reflect.TypeOf(VMIntTPDblInt{})
	typeRegistry["mids"] = reflect.TypeOf(VMIntTPDblStr{})
	typeRegistry["VMIntTPDblStr"] = reflect.TypeOf(VMIntTPDblStr{})
	typeRegistry["mii"] = reflect.TypeOf(VMIntInt{})
	typeRegistry["VMIntInt"] = reflect.TypeOf(VMIntInt{})
	typeRegistry["miid"] = reflect.TypeOf(VMIntTPIntDbl{})
	typeRegistry["VMIntTPIntDbl"] = reflect.TypeOf(VMIntTPIntDbl{})
	typeRegistry["miii"] = reflect.TypeOf(VMIntTPIntInt{})
	typeRegistry["VMIntTPIntInt"] = reflect.TypeOf(VMIntTPIntInt{})
	typeRegistry["miis"] = reflect.TypeOf(VMIntTPIntStr{})
	typeRegistry["VMIntTPIntStr"] = reflect.TypeOf(VMIntTPIntStr{})
	typeRegistry["mis"] = reflect.TypeOf(VMIntStr{})
	typeRegistry["VMIntStr"] = reflect.TypeOf(VMIntStr{})
	typeRegistry["misd"] = reflect.TypeOf(VMIntTPStrDbl{})
	typeRegistry["VMIntTPStrDbl"] = reflect.TypeOf(VMIntTPStrDbl{})
	typeRegistry["misi"] = reflect.TypeOf(VMIntTPStrInt{})
	typeRegistry["VMIntTPStrInt"] = reflect.TypeOf(VMIntTPStrInt{})
	typeRegistry["miss"] = reflect.TypeOf(VMIntTPStrStr{})
	typeRegistry["VMIntTPStrStr"] = reflect.TypeOf(VMIntTPStrStr{})
	typeRegistry["msd"] = reflect.TypeOf(VMStrDbl{})
	typeRegistry["VMStrDbl"] = reflect.TypeOf(VMStrDbl{})
	typeRegistry["msdd"] = reflect.TypeOf(VMStrTPDblDbl{})
	typeRegistry["VMStrTPDblDbl"] = reflect.TypeOf(VMStrTPDblDbl{})
	typeRegistry["msdi"] = reflect.TypeOf(VMStrTPDblInt{})
	typeRegistry["VMStrTPDblInt"] = reflect.TypeOf(VMStrTPDblInt{})
	typeRegistry["msds"] = reflect.TypeOf(VMStrTPDblStr{})
	typeRegistry["VMStrTPDblStr"] = reflect.TypeOf(VMStrTPDblStr{})
	typeRegistry["msi"] = reflect.TypeOf(VMStrInt{})
	typeRegistry["VMStrInt"] = reflect.TypeOf(VMStrInt{})
	typeRegistry["msid"] = reflect.TypeOf(VMStrTPIntDbl{})
	typeRegistry["VMStrTPIntDbl"] = reflect.TypeOf(VMStrTPIntDbl{})
	typeRegistry["msii"] = reflect.TypeOf(VMStrTPIntInt{})
	typeRegistry["VMStrTPIntInt"] = reflect.TypeOf(VMStrTPIntInt{})
	typeRegistry["msis"] = reflect.TypeOf(VMStrTPIntStr{})
	typeRegistry["VMStrTPIntStr"] = reflect.TypeOf(VMStrTPIntStr{})
	typeRegistry["mss"] = reflect.TypeOf(VMStrStr{})
	typeRegistry["VMStrStr"] = reflect.TypeOf(VMStrStr{})
	typeRegistry["mssd"] = reflect.TypeOf(VMStrTPStrDbl{})
	typeRegistry["VMStrTPStrDbl"] = reflect.TypeOf(VMStrTPStrDbl{})
	typeRegistry["mssi"] = reflect.TypeOf(VMStrTPStrInt{})
	typeRegistry["VMStrTPStrInt"] = reflect.TypeOf(VMStrTPStrInt{})
	typeRegistry["msss"] = reflect.TypeOf(VMStrTPStrStr{})
	typeRegistry["VMStrTPStrStr"] = reflect.TypeOf(VMStrTPStrStr{})
	typeRegistry["sd"] = reflect.TypeOf(VSDbl{})
	typeRegistry["VSDbl"] = reflect.TypeOf(VSDbl{})
	typeRegistry["sdd"] = reflect.TypeOf(VSDblDbl{})
	typeRegistry["VSDblDbl"] = reflect.TypeOf(VSDblDbl{})
	typeRegistry["sdi"] = reflect.TypeOf(VSDblInt{})
	typeRegistry["VSDblInt"] = reflect.TypeOf(VSDblInt{})
	typeRegistry["sds"] = reflect.TypeOf(VSDblStr{})
	typeRegistry["VSDblStr"] = reflect.TypeOf(VSDblStr{})
	typeRegistry["si"] = reflect.TypeOf(VSInt{})
	typeRegistry["VSInt"] = reflect.TypeOf(VSInt{})
	typeRegistry["sid"] = reflect.TypeOf(VSIntDbl{})
	typeRegistry["VSIntDbl"] = reflect.TypeOf(VSIntDbl{})
	typeRegistry["sii"] = reflect.TypeOf(VSIntInt{})
	typeRegistry["VSIntInt"] = reflect.TypeOf(VSIntInt{})
	typeRegistry["sis"] = reflect.TypeOf(VSIntStr{})
	typeRegistry["VSIntStr"] = reflect.TypeOf(VSIntStr{})
	typeRegistry["ss"] = reflect.TypeOf(VSStr{})
	typeRegistry["VSStr"] = reflect.TypeOf(VSStr{})
	typeRegistry["ssd"] = reflect.TypeOf(VSStrDbl{})
	typeRegistry["VSStrDbl"] = reflect.TypeOf(VSStrDbl{})
	typeRegistry["ssi"] = reflect.TypeOf(VSStrInt{})
	typeRegistry["VSStrInt"] = reflect.TypeOf(VSStrInt{})
	typeRegistry["sss"] = reflect.TypeOf(VSStrStr{})
	typeRegistry["VSStrStr"] = reflect.TypeOf(VSStrStr{})
}
