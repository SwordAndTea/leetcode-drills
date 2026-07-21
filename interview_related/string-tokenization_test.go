package interview_related

import "testing"

func TestStringTokenization(t *testing.T) {
	t.Logf("%+v", stringTokenization("applepiepear", []string{"app:10", "apple:20", "pie:30"}))
	t.Logf("%+v", stringTokenization("acdebe", []string{"a:1", "b:2", "cd:3"}))
	t.Logf("%+v", stringTokenization("programmingprogrampropro", []string{"pro:1", "program:2", "programming:3", "gram:4", "ming:5", "pr:6", "og:7"}))
	t.Logf("%+v", stringTokenization("helloworldhelloxyzworld", []string{"hello:100", "world:200", "hel:50", "lo:51", "wor:52", "ld:53"}))
	t.Logf("%+v", stringTokenization("abcabcdabcdeabcdef", []string{"a:1", "ab:2", "abc:3", "abcd:4", "abcde:5", "abcdef:6", "bc:7", "cd:8", "de:9", "ef:10"}))
	t.Logf("%+v", stringTokenization("", []string{"a:1", "b:2"}))
	t.Logf("%+v", stringTokenization("abc", []string{}))
	t.Logf("%+v", stringTokenization("thetheatertheme", []string{"the:1", "theater:2", "theme:3", "heat:4", "at:5", "er:6", "he:7"}))
}
