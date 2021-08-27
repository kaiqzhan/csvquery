package datatable_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableOrderBy(t *testing.T) {
	fromTable := newTestTable()

	expect := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "Column2"},
			{Name: "Column3"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "C"},
					{Data: "D"},
					{Data: "2"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "B"},
					{Data: "1"},
				},
			},
		},
	}

	actual, err := fromTable.OrderBy("Column3")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableOrderByInvalidColumnName(t *testing.T) {
	fromTable := newTestTable()

	_, err := fromTable.OrderBy("InvalidColumnName")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotFound, err)
	}
}

func TestDataTableOrderByNonNumericColumn(t *testing.T) {
	fromTable := newTestTable()

	_, err := fromTable.OrderBy("Column1")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotNumeric) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotNumeric, err)
	}
}
