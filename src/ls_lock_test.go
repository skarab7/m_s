package main

import "testing"

func TestGetInodesInDirectory(t *testing.T) {
	c, _ := GetInodesInDirectory(".")
	if len(c.iNodeToFile) == 0 {
		t.Errorf("expected to get inodes for the current directory")
	}
}
