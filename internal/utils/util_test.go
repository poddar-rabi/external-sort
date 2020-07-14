package utils

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFiles(t *testing.T) {
	var files []string
	f, _ := ioutil.TempFile("", "temp.*.log")
	defer os.Remove(f.Name())
	files = append(files, f.Name())
	err := DeleteFiles(files)
	assert.NoError(t, err)
}

func TestIsDirExist(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	empty := IsDirExist(f)
	assert.True(t, empty)
}

func TestIsDirEmpty(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	empty, err := IsDirEmpty(f)
	assert.NoError(t, err)
	assert.True(t, empty)
}

func TestCreateTempFile(t *testing.T) {
	f, err := CreateTempFile("")
	defer os.Remove(f.Name())
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

func TestIsEmptyFile(t *testing.T) {
	f, _ := CreateTempFile("")
	defer os.Remove(f.Name())
	empty := IsEmptyFile(f.Name())
	assert.True(t, empty)
}
