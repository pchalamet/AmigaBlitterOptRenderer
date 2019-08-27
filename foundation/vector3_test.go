package foundation


import "testing"


func TestVector3Builder(t *testing.T) {
	v := NewVector3(1, 2, 3)
	if v.X != 1 || v.Y != 2 || v.Z != 3 {
		t.Error()
	}
}