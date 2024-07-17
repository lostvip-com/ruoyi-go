package dtos

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type ImportFile struct {
	Excel *excelize.File
}

func NewImportFile(f io.Reader) (*ImportFile, error) {
	file, err := excelize.OpenReader(f)
	if err != nil {
		return nil, err
	}
	return &ImportFile{
		Excel: file,
	}, nil
}
