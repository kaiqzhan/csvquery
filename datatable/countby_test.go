package datatable_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableCountBy(t *testing.T) {
	fromTable := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "Column2"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "10"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "20"},
				},
			},
		},
	}

	expect := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "count"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "2"},
				},
			},
		},
	}

	actual, err := fromTable.CountBy("Column1")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableCountByInvalidColumnName(t *testing.T) {
	fromTable := newTestTable()

	_, err := fromTable.CountBy("InvalidColumnName")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotFound, err)
	}
}
