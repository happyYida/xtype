package xtype

import (
	"encoding/json"
	"testing"
)

func TestNewXInt64(t *testing.T) {
	v := NewXInt64(20)
	t.Logf("%#v\n", v)
}

func TestXInt64_MarshalJSON0(t *testing.T) {
	type Person struct {
		Age XInt64
	}

	p0 := Person{Age: NewXInt64(20)}

	if b, e := json.Marshal(p0); e != nil {
		t.Error(e)
	}else {
		t.Logf("not null --> %s\n", b)
	}

	p1 := Person{}

	if b, e := json.Marshal(p1); e != nil {
		t.Error(e)
	}else {
		t.Logf("null --> %s\n", b)
	}
}

func TestXInt64_MarshalJSON1(t *testing.T) {
	type Person struct {
		Age *XInt64
	}

	p0 := Person{Age: &XInt64{Int64:20, Valid:true}}

	if b, e := json.Marshal(p0); e != nil {
		t.Error(e)
	}else {
		t.Logf("not null --> %s\n", b)
	}

	p1 := Person{}

	if b, e := json.Marshal(p1); e != nil {
		t.Error(e)
	}else {
		t.Logf("null --> %s\n", b)
	}

	p2 := Person{Age: &XInt64{Valid:false}}

	if b, e := json.Marshal(p2); e != nil {
		t.Error(e)
	}else {
		t.Logf("null --> %s\n", b)
	}
}

func TestXInt64_UnmarshalJSON0(t *testing.T) {
	type Person struct {
		Age XInt64
	}

	json_bytes := []byte(`{"Age": "20"}`)
	p := new(Person)
	if e := json.Unmarshal(json_bytes, p); e != nil {
		t.Log("ERROR ", e)
	} else {
		t.Logf("%#v\n", *p)
	}

	json_bytes = []byte(`{"Age": 20.0}`)
	p = new(Person)
	if e := json.Unmarshal(json_bytes, p); e != nil {
		t.Log("ERROR ", e)
	} else {
		t.Logf("%#v\n", *p)
	}

	json_bytes = []byte(`{"Age": null}`)
	p = new(Person)
	if e := json.Unmarshal(json_bytes, p); e != nil {
		t.Log("ERROR ", e)
	} else {
		t.Logf("%#v\n", *p)
	}

	json_bytes = []byte(`{"Age": 30}`)
	p = new(Person)
	if e := json.Unmarshal(json_bytes, p); e != nil {
		t.Log("ERROR ", e)
	} else {
		t.Logf("%#v\n", *p)
	}
}
