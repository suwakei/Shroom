package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

type DictKey struct {
	Type  ObjectType
	Value uint64
}

func (boolean *Boolean) DictKey() DictKey {
	var value uint64

	if boolean.Value {
		value = 1
	} else {
		value = 0
	}

	return DictKey{Type: boolean.Type(), Value: value}
}

func (integer *Integer) DictKey() DictKey {
	return DictKey{Type: integer.Type(), Value: uint64(integer.Value)}
}

func (str *String) DictKey() DictKey {
	d := fnv.New64a()
	d.Write([]byte(str.Value))

	return DictKey{Type: str.Type(), Value: d.Sum64()}
}

type DictPair struct {
	Key   Object
	Value Object
}

type Dict struct {
	Pairs map[DictKey]DictPair
}

type HashableDict interface {
	DictKey() DictKey
}

func (d *Dict) Type() ObjectType { return DICT_OBJ }

func (d *Dict) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range d.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
