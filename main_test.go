package main

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	t.Log(fibonacci(30))
}

func Test_fibonacci2(t *testing.T) {
	t.Log(fibonacci2(30))
}

func TestNegativeMod(t *testing.T) {
	t.Log((-4) % 5)
}
