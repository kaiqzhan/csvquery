package datatable_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableJoin(t *testing.T) {
	leftTable := newTestTable()

	rightTable := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column4"},
			{Name: "Column2"},
			{Name: "Column5"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "10"},
					{Data: "B"},
					{Data: "E"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "20"},
					{Data: "D"},
					{Data: "F"},
				},
			},
		},
	}

	expect := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "Column2"},
			{Name: "Column3"},
			{Name: "Column4"},
			{Name: "Column5"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "B"},
					{Data: "1"},
					{Data: "10"},
					{Data: "E"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "C"},
					{Data: "D"},
					{Data: "2"},
					{Data: "20"},
					{Data: "F"},
				},
			},
		},
	}

	actual, err := leftTable.Join(rightTable, "Column2")
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableJoinColumnMissingOnLeftTable(t *testing.T) {
	leftTable := newTestTable()

	rightTable := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column4"},
			{Name: "Column2"},
			{Name: "Column5"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "10"},
					{Data: "B"},
					{Data: "E"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "20"},
					{Data: "D"},
					{Data: "F"},
				},
			},
		},
	}

	_, err := leftTable.Join(rightTable, "Column4")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotFound, err)
	}
}

func TestDataTableJoinColumnMissingOnRightTable(t *testing.T) {
	leftTable := newTestTable()

	rightTable := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column4"},
			{Name: "Column2"},
			{Name: "Column5"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "10"},
					{Data: "B"},
					{Data: "E"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "20"},
					{Data: "D"},
					{Data: "F"},
				},
			},
		},
	}

	_, err := leftTable.Join(rightTable, "Column1")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotFound, err)
	}
}

func TestDataTableJoinNoMatchingOnRightTable(t *testing.T) {
	leftTable := newTestTable()

	rightTable := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column4"},
			{Name: "Column2"},
			{Name: "Column5"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "10"},
					{Data: "X"},
					{Data: "E"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "20"},
					{Data: "D"},
					{Data: "F"},
				},
			},
		},
	}

	_, err := leftTable.Join(rightTable, "Column2")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrDataNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrDataNotFound, err)
	}
}
