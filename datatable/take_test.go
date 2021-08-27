package datatable_test

import (
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/datatable"
)

func TestDataTableTake(t *testing.T) {
	fromTable := newTestTable()

	expect := datatable.DataTable{
		Columns: fromTable.Columns,
		Rows:    fromTable.Rows[:1],
	}

	actual := fromTable.Take(1)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableTakeZero(t *testing.T) {
	fromTable := newTestTable()

	expect := datatable.DataTable{
		Columns: fromTable.Columns,
		Rows:    []datatable.DataRow{},
	}

	actual := fromTable.Take(0)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}

func TestDataTableTakeMoreThanAll(t *testing.T) {
	fromTable := newTestTable()

	actual := fromTable.Take(len(fromTable.Rows) + 1)

	if !reflect.DeepEqual(fromTable, actual) {
		t.Errorf("expected %+v, got %+v", fromTable, actual)
	}
}
