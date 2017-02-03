#!env python3

"""
Generate some protobuf things and some go code from a boiler plate set of types
tuples, map, sets of basic object for the many types
"""

from string import Template

PROTOFILE = open("fvec.proto", "w")
GOFILE = open("fvec_boiler.go", "w")

HEAD ="""
// GENERATED with python genproto.py
// Edit at your own risk
syntax = 'proto3';

package github.com.wyndhblb.fvec;

option go_package = "fvec";

import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "github.com/wyndhblb/timeslab/timeslab.proto";

option (gogoproto.goproto_getters_all) = false;
option (gogoproto.benchgen_all) = true;
option (gogoproto.unmarshaler_all) = true;
option (gogoproto.marshaler_all) = true;
option (gogoproto.sizer_all) = true;
option (gogoproto.populate_all) = true;


// easyjson:json
message Tag {
	string name = 1;
	string value = 2;
}


// VN the base name for a given indexible element
// easyjson:json
message VName {
    // @inject_tag: cql:"key" msg:"key"
	string key = 1;
    // @inject_tag: cql:"res" msg:"res"
	github.com.wyndhblb.timeslab.Resolution resolution = 2;
    // @inject_tag: cql:"ttl" msg:"ttl"
	uint32 ttl = 3;
    // @inject_tag: cql:"-" msg:"-"
	uint64 unique_id = 4 [(gogoproto.customname) = "XUniqueId", (gogoproto.jsontag) = "-"];
    // @inject_tag: cql:"uid" msg:"uid"
	string unique_str = 5 [(gogoproto.customname) = "XUniqueStr", (gogoproto.jsontag) = "-"];
    // @inject_tag: cql:"time" msg:"time"
    int64 time = 6;
    // @inject_tag: cql:"tags" msg:"tags"
	repeated Tag tags = 11;

}


"""

print(HEAD, file=PROTOFILE)

MAPTYPES = {"string":"Str", "int64":"Int", "double":"Dbl"}
GO_TYPES = {"string":"string", "int64":"int64", "double":"float64", "float64":"float64", "Int": "int64", "Dbl":"float64", "Str":"string", "":""}
CONV_FUNCS = {"string": "stringToString", "float64":"stringToFloat", "double":"stringToFloat", "int64":"stringToInt", "Int": "stringToInt", "Dbl":"stringToFloat", "Str":"stringToString", "":""}
JAVA_TYPES = {"string":"String", "int64":"Integer", "double":"Double", "varchar":"String", "Dbl": "Double", "bigint":"Integer", "Int":"Integer", "Str":"String", "":""}
SPRINTTYPES = {"string":"%s", "int64":"%d", "double": "%v", "float64":"%v", "":""}
CASSTYPES = {"string":"varchar", "int64":"bigint", "double":"double"}
KEYTYPE = ["string", "int64"]


# the map of cassandra types to objects
CassandraTypes = {}

################################################################################
############### Tuples
################################################################################


tupleGEN = """
// VT$fo$fi Tuple Type key=$to value=$ti
// easyjson:json
message VT$fo$fi{
    // @inject_tag: cql:"k" msg:"key"
    $to key = 1;
    // @inject_tag: cql:"v" msg:"value"
    $ti value = 2;
}
"""

Tuples = []

## tuple gen
for to,fo in MAPTYPES.items():
    for ti, fi in MAPTYPES.items():
        tp = Template(tupleGEN).substitute(fo=fo, fi=fi, to=to, ti=ti)
        print("Tuple: VT" + fo + fi)
        Tuples.append({
            "to": to, "ti": ti, "fo": fo, "fi": fi, "co": CASSTYPES[to], "ci":CASSTYPES[ti],  "tpl": tp,
            "name": "VT" + fo + fi
        })
        print(tp, file=PROTOFILE)


########### Lists

Lists = []

listOneGen = """
// VL$fo list
// easyjson:json
message VL$fo {
    // @inject_tag: cql:"vec" msg:"vec"
    repeated $to vec = 1;
}
"""
for to,fo in MAPTYPES.items():
    tp = Template(listOneGen).substitute(fo=fo, to=to)
    print("List: VL" + fo)
    ts={
        "to": to, "ti": "", "fo": fo, "fi": "", "co": CASSTYPES[to], "ci":"", "tpl": tp
    }
    ts["name"] = "VL"+fo
    Lists.append((ts))
    print(tp, file=PROTOFILE)


listTpGen = """
// VL$fo$fi list
// easyjson:json
message VL$fo$fi {
    // @inject_tag: cql:"vec" msg:"vec"
    repeated VT$fo$fi vec = 1;
}
"""
for tps in Tuples:
    tp = Template(listTpGen).substitute(**tps)
    ts = tps.copy()
    ts["tpl"] = tp
    ts["name"] = "VL" + tps['fo'] + tps['fi']
    print("List: VL" +  tps['fo'] + tps['fi'])
    Lists.append(ts)
    print(tp, file=PROTOFILE)



################################################################################
############### SETS
################################################################################

Sets = []

setOneGen = """
// VS$fo set
// easyjson:json
message VS$fo {
    // @inject_tag: cql:"vec" msg:"vec"
    repeated $to vec = 1;
}
"""
for to,fo in MAPTYPES.items():
    tp = Template(setOneGen).substitute(fo=fo, to=to)
    ts = {
        "to": to, "ti": "", "fo": fo, "fi": "", "co": CASSTYPES[to], "ci":"", "tpl": tp,
        "name": "VS"+fo
    }
    Sets.append(ts)
    print("Set: VS" + fo)
    print(tp, file=PROTOFILE)

setTpGen = """
// VS${fo}$fi set
// easyjson:json
message VS${fo}$fi {
    // @inject_tag: cql:"vec" msg:"vec"
    repeated VT$fo$fi vec = 1;
}
"""
for tps in Tuples:
    tp = Template(setTpGen).substitute(**tps)
    ts = tps.copy()
    ts["tpl"] = tp
    ts["name"] = "VS" +tps['fo']  +tps['fi']
    Sets.append(ts)
    print("Set: VS" + tps['fo']  +tps['fi'])
    print(tp, file=PROTOFILE)


################################################################################
############### Maps
################################################################################

# NOTE: cannot use "doubles" as keys for maps
# NOTE: a special msgpack encoder is needed for map[int64]

Maps = []

mapSiGen = """
// VM$fo$fi map of $to -> $ti
$easyjson
message VM$fo$fi {
    // @inject_tag: cql:"vec" msg:"vec${ext}"
    map<$to, $ti> vec = 1;
}
"""
_mdid = []

for to,fo in MAPTYPES.items():
    if to == "double":
            continue
    eas ="// easyjson:json"
    ext =""
    if to == "int64":
        eas = ""
        ext = ",extention"

    for ti, fi in MAPTYPES.items():

        tp = Template(mapSiGen).substitute(fo=fo, fi=fi, to=to, ti=ti, easyjson=eas, ext=ext)
        ts={
            "to": to, "ti": ti, "fo": fo, "fi": fi, "tpl": tp, "co": CASSTYPES[to], "ci":CASSTYPES[ti], "mapsing" : True
        }
        ts["name"] = "VM"+fo + fi
        _mdid.append(ts["name"])
        Maps.append(ts)
        print("Map: VM" + fo + fi)
        print(tp, file=PROTOFILE)

mapTPGen = """
// VM${bo}TP$fo$fi map of $to -> set($to, $ti)
$easyjson
message VM${bo}TP$fo$fi {
    // @inject_tag: cql:"vec" msg:"vec${ext}"
    map<$t2, VT$fo$fi> vec = 1;
}
"""

## cannot map doubles
for t2,bo in MAPTYPES.items():
    if bo == "Dbl":
        continue

    for tps in Tuples:
        ext = ""
        eas ="// easyjson:json"
        if t2 == "int64":
            eas = ""
            ext = ",extention"
        tpl = tps.copy()
        tpl["bo"] = bo
        tpl["t2"] = t2
        tpl["name"] = "VM"+bo + "TP" + tps['fo'] + tps['fi']
        tpl["mapsing"] = False
        if tpl["name"] in _mdid:
            continue

        tp = Template(mapTPGen).substitute(**tpl, easyjson=eas, ext=ext)
        tpl["tpl"] = tp
        _mdid.append(ts["name"])

        print("Map: "+tpl["name"])
        Maps.append(tpl)
        print(tp, file=PROTOFILE)

PROTOFILE.close()

###################### BOILER PLATE CODE
## For all those maps and sets we need to write "boiler plate" functions for them all

BOILHEAD ="""
// GENERATED with python3 genproto.py
// Edit at your own risk

package fvec

import "reflect"
import "fmt"
import "strings"
import "errors"
import "strconv"

var ErrorInvalidRedisValue = errors.New("Invalid redis value, cannot parse")

func stringToFloat(tp string) (float64, error){
    return strconv.ParseFloat(tp, 64)
}

func stringToInt(tp string) (int64, error){
    return strconv.ParseInt(tp, 10, 64)
}

func stringToString(tp string) (string, error){
    return tp, nil
}

"""
print(BOILHEAD, file=GOFILE)

allBoilTpl = """

// IsVector more for interface acceptance
func (t *$name) IsVector() bool{
    return true
}

// Name the type name for ease
func (t *$name) TypeName() string{
    return "$name"
}

// GoType the type of object in go
func (t *$name) GoType() string{
    return "$name"
}
"""

############ Tuple Boiler


tplFuncCassType = """
//**************** Tuple: $to $ti VT$fo$fi **********************/

// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".$name ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the tuple
func (t *$name) CassandraType() string{
    return "$name";
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "LPUSH " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisValue() string{
    return fmt.Sprintf("$spr1 $spr2", t.Key, t.Value);
}

// JavaType the type of object in java
// org.apache.commons.lang3.tuple
func (t *$name) JavaType() string{
    return "Pair<$jo,$ji>"
}

""" + allBoilTpl

for tpl in Tuples:
    tpl["spr1"] = SPRINTTYPES[tpl["to"]]
    tpl["spr2"] = SPRINTTYPES[tpl["ti"]]
    tpl["jo"] = JAVA_TYPES[tpl["to"]]
    tpl["ji"] = JAVA_TYPES[tpl["ti"]]
    tpl["go"] = GO_TYPES[tpl["to"]]
    tpl["gi"] = GO_TYPES[tpl["ti"]]
    tpl["convo"] = CONV_FUNCS[tpl["to"]]
    tpl["convi"] = CONV_FUNCS[tpl["ti"]]

    tp = Template(tplFuncCassType).substitute(**tpl)
    print(tp, file=GOFILE)

########################################################
############ Set Boiler
########################################################

setBase ="""
//**************** Set: $to $ti VS$fo$fi **********************/

// VecType the type of vector (set)
func (t *$name) VecType() string{
    return "set"
}

// IsSet is a set type
func (t *$name) IsSet() bool{
    return true
}


""" + allBoilTpl

setOneFuncCassType = """
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "";
}

// CassandraType the matching types in cassandra for the set
func (t *$name) CassandraType() string{
    return "set<$co>";
}

// JavaType the type of object in java
// java.util
func (t *$name) JavaType() string{
    return "Set<$jo>"
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *$name) RedisRemoveCmd(key string) string{
    return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisValue(v $go) string{
    return fmt.Sprintf("$spr1", v);
}

// FromRedisValue given the redis value, make it into a proper $go
func (t *$name) FromRedisValue(i string) (v $go, err error){
    return $convo(i)
}
"""

setTPFuncCassType = """
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".VT$fo$fi ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the set
func (t *$name) CassandraType() string{
    return "set<frozen<VT$fo$fi>>";
}

// JavaType the type of object in java
// java.util
// org.apache.commons.lang3.tuple
func (t *$name) JavaType() string{
    return "Set<Pair<$jo,$ji>>"
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "SADD " + key
}

// RedisRemoveCmd returns the redis add command
func (t *$name) RedisRemoveCmd(key string) string{
    return "SREM " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisInsertValue(v VT$fo$fi) string{
    return fmt.Sprintf("$spr1:$spr2", v.Key, v.Value);
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *$name) FromRedisValue(i string) (v VT$fo$fi, err error){
    spl := strings.Split(i, ":")
    if len(spl) != 2{
        return v, ErrorInvalidRedisValue
    }
    v.Key, err = $convo(spl[0])
    v.Value, err = $convi(spl[1])
    return v, err
}
"""

for tpl in Sets:
    print(Template(setBase).substitute(**tpl), file=GOFILE)

    tpl["jo"] = JAVA_TYPES[tpl["to"]]
    tpl["ji"] = JAVA_TYPES[tpl["ti"]]
    tpl["spr1"] = SPRINTTYPES[tpl["to"]]
    tpl["spr2"] = SPRINTTYPES[tpl["ti"]]
    tpl["go"] = GO_TYPES[tpl["to"]]
    tpl["gi"] = GO_TYPES[tpl["ti"]]
    tpl["convo"] = CONV_FUNCS[tpl["to"]]
    tpl["convi"] = CONV_FUNCS[tpl["ti"]]

    if tpl["ti"] == "":
        CassandraTypes["set<" + tpl['co'] + ">"] = tpl["name"]
        tp = Template(setOneFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)
    else:
        CassandraTypes["set<frozen<VT" + tpl['fo'] + tpl['fi'] + ">>"] = tpl["name"]
        tp = Template(setTPFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)


########################################################
############ List Boiler
########################################################


listBase ="""
//**************** List: $to $ti VL$fo$fi **********************/

// VecType the type of vector (list)
func (t *$name) VecType() string{
    return "list"
}

// IsList this is a list type
func (t *$name) IsList() bool{
    return true
}

""" + allBoilTpl

listOneFuncCassType = """
// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "";
}

// CassandraType the matching types in cassandra for the list
func (t *$name) CassandraType() string{
    return "list<$co>";
}

// JavaType the java type for the object
func (t* $name) JavaType() string{
    return "List<$jo>";
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *$name) RedisRemoveCmd(key string) string{
    return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisInsertValue(v $go) string{
    return fmt.Sprintf("$spr1", v);
}

// FromRedisValue given the redis value, make it into a proper $go
func (t *$name) FromRedisValue(i string) (v $go, err error){
    return $convo(i)
}
"""

listTPFuncCassType = """
// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".VT$fo$fi ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the list
func (t *$name) CassandraType() string{
    return "list<frozen<VT$fo$fi>>";
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t* $name) JavaType() string{
    return "List<Pair<$jo,$ji>>";
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "LADD " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *$name) RedisRemoveCmd(key string) string{
    return ""
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisInsertValue(v VT$fo$fi) string{
    return fmt.Sprintf("$spr1:$spr2", v.Key, v.Value);
}

// FromRedisValue given the redis value, make it into a proper tuple
func (t *$name) FromRedisValue(i string) (v VT$fo$fi, err error){
    spl := strings.Split(i, ":")
    if len(spl) != 2{
        return v, ErrorInvalidRedisValue
    }
    v.Key, err = $convo(spl[0])
    v.Value, err = $convi(spl[1])
    return v, err
}
"""

for tpl in Lists:
    tpl["jo"] = JAVA_TYPES[tpl["to"]]
    tpl["ji"] = JAVA_TYPES[tpl["ti"]]
    tpl["spr1"] = SPRINTTYPES[tpl["to"]]
    tpl["spr2"] = SPRINTTYPES[tpl["ti"]]
    tpl["go"] = GO_TYPES[tpl["to"]]
    tpl["gi"] = GO_TYPES[tpl["ti"]]
    tpl["convo"] = CONV_FUNCS[tpl["to"]]
    tpl["convi"] = CONV_FUNCS[tpl["ti"]]

    print(Template(listBase).substitute(**tpl), file=GOFILE)
    if tpl["ti"] == "":
        CassandraTypes["list<" + tpl['co'] + ">"] = tpl["name"]
        tp = Template(listOneFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)
    else:
        CassandraTypes["list<frozen<VT" + tpl['fo'] + tpl['fi'] + ">>"] = tpl["name"]
        tp = Template(listTPFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)



########################################################
############ Map Boiler
########################################################

mapOneFuncCassType = """
// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "";
}

// CassandraType the matching types in cassandra for the map
func (t *$name) CassandraType() string{
    return "map<$co,$ci>";
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t* $name) JavaType() string{
    return "Map<$jo,$ji>";
}

// RedisInsertCmd returns the redis add command
func (t *$name) RedisInsertCmd(key string) string{
    return "HMSET " + key
}

// RedisRemoveCmd returns the redis add command
// (cannot remove from lists) this is a blank command
func (t *$name) RedisRemoveCmd(key string) string{
    return "HDEL " + key
}

// RedisInsertValue returns what the value string would be for a redis command
func (t *$name) RedisInsertValue(v $go) string{
    return fmt.Sprintf("$spr1", v);
}

// FromRedisValue given the redis value, make it into a proper $go
func (t *$name) FromRedisValue(i string) (v $go, err error){
    return $convo(i)
}
"""

mapTPFuncCassType = """
// CassandraCreateType string for the create type (if necessary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".VT$fo$fi ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the map
func (t *$name) CassandraType() string{
    return "map<$mtyp,frozen<VT$fo$fi>>";
}

// JavaType the java type for the object
// org.apache.commons.lang3.tuple
func (t* $name) JavaType() string{
    return "Map<$joo,Pair<$jo,$ji>>";
}

"""

mapBase ="""
//**************** Map: $to $ti $name **********************/

// VecType the type of vector (map)
func (t *$name) VecType() string{
    return "map"
}

// IsMap is a set type
func (t *$name) IsMap() bool{
    return true
}

""" + allBoilTpl

for tpl in Maps:
    tpl["jo"] = JAVA_TYPES[tpl["to"]]
    tpl["ji"] = JAVA_TYPES[tpl["ti"]]
    tpl["joo"] = JAVA_TYPES[tpl.get('t2', '')]
    tpl["mtyp"] = CASSTYPES.get(tpl.get('t2', ''), "")
    tpl["convo"] = CONV_FUNCS[tpl["to"]]
    tpl["convi"] = CONV_FUNCS[tpl["ti"]]
    tpl["spr1"] = SPRINTTYPES[tpl["to"]]
    tpl["spr2"] = SPRINTTYPES[tpl["ti"]]
    tpl["go"] = GO_TYPES[tpl["to"]]
    tpl["gi"] = GO_TYPES[tpl["ti"]]

    print(tpl["joo"], tpl["mtyp"])
    print(Template(mapBase).substitute(**tpl), file=GOFILE)
    if tpl.get("mapsing", False):
        CassandraTypes["map<" + tpl['co'] + "," + tpl['ci'] + ">"] = tpl["name"]
        tp = Template(mapOneFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)
    else:
        CassandraTypes["map<" + tpl['co'] + ",frozen<VT" + tpl['fo'] + tpl['fi'] + ">>"] = tpl["name"]
        tp = Template(mapTPFuncCassType).substitute(**tpl)
        print(tp, file=GOFILE)

#### utilities mappings and the like

mpl = []
for tpl in Tuples:
    mpl.append('"' + tpl["name"] +'"')

print("var TupleNames = [...]string{" + ", ".join(mpl) + "}", file=GOFILE)

mpl = []
for tpl in Lists:
    mpl.append('"' + tpl["name"] +'"')

print("var ListNames = [...]string{" + ", ".join(mpl) + "}", file=GOFILE)

mpl = []
for tpl in Sets:
    mpl.append('"' + tpl["name"] +'"')

print("var SetNames = [...]string{" + ", ".join(mpl) + "}", file=GOFILE)


mpl = []
for tpl in Maps:
    mpl.append('"' + tpl["name"] +'"')

print("var MapNames = [...]string{" + ", ".join(mpl) + "}", file=GOFILE)


def sMapping(s):
    if s == "i":
        return "Int"
    elif s == "s":
        return "Str"
    else:
        return "Dbl"

shortMappings = {}
for i in ("l", "s", "m"):
    for t1 in ("s", "i", "d"):
        nNM = ""
        if i == "l":
            nNM = "VL"
        elif i=="s":
            nNM = "VS"
        else:
            nNM = "VM"
        if i != "m":
            shortMappings[i+t1] = nNM + sMapping(t1)

        for t2 in ("s", "i", "d"):
            if i == "m" and t1 == "d":
                continue
            shortMappings[i+t1+t2] = nNM + sMapping(t1) + sMapping(t2)
            if i == "m":
                for t3 in ("s", "i", "d"):
                    shortMappings[i+t1+t2+t3] = nNM + sMapping(t1) + "TP" + sMapping(t2) + sMapping(t3)

print("var SHORT_NAME_MAP = map[string]string{", file=GOFILE)
ks = list(shortMappings.keys())
ks.sort()
for k in ks:
    print('\t"' + k +'": "' + shortMappings[k] + '",', file=GOFILE)

print("}", file=GOFILE)

## cass types list
print("var CASSANDRA_TYPE_MAP = map[string]string{", file=GOFILE)

cs = list(CassandraTypes.keys())
cs.sort()
for k in cs:
    print('\t"' + k +'": "' + CassandraTypes[k] + '",', file=GOFILE)

print("}", file=GOFILE)

# switch string to new vector object
print("func stringToVector(vec string) Vector {", file=GOFILE)
print("\tswitch vec {", file=GOFILE)

for k in ks:
    print('\t\tcase "' + k + '":', file=GOFILE)
    print('\t\t\treturn new(' + shortMappings[k] + ')', file=GOFILE)
    print('\t\tcase "' + shortMappings[k] + '":', file=GOFILE)
    print('\t\t\treturn new(' + shortMappings[k] + ')', file=GOFILE)

print("\t}", file=GOFILE)

print("\treturn nil", file=GOFILE)
print("}", file=GOFILE)


print("var typeRegistry = make(map[string]reflect.Type)", file=GOFILE)


print("func init(){", file=GOFILE)

for k in ks:
    print('\ttypeRegistry["' + k +'"] = reflect.TypeOf(' + shortMappings[k] + '{})', file=GOFILE)
    print('\ttypeRegistry["' + shortMappings[k] +'"] = reflect.TypeOf(' + shortMappings[k] + '{})', file=GOFILE)

print("}", file=GOFILE)

GOFILE.close()