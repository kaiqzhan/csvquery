package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kaiqzhan/csvquery/cli"
	"github.com/kaiqzhan/csvquery/csv"
)

func TestQuery(t *testing.T) {
	c := cli.CLI{
		DataPath: "../data",
	}

	testCases := []struct {
		cmd        string
		expectFile string
	}{
		{"FROM city.csv", "../data/city.csv"},
		{"FROM country.csv", "../data/country.csv"},
		{"FROM language.csv", "../data/language.csv"},
		{"FROM city.csv SELECT CityName", "testdata/test1.csv"},
		{"FROM country.csv SELECT CountryCode,Continent,CountryPop", "testdata/test2.csv"},
		{"FROM city.csv TAKE 5", "testdata/test3.csv"},
		{"FROM city.csv ORDERBY CityPop TAKE 10", "testdata/test4.csv"},
		{"FROM city.csv JOIN country.csv CountryCode", "testdata/test5.csv"},
		{"FROM city.csv JOIN country.csv CountryCode JOIN language.csv CountryCode", "testdata/test6.csv"},
		{"FROM language.csv COUNTBY Language ORDERBY count TAKE 4", "testdata/test7.csv"},
	}

	for _, testCase := range testCases {
		expect, err := csv.Import(testCase.expectFile)
		if err != nil {
			t.Errorf("Run query: '%s'. Import expect csv file got error: %v", testCase.cmd, err)
		}

		actual, err := c.Query(testCase.cmd)
		if err != nil {
			t.Errorf("Run query: '%s'. Got error: %v", testCase.cmd, err)
		}

		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("Run query '%s' failed. Expected %+v, got %+v", testCase.cmd, expect, actual)
		}
	}
}

func TestQueryNotStartsWithFrom(t *testing.T) {
	c := cli.CLI{
		DataPath: "../data",
	}

	_, err := c.Query("TAKE 10 FROM city.csv")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, cli.ErrInvalidCommand) {
		t.Errorf("expect error: %v, got: %v", cli.ErrInvalidCommand, err)
	}
}

func TestQueryInvalidCommand(t *testing.T) {
	c := cli.CLI{
		DataPath: "../data",
	}

	_, err := c.Query("FROM city.csv LIMIT 10")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, cli.ErrInvalidCommand) {
		t.Errorf("expect error: %v, got: %v", cli.ErrInvalidCommand, err)
	}
}

func TestQueryInsufficientParameters(t *testing.T) {
	c := cli.CLI{
		DataPath: "../data",
	}

	_, err := c.Query("FROM city.csv JOIN country.csv")

	if err == nil {
		t.Errorf("expect error but got nil")
	}

	if !errors.Is(err, cli.ErrInvalidCommand) {
		t.Errorf("expect error: %v, got: %v", cli.ErrInvalidCommand, err)
	}
}
