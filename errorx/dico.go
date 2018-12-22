package errorx

import (
	"math"
)

// ErrNotDefine is a reserved error on the MaxUint64 value
// catch an undefine case
const (
	indexReserved uint64 = math.MaxUint64
	ErrNotDefine         = "error not recorded in dictionnary errors"
)

// Dico describe a combo type to report an error message
type Dico map[uint64]string

// NewDico return a fresh and empty Dico
func NewDico() Dico {
	d := make(map[uint64]string)
	d[indexReserved] = ErrNotDefine
	return d
}

// SetEntry record a new message to a given code
// You can erase the reserved index to personalize the relative message
func (d *Dico) SetEntry(code uint64, message string) {
	map[uint64]string(*d)[code] = message
}

// Error create a new error from the dictionnary code
func (d Dico) Error(code uint64) *Errorx {
	message, ok := d[code]
	if !ok {
		return New(indexReserved, d[indexReserved])
	}
	return New(code, message)
}
