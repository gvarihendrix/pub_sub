package pub_sub

import "testing"

func TestAdd(t *testing.T) {
	ble := Add(1, 2)
	if ble != 3 {
		t.Fatalf("Adding 1 + 2 should be 3 but got: %d", ble)
	}
}

func BadTest(t *testing.T) {
	num := Add(2, 2)
	if num != 5 {
		t.Fatalf("Adding 2 + 2 should be 4 but got: %d", num)
	}
}
