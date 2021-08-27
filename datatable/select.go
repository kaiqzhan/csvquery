package datatable

// Select returns a data table with only the picked columns and keeps the same row orders.
func (dt *DataTable) Select(columnNames []string) (DataTable, error) {
	if len(columnNames) == 0 {
		return DataTable{}, ErrEmptyColumnNames
	}

	// get columns indices
	columnIndices, err := dt.columnIndicesByNames(columnNames)
	if err != nil {
		return DataTable{}, err
	}

	// select columns
	outColumns := make([]DataColumn, 0, len(columnIndices))
	for _, index := range columnIndices {
		outColumns = append(outColumns, dt.Columns[index])
	}

	// select rows
	outRows := make([]DataRow, 0, len(dt.Rows))
	for _, row := range dt.Rows {
		outRows = append(outRows, row.selectByIndices(columnIndices))
	}

	return DataTable{
		Columns: outColumns,
		Rows:    outRows,
	}, nil
}
