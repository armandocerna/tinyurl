package tinyurl

import "testing"

func TestCreateRandomString(t *testing.T) {
	r := CreateRandomString()
	if len(r) != 8 {
		t.Fatalf("Got %d, length instead of 8: %s", len(r), r)
	}
	t.Log(r)
}
