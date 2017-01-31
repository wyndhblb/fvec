// GENERATED with python3 genproto.py
// Edit at your own risk

package fvec

import "reflect"
import "fmt"
import "strings"
import "errors"
import "strconv"

var ErrorInvalidRedisValue = errors.New("Invalid redis value, cannot parse")

func stringToFloat(tp string) (float64, error) {
	return strconv.ParseFloat(tp, 64)
}

func stringToInt(tp string) (int64, error) {
	return strconv.ParseInt(tp, 10, 64)
}

func stringToString(tp string) (string, error) {
	return tp, nil
}

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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTDblDbl) JavaType() string {
	return "Pair<Double,Double>"
}

// IsVector more for interface acceptance
func (t *VTDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTDblDbl) TypeName() string {
	return "VTDblDbl"
}

// GoType the type of object in go
func (t *VTDblDbl) GoType() string {
	return "VTDblDbl"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTDblStr) JavaType() string {
	return "Pair<Double,String>"
}

// IsVector more for interface acceptance
func (t *VTDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTDblStr) TypeName() string {
	return "VTDblStr"
}

// GoType the type of object in go
func (t *VTDblStr) GoType() string {
	return "VTDblStr"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTDblInt) JavaType() string {
	return "Pair<Double,Integer>"
}

// IsVector more for interface acceptance
func (t *VTDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTDblInt) TypeName() string {
	return "VTDblInt"
}

// GoType the type of object in go
func (t *VTDblInt) GoType() string {
	return "VTDblInt"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTStrDbl) JavaType() string {
	return "Pair<String,Double>"
}

// IsVector more for interface acceptance
func (t *VTStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTStrDbl) TypeName() string {
	return "VTStrDbl"
}

// GoType the type of object in go
func (t *VTStrDbl) GoType() string {
	return "VTStrDbl"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTStrStr) JavaType() string {
	return "Pair<String,String>"
}

// IsVector more for interface acceptance
func (t *VTStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTStrStr) TypeName() string {
	return "VTStrStr"
}

// GoType the type of object in go
func (t *VTStrStr) GoType() string {
	return "VTStrStr"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTStrInt) JavaType() string {
	return "Pair<String,Integer>"
}

// IsVector more for interface acceptance
func (t *VTStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTStrInt) TypeName() string {
	return "VTStrInt"
}

// GoType the type of object in go
func (t *VTStrInt) GoType() string {
	return "VTStrInt"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTIntDbl) JavaType() string {
	return "Pair<Integer,Double>"
}

// IsVector more for interface acceptance
func (t *VTIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTIntDbl) TypeName() string {
	return "VTIntDbl"
}

// GoType the type of object in go
func (t *VTIntDbl) GoType() string {
	return "VTIntDbl"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTIntStr) JavaType() string {
	return "Pair<Integer,String>"
}

// IsVector more for interface acceptance
func (t *VTIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTIntStr) TypeName() string {
	return "VTIntStr"
}

// GoType the type of object in go
func (t *VTIntStr) GoType() string {
	return "VTIntStr"
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

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *VTIntInt) JavaType() string {
	return "Pair<Integer,Integer>"
}

// IsVector more for interface acceptance
func (t *VTIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VTIntInt) TypeName() string {
	return "VTIntInt"
}

// GoType the type of object in go
func (t *VTIntInt) GoType() string {
	return "VTIntInt"
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

// IsVector more for interface acceptance
func (t *VSDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDbl) TypeName() string {
	return "VSDbl"
}

// GoType the type of object in go
func (t *VSDbl) GoType() string {
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

// JavaType the type of object in java
// java.util
func (t *VSDbl) JavaType() string {
	return "Set<Double>"
}

// RedisInsertCmd returns the redis add command
func (t *VSDbl) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSDbl) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSDbl) RedisValue(v float64) string {
	return fmt.Sprintf("%v", v)
}

// FromRedisValue given the redis value, make it into a proper float64
func (t *VSDbl) FromRedisValue(i string) (v float64, err error) {
	return stringToFloat(i)
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

// IsVector more for interface acceptance
func (t *VSStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStr) TypeName() string {
	return "VSStr"
}

// GoType the type of object in go
func (t *VSStr) GoType() string {
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

// JavaType the type of object in java
// java.util
func (t *VSStr) JavaType() string {
	return "Set<String>"
}

// RedisInsertCmd returns the redis add command
func (t *VSStr) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSStr) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSStr) RedisValue(v string) string {
	return fmt.Sprintf("%s", v)
}

// FromRedisValue given the redis value, make it into a proper string
func (t *VSStr) FromRedisValue(i string) (v string, err error) {
	return stringToString(i)
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

// IsVector more for interface acceptance
func (t *VSInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSInt) TypeName() string {
	return "VSInt"
}

// GoType the type of object in go
func (t *VSInt) GoType() string {
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

// JavaType the type of object in java
// java.util
func (t *VSInt) JavaType() string {
	return "Set<Integer>"
}

// RedisInsertCmd returns the redis add command
func (t *VSInt) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSInt) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSInt) RedisValue(v int64) string {
	return fmt.Sprintf("%d", v)
}

// FromRedisValue given the redis value, make it into a proper int64
func (t *VSInt) FromRedisValue(i string) (v int64, err error) {
	return stringToInt(i)
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

// IsVector more for interface acceptance
func (t *VSDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblDbl) TypeName() string {
	return "VSDblDbl"
}

// GoType the type of object in go
func (t *VSDblDbl) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSDblDbl) JavaType() string {
	return "Set<Pair<Double,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSDblDbl) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSDblDbl) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSDblDbl) RedisInsertValue(v VTDblDbl) string {
	return fmt.Sprintf("%v:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSDblDbl) FromRedisValue(i string) (v VTDblDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblStr) TypeName() string {
	return "VSDblStr"
}

// GoType the type of object in go
func (t *VSDblStr) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSDblStr) JavaType() string {
	return "Set<Pair<Double,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSDblStr) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSDblStr) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSDblStr) RedisInsertValue(v VTDblStr) string {
	return fmt.Sprintf("%v:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSDblStr) FromRedisValue(i string) (v VTDblStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSDblInt) TypeName() string {
	return "VSDblInt"
}

// GoType the type of object in go
func (t *VSDblInt) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSDblInt) JavaType() string {
	return "Set<Pair<Double,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSDblInt) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSDblInt) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSDblInt) RedisInsertValue(v VTDblInt) string {
	return fmt.Sprintf("%v:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSDblInt) FromRedisValue(i string) (v VTDblInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrDbl) TypeName() string {
	return "VSStrDbl"
}

// GoType the type of object in go
func (t *VSStrDbl) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSStrDbl) JavaType() string {
	return "Set<Pair<String,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSStrDbl) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSStrDbl) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSStrDbl) RedisInsertValue(v VTStrDbl) string {
	return fmt.Sprintf("%s:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSStrDbl) FromRedisValue(i string) (v VTStrDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrStr) TypeName() string {
	return "VSStrStr"
}

// GoType the type of object in go
func (t *VSStrStr) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSStrStr) JavaType() string {
	return "Set<Pair<String,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSStrStr) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSStrStr) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSStrStr) RedisInsertValue(v VTStrStr) string {
	return fmt.Sprintf("%s:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSStrStr) FromRedisValue(i string) (v VTStrStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSStrInt) TypeName() string {
	return "VSStrInt"
}

// GoType the type of object in go
func (t *VSStrInt) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSStrInt) JavaType() string {
	return "Set<Pair<String,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSStrInt) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSStrInt) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSStrInt) RedisInsertValue(v VTStrInt) string {
	return fmt.Sprintf("%s:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSStrInt) FromRedisValue(i string) (v VTStrInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntDbl) TypeName() string {
	return "VSIntDbl"
}

// GoType the type of object in go
func (t *VSIntDbl) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSIntDbl) JavaType() string {
	return "Set<Pair<Integer,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSIntDbl) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSIntDbl) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSIntDbl) RedisInsertValue(v VTIntDbl) string {
	return fmt.Sprintf("%d:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSIntDbl) FromRedisValue(i string) (v VTIntDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntStr) TypeName() string {
	return "VSIntStr"
}

// GoType the type of object in go
func (t *VSIntStr) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSIntStr) JavaType() string {
	return "Set<Pair<Integer,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSIntStr) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSIntStr) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSIntStr) RedisInsertValue(v VTIntStr) string {
	return fmt.Sprintf("%d:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSIntStr) FromRedisValue(i string) (v VTIntStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VSIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VSIntInt) TypeName() string {
	return "VSIntInt"
}

// GoType the type of object in go
func (t *VSIntInt) GoType() string {
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

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *VSIntInt) JavaType() string {
	return "Set<Pair<Integer,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VSIntInt) RedisInsertCmd(key string) string {
	return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *VSIntInt) RedisRemoveCmd(key string) string {
	return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VSIntInt) RedisInsertValue(v VTIntInt) string {
	return fmt.Sprintf("%d:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VSIntInt) FromRedisValue(i string) (v VTIntInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDbl) TypeName() string {
	return "VLDbl"
}

// GoType the type of object in go
func (t *VLDbl) GoType() string {
	return "VLDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VLDbl) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLDbl) CassandraType() string {
	return "list<double>"
}

// JavaType the java type for the object
func (t *VLDbl) JavaType() string {
	return "List<Double>"
}

// RedisInsertCmd returns the redis add command
func (t *VLDbl) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLDbl) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLDbl) RedisInsertValue(v float64) string {
	return fmt.Sprintf("%v", v)
}

// FromRedisValue given the redis value, make it into a proper float64
func (t *VLDbl) FromRedisValue(i string) (v float64, err error) {
	return stringToFloat(i)
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

// IsVector more for interface acceptance
func (t *VLStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStr) TypeName() string {
	return "VLStr"
}

// GoType the type of object in go
func (t *VLStr) GoType() string {
	return "VLStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VLStr) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLStr) CassandraType() string {
	return "list<varchar>"
}

// JavaType the java type for the object
func (t *VLStr) JavaType() string {
	return "List<String>"
}

// RedisInsertCmd returns the redis add command
func (t *VLStr) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLStr) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLStr) RedisInsertValue(v string) string {
	return fmt.Sprintf("%s", v)
}

// FromRedisValue given the redis value, make it into a proper string
func (t *VLStr) FromRedisValue(i string) (v string, err error) {
	return stringToString(i)
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

// IsVector more for interface acceptance
func (t *VLInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLInt) TypeName() string {
	return "VLInt"
}

// GoType the type of object in go
func (t *VLInt) GoType() string {
	return "VLInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VLInt) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the list
func (t *VLInt) CassandraType() string {
	return "list<bigint>"
}

// JavaType the java type for the object
func (t *VLInt) JavaType() string {
	return "List<Integer>"
}

// RedisInsertCmd returns the redis add command
func (t *VLInt) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLInt) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLInt) RedisInsertValue(v int64) string {
	return fmt.Sprintf("%d", v)
}

// FromRedisValue given the redis value, make it into a proper int64
func (t *VLInt) FromRedisValue(i string) (v int64, err error) {
	return stringToInt(i)
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

// IsVector more for interface acceptance
func (t *VLDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblDbl) TypeName() string {
	return "VLDblDbl"
}

// GoType the type of object in go
func (t *VLDblDbl) GoType() string {
	return "VLDblDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblDbl) CassandraType() string {
	return "list<frozen<VTDblDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLDblDbl) JavaType() string {
	return "List<Pair<Double,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLDblDbl) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLDblDbl) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLDblDbl) RedisInsertValue(v VTDblDbl) string {
	return fmt.Sprintf("%v:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLDblDbl) FromRedisValue(i string) (v VTDblDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblStr) TypeName() string {
	return "VLDblStr"
}

// GoType the type of object in go
func (t *VLDblStr) GoType() string {
	return "VLDblStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblStr) CassandraType() string {
	return "list<frozen<VTDblStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLDblStr) JavaType() string {
	return "List<Pair<Double,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLDblStr) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLDblStr) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLDblStr) RedisInsertValue(v VTDblStr) string {
	return fmt.Sprintf("%v:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLDblStr) FromRedisValue(i string) (v VTDblStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLDblInt) TypeName() string {
	return "VLDblInt"
}

// GoType the type of object in go
func (t *VLDblInt) GoType() string {
	return "VLDblInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLDblInt) CassandraType() string {
	return "list<frozen<VTDblInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLDblInt) JavaType() string {
	return "List<Pair<Double,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLDblInt) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLDblInt) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLDblInt) RedisInsertValue(v VTDblInt) string {
	return fmt.Sprintf("%v:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLDblInt) FromRedisValue(i string) (v VTDblInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToFloat(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrDbl) TypeName() string {
	return "VLStrDbl"
}

// GoType the type of object in go
func (t *VLStrDbl) GoType() string {
	return "VLStrDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrDbl) CassandraType() string {
	return "list<frozen<VTStrDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLStrDbl) JavaType() string {
	return "List<Pair<String,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLStrDbl) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLStrDbl) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLStrDbl) RedisInsertValue(v VTStrDbl) string {
	return fmt.Sprintf("%s:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLStrDbl) FromRedisValue(i string) (v VTStrDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrStr) TypeName() string {
	return "VLStrStr"
}

// GoType the type of object in go
func (t *VLStrStr) GoType() string {
	return "VLStrStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrStr) CassandraType() string {
	return "list<frozen<VTStrStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLStrStr) JavaType() string {
	return "List<Pair<String,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLStrStr) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLStrStr) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLStrStr) RedisInsertValue(v VTStrStr) string {
	return fmt.Sprintf("%s:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLStrStr) FromRedisValue(i string) (v VTStrStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLStrInt) TypeName() string {
	return "VLStrInt"
}

// GoType the type of object in go
func (t *VLStrInt) GoType() string {
	return "VLStrInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLStrInt) CassandraType() string {
	return "list<frozen<VTStrInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLStrInt) JavaType() string {
	return "List<Pair<String,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLStrInt) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLStrInt) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLStrInt) RedisInsertValue(v VTStrInt) string {
	return fmt.Sprintf("%s:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLStrInt) FromRedisValue(i string) (v VTStrInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToString(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntDbl) TypeName() string {
	return "VLIntDbl"
}

// GoType the type of object in go
func (t *VLIntDbl) GoType() string {
	return "VLIntDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntDbl) CassandraType() string {
	return "list<frozen<VTIntDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLIntDbl) JavaType() string {
	return "List<Pair<Integer,Double>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLIntDbl) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLIntDbl) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLIntDbl) RedisInsertValue(v VTIntDbl) string {
	return fmt.Sprintf("%d:%v", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLIntDbl) FromRedisValue(i string) (v VTIntDbl, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToFloat(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntStr) TypeName() string {
	return "VLIntStr"
}

// GoType the type of object in go
func (t *VLIntStr) GoType() string {
	return "VLIntStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntStr) CassandraType() string {
	return "list<frozen<VTIntStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLIntStr) JavaType() string {
	return "List<Pair<Integer,String>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLIntStr) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLIntStr) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLIntStr) RedisInsertValue(v VTIntStr) string {
	return fmt.Sprintf("%d:%s", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLIntStr) FromRedisValue(i string) (v VTIntStr, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToString(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VLIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VLIntInt) TypeName() string {
	return "VLIntInt"
}

// GoType the type of object in go
func (t *VLIntInt) GoType() string {
	return "VLIntInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VLIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the list
func (t *VLIntInt) CassandraType() string {
	return "list<frozen<VTIntInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VLIntInt) JavaType() string {
	return "List<Pair<Integer,Integer>>"
}

// RedisInsertCmd returns the redis add command
func (t *VLIntInt) RedisInsertCmd(key string) string {
	return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VLIntInt) RedisRemoveCmd(key string) string {
	return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VLIntInt) RedisInsertValue(v VTIntInt) string {
	return fmt.Sprintf("%d:%d", v.Key, v.Value)
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *VLIntInt) FromRedisValue(i string) (v VTIntInt, err error) {
	spl := strings.Split(i, ":")
	if len(spl) != 2 {
		return v, ErrorInvalidRedisValue
	}
	v.Key, err = stringToInt(spl[0])
	v.Value, err = stringToInt(spl[1])
	return v, err
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

// IsVector more for interface acceptance
func (t *VMStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrDbl) TypeName() string {
	return "VMStrDbl"
}

// GoType the type of object in go
func (t *VMStrDbl) GoType() string {
	return "VMStrDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMStrDbl) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrDbl) CassandraType() string {
	return "map<varchar,double>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrDbl) JavaType() string {
	return "Map<String,Double>"
}

// RedisInsertCmd returns the redis add command
func (t *VMStrDbl) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMStrDbl) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMStrDbl) RedisInsertValue(v string) string {
	return fmt.Sprintf("%s", v)
}

// FromRedisValue given the redis value, make it into a proper string
func (t *VMStrDbl) FromRedisValue(i string) (v string, err error) {
	return stringToString(i)
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

// IsVector more for interface acceptance
func (t *VMStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrStr) TypeName() string {
	return "VMStrStr"
}

// GoType the type of object in go
func (t *VMStrStr) GoType() string {
	return "VMStrStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMStrStr) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrStr) CassandraType() string {
	return "map<varchar,varchar>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrStr) JavaType() string {
	return "Map<String,String>"
}

// RedisInsertCmd returns the redis add command
func (t *VMStrStr) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMStrStr) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMStrStr) RedisInsertValue(v string) string {
	return fmt.Sprintf("%s", v)
}

// FromRedisValue given the redis value, make it into a proper string
func (t *VMStrStr) FromRedisValue(i string) (v string, err error) {
	return stringToString(i)
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

// IsVector more for interface acceptance
func (t *VMStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrInt) TypeName() string {
	return "VMStrInt"
}

// GoType the type of object in go
func (t *VMStrInt) GoType() string {
	return "VMStrInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMStrInt) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrInt) CassandraType() string {
	return "map<varchar,bigint>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrInt) JavaType() string {
	return "Map<String,Integer>"
}

// RedisInsertCmd returns the redis add command
func (t *VMStrInt) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMStrInt) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMStrInt) RedisInsertValue(v string) string {
	return fmt.Sprintf("%s", v)
}

// FromRedisValue given the redis value, make it into a proper string
func (t *VMStrInt) FromRedisValue(i string) (v string, err error) {
	return stringToString(i)
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

// IsVector more for interface acceptance
func (t *VMIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntDbl) TypeName() string {
	return "VMIntDbl"
}

// GoType the type of object in go
func (t *VMIntDbl) GoType() string {
	return "VMIntDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMIntDbl) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntDbl) CassandraType() string {
	return "map<bigint,double>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntDbl) JavaType() string {
	return "Map<Integer,Double>"
}

// RedisInsertCmd returns the redis add command
func (t *VMIntDbl) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMIntDbl) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMIntDbl) RedisInsertValue(v int64) string {
	return fmt.Sprintf("%d", v)
}

// FromRedisValue given the redis value, make it into a proper int64
func (t *VMIntDbl) FromRedisValue(i string) (v int64, err error) {
	return stringToInt(i)
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

// IsVector more for interface acceptance
func (t *VMIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntStr) TypeName() string {
	return "VMIntStr"
}

// GoType the type of object in go
func (t *VMIntStr) GoType() string {
	return "VMIntStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMIntStr) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntStr) CassandraType() string {
	return "map<bigint,varchar>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntStr) JavaType() string {
	return "Map<Integer,String>"
}

// RedisInsertCmd returns the redis add command
func (t *VMIntStr) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMIntStr) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMIntStr) RedisInsertValue(v int64) string {
	return fmt.Sprintf("%d", v)
}

// FromRedisValue given the redis value, make it into a proper int64
func (t *VMIntStr) FromRedisValue(i string) (v int64, err error) {
	return stringToInt(i)
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

// IsVector more for interface acceptance
func (t *VMIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntInt) TypeName() string {
	return "VMIntInt"
}

// GoType the type of object in go
func (t *VMIntInt) GoType() string {
	return "VMIntInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *VMIntInt) CassandraCreateType(keyspace string) string {
	return ""
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntInt) CassandraType() string {
	return "map<bigint,bigint>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntInt) JavaType() string {
	return "Map<Integer,Integer>"
}

// RedisInsertCmd returns the redis add command
func (t *VMIntInt) RedisInsertCmd(key string) string {
	return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *VMIntInt) RedisRemoveCmd(key string) string {
	return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *VMIntInt) RedisInsertValue(v int64) string {
	return fmt.Sprintf("%d", v)
}

// FromRedisValue given the redis value, make it into a proper int64
func (t *VMIntInt) FromRedisValue(i string) (v int64, err error) {
	return stringToInt(i)
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

// IsVector more for interface acceptance
func (t *VMStrTPDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblDbl) TypeName() string {
	return "VMStrTPDblDbl"
}

// GoType the type of object in go
func (t *VMStrTPDblDbl) GoType() string {
	return "VMStrTPDblDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblDbl) CassandraType() string {
	return "map<varchar,frozen<VTDblDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPDblDbl) JavaType() string {
	return "Map<String,Pair<Double,Double>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblStr) TypeName() string {
	return "VMStrTPDblStr"
}

// GoType the type of object in go
func (t *VMStrTPDblStr) GoType() string {
	return "VMStrTPDblStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblStr) CassandraType() string {
	return "map<varchar,frozen<VTDblStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPDblStr) JavaType() string {
	return "Map<String,Pair<Double,String>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPDblInt) TypeName() string {
	return "VMStrTPDblInt"
}

// GoType the type of object in go
func (t *VMStrTPDblInt) GoType() string {
	return "VMStrTPDblInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPDblInt) CassandraType() string {
	return "map<varchar,frozen<VTDblInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPDblInt) JavaType() string {
	return "Map<String,Pair<Double,Integer>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrDbl) TypeName() string {
	return "VMStrTPStrDbl"
}

// GoType the type of object in go
func (t *VMStrTPStrDbl) GoType() string {
	return "VMStrTPStrDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrDbl) CassandraType() string {
	return "map<varchar,frozen<VTStrDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPStrDbl) JavaType() string {
	return "Map<String,Pair<String,Double>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrStr) TypeName() string {
	return "VMStrTPStrStr"
}

// GoType the type of object in go
func (t *VMStrTPStrStr) GoType() string {
	return "VMStrTPStrStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrStr) CassandraType() string {
	return "map<varchar,frozen<VTStrStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPStrStr) JavaType() string {
	return "Map<String,Pair<String,String>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPStrInt) TypeName() string {
	return "VMStrTPStrInt"
}

// GoType the type of object in go
func (t *VMStrTPStrInt) GoType() string {
	return "VMStrTPStrInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPStrInt) CassandraType() string {
	return "map<varchar,frozen<VTStrInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPStrInt) JavaType() string {
	return "Map<String,Pair<String,Integer>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntDbl) TypeName() string {
	return "VMStrTPIntDbl"
}

// GoType the type of object in go
func (t *VMStrTPIntDbl) GoType() string {
	return "VMStrTPIntDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntDbl) CassandraType() string {
	return "map<varchar,frozen<VTIntDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPIntDbl) JavaType() string {
	return "Map<String,Pair<Integer,Double>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntStr) TypeName() string {
	return "VMStrTPIntStr"
}

// GoType the type of object in go
func (t *VMStrTPIntStr) GoType() string {
	return "VMStrTPIntStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntStr) CassandraType() string {
	return "map<varchar,frozen<VTIntStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPIntStr) JavaType() string {
	return "Map<String,Pair<Integer,String>>"
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

// IsVector more for interface acceptance
func (t *VMStrTPIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMStrTPIntInt) TypeName() string {
	return "VMStrTPIntInt"
}

// GoType the type of object in go
func (t *VMStrTPIntInt) GoType() string {
	return "VMStrTPIntInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMStrTPIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMStrTPIntInt) CassandraType() string {
	return "map<varchar,frozen<VTIntInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMStrTPIntInt) JavaType() string {
	return "Map<String,Pair<Integer,Integer>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPDblDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblDbl) TypeName() string {
	return "VMIntTPDblDbl"
}

// GoType the type of object in go
func (t *VMIntTPDblDbl) GoType() string {
	return "VMIntTPDblDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblDbl ( k double, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblDbl) CassandraType() string {
	return "map<bigint,frozen<VTDblDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPDblDbl) JavaType() string {
	return "Map<Integer,Pair<Double,Double>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPDblStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblStr) TypeName() string {
	return "VMIntTPDblStr"
}

// GoType the type of object in go
func (t *VMIntTPDblStr) GoType() string {
	return "VMIntTPDblStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblStr ( k double, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblStr) CassandraType() string {
	return "map<bigint,frozen<VTDblStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPDblStr) JavaType() string {
	return "Map<Integer,Pair<Double,String>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPDblInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPDblInt) TypeName() string {
	return "VMIntTPDblInt"
}

// GoType the type of object in go
func (t *VMIntTPDblInt) GoType() string {
	return "VMIntTPDblInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPDblInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTDblInt ( k double, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPDblInt) CassandraType() string {
	return "map<bigint,frozen<VTDblInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPDblInt) JavaType() string {
	return "Map<Integer,Pair<Double,Integer>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPStrDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrDbl) TypeName() string {
	return "VMIntTPStrDbl"
}

// GoType the type of object in go
func (t *VMIntTPStrDbl) GoType() string {
	return "VMIntTPStrDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrDbl ( k varchar, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrDbl) CassandraType() string {
	return "map<bigint,frozen<VTStrDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPStrDbl) JavaType() string {
	return "Map<Integer,Pair<String,Double>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPStrStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrStr) TypeName() string {
	return "VMIntTPStrStr"
}

// GoType the type of object in go
func (t *VMIntTPStrStr) GoType() string {
	return "VMIntTPStrStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrStr ( k varchar, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrStr) CassandraType() string {
	return "map<bigint,frozen<VTStrStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPStrStr) JavaType() string {
	return "Map<Integer,Pair<String,String>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPStrInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPStrInt) TypeName() string {
	return "VMIntTPStrInt"
}

// GoType the type of object in go
func (t *VMIntTPStrInt) GoType() string {
	return "VMIntTPStrInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPStrInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTStrInt ( k varchar, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPStrInt) CassandraType() string {
	return "map<bigint,frozen<VTStrInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPStrInt) JavaType() string {
	return "Map<Integer,Pair<String,Integer>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPIntDbl) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntDbl) TypeName() string {
	return "VMIntTPIntDbl"
}

// GoType the type of object in go
func (t *VMIntTPIntDbl) GoType() string {
	return "VMIntTPIntDbl"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntDbl) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntDbl ( k bigint, v double );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntDbl) CassandraType() string {
	return "map<bigint,frozen<VTIntDbl>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPIntDbl) JavaType() string {
	return "Map<Integer,Pair<Integer,Double>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPIntStr) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntStr) TypeName() string {
	return "VMIntTPIntStr"
}

// GoType the type of object in go
func (t *VMIntTPIntStr) GoType() string {
	return "VMIntTPIntStr"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntStr) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntStr ( k bigint, v varchar );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntStr) CassandraType() string {
	return "map<bigint,frozen<VTIntStr>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPIntStr) JavaType() string {
	return "Map<Integer,Pair<Integer,String>>"
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

// IsVector more for interface acceptance
func (t *VMIntTPIntInt) IsVector() bool {
	return true
}

// Name the type name for ease
func (t *VMIntTPIntInt) TypeName() string {
	return "VMIntTPIntInt"
}

// GoType the type of object in go
func (t *VMIntTPIntInt) GoType() string {
	return "VMIntTPIntInt"
}

// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *VMIntTPIntInt) CassandraCreateType(keyspace string) string {
	return "CREATE TYPE IF NOT EXISTS " + keyspace + ".VTIntInt ( k bigint, v bigint );"
}

// CassandraType the matching types in cassandra for the map
func (t *VMIntTPIntInt) CassandraType() string {
	return "map<bigint,frozen<VTIntInt>>"
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t *VMIntTPIntInt) JavaType() string {
	return "Map<Integer,Pair<Integer,Integer>>"
}

var TupleNames = [...]string{"VTDblDbl", "VTDblStr", "VTDblInt", "VTStrDbl", "VTStrStr", "VTStrInt", "VTIntDbl", "VTIntStr", "VTIntInt"}
var ListNames = [...]string{"VLDbl", "VLStr", "VLInt", "VLDblDbl", "VLDblStr", "VLDblInt", "VLStrDbl", "VLStrStr", "VLStrInt", "VLIntDbl", "VLIntStr", "VLIntInt"}
var SetNames = [...]string{"VSDbl", "VSStr", "VSInt", "VSDblDbl", "VSDblStr", "VSDblInt", "VSStrDbl", "VSStrStr", "VSStrInt", "VSIntDbl", "VSIntStr", "VSIntInt"}
var MapNames = [...]string{"VMStrDbl", "VMStrStr", "VMStrInt", "VMIntDbl", "VMIntStr", "VMIntInt", "VMStrTPDblDbl", "VMStrTPDblStr", "VMStrTPDblInt", "VMStrTPStrDbl", "VMStrTPStrStr", "VMStrTPStrInt", "VMStrTPIntDbl", "VMStrTPIntStr", "VMStrTPIntInt", "VMIntTPDblDbl", "VMIntTPDblStr", "VMIntTPDblInt", "VMIntTPStrDbl", "VMIntTPStrStr", "VMIntTPStrInt", "VMIntTPIntDbl", "VMIntTPIntStr", "VMIntTPIntInt"}
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
	"map<bigint,bigint>":            "VMIntInt",
	"map<bigint,double>":            "VMIntDbl",
	"map<bigint,frozen<VTIntDbl>>":  "VMIntTPIntDbl",
	"map<bigint,frozen<VTIntInt>>":  "VMIntTPIntInt",
	"map<bigint,frozen<VTIntStr>>":  "VMIntTPIntStr",
	"map<bigint,varchar>":           "VMIntStr",
	"map<double,frozen<VTDblDbl>>":  "VMIntTPDblDbl",
	"map<double,frozen<VTDblInt>>":  "VMIntTPDblInt",
	"map<double,frozen<VTDblStr>>":  "VMIntTPDblStr",
	"map<varchar,bigint>":           "VMStrInt",
	"map<varchar,double>":           "VMStrDbl",
	"map<varchar,frozen<VTStrDbl>>": "VMIntTPStrDbl",
	"map<varchar,frozen<VTStrInt>>": "VMIntTPStrInt",
	"map<varchar,frozen<VTStrStr>>": "VMIntTPStrStr",
	"map<varchar,varchar>":          "VMStrStr",
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
