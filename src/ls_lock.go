package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

const LOCK_PROC_FILE string = "/proc/locks"
const FIELD_NUM_LC_TYPE int32 = 1
const FIELD_NUM_INODE = 6
const FLOCK string = "FLOCK"

func get_inodes_in_directory(target_dir string) (FileInodeCollector, error) {
	var i FileInodeCollector
	i.inode_to_file = make(map[uint64]string)
	err := filepath.Walk(target_dir, i.extract_inode)
	return i, err
}

type FileInodeCollector struct {
	inode_to_file map[uint64]string
}

func (in *FileInodeCollector) extract_inode(path string, f os.FileInfo, err error) error {
	stat, _ := f.Sys().(*syscall.Stat_t)
	in.inode_to_file[stat.Ino] = path
	return nil
}

//
// Get the target directory
//
func get_target_directory() (string, int) {
	flag.Parse()
	d := flag.Arg(0)
	if len(d) == 0 {
		return "", 1
	}
	return d, 0
}

// Complexity
// O(n)
func get_indes_used_in_flocks(os_lock_file string) ([]uint64, error) {

	d, e := ioutil.ReadFile(os_lock_file)
	check(e)
	lines := strings.Split(string(d), "\n")

	// ensure that we do not need to reallocate
	// trade memory for complexity
	var lockedInodes []uint64 = make([]uint64, len(lines), len(lines))

	// O(n)
	for _, l := range lines {
		if len(l) > 0 {
			fields := strings.Split(l, " ")
			if fields[FIELD_NUM_LC_TYPE] == FLOCK {
				inode := fields[FIELD_NUM_INODE]
				fmt.Printf("%v\n", inode)
				v, err := strconv.Atoi(inode)
				if err != nil {
					return nil, err
				}
				// we allocate eagerly, so no
				// resize should happen
				lockedInodes = append(lockedInodes, uint64(v))
			}
		}
	}
	return lockedInodes, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var target_dir string
	target_dir, status := get_target_directory()
	if status != 0 {
		fmt.Printf("Provide target dir\n")
		os.Exit(status)
	}
	collector, _ := get_inodes_in_directory(target_dir)
	inodes, _ := get_indes_used_in_flocks("example.txt")
	for _, i := range inodes {

		if len(collector.inode_to_file[i]) != 0 {
			// DO STH
		} else {
			// DO STH ELESE
		}
	}
}
