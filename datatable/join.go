package datatable

import "fmt"

// Join combines two data tables based on the same value on the specified on column.
// Column ordering should be preserved as if the first and second table had their columns concatenated,
// with the matching column name removed from the second column.
func (dt *DataTable) Join(rightTable DataTable, onColumnName string) (DataTable, error) {
	// get on column index of both tables
	onColumnIndexLeft, err := dt.columnIndexByName(onColumnName)
	if err != nil {
		return DataTable{}, err
	}

	onColumnIndexRight, err := rightTable.columnIndexByName(onColumnName)
	if err != nil {
		return DataTable{}, err
	}

	// join columns
	outColumns := joinColumns(dt.Columns, rightTable.Columns, onColumnIndexRight)

	// build lookup index for the on column of right table
	rowIndexMapRight := rightTable.rowIndexMapByColumn(onColumnIndexRight)

	// join rows
	outRows := make([]DataRow, 0, len(dt.Rows))
	for _, leftRow := range dt.Rows {
		// find corresponding row in the right table
		onDataLeft := leftRow.Items[onColumnIndexLeft].Data
		rightRowIndex, ok := rowIndexMapRight[onDataLeft]
		if !ok {
			return DataTable{}, fmt.Errorf("search data '%s' on the right table on column '%s': %w", onDataLeft, onColumnName, ErrDataNotFound)
		}
		rightRow := rightTable.Rows[rightRowIndex]

		// append the joined row
		outRow := leftRow.joinOnColumnIndexRight(rightRow, onColumnIndexRight)
		outRows = append(outRows, outRow)
	}

	return DataTable{
		Columns: outColumns,
		Rows:    outRows,
	}, nil
}

// joinColumns combines the left columns and right columns, removing the on column of the right row.
func joinColumns(leftColumns, rightColumns []DataColumn, onColumnIndexRight int) []DataColumn {
	outColumns := make([]DataColumn, 0, len(leftColumns)+len(rightColumns)-1)
	// append left columns
	outColumns = append(outColumns, leftColumns...)
	// append right columns except on column
	outColumns = append(outColumns, rightColumns[:onColumnIndexRight]...)
	outColumns = append(outColumns, rightColumns[onColumnIndexRight+1:]...)

	return outColumns
}
