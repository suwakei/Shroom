package object_test

import (
	"Shroom/object"
	"fmt"
	"testing"
)

func TestStringDictKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "I am keito"}
	diff2 := &object.String{Value: "I am keito"}

	if hello1.DictKey() != hello2.DictKey() {
		t.Errorf("strings with same content have diffrent Dict keys")
	}

	if diff1.DictKey() != diff2.DictKey() {
		t.Errorf("strings with same content have diffrent Dict keys")
	}

	if hello1.DictKey() == diff1.DictKey() {
		t.Errorf("strings with diffrent content have same Dict keys")
	}

	name1 := &object.String{Value: "name"}
	shroom := &object.String{Value: "Shroom"}

	pairs := map[object.DictKey]object.Object{}
	pairs[name1.DictKey()] = shroom

	fmt.Printf("pairs[name1.DictKey()]=%+v\n", pairs[name1.DictKey()])
	// => pairs[name1.DictKeys]=&{Value: Shroom}

	name2 := &object.String{Value: "name"}

	fmt.Printf("pairs[name2.DictKey()]=%+v\n", pairs[name2.DictKey()])
	// => pairs[name2.DictKeys]=&{Value: Shroom}

	fmt.Printf("(name1 == name2)=%t\n", name1 == name2)
	// => (name1 == name2)=false

	fmt.Printf("(name1.DictKey() == name2.DictKey())=%t\n", name1.DictKey() == name2.DictKey())
	// => (name1.DictKey() == name2,DictKey())=true
}
