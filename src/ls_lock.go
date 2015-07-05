package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

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

func main() {
	var target_dir string
	target_dir, status := get_target_directory()
	if status != 0 {
		fmt.Printf("Provide target dir\n")
		os.Exit(status)
	}
	files_in_dir, _ := get_inodes_in_directory(target_dir)
	fmt.Printf("%#v", files_in_dir)
	fmt.Printf(target_dir)
}
