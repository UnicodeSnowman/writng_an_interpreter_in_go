package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "Other One"}
	diff2 := &String{Value: "Other One"}
	boolTrue1 := &Boolean{Value: true}
	boolTrue2 := &Boolean{Value: true}
	boolFalse1 := &Boolean{Value: false}
	boolFalse2 := &Boolean{Value: false}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content must have same hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content must have same hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content must have different hash keys")
	}

	if boolTrue1.HashKey() != boolTrue2.HashKey() {
		t.Errorf("bools with same content must have same hash keys")
	}

	if boolFalse1.HashKey() != boolFalse2.HashKey() {
		t.Errorf("bools with same content must have same hash keys")
	}

	if boolTrue1.HashKey() == boolFalse1.HashKey() {
		t.Errorf("bools with different content must have different hash keys")
	}
}
