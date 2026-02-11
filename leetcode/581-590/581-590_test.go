package _581_590

import "testing"

func TestFileSystem(t *testing.T) {
	fs := Constructor()
	t.Logf("%d, %+v", len(fs.Ls("/")), fs.Ls("/"))
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	t.Logf("%d, %+v", len(fs.Ls("/")), fs.Ls("/"))
	t.Logf("%s", fs.ReadContentFromFile("/a/b/c/d"))
}
