package xtype

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type XString struct {
	String string
	Valid  bool
}

func NewXString(s string) XString {
	return XString{
		String: s,
		Valid:  true,
	}
}

//json编码
func (x XString) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return []byte(strconv.Quote(x.String)), nil
	} else {
		return jsonNull, nil
	}
}

//json解码
//严格遵循双引号限制，双引号内为字符串内容
func (x *XString) UnmarshalJSON(b []byte) error {

	if bytes.Equal(b, jsonNull) {
		x.String, x.Valid = "", false
		return nil
	} else {
		if bytes.HasPrefix(b, quoteBytes) && bytes.HasSuffix(b, quoteBytes) {

			if len(b) == 2 {
				x.String, x.Valid = "", true
			} else {
				x.String, x.Valid = string(b[1:len(b)-1]), true
			}

			return nil

		} else {
			x.String, x.Valid = "", false
			return errors.New(fmt.Sprint(string(b), " XString unmarshal json fail."))
		}
	}

}
