package csv_test

import (
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/csv"
	"github.com/kaiqzhan/csvquery/datatable"
)

func TestImport(t *testing.T) {
	expect := datatable.DataTable{
		Columns: []datatable.DataColumn{
			{Name: "Column1"},
			{Name: "Column2"},
		},
		Rows: []datatable.DataRow{
			{
				Items: []datatable.DataItem{
					{Data: "A"},
					{Data: "B"},
				},
			},
			{
				Items: []datatable.DataItem{
					{Data: "1"},
					{Data: "2"},
				},
			},
		},
	}

	actual, err := csv.Import("./testdata/data.csv")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected %+v, got %+v", expect, actual)
	}
}
