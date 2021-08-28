# CSV Query
A small command-line program for querying CSV files

## Quickstart
### Prerequisite
docker or go 1.17

### Run Program
```
make run
```

### Build and Run in Container
```
docker build . -t csvquery
docker run -it csvquery
```

### Run Unit Tests
```
make test
```

### Dependency
Only depends on the go built-in library

## Code Structure
```
/csvquery
    /cli        # parse command line and run
    /csv        # import table from csv file
    /data       # data source csv files
    /datatable  # core engine to store and operate on table
    main.go     # get input from stdin and run query
```

## How It Works
The program splits and groups the input command line to sub commands. Each sub command contains a keyword and 1 or 2 parameters. For example,
```
FROM language.csv COUNTBY Language ORDERBY count TAKE 4
```
can be grouped to:
```
(FROM language.csv) -> (COUNTBY Language) -> (ORDERBY count) -> (TAKE 4)
```
Each group takes an input table and produces an output table, except `FROM` which does not take input but read directly from file.

The core engine of this program is the `DataTable` struct, which stores all the content of a table in memory. The data structure supports `Select()`, `Take()`, `OrderBy()`, `Join()`, `CountBy()` methods which transforms the table to the desired output.

At the end, it just print out the data table in a formatted way.

## Design Decisions
- Schema vs No schema
    - In SQL we defines the table schema when creating a table so that title and data type of each column is determistic. I chose no schema since it's more flexible considering we're loading table from CSV files. The downside is validating data type at the runtime is required for `ORDERBY` operation.
- Store int vs Store string
    - Since non-schema solution was chosen, we're not able to know the data type of the entries while loading the data. Just store the entries as the raw string and try convert at the query runtime for numeric columns.
- Build index at creation vs Build Index at query
    - To implement the `JOIN` feature, we need to seach a specific value on the right table to find the row id. This usually require a index since nested search is too time consuming. One option is to build the index for each column when/after creating the table, so that we only build index once. It takes some memory to store the index. The other option is to build the index for a specific column when calling `JOIN`, it adds duplicated work to each query but is easier to implement and less memory consuming.
- Hash-based index vs Tree-based index
    - Tree-based index is widely used in real world databases since it's more friendly for persistence, it also preserve all the row id which contain the same value. Hash-based is easier to implement, considering our use case is in memory only, I chose the hash-based index. In my implementation only the first occurance of the value is counted when `JOIN` the tables.
- Load CSV files at beginning vs Load CSV files at query
    - Loading all the CSV files to memory when program started can save many time if we're running a lot of queries. For simplicity, I only implemented loading at each query. This is a place can be improved in the future.
- Maximum table size
    - Maximum table size was calculated based on memory size. Assume each entry is a string less or equal to 5 characters, so it concumes ~10 bytes. `10G * 100 cols * 100000 rows = 1GB`, which can fit into almost any morden computers.
