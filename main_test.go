package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const duplicateFilesDirectory = "duplicates_files_directory"

// Approach-2 with go testing library with multiple multiple tests scenerios(Table Driven Test)
type toReadableSizeTest struct {
	nbytes   int
	expected string
}

var toReadableSizeTests = []toReadableSizeTest{
	{1000, "1000 B"},
	{1000 * 1000, "1000 KB"},
	{1000 * 1000 * 1000, "1000 MB"},
	{1000 * 1000 * 1000 * 1000, "1000 GB"},
}

func TestToReadableSizeMultiple(t *testing.T) {

	for _, test := range toReadableSizeTests {
		if output := toReadableSize(int64(test.nbytes)); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// TestTraverseDir: UT to test traverseDir method.
func TestTraverseDir(t *testing.T) {
	hashes := map[string]string{}
	duplicates := map[string]string{}
	var dupeSize int64

	cwd, err := os.Getwd()
	assert.Nil(t, err)

	directory := fmt.Sprintf("%s/%s", cwd, duplicateFilesDirectory)
	dirFiles, err := ioutil.ReadDir(directory)
	assert.Nil(t, err)

	traverseDir(hashes, duplicates, &dupeSize, dirFiles, directory)

	assert.Equal(t, 3, len(hashes))
	assert.Equal(t, 3, len(duplicates))

	assert.Equal(t, int64(132), dupeSize)
}
