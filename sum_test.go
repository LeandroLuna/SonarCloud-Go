package main

import "testing"

func TestSum(t *testing.T) {
	if sum(5, 10) != 15 {
		t.Fail()
	}
}
