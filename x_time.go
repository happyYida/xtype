package xtype

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type timeFormat int

func (f timeFormat) String() string {
	switch f {
	case FormatTimestamp:
		return "timestamp"
	case FormatTime:
		return "time"
	case FormatDate:
		return "date"
	case FormatDatetime:
		return "datetime"
	case FormatYear:
		return "year"
	}
	return ""
}

const (
	FormatTimestamp timeFormat = iota  // 毫秒时间戳
	FormatYear
	FormatDatetime
	FormatDate
	FormatTime
)

const timeBase = "2006-01-02 15:04:05"

type XTime struct {
	Time	time.Time
	Valid	bool
	Format	timeFormat
}

func NewXTime(t time.Time, f timeFormat) XTime {
	return XTime{
		Time: 	t,
		Valid:	true,
		Format:	f,
	}
}

//json编码
//FormatTimestamp返回number，其他返回string
func (x XTime) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return jsonNull, nil
	} else {
		switch x.Format {
		case FormatTimestamp:
			return []byte(strconv.FormatInt(x.Time.Unix() * 1000, 10)), nil
		case FormatYear:
			return []byte(strconv.Quote(x.Time.Format(timeBase[:4]))), nil
		case FormatDatetime:
			return []byte(strconv.Quote(x.Time.Format(timeBase))), nil
		case FormatDate:
			return []byte(strconv.Quote(x.Time.Format(timeBase[:10]))), nil
		case FormatTime:
			return []byte(strconv.Quote(x.Time.Format(timeBase[11:]))), nil
		}
	}
	return nil, nil
}

//json解码
//支持FormatTimestamp, FormatYear, FormatDatetime, FormatDate格式
//FormatTimestamp只接收number，其他只接收string
//解码成功后，XTime.Time.loc == time.Local
//默认解码类型为 FormatTimestamp
func (x *XTime) UnmarshalJSON(b []byte) error {

	b = bytes.TrimSpace(b)

	if len(b) == 8 {
		return errors.New("XType unmarshal json not support FormatTime. ")
	}

	if bytes.Equal(b, jsonNull) {
		x.Valid = false
		return nil
	} else {
		if bytes.HasPrefix(b, quoteBytes) && bytes.HasSuffix(b, quoteBytes) {

			s, e := strconv.Unquote(string(b))

			if e != nil {
				x.Valid = false
				return errors.New(fmt.Sprint(string(b), " XTime unmarshal json invalid, ", e.Error()))
			} else {
				switch len(s) {
				case 4, 10, 19:
					t, e := time.ParseInLocation(timeBase[:len(s)], s, time.Local)
					if e != nil {
						x.Valid = false
						return errors.New(fmt.Sprint(string(b), " XTime unmarshal json invalid, ", e.Error()))
					} else {
						x.Time, x.Valid = t, true
						switch len(s) {
						case 4:
							x.Format = FormatYear
						case 10:
							x.Format = FormatDate
						case 19:
							x.Format = FormatDatetime
						}
						return nil
					}
				default:
					x.Valid = false
					return errors.New(fmt.Sprint(string(b), " XTime unmarshal json invalid"))
				}
			}

		} else {
			if len(b) == 13 {
				timestamp, e := strconv.ParseInt(string(b), 10, 64)
				if e != nil {
					x.Valid = false
					return errors.New(fmt.Sprint(string(b), " XTime unmarshal json invalid."))
				} else {
					x.Time = time.Unix(0, timestamp * int64(time.Millisecond))
					x.Valid, x.Format = true, FormatTimestamp
					return nil
				}

			} else {
				x.Valid = false
				return errors.New(fmt.Sprint(string(b), " XTime unmarshal json invalid."))
			}
		}
	}

}
