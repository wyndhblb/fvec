# Feature Vectors

    go get github.com/wyndhblb/fvec
    
### define: (to quote :: https://en.wikipedia.org/wiki/Feature_vector)

In pattern recognition and machine learning, a feature vector is an n-dimensional vector of numerical features that represent some object. 
Many algorithms in machine learning require a numerical representation of objects, 
since such representations facilitate processing and statistical analysis.


## Purpose

Forms the base objects for a feature vector over time (or all time).  
Time series are just one of the possible vectors (i.e. `set/list(int64, double)`, the `int64` being the epoch time).
Sometimes one wishes for vectors/series over a given time period (hour, day, month, year, all-the-time, etc) and not all vectors are the same type.
For example we may want a map of strings to strings or a set of ints, etc.  
Basically something "what's this category of things a user registered/saw/saved in the last day".

For a typed language and performance critical systems, it is much better to NOT use `interface{}` and then attempt to deal with 
all the type conversion/branch stuff that can happen.

If iterating, network passing, using, etc many vectors of many types without proper type information GC is highly effected as well as CPU cycles just dealing
with the type branch conversions as well as the fact there is no protobuf "generic type".

GoLang does not have generics (not starting a generics battle here, just a fact), 
so this is simple a generator of vectors types (list, set, maps) based on a a few base types (string, int64, float64) and tuples of them as well.

As well as all the protobufs, msgpack, and easyjson encoder/decoder goodies.

This is mainly meant for a Cassandra/Redis DB, where lists, sets, maps of basic types (and more complex types) are allowed.

The basic types vector includes a "Name" and a "list/map of stuff".

This forms the basis of a much large vector machine storage engine.  

### warning

I do NOT recommend this style for really high velocity distinct time series (aka data that is coming in at or less then 1 sec intervals, every second, all the time
things like CPU(s) usage for a given machine). 

For such things I recommend cadent (https://github.com/wyndhblb/cadent) or other related projects 


## To Generate

python3 is required for the generation as well, to laydown the initial boilerplate and protobuf file, but it's as easy as 

    cd wyndhblb/fvec
    go generate


## Name

A rather generic key rep w/ tags resolution

    Key string
    Tags = [](key, value)
    Resolution 
    TTL

### UniqueId

There is a nice utility method for generating "unique" IDs for a given Name which is 

    fnv64a(key + ":" + SortedByName(tags))
    
and a Unique "string" format which is

    base36(unique_id)
 
This is needed as many DBs overflow on a `unit64` (the fnv64a type) (aka the DB we're targeting cassandra, elastic, redis) only support
`int64`, thus a `uint64` will overflow that for 50% of the time.
 
 
### Resolution

Resolution is a concept for "time slabbing" things.  A time slab is a string representation of time in 
 various buckets (see https://github.com/wyndhblb/timeslab).  
 
 
        Resolution_MIN: 200601021504
        Resolution_MIN10: 2006010215I10{min / 10}
        Resolution_MIN30: 2006010215I30{min / 30}
        Resolution_HOUR: 2006010215
        Resolution_DAY: 20060102
        Resolution_WEEK: 2006{iso week: 1->53}
        Resolution_MONTH: 200601
        Resolution_MONTH2: 2006M2{month / 2}
        Resolution_MONTH3: 2006M3{month / 3}
        Resolution_MONTH6: 2006M6{month / 6}
        Resolution_YEAR: 2006
        Resolution_ALL: ALL
        
## Vector Permutations

    Vector{
        {list, set, or map of things}
    }

### Base Types

A basic vector has some type it's a vector of.

    int64
    float64
    string

And the tuples of those basic types (9 total permutations)

    (string|int64|double, string|int64|double)
    
The tuple object is a basic 2 element object of "key" and "value" which is designated by `TP`

The labels

    Str -> string
    Int -> int64
    Dbl -> float64

Make up the parts of the object names.  

So a Tuple type looks like `VT{int64|string|double}{int64|string|double}` in the code you see but really is just a two element object
of the form 

    VT{t1}{t2}{
        key {t1}
        value {t2}
    }

where `t[1-2]` are base types.


### Object naming conventions

Each vector object starts with it's root type prefix

    VL -> lists
    VS -> sets
    VM -> map
    VT -> tuple

### Lists + Sets

Have 12 variations each

    VL{Int|Str|Dbl} + VL{Int|Str|Dbl}{Int|Str|Dbl}
    VS{Int|Str|Dbl} + VS{Int|Str|Dbl}{Int|Str|Dbl}

#### A note about `Sets`

a `SET` is not really a "true set" internal to golang (as there is not a `set` base type), instead the cassandra/redis `type` is a set, 
but internal golang type to go is a simple array/slice. 
 
The Cassandra query
 
    UPDATE {keyspace}.{table}.{column} = {keyspace}.{table}.{column} + {vector_element} 
    
Will not do anything if the `{vector_element}` already exists inside a `set<vector_element>` type

Same is true for redis where

    SADD {key} {vector_element} 

won't do anything if the `vector_element` is already there.


When reading in a `set type` from cassandra/redis one can save many CPU cycles assuming the underlying database has made the `vector_element`s unique.


### Maps

Have a 3rd dimension, the `map key`, so they expand larger, but note that json + protobufs cannot encode "doubles" as keys (by their spec), so there are no
`map[double]stuff` types.  Also not json and msgpack cannot encode `int` based map keys either, however protobufs can.

    VM{Str|Int}{Str|Int|Dbl} + VM{Str|Int}TP{Str|Int|Dbl}{Str|Int|Dbl}


### ShortHand

There is also a short hand notation to describe these objects.  Most of the time you can use `GetVector` to get a new
vector.  

    GetVector(basetype byte, subtypes string)
    
Here `basetype`
    
    l -> list
    s -> set
    m -> map
    
the subtypes are represented as
 
    s -> string
    i -> int
    d -> double

Some examples:

- `m`, `si` is a named map vector with the vector object as `map[string]int64`
- `m`, `ssi` is a named map vector with the vector object as `map[string]Tuple{string, int64}`
- `s`, `s` is a named set vector with the vector object as `[]string`
- `l`, `d` is a named set vector with the vector object as `[]double`
- `l`, `sd` is a named set vector with the vector object as `[]Tuple{string, double}`


Another method is  

    GetVectorFromString(string)
    
Which will take the ShortHand or "Full struct form"

    msss which is a VMStrTPStrStr which is map<string, VTStrStr> which is map<string, <string, string>
    sis which is VSStrStr which is set<VTStrStr> which is set<string, string>
    VMStrTPStrStr
    VSStrStr
    li which is VLInt which is list<bigint>
    VLint
    
*note* would not recommend using this for really high volume object initialization as the branch/map and string GC issues can cause performance issue.
If the volume is low (1000s/Second) the extra GC/CPU cycles is probably ok.

## Interface

The basic interface is a `Vector` which can be used as the basis for usage in other things.

Other interfaces are `VectorMap`, `VectorList`, `VectorSet` for each sub type.

The basic interface for scalar values is a `Scalar`, these are effectively wrappers around the base types (int64, string, float64)
with a special struct for Counters as they are treated very differently in cassandra then the other types

## FVecGen

I included a simple golang class generator for creating "complex" objects.  The basic idea is that in a cassandra table, we collect
a bunch of "vectors and scalars" together as a typical feature vector can contain many parts and all parts tend to be needed
for evaluation in models.

THe basic syntax so far is 

    fvecgen --classname=CLASSS --pkgname=PACKAGE --format={VarName}={shorthand},{VarName}={shorthand}...
    
The `VarName` will end as a public struct attribute (CamelCase) and the name of a cassandra column (snake_case).

For example

    fvecgen --classname=GenTT --pkgname=main --format=col1=s,col2=mssi,col3=c,col4=f,col5=si,col6=msss,col7=li,col8=b,col9=ss,col10=msd,col11=c,col12=sid 
    
Will generate a struct that looks like
    
    type GenTT struct {
    	Name *fvec.VName
    
    	Col1 *fvec.StringType `json:"col1" cql:"col1" msg:"col1"`
    
    	Col2 *fvec.VMStrTPStrInt `json:"col2" cql:"col2" msg:"col2"`
    
    	Col3 *fvec.CounterType `json:"col3" cql:"col3" msg:"col3"`
    
    	Col4 *fvec.DoubleType `json:"col4" cql:"col4" msg:"col4"`
    
    	Col5 *fvec.VSInt `json:"col5" cql:"col5" msg:"col5"`
    
    	Col6 *fvec.VMStrTPStrStr `json:"col6" cql:"col6" msg:"col6"`
    
    	Col7 *fvec.VLInt `json:"col7" cql:"col7" msg:"col7"`
    
    	Col8 *fvec.ByteType `json:"col8" cql:"col8" msg:"col8"`
    
    	Col9 *fvec.VSStr `json:"col9" cql:"col9" msg:"col9"`
    
    	Col10 *fvec.VMStrDbl `json:"col10" cql:"col10" msg:"col10"`
    
    	Col11 *fvec.CounterType `json:"col11" cql:"col11" msg:"col11"`
    
    	Col12 *fvec.VSIntDbl `json:"col12" cql:"col12" msg:"col12"`
    }

and come with a bunch of helper functions.  Since there are counters here, the generated cassandra "create" structure will 
actually have 2 tables, one for the counters and one for everything else.  the counter table is prefixed with `_counters`

    CREATE TYPE IF NOT EXISTS mykey.VTStrInt ( k varchar, v bigint );
    CREATE TYPE IF NOT EXISTS mykey.VTStrStr ( k varchar, v varchar );
    CREATE TYPE IF NOT EXISTS mykey.VTIntDbl ( k bigint, v double );
    CREATE TABLE IF NOT EXISTS mykey.mtable(
        uid ascii, 
        slab ascii, 
        ord ascii, 
        col1 varchar, 
        col2 map<varchar,frozen<VTStrInt>>, 
        col4 double, 
        col5 set<bigint>, 
        col6 map<varchar,frozen<VTStrStr>>, 
        col7 list<bigint>, 
        col8 blob, 
        col9 set<varchar>, 
        col10 map<varchar,double>, 
        col12 set<frozen<VTIntDbl>>,
        PRIMARY KEY ((uid, slab), ord)
        ) WITH CLUSTERING ORDER BY (ord ASC) AND
            compaction = {
            'class': 'TimeWindowCompactionStrategy',
            'compaction_window_unit': 'DAYS',
            'compaction_window_size': '1'
        }
        AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'};
            
    CREATE TABLE IF NOT EXISTS mykey.mtable_counters(
        uid ascii, 
        slab ascii, 
        ord ascii, 
        col3 counter, 
        col11 counter, 
        PRIMARY KEY ((uid, slab), ord)
        ) WITH compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
        AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'};

