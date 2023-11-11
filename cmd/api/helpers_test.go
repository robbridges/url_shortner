package main

import "testing"

func TestStringWithCharset(t *testing.T) {
	string := StringWithCharset(5, charset)
	if len(string) != 5 {
		t.Errorf("Expected string length of 5, but got %d", len(string))
	}

}
