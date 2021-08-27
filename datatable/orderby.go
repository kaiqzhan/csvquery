package datatable

import (
	"fmt"
	"sort"
	"strconv"
)

// OrderBy returns a data table sorted by the specified column name in descending order.
func (dt *DataTable) OrderBy(columnName string) (DataTable, error) {
	// get column index
	columnIndex, err := dt.columnIndexByName(columnName)
	if err != nil {
		return DataTable{}, err
	}

	// check if column only contains numeric data
	if !dt.isColumnNumeric(columnIndex) {
		return DataTable{}, fmt.Errorf("order by column '%s': %w", columnName, ErrColumnNotNumeric)
	}

	// make a copy of the original rows
	outRows := append([]DataRow(nil), dt.Rows...)

	// sort rows by the specific column
	sort.Slice(outRows, func(i, j int) bool {
		// convert data to int
		// ignored error since it has been checked in isColumnNumeric()
		iNum, _ := strconv.Atoi(outRows[i].Items[columnIndex].Data)
		jNum, _ := strconv.Atoi(outRows[j].Items[columnIndex].Data)
		return jNum < iNum // descending order
	})

	return DataTable{
		Columns: dt.Columns,
		Rows:    outRows,
	}, nil
}
