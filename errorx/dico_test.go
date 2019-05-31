package errorx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDico(t *testing.T) {
	// default: should be ok
	func() {
		// init
		d := NewDico()

		// assert
		if assert.NotNil(t, d) {
			assert.Len(t, d, 1)
			assert.Equal(t, ErrNotDefine, d[indexReserved])
		}
	}()
}

func TestSetEntry(t *testing.T) {
	code := uint64(42)
	message := "error example"

	// default: should be ok
	func() {
		// init
		d := NewDico()
		d.SetEntry(code, message)

		// assert
		assert.Len(t, d, 2)
		assert.Equal(t, message, d[code])
		assert.Equal(t, ErrNotDefine, d[indexReserved])
	}()
}

func TestDico_Error(t *testing.T) {
	code0 := uint64(8)
	code1 := uint64(42)
	code2 := uint64(43)
	code3 := uint64(44)
	message1 := "error example_1"
	message2 := "error example_2"
	message3 := "error example_3"

	// default: should be ok
	func() {
		// init
		d := NewDico()
		d.SetEntry(code1, message1)
		d.SetEntry(code2, message2)
		d.SetEntry(code3, message3)

		// assert
		assert.Len(t, d, 4)
		assert.Equal(t, message1, d.Error(code1).Error())
		assert.Equal(t, message2, d.Error(code2).Error())
		assert.Equal(t, message3, d.Error(code3).Error())
		assert.Equal(t, ErrNotDefine, d.Error(code0).Error())
	}()
}
