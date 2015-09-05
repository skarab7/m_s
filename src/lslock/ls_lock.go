package lslock

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

const ProcLocksFile string = "/proc/locks"
const FieldNumLockType int32 = 1
const FieldNumINode = 6
const FileLockName string = "FLOCK"

type FileInodeCollector struct {
	iNodeToFile map[uint64]string
}

func GetInodesInDirectory(targetDir string) (FileInodeCollector, error) {
	var i FileInodeCollector
	i.iNodeToFile = make(map[uint64]string)
	err := filepath.Walk(targetDir, i.ExtractInode)
	return i, err
}

func (in *FileInodeCollector) ExtractInode(path string, f os.FileInfo, err error) error {
	stat, _ := f.Sys().(*syscall.Stat_t)
	in.iNodeToFile[stat.Ino] = path
	return nil
}

//
// Get the target directory
//
func GetTargetDirectory() (string, int) {
	flag.Parse()
	d := flag.Arg(0)
	if len(d) == 0 {
		return "", 1
	}
	return d, 0
}

// Complexity
// O(n)
func GetFlocksInodes(content string) ([]uint64, error) {
	lines := strings.Split(content, "\n")

	// ensure that we do not need to reallocate
	// trade memory for complexity
	var lockedInodes []uint64 = make([]uint64, len(lines), len(lines))

	// O(n)
	for _, l := range lines {
		if len(l) > 0 {
			fields := strings.Split(l, " ")
			if fields[FieldNumLockType] == FileLockName {
				inode := fields[FieldNumINode]
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

func GetLockFileContent() (string, error) {
	d, e := ioutil.ReadFile(ProcLocksFile)
	return string(d), e
}

func ExitIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func findLockedPaths(inodes []uint64, inodeToFile map[uint64]string)  map[string]uint64 {
	var r map[string] uint64 = make(map[string]uint64)
	// merge sort would be the way to go
	for _, i := range inodes {
		if len(inodeToFile[i]) != 0 {
			i2f := inodeToFile[i]
			r[i2f] = i
		}
	}
	return r
}

func main() {
	var targetDir string
	targetDir, status := GetTargetDirectory()
	if status != 0 {
		fmt.Printf("Provide target dir\n")
		os.Exit(status)
	}
	collector, _ := GetInodesInDirectory(targetDir)
	lockFileContent, e := GetLockFileContent()
	ExitIfError(e)
	inodes, _ := GetFlocksInodes(lockFileContent)

	lockedPaths := findLockedPaths(inodes, collector.iNodeToFile)

	for path, inode := range lockedPaths {
		fmt.Println("Path: ", path, " INODE: ", inode)
	}
}
