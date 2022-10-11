package validators

import "testing"

func TestIsIntEmpty(t *testing.T) {
	emptyValue := 0
	isEmptyInt := IsIntEmpty(emptyValue)

	if !isEmptyInt {
		t.Errorf("value is not empty, expected: %v got %v", true, isEmptyInt)
	}

	notEmptyValue := 3
	isEmptyInt = IsIntEmpty(notEmptyValue)

	if isEmptyInt {
		t.Errorf("value is empty, expected: %v got %v", true, isEmptyInt)
	}
}

func TestIsUintEmpty(t *testing.T) {
	emptyValue := 0
	isEmptyInt := IsUintEmpty(uint(emptyValue))

	if !isEmptyInt {
		t.Errorf("value is not empty, expected: %v got %v", true, isEmptyInt)
	}

	notEmptyValue := 3
	isEmptyInt = IsUintEmpty(uint(notEmptyValue))

	if isEmptyInt {
		t.Errorf("value is empty, expected: %v got %v", true, isEmptyInt)
	}
}
