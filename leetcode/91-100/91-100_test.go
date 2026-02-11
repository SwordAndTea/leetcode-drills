package _91_100

import "testing"

func TestTmp(t *testing.T) {
	t.Log(restoreIpAddresses("25525511135"))
}

func Test_generateTrees(t *testing.T) {
	t.Log(generateTrees(3))
}

func Test_isInterleave(t *testing.T) {
	t.Log(isInterleave(
		"bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa",
		"babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab",
		"babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"))
}
