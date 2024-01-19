package helper

import "testing"

func TestNegative(t *testing.T) {
	result := MakeIndex(-1, 4)
	expected := 3

	if result != expected {
		t.Errorf("negative number does not turn to right array number")
	}

}

func TestOutOfRangeOne(t *testing.T) {
	result := MakeIndex(4, 4)
	expected := 0

	if result != expected {
		t.Errorf("out of range number does not turn to right array number")
	}
}

func TestOutOfRangeMoreThanOne(t *testing.T) {
	result := MakeIndex(5, 4)
	expected := 1

	if result != expected {
		t.Errorf("out of range number does not turn to right array number")
	}
}
