package main

import "testing"

func TestCreateNum(t *testing.T) {
	n := createNum()

	if len(n) != 11 {
		t.Errorf("Expected length of num is 11, but got %v", len(n))
	}
	if n[0] != 0 {
		t.Errorf("Expected first number of num is 0, but got %v", n[0])
	}
	if n[len(n)-1] != 10 {
		t.Errorf("Expected last number of num is 10, but got %v", n[len(n)-1])
	}
}
