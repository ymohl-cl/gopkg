package postgres

import "strconv"

// QuerySelect type
type QuerySelect interface {
	From(tableName string) QuerySelect
	Where(key string, value interface{}) QuerySelect
	String() string
	Values() []interface{}
}

type querySelect struct {
	str       string
	filters   []interface{}
	nbFilters int16
}

// NewQuerySelect return
func NewQuerySelect(keys []string) QuerySelect {
	q := querySelect{str: "SELECT"}
	size := len(keys)
	for i, v := range keys {
		if i+1 < size {
			q.str += " " + v + ","
		} else {
			q.str += " " + v
		}
	}
	return &q
}

// String getter to the query string
func (q querySelect) String() string {
	return q.str
}

// Values getter to the values describe in the query string
func (q querySelect) Values() []interface{} {
	return q.filters
}

// From attach the table name to the query
func (q *querySelect) From(tableName string) QuerySelect {
	q.str += " FROM " + tableName
	return q
}

// Where attach filter condition in the query
func (q *querySelect) Where(key string, value interface{}) QuerySelect {
	if q.nbFilters == 0 {
		q.str += " WHERE"
	} else {
		q.str += " AND"
	}
	q.nbFilters++
	q.filters = append(q.filters, value)
	q.str += " " + key + " = $" + strconv.Itoa(int(q.nbFilters))
	return q
}

// QueryInsert type
type QueryInsert interface {
	Into(tableName string) QueryInsert
	WithReturn(key string) QueryInsert
	WithValues(values []interface{}) QueryInsert
	String() string
	Values() []interface{}
}

type queryInsert struct {
	str       string
	filters   []interface{}
	nbFilters int16
}

// NewQueryInsert return
func NewQueryInsert(keys []string) QueryInsert {
	q := queryInsert{}
	q.nbFilters = int16(len(keys))
	for i, v := range keys {
		if i+1 < int(q.nbFilters) {
			q.str += v + ", "
		} else {
			q.str += v
		}
	}
	q.str += ") VALUES("
	for i := 0; i < int(q.nbFilters); i++ {
		if i+1 < int(q.nbFilters) {
			q.str += "$" + strconv.Itoa(i+1) + ","
		} else {
			q.str += "$" + strconv.Itoa(i+1)
		}
	}
	q.str += ")"
	return &q
}

// Into define the table name to the query
func (q *queryInsert) Into(tableName string) QueryInsert {
	q.str = "INSERT INTO " + tableName + "(" + q.str
	return q
}

// WithReturn define the key returning after intertion
func (q *queryInsert) WithReturn(key string) QueryInsert {
	q.str += " RETURNING " + key
	return q
}

// WithValues attachs values describe by the query string
func (q *queryInsert) WithValues(values []interface{}) QueryInsert {
	q.filters = append(q.filters, values...)
	return q
}

// String getter to the query string
func (q queryInsert) String() string {
	return q.str
}

// Values getter to the values describe in the query string
func (q queryInsert) Values() []interface{} {
	return q.filters
}
