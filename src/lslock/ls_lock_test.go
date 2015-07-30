package main

import (
	"ls_lock"
	"testing"
)

func Test(t *testing.T) {
	c, _ := ls_lock.get_inodes_in_directory(".")
}
