package public

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSorterFailsInputDirDoesntExist(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	client, err := NewSorter("dummy", f, 5)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewSorterFailsInputDirEmpty(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)

	client, err := NewSorter(f, "dummy", 5)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewSorterFailsOutputDirDoesntExist(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	file, _ := ioutil.TempFile(f, "temp.*.log")
	defer os.Remove(file.Name())
	client, err := NewSorter(f, "dummy",  5)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewSorterFailsLimitNotMinimum(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	file, _ := ioutil.TempFile(f, "temp.*.log")
	defer os.Remove(file.Name())
	f1, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f1)
	client, err := NewSorter(f, f1,  1)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewSorter(t *testing.T) {
	f, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f)
	file, _ := ioutil.TempFile(f, "temp.*.log")
	defer os.Remove(file.Name())
	f1, _ := ioutil.TempDir("", "temp"+string(time.Now().Unix()))
	defer os.Remove(f1)
	client, err := NewSorter(f, f1,  2)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}
