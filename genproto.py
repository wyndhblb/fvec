#!env python3

"""
Generate some protobuf things and some go code from a boiler plate set of types
tuples, map, sets of basic object for the many types
"""

from string import Template

PROTOFILE = open("vepr.proto", "w")
GOFILE = open("vepr_boiler.go", "w")

HEAD ="""
// GENERATED3 with python genproto.py
// Edit at your own risk
syntax = 'proto3';

package github.com.wyndhblb.fvec;

option go_package = "vepr";

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
    // @inject_tag: cql:"tags" msg:"tags"
	repeated Tag tags = 11;

}


"""

print(HEAD, file=PROTOFILE)

MAPTYPES = {"string":"Str", "int64":"Int", "double":"Dbl"}
SPRINTTYPES = {"string":"%s", "int64":"%d", "double":"%v"}
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
    // @inject_tag: cql:"name" msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec" msg:"vec"
    repeated $to vec = 2;
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
    // @inject_tag: cql:"name" msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec" msg:"vec"
    repeated VT$fo$fi vec = 2;
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



########### Sets

Sets = []

setOneGen = """
// VS$fo set
// easyjson:json
message VS$fo {
    // @inject_tag: cql:"name" msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec" msg:"vec"
    repeated $to vec = 2;
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
    // @inject_tag: cql:"name" msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec" msg:"vec"
    repeated VT$fo$fi vec = 2;
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


########### Maps

# NOTE: cannot use "doubles" as keys for maps

Maps = []

mapSiGen = """
// VM$fo$fi map of $to -> $ti
$easyjson
message VM$fo$fi {
    // @inject_tag: cql:"name" msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec" msg:"vec"
    map<$to, $ti> vec = 2;
}
"""
_mdid = []

for to,fo in MAPTYPES.items():
    if to == "double":
            continue
    eas ="// easyjson:json"
    if to == "int64":
        eas = ""
    for ti, fi in MAPTYPES.items():

        tp = Template(mapSiGen).substitute(fo=fo, fi=fi, to=to, ti=ti, easyjson=eas)
        ts={
            "to": to, "ti": ti, "fo": fo, "fi": fi, "tpl": tp, "co": CASSTYPES[to], "ci":CASSTYPES[ti],
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
    // @inject_tag: cql:"name", msg:"name"
    VName name = 1;
    // @inject_tag: cql:"vec", msg:"vec"
    map<$t2, VT$fo$fi> vec = 2;
}
"""

## cannot map doubles
for t2,bo in MAPTYPES.items():
    if bo == "Dbl":
        continue

    for tps in Tuples:

        eas ="// easyjson:json"
        if t2 == "int64":
            eas = ""
        tpl = tps.copy()
        tpl["bo"] = bo
        tpl["t2"] = t2
        tpl["name"] = "VM"+bo + "TP" + tps['fo'] + tps['fi']
        if tpl["name"] in _mdid:
            continue

        tp = Template(mapTPGen).substitute(**tpl, easyjson=eas)
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

package vepr

import "reflect"
import "fmt"

"""
print(BOILHEAD, file=GOFILE)

allBoilTpl = """

// GetName returns the Name of the vector
func (t *$name) GetName() *VName{
    return t.Name
}

// Key returns the key of the vector
func (t *$name) Key() string{
    return t.Name.Key
}

// Tags returns the tags of the vector
func (t *$name) Tags() Tags{
    return t.Name.Tags
}

// UniqueId returns the tags of the vector
func (t *$name) UniqueId() uint64{
    return t.Name.UniqueId()
}

// UniqueIdString returns the tags of the vector
func (t *$name) UniqueIdString() string{
    return t.Name.UniqueIdString()
}

// IsVector more for interface acceptance
func (t *$name) IsVector() bool{
    return true
}

// Name the type name for ease
func (t *$name) TypeName() string{
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

"""

for tpl in Tuples:
    tpl["spr1"] = SPRINTTYPES[tpl["to"]]
    tpl["spr2"] = SPRINTTYPES[tpl["ti"]]

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

"""

for tpl in Sets:
    print(Template(setBase).substitute(**tpl), file=GOFILE)

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
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "";
}

// CassandraType the matching types in cassandra for the list
func (t *$name) CassandraType() string{
    return "list<$co>";
}


"""

listTPFuncCassType = """
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".VT$fo$fi ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the list
func (t *$name) CassandraType() string{
    return "list<frozen<VT$fo$fi>>";
}


"""

for tpl in Lists:
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
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
func (t *$name) CassandraCreateType(keyspace string) string{
    return "";
}

// CassandraType the matching types in cassandra for the map
func (t *$name) CassandraType() string{
    return "map<$co, $ci>";
}
"""

mapTPFuncCassType = """
// CassandraCreateType string for the create type (if nessesary)
// the string will be blank if no create is needed
// the set of of the tpl types
func (t *$name) CassandraCreateType(keyspace string) string{
    return "CREATE TYPE IF NOT EXISTS " + keyspace +".VT$fo$fi ( k $co, v $ci );";
}

// CassandraType the matching types in cassandra for the map
func (t *$name) CassandraType() string{
    return "map<$co, frozen<VT$fo$fi>>";
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
    print(Template(mapBase).substitute(**tpl), file=GOFILE)
    if tpl["ti"] == "":
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

