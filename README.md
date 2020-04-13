# Go QueryFX 

## Overview
The Go QueryFX is an SQL utility for escaping, formattting SQL statements 
before its executed at the database end. Its main functionality is to parse 
placeholders in a given SQL statement, validate the data provided for replacing 
the placeholders and then building and returning an escaped SQL query from the 
query placeholders and their corresponding replacement data to return an 
SQL statement that can executed safely against an SQL database.

## Installation

```bash
go get github.com/donama/go-queryfx
```

## License

Apache 2.0
