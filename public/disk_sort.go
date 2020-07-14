package public

import (
	"disksort/internal/sort"
	"disksort/internal/utils"
	"errors"
)

type Sort interface {
	Sort() (*string, error)
}

func NewSorter(inputFileDir, outputDir string, limit int) (Sort, error) {
	if !utils.IsDirExist(inputFileDir) {
		return nil, errors.New("please specify an existing input directory")
	}

	empty, err := utils.IsDirEmpty(inputFileDir)
	if err != nil {
		return nil, err
	}
	if empty {
		return nil, errors.New("please specify an input directory containing " +
			"only the files to be sorted")
	}

	if !utils.IsDirExist(outputDir) {
		return nil, errors.New("please specify an existing temp directory")
	}

	empty, err = utils.IsDirEmpty(outputDir)
	if err != nil {
		return nil, err
	}
	if !empty {
		return nil, errors.New("please specify an empty temp dir")
	}

	if limit < 2 {
		return nil, errors.New("please specify minimum limit 2")
	}

	return &sort.DiskSort{
		InputFileDir: inputFileDir,
		OutputDir:    outputDir,
		Limit:        limit,
	}, nil
}
