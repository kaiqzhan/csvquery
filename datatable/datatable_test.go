package datatable_test

import (
	"errors"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableInit(t *testing.T) {
	table := &datatable.DataTable{}
	columns := []datatable.DataColumn{
		{Name: "Column1"},
		{Name: "Column2"},
	}

	err := table.Init(columns)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
}

func TestDataTableInitZeroColumns(t *testing.T) {
	table := &datatable.DataTable{}
	columns := []datatable.DataColumn{}

	err := table.Init(columns)

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrEmptyColumns) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrEmptyColumns, err)
	}
}

func TestDataTableInitOverMaxColumns(t *testing.T) {
	table := &datatable.DataTable{}
	columns := make([]datatable.DataColumn, datatable.MaxColumns+1)

	err := table.Init(columns)

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrOverMaxColumns) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrOverMaxColumns, err)
	}
}

func TestDataTableAppendRow(t *testing.T) {
	fromTable := newTestTable()
	row := datatable.DataRow{
		Items: []datatable.DataItem{
			{Data: "E"},
			{Data: "F"},
			{Data: "3"},
		},
	}

	err := fromTable.AppendRow(row)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
}

func TestDataTableAppendInvalidRowSize(t *testing.T) {
	fromTable := newTestTable()
	row := datatable.DataRow{
		Items: []datatable.DataItem{
			{Data: "E"},
			{Data: "F"},
		},
	}

	err := fromTable.AppendRow(row)

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrInvalidRowSize) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrInvalidRowSize, err)
	}
}

func TestDataTableAppendOverMaxRows(t *testing.T) {
	fromTable := newTestTable()
	row := datatable.DataRow{
		Items: []datatable.DataItem{
			{Data: "E"},
			{Data: "F"},
			{Data: "3"},
		},
	}

	for i := 0; i < datatable.MaxRows-2; i++ {
		err := fromTable.AppendRow(row)
		if err != nil {
			t.Errorf("got error: %s", err)
		}
	}

	// rows is maximized, should return error
	err := fromTable.AppendRow(row)

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrOverMaxRows) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrOverMaxRows, err)
	}
}

func TestDataRowAppendItem(t *testing.T) {
	fromRow := datatable.DataRow{}

	err := fromRow.AppendItem(datatable.DataItem{Data: "A"})

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if len(fromRow.Items) != 1 {
		t.Errorf("expect 1 item in the row, got: %d", len(fromRow.Items))
	}
}

func TestDataRowAppendItemOverMaxColumns(t *testing.T) {
	fromRow := datatable.DataRow{
		Items: make([]datatable.DataItem, datatable.MaxColumns),
	}

	err := fromRow.AppendItem(datatable.DataItem{Data: "A"})

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrOverMaxColumns) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrOverMaxColumns, err)
	}
}

func newTestTable() datatable.DataTable {
	return datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "Column2"},
			{Name: "Column3"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "B"},
					{Data: "1"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "C"},
					{Data: "D"},
					{Data: "2"},
				},
			},
		},
	}
}
