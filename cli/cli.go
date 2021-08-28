package cli

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/kaiqzhan/csvquery/csv"
	"github.com/kaiqzhan/csvquery/datatable"
)

type CLI struct {
	DataPath string
}

// Query performs query based on command and returns a data table as result.
// dataPath is the path of the source data csv files.
func (c *CLI) Query(commond string) (datatable.DataTable, error) {
	// tokenize command
	tokens := strings.Split(commond, " ")

	// the first and only the first command should be "FROM", and it has 1 parameters
	if len(tokens) < 2 || tokens[0] != "FROM" {
		return datatable.DataTable{}, fmt.Errorf("command does not start with FROM and its parameter: %w", ErrInvalidCommand)
	}

	table, err := c.execFrom(tokens[1])
	if err != nil {
		return datatable.DataTable{}, err
	}

	// run the following sub commands in pipeline
	for i := 2; i < len(tokens); i++ {
		var paramCount int
		var execFunc func(datatable.DataTable, ...string) (datatable.DataTable, error)

		switch tokens[i] {
		case "SELECT":
			paramCount = 1
			execFunc = c.execSelect
		case "TAKE":
			paramCount = 1
			execFunc = c.execTake
		case "ORDERBY":
			paramCount = 1
			execFunc = c.execOrderBy
		case "JOIN":
			paramCount = 2
			execFunc = c.execJoin
		case "COUNTBY":
			paramCount = 1
			execFunc = c.execCountBy
		default:
			return datatable.DataTable{}, fmt.Errorf("invalid sub command '%s': %w", tokens[i], ErrInvalidCommand)
		}

		// get params
		if i+paramCount >= len(tokens) {
			return datatable.DataTable{}, fmt.Errorf("insuffient parameters for sub command '%s': %w", tokens[i], ErrInvalidCommand)
		}
		params := tokens[i+1 : i+1+paramCount]
		i += paramCount

		// execute sub command
		table, err = execFunc(table, params...)
		if err != nil {
			return datatable.DataTable{}, err
		}
	}

	return table, nil
}

// execFrom executes the FROM command
func (c *CLI) execFrom(fileName string) (datatable.DataTable, error) {
	filePath := path.Join(c.DataPath, fileName)

	table, err := csv.Import(filePath)
	if err != nil {
		return datatable.DataTable{}, err
	}

	return table, err
}

// execSelect executes the SELECT command
func (c *CLI) execSelect(fromTable datatable.DataTable, params ...string) (datatable.DataTable, error) {
	columnNames := strings.Split(params[0], ",")
	return fromTable.Select(columnNames)
}

// execTake executes the TAKE command
func (c *CLI) execTake(fromTable datatable.DataTable, params ...string) (datatable.DataTable, error) {
	n, err := strconv.Atoi(params[0])
	if err != nil {
		return datatable.DataTable{}, fmt.Errorf("parameter of TAKE is not a number: %s :%w", err, ErrInvalidParameter)
	}
	return fromTable.Take(n), nil
}

// execOrderBy executes the ORDERBY command
func (c *CLI) execOrderBy(fromTable datatable.DataTable, params ...string) (datatable.DataTable, error) {
	return fromTable.OrderBy(params[0])
}

// execJoin executes the JOIN command
func (c *CLI) execJoin(fromTable datatable.DataTable, params ...string) (datatable.DataTable, error) {
	rightFilePath := path.Join(c.DataPath, params[0])
	rightTable, err := csv.Import(rightFilePath)
	if err != nil {
		return datatable.DataTable{}, err
	}
	return fromTable.Join(rightTable, params[1])
}

// execCountBy executes the COUNTBY command
func (c *CLI) execCountBy(fromTable datatable.DataTable, params ...string) (datatable.DataTable, error) {
	return fromTable.CountBy(params[0])
}
