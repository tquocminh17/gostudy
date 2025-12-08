package testmod_test

import (
	"gostudy/testmod"
	"testing"
)

func TestTest(t *testing.T) {
	if testmod.Aaa() != "Zek" {
		t.Fatal("Wrong answer")
	}
}
