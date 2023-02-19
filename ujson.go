package ujson

import (
	"encoding/json"
	"sort"
)

func Unmarshal(data []byte) (Any, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return decode(v, false), nil
}

func UnmarshalWithSort(data []byte) (Any, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return decode(v, true), nil
}

func decode(data interface{}, needSort bool) Any {
	switch v := data.(type) {
	case map[string]interface{}:
		a := Object{
			data: make(map[string]Any),
		}
		for key, value := range v {
			a.data[key] = decode(value, needSort)
		}
		return a
	case []interface{}:
		a := Array{}
		for i := range v {
			a.Append(decode(v[i], needSort))
		}
		if needSort {
			a.Sort()
		}
		return a
	case bool:
		return Bool{data: v}
	case string:
		return String{data: v}
	case int64:
		return NumberInt{data: v}
	case uint64:
		return NumberUint{data: v}
	case float64:
		return NumberFloat{data: v}
	case nil:
		return Null{}
	}
	return Null{}
}

type T string

const (
	TObject      T = "Object"
	TArray       T = "Array"
	TNumber      T = "Number"
	TNumberInt   T = "NumberInt"
	TNumberUint  T = "NumberUint"
	TNumberFloat T = "NumberFloat"
	TString      T = "String"
	TBool        T = "Bool"
	TNull        T = "Null"
)

func (t *T) IsNumber() bool {
	if *t == TNumberInt || *t == TNumberUint || *t == TNumberFloat {
		return true
	}
	return false
}

type Any interface {
	T() T
	String() string
	MarshalJSON() ([]byte, error)
}

type Object struct {
	data map[string]Any
}

func (o Object) T() T     { return TObject }
func (o Object) Len() int { return len(o.data) }
func (o Object) Keys() []string {
	keys := make([]string, o.Len())
	i := 0
	for k := range o.data {
		keys[i] = k
		i++
	}
	return keys
}
func (o Object) Value(key string) (v Any, ok bool) {
	v, ok = o.data[key]
	return
}
func (o Object) String() string {
	v, _ := json.Marshal(o.data)
	return string(v)
}
func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.data)
}

type Array struct {
	data []Any
}

func (a Array) T() T { return TArray }
func (a *Array) Append(v Any) {
	a.data = append(a.data, v)
}
func (a *Array) Len() int { return len(a.data) }
func (a *Array) Index(i int) Any {
	return a.data[i]
}
func (a Array) String() string {
	v, _ := json.Marshal(a.data)
	return string(v)
}
func (a Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.data)
}
func (a *Array) Sort() {
	sort.SliceStable(a.data, func(i, j int) bool {
		ai, _ := json.Marshal(a.data[i])
		aj, _ := json.Marshal(a.data[j])
		return string(ai) < string(aj)
	})
}

type Number interface {
	T() T
	String() string
	MarshalJSON() ([]byte, error)
}

type NumberInt struct {
	data int64
}

func (n NumberInt) T() T         { return TNumberInt }
func (n NumberInt) Int64() int64 { return n.data }
func (n NumberInt) String() string {
	v, _ := json.Marshal(n.data)
	return string(v)
}
func (n NumberInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.data)
}

type NumberUint struct {
	data uint64
}

func (n NumberUint) T() T           { return TNumberUint }
func (n NumberUint) Uint64() uint64 { return n.data }
func (n NumberUint) String() string {
	v, _ := json.Marshal(n.data)
	return string(v)
}
func (n NumberUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.data)
}

type NumberFloat struct {
	data float64
}

func (n NumberFloat) T() T             { return TNumberFloat }
func (n NumberFloat) Float64() float64 { return n.data }
func (n NumberFloat) String() string {
	v, _ := json.Marshal(n.data)
	return string(v)
}
func (n NumberFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.data)
}

type String struct {
	data string
}

func (s String) T() T           { return TString }
func (s String) String() string { return s.data }
func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.data)
}

type Bool struct {
	data bool
}

func (b Bool) T() T       { return TBool }
func (b Bool) Bool() bool { return b.data }
func (b Bool) String() string {
	v, _ := json.Marshal(b.data)
	return string(v)
}
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.data)
}

type Null struct {
}

func (n Null) T() T { return TNull }
func (n Null) String() string {
	return ""
}
func (n Null) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func ObjectOK(a Any) (o Object, ok bool) {
	o, ok = a.(Object)
	return
}
func ArrayOK(a Any) (o Array, ok bool) {
	o, ok = a.(Array)
	return
}
func NumberIntOK(a Any) (o NumberInt, ok bool) {
	o, ok = a.(NumberInt)
	return
}
func NumberUintOK(a Any) (o NumberUint, ok bool) {
	o, ok = a.(NumberUint)
	return
}
func NumberFloatOK(a Any) (o NumberFloat, ok bool) {
	o, ok = a.(NumberFloat)
	return
}
func StringOK(a Any) (o String, ok bool) {
	o, ok = a.(String)
	return
}
func BoolOK(a Any) (o Bool, ok bool) {
	o, ok = a.(Bool)
	return
}
func NullOK(a Any) (o Null, ok bool) {
	o, ok = a.(Null)
	return
}
