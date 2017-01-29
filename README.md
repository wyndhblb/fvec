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

For such things I recommend cadent (https://github.com/wyndhblb/cadent).


## To Generate

python3 is required for the generation as well, to laydown the inital boilerplate and protobut file, but it's as easy as 

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
    
### Resolution

Resolution is a concept for "time slabbing" things.  A time slab is a string representation of time in 
 various buckets.  
 
 
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
        Name
        {[], or map of things}
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
        vallue {t2}
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
`map[double]stuff` types.

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