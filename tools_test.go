package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	tool := Tools{}
	r := tool.RandomString(10)

	if len(r) != 10 {
		t.Error("Wrong lenght!")
	}
}