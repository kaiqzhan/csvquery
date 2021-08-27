package datatable_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableSelect(t *testing.T) {
	fromTable := newTestTable()

	expect := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column3"},
			{Name: "Column1"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "1"},
					{Data: "A"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "2"},
					{Data: "C"},
				},
			},
		},
	}

	actual, err := fromTable.Select([]string{"Column3", "Column1"})
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableSelectInvalidColumnName(t *testing.T) {
	fromTable := newTestTable()

	_, err := fromTable.Select([]string{"InvalidColumn"})

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, datatable.ErrColumnNotFound) {
		t.Errorf("expect error: %v, got: %v", datatable.ErrColumnNotFound, err)
	}
}

func TestDataTableSelectEmptyColumnNames(t *testing.T) {
	fromTable := newTestTable()

	_, err := fromTable.Select(nil)

	if err == nil {
		t.Errorf("expect error but got nil")
	}
}
