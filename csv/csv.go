package csv

import (
	"bufio"
	"os"
	"strings"

	"github.com/kaiqzhan/csvquery/datatable"
)

// Import imports a data table from a CSV file on 'filePath'.
// The first row of the CSV must contains the title of each column.
// The title and data entries are concated by commas.
// The entries should not contains any spaces, commas, or other special characters.
func Import(filePath string) (datatable.DataTable, error) {
	table := datatable.DataTable{}

	// open/close file
	f, err := os.Open(filePath)
	if err != nil {
		return table, err
	}
	defer f.Close()

	// read file line by line and convert to data table
	isFirstLine := true
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()             // read text line
		tokens := strings.Split(line, ",") // tokenize text line

		if isFirstLine { // create columns based on the first row
			err := table.Init(toColumns(tokens))
			if err != nil {
				return table, err
			}
			isFirstLine = false
		} else { // create rows for other rows
			err := table.AppendRow(toRow(tokens))
			if err != nil {
				return table, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return table, err
	}

	return table, nil
}

// toColumns converts column title tokens to data columns.
func toColumns(tokens []string) []datatable.DataColumn {
	var columns []datatable.DataColumn
	for _, token := range tokens {
		columns = append(columns, datatable.DataColumn{Name: token})
	}
	return columns
}

// toRow converts data entry tokens to a data row.
func toRow(tokens []string) datatable.DataRow {
	row := datatable.DataRow{}
	for _, token := range tokens {
		row.AppendItem(datatable.DataItem{Data: token})
	}
	return row
}
