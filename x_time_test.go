package xtype

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewXTime(t *testing.T) {
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-10 10:10:10", time.Local)

	a := NewXTime(p, FormatTimestamp)
	t.Logf("%+v", a)

	a = NewXTime(p, FormatYear)
	t.Logf("%+v", a)

	a = NewXTime(p, FormatDatetime)
	t.Logf("%+v", a)

	a = NewXTime(p, FormatDate)
	t.Logf("%+v", a)

	a = NewXTime(p, FormatTime)
	t.Logf("%+v", a)
}

func TestXTime_MarshalJSON(t *testing.T) {

	type Person struct {
		Birthday XTime
	}

	f, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-10 10:10:10", time.Local)

	p := Person{
		Birthday:NewXTime(f, FormatDatetime),
	}
	b, e := json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}


	p = Person{
		Birthday:NewXTime(f, FormatDate),
	}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{
		Birthday:NewXTime(f, FormatTime),
	}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{
		Birthday:NewXTime(f, FormatYear),
	}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{
		Birthday:NewXTime(f, FormatTimestamp),
	}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

}

func TestXTime_UnmarshalJSON(t *testing.T) {
	type Range struct {
		End XTime
	}
	r := new(Range)

	jsonBuf := []byte(`{"End": null}`)
	err := json.Unmarshal(jsonBuf, r)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%#v", *r)
	}

	jsonBuf = []byte(`{"End": 1556205895873}`)
	err = json.Unmarshal(jsonBuf, r)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", *r)
	}

	jsonBuf = []byte(`{"End": "1991"}`)
	err = json.Unmarshal(jsonBuf, r)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", *r)
	}

	jsonBuf = []byte(`{"End": "1992-12-13"}`)
	err = json.Unmarshal(jsonBuf, r)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", *r)
	}

	jsonBuf = []byte(`{"End": "1993-01-20 12:13:30"}`)
	err = json.Unmarshal(jsonBuf, r)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", *r)
	}


}
