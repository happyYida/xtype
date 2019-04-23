package xtype

import (
	"encoding/json"
	"testing"
)

func TestNewXBool(t *testing.T) {
	b := NewXBool(false)
	t.Logf("%#v", b)

	b = NewXBool(true)
	t.Logf("%#v", b)
}

func TestXBool_MarshalJSON(t *testing.T) {
	type Person struct {
		IsMan XBool
	}

	var p Person

	p = Person{IsMan:XBool{Bool:false, Valid:true}}
	if b, e := json.Marshal(p); e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{IsMan:XBool{Bool:true, Valid:false}}
	if b, e := json.Marshal(p); e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{IsMan:XBool{Bool:true, Valid:true}}
	if b, e := json.Marshal(p); e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}
}

func TestXBool_MarshalJSON2(t *testing.T) {
	type Person struct {
		IsMan *XBool
	}

	bf, e := json.Marshal(Person{IsMan: &XBool{Bool:true, Valid:true}})
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", bf)
	}

	bf, e = json.Marshal(Person{IsMan: &XBool{Bool:false, Valid:true}})
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", bf)
	}

	bf, e = json.Marshal(Person{IsMan: &XBool{Bool:true, Valid:false}})
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", bf)
	}

	bf, e = json.Marshal(Person{IsMan: &XBool{Bool:false, Valid:false}})
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", bf)
	}


}

func TestXBool_UnmarshalJSON(t *testing.T) {
	type Person struct {
		IsMan XBool
	}

	bf := []byte(`{"IsMan":true}`)
	p := new(Person)
	e := json.Unmarshal(bf, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}

	bf = []byte(`{"IsMan":false}`)
	p = new(Person)
	e = json.Unmarshal(bf, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}

	bf = []byte(`{"IsMan":null}`)
	p = new(Person)
	e = json.Unmarshal(bf, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}
}

func TestXBool_UnmarshalJSON2(t *testing.T) {
	type Person struct {
		IsMan *XBool
	}

	var p Person
	bf := []byte(`{"IsMan":false}`)

	e := json.Unmarshal(bf, &p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p.IsMan)
	}
}

func TestXBool_UnmarshalJSON3(t *testing.T) {
	type Person struct {
		IsMan XBool
	}

	bf := []byte(`{"IsMan":123}`)
	p := new(Person)
	e := json.Unmarshal(bf, p)
	if e != nil {
		t.Log(e)
	} else {
		t.Errorf("%#v", *p)
	}
}
