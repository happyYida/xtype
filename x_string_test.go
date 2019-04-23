package xtype

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNewXString(t *testing.T) {
	xs := NewXString("Hello")
	t.Logf("%#v", xs)

	xs = NewXString("")
	t.Logf("%#v", xs)
}

func TestXString_MarshalJSON(t *testing.T) {

	type Person struct {
		Name XString
	}

	p := Person{Name: XString{String:"", Valid:true}}
	b, e := json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{Name: XString{String:"", Valid:false}}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{Name: XString{String:"hello", Valid:false}}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{Name: XString{String:"hello", Valid:true}}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	p = Person{Name: XString{String:"<xml>ok</xml>", Valid:true}}
	b, e = json.Marshal(p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%s", b)
	}

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(p)
	t.Log(buffer.String())
}

func TestXString_UnmarshalJSON(t *testing.T) {
	type Person struct {
		Name XString
	}

	j := []byte(`{"Name": ""}`)
	p := new(Person)
	e := json.Unmarshal(j, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}


	j = []byte(`{}`)
	p = new(Person)
	e = json.Unmarshal(j, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}

	j = []byte(`{"Name": null}`)
	p = new(Person)
	e = json.Unmarshal(j, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}

	j = []byte(`{"Name": "yangyingjun"}`)
	p = new(Person)
	e = json.Unmarshal(j, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
	}

	j = []byte(`{"Name": "1233333"}`)
	p = new(Person)
	e = json.Unmarshal(j, p)
	if e != nil {
		t.Error(e)
	} else {
		t.Logf("%#v", *p)
		t.Logf("%v", p.Name)
	}
}
