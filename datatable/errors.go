package datatable

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyColumnNames = errors.New("empty column names")
	ErrColumnNotFound   = errors.New("column not found")
	ErrColumnNotNumeric = errors.New("column is not numeric")
	ErrDataNotFound     = errors.New("data not found")
	ErrEmptyColumns     = errors.New("number of columns must be greater than 0")
	ErrOverMaxColumns   = fmt.Errorf("number of columns must be smaller or equal to %d", MaxColumns)
	ErrOverMaxRows      = fmt.Errorf("number of rows must be smaller or equal to %d", MaxRows)
	ErrInvalidRowSize   = errors.New("row size does not equal to number of columns")
)
