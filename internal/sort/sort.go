package sort

import (
	"bufio"
	"disksort/internal/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"unicode/utf8"
)

type DiskSort struct {
	InputFileDir string
	OutputDir    string
	Limit        int
}

func (d *DiskSort) Sort() (*string, error) {
	if err := d.splitFiles(); err != nil {
		return nil, err
	}

	if err := d.mergeFiles(); err != nil {
		return nil, err
	}
	var files []string
	if err := filepath.Walk(d.OutputDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(files) > 1 {
		return nil, errors.New("something went wrong. more than one output file")
	}

	return &files[0], nil
}

// splitFiles splits the input files into multiple smaller and sorted files of size Limit
func (d *DiskSort) splitFiles() error {
	var files []string
	if err := filepath.Walk(d.InputFileDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}); err != nil {
		return err
	}

	createNewFile := true
	var logs []string
	var tmpFile *os.File
	for _, inputFile := range files {
		file, err := os.Open(inputFile)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			var err error
			if logs == nil && createNewFile {
				if tmpFile, err = utils.CreateTempFile(d.OutputDir); err != nil {
					return err
				}
			}
			if len(scanner.Text()) != 0 {
				logs = append(logs, scanner.Text())
			}
			if len(logs) == d.Limit {
				sort.SliceStable(logs, func(i, j int) bool {
					return logs[i] < logs[j]
				})
				for _, value := range logs {
					if utf8.ValidString(value) {
						if _, err = fmt.Fprintln(tmpFile, value); err != nil {
							return err
						}
					}
				}
				createNewFile = true
				logs = nil
			}
		}
		_ = file.Close()
	}
	if len(logs) != 0 {
		sort.SliceStable(logs, func(i, j int) bool {
			return logs[i] < logs[j]
		})
		for _, value := range logs {
			if utf8.ValidString(value) {
				if _, err := fmt.Fprintln(tmpFile, value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (d *DiskSort) mergeFiles() error {
	var files []string
	if err := filepath.Walk(d.OutputDir, func(path string, info os.FileInfo, err error) error {
		if !utils.IsEmptyFile(path) && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return err
	}
	var selectFiles []string
	for len(files) > 1 {
		if len(files) < d.Limit {
			selectFiles = files
			files = nil
		} else {
			selectFiles = files[:d.Limit]
			files = files[d.Limit:]
		}
		if err := d.mergeKFiles(selectFiles); err != nil {
			return err
		}
		files = nil
		if err := filepath.Walk(d.OutputDir, func(path string, info os.FileInfo, err error) error {
			if !utils.IsEmptyFile(path) && !info.IsDir() {
				files = append(files, path)
			}
			return nil
		}); err != nil {
			return err
		}

	}
	return nil
}

type entry struct {
	log   string
	index int
}

func (d *DiskSort) mergeKFiles(files []string) error {
	var fileScanner []*bufio.Scanner
	var tmpFile *os.File
	var err error
	if tmpFile, err = utils.CreateTempFile(d.OutputDir); err != nil {
		return err
	}

	var logs []entry
	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		defer file.Close()
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		fileScanner = append(fileScanner, scanner)
		if fileScanner[i].Scan() {
			logs = append(logs, entry{
				log:   fileScanner[i].Text(),
				index: i,
			})

		}
	}

	count := 0
	for count < len(files) {
		index := 0
		if len(logs) == (len(files) - count) {
			sort.SliceStable(logs, func(i, j int) bool {
				return logs[i].log < logs[j].log
			})
			if utf8.ValidString(logs[0].log) {
				if _, err := fmt.Fprintln(tmpFile, logs[0].log); err != nil {
					return err
				} else {
					index = logs[0].index
					logs = logs[1:]

				}
			}

		}

		if fileScanner[index].Scan() {
			logs = append(logs, entry{
				log:   fileScanner[index].Text(),
				index: index,
			})

		} else {
			count++
		}
	}

	if len(logs) != 0 {
		sort.SliceStable(logs, func(i, j int) bool {
			return logs[i].log < logs[j].log
		})
		for _, value := range logs {
			if utf8.ValidString(value.log) {
				if _, err := fmt.Fprintln(tmpFile, value.log); err != nil {
					return err
				}
			}
		}

	}

	if err := utils.DeleteFiles(files); err != nil {
		return err
	}
	return nil
}
