package lslock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInodesInDirectory(t *testing.T) {
	c, _ := GetInodesInDirectory(".")
	assert.NotEqual(t, c.iNodeToFile, 0, "We should have found some inodes")
}

func TestGetFlocksInodes(t *testing.T) {
	content := `88: POSIX  ADVISORY  WRITE 16967 08:05:134864 0 EOF
89: FLOCK  ADVISORY  WRITE 6606 08:05:142013 0 EOF
90: FLOCK  ADVISORY  WRITE 49205350 00:1f:183 0 EOF
91: FLOCK  ADVISORY  WRITE 6601 00:1f:182 0 EOF
92: FLOCK  ADVISORY  WRITE 49205353 00:1f:124 0 EOF
93: FLOCK  ADVISORY  WRITE 49205354 00:1f:115 0 EOF
94: FLOCK  ADVISORY  WRITE 2088 00:10:13759 0 EOF
95: FLOCK  ADVISORY  WRITE 3932 00:42:76392 0 EOF`
	inodes, _ := GetFlocksInodes(content)
	assert.NotEqual(t, inodes, 0, "")
}

func TestFindCommon(t *testing.T) {

	
}
