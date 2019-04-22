package xtype

import (
	"errors"
	"fmt"
	"strconv"
)

const null = "null"

var jsonNull = []byte(null)

type XInt64 struct {
	Int64 int64
	Valid bool
}

func NewXInt64(Int64 int64) XInt64 {
	return XInt64{
		Int64: Int64,
		Valid: true,
	}
}

// json编码
func (x XInt64) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return []byte(strconv.FormatInt(x.Int64, 10)), nil
	} else {
		return jsonNull, nil
	}
}

// json解码
func (x *XInt64) UnmarshalJSON(b []byte) error {
	s := string(b)

	if s == null {
		x.Int64, x.Valid = 0, false
		return nil
	} else {
		i, e := strconv.ParseInt(s, 10, 64)
		if e != nil {
			x.Int64, x.Valid = 0, false
			return errors.New(fmt.Sprint(s, " XInt64 unmarshal json invalid, ", e.Error()))
		} else {
			x.Int64, x.Valid = i, true
			return nil
		}
	}

}


