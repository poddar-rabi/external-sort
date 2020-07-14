package sort

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiskSort_Sort(t *testing.T) {
	inputDir, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(inputDir)

	file, _ := ioutil.TempFile(inputDir, "temp.*.log")
	defer os.Remove(file.Name())

	_, _ = fmt.Fprintln(file, "dddd")
	_, _ = fmt.Fprintln(file, "aaaa")
	_, _ = fmt.Fprintln(file, "cccc")
	_, _ = fmt.Fprintln(file, "bbbb")

	outputDir, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(outputDir)

	sorter := DiskSort{
		InputFileDir: inputDir,
		OutputDir:    outputDir,
		Limit:        2,
	}
	sortedFile, err := sorter.Sort()
	sFile, err := os.Open(*sortedFile)
	scanner := bufio.NewScanner(sFile)
	scanner.Split(bufio.ScanLines)
	assert.True(t, scanner.Scan())
	assert.Equal(t, scanner.Text(), "aaaa")
	assert.True(t, scanner.Scan())
	assert.Equal(t, scanner.Text(), "bbbb")
	assert.True(t, scanner.Scan())
	assert.Equal(t, scanner.Text(), "cccc")
	assert.True(t, scanner.Scan())
	assert.Equal(t, scanner.Text(), "dddd")
	assert.Nil(t, err)
}
