package datatable

// Take returns a data table with the first n rows from the original table.
// If the number of existing rows is smaller than n, it returns all the rows.
func (dt *DataTable) Take(n int) DataTable {
	if n > len(dt.Rows) {
		n = len(dt.Rows)
	}

	return DataTable{
		Columns: dt.Columns,
		Rows:    dt.Rows[:n],
	}
}
