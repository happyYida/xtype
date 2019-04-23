package xtype

import (
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

// json编码
func (x XBool) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return []byte(strconv.FormatBool(x.Bool)), nil
	} else {
		return jsonNull, nil
	}
}

// json解码
func (x *XBool) UnmarshalJSON(b []byte) error {
	s := string(b)

	if s == null {
		x.Bool, x.Valid = false, false
		return nil
	} else {
		if b, e := strconv.ParseBool(s); e != nil {
			x.Bool, x.Valid = false, false
			return errors.New(fmt.Sprint(s, " XBool unmarshal json invalid, ", e.Error()))
		} else {
			x.Bool, x.Valid = b, true
			return nil
		}
	}
}
