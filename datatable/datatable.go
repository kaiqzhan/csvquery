package datatable

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// DataTable stores the column titles and rows of data.
type DataTable struct {
	Columns []DataColumn
	Rows    []DataRow
}

// Init initializes the data table columns.
func (dt *DataTable) Init(columns []DataColumn) error {
	if len(columns) == 0 {
		return ErrEmptyColumns
	}

	if len(columns) > MaxColumns {
		return ErrOverMaxColumns
	}

	dt.Columns = columns
	return nil
}

// AppendRow appends a row to the data table. The number of items in the row must equals to the number of columns.
func (dt *DataTable) AppendRow(row DataRow) error {
	if len(row.Items) != len(dt.Columns) {
		return ErrInvalidRowSize
	}

	if len(dt.Rows) >= MaxRows {
		return ErrOverMaxRows
	}

	dt.Rows = append(dt.Rows, row)
	return nil
}

// String return the content of the data table in csv format
func (dt *DataTable) String() string {
	var buffer bytes.Buffer

	// print columns
	columnNames := make([]string, 0, len(dt.Columns))
	for _, col := range dt.Columns {
		columnNames = append(columnNames, col.Name)
	}
	buffer.WriteString(strings.Join(columnNames, ","))
	buffer.WriteString("\n")

	// print rows
	for _, row := range dt.Rows {
		datas := make([]string, 0, len(row.Items))
		for _, item := range row.Items {
			datas = append(datas, item.Data)
		}
		buffer.WriteString(strings.Join(datas, ","))
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// columnIndexByName returns the index of the column with the title columnName.
func (dt *DataTable) columnIndexByName(columnName string) (int, error) {
	for i, col := range dt.Columns {
		if col.Name == columnName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("get index of column '%s': %w", columnName, ErrColumnNotFound)
}

// columnIndicesByNames returns the indices of the columns for each column with title in columnNames.
func (dt *DataTable) columnIndicesByNames(columnNames []string) ([]int, error) {
	columnNameIndexMap := make(map[string]int)
	for i, col := range dt.Columns {
		columnNameIndexMap[col.Name] = i
	}

	indices := make([]int, len(columnNames))
	for i, name := range columnNames {
		index, ok := columnNameIndexMap[name]
		if !ok {
			return nil, fmt.Errorf("get index of column '%s': %w", name, ErrColumnNotFound)
		}
		indices[i] = index
	}
	return indices, nil
}

// isColumnNumeric detects if all the data entries in the column in numeric.
func (dt *DataTable) isColumnNumeric(columnIndex int) bool {
	for _, row := range dt.Rows {
		item := row.Items[columnIndex]
		if _, err := strconv.Atoi(item.Data); err != nil {
			return false
		}
	}
	return true
}

// rowIndexMapByColumn return the map of columnName-columnIndex pairs for index lookup purpose.
func (dt *DataTable) rowIndexMapByColumn(columnIndex int) map[string]int {
	rowIndexMap := make(map[string]int)

	for i, row := range dt.Rows {
		data := row.Items[columnIndex].Data
		if _, ok := rowIndexMap[data]; !ok {
			rowIndexMap[data] = i // only add the first occurance
		}
	}

	return rowIndexMap
}

// DataColumn stores the column titles of data table.
type DataColumn struct {
	Name string
}

// DataRow stores the row of data table. It contains data items for each column.
type DataRow struct {
	Items []DataItem
}

// AppendItem appends a data item to the data row.
func (dr *DataRow) AppendItem(item DataItem) error {
	if len(dr.Items) >= MaxColumns {
		return fmt.Errorf("append item: %w", ErrOverMaxColumns)
	}

	dr.Items = append(dr.Items, item)
	return nil
}

// selectByIndices returns a new data row with selected data items by the column indices.
func (dr *DataRow) selectByIndices(columnIndices []int) DataRow {
	outItems := make([]DataItem, 0, len(columnIndices))

	for _, index := range columnIndices {
		outItems = append(outItems, dr.Items[index])
	}

	return DataRow{
		Items: outItems,
	}
}

// joinOnColumnIndexRight combines the left row and right row, removing the on column of the right row.
func (dr *DataRow) joinOnColumnIndexRight(rightRow DataRow, onColumnIndexRight int) DataRow {
	outItems := make([]DataItem, 0, len(dr.Items)+len(rightRow.Items)-1)
	// append left row
	outItems = append(outItems, dr.Items...)
	// append right row except on column
	outItems = append(outItems, rightRow.Items[:onColumnIndexRight]...)
	outItems = append(outItems, rightRow.Items[onColumnIndexRight+1:]...)

	return DataRow{
		Items: outItems,
	}
}

// DataItem stores a data entry. Data is stored as raw string.
type DataItem struct {
	Data string
}
