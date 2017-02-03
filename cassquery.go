/**
Since we are generating things to go into cassandra, this object is for each subobject to pass around
their various WHEREs args, and other things to be able to generate the final query
*/

package fvec

type CassQuery interface {
	// GetWhere .. a string that is of the form "column=? AND column2=?"
	GetWhere() string

	// GetUpdateArgs a list of the objects that are to be the markers for the "?" in the Where and Update/Insert
	GetWhereArgs() []interface{}

	// GetUpdateArgs a list of the objects that are to be the markers for the "?" in the Where and Update/Insert
	GetUpdateArgs() []interface{}

	// GetForUpdate in cassandra updates and inserts are both "upserts" so we just need the update command
	// this will be the UPDATE {GetForUpdate} WHERE {GetWhere} in a given query
	GetForUpdate() []string
}

// CassBaseQuery the simplest implementation of the interface
type CassBaseQuery struct {
	whereArgs  []interface{}
	updateArgs []interface{}
	where      string
	forUpdate  string
}

func (c *CassBaseQuery) GetWhere() string {
	return c.where
}

func (c *CassBaseQuery) GetWhereArgs() []interface{} {
	return c.whereArgs
}

func (c *CassBaseQuery) GetUpdateArgs() []interface{} {
	return c.updateArgs
}

func (c *CassBaseQuery) GetForUpdate() string {
	return c.forUpdate
}
