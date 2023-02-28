package ascii

import (
	"testing"
)

func Test_getID(t *testing.T) {
	str, err := GetID("th301_giangvien")
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(str)
}
