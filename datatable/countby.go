package datatable

import (
	"strconv"
)

const (
	countColumnName = "count"
)

// CountBy takes a single column name and returns a data table with two columns.
// The first column is the specified column, the second column is "count" which represents
// the number of times that value of the specified column appears.
func (dt *DataTable) CountBy(columnName string) (DataTable, error) {
	columnIndex, err := dt.columnIndexByName(columnName)
	if err != nil {
		return DataTable{}, err
	}

	// count using hash map
	countMap := make(map[string]int)
	for _, row := range dt.Rows {
		item := row.Items[columnIndex]
		if _, ok := countMap[item.Data]; !ok {
			countMap[item.Data] = 0
		}
		countMap[item.Data]++
	}

	// generate rows from hash map
	outRows := make([]DataRow, 0, len(countMap))
	for data, count := range countMap {
		row := DataRow{
			Items: []DataItem{
				{Data: data},
				{Data: strconv.Itoa(count)},
			},
		}
		outRows = append(outRows, row)
	}

	return DataTable{
		Columns: []DataColumn{
			dt.Columns[columnIndex],
			{Name: countColumnName},
		},
		Rows: outRows,
	}, nil
}
