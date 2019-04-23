package xtype

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type XBool struct {
	Bool  bool
	Valid bool
}

func NewXBool (Bool bool) XBool {
	return XBool{
		Bool:Bool,
		Valid:true,
	}
}

//json编码
func (x XBool) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return []byte(strconv.FormatBool(x.Bool)), nil
	} else {
		return jsonNull, nil
	}
}

//json解码
//严格遵守值为true或false才能进行转换
func (x *XBool) UnmarshalJSON(b []byte) error {

	if bytes.Equal(b, jsonNull) {
		x.Bool, x.Valid = false, false
		return nil
	} else {

		switch {
		case bytes.Equal(b, jsonTrue):
			x.Bool, x.Valid = true, true
			return nil
		case bytes.Equal(b, jsonFalse):
			x.Bool, x.Valid = false, true
			return nil
		default:
			x.Bool, x.Valid = false, false
			return errors.New(fmt.Sprint(string(b), " XBool unmarshal json invalid."))
		}
	}
}
