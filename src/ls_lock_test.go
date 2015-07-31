package main

import "testing"

func TestGetInodesInDirectory(t *testing.T) {
	c, _ := get_inodes_in_directory(".")
	if len(c.inode_to_file) == 0 {
		t.Error("Expected to get inodes for the current directory")
	}
}
