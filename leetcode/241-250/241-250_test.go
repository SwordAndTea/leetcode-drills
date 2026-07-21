package _241_250

import "testing"

func TestGroupStrings(t *testing.T) {
	t.Logf("%#v", groupStrings([]string{
		"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z",
	}))

	t.Logf("%#v", groupStrings([]string{
		"acb", "bdc", "acef", "xyz", "az", "ba", "a", "z",
	}))
}
