package pub_sub

import "testing"

func TestAdd(t *testing.T) {
	ble := Add(1, 2)
	if ble != 3 {
		t.Fatalf("Adding 1 + 2 should be 3 but got: %d", ble)
	}
}
