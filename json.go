package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"
)

var replaceSpaces *regexp.Regexp = regexp.MustCompile(` `)

func (table *Table) getTableAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{"%s":`, table.tableName))
	buf.WriteByte('[')
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		buf.WriteByte(123) // Opening brace
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteByte('"')
			buf.WriteString(table.colNames[colIndex])
			buf.WriteString(`":`)
			var val interface{}
			val, err = table.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				return "", err
			}
			switch val.(type) {
			case string:
				buf.WriteString(`"` + val.(string) + `"`)
			case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return "", err
				}
				buf.WriteString(valStr)
			case bool:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return "", err
				}
				buf.WriteString(valStr)
			case []byte:
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return "", err
				}
				// Insert comma delimiters between slice elements.
//				valStr = strings.ReplaceAll(valStr, " ", ",")	// New in Go 1.11?
				valStr = replaceSpaces.ReplaceAllString(valStr, ",")
				buf.WriteString(valStr)
			default:
				buf.WriteString(`"TYPE UNKNOWN"`)
			}
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(125)	// Closing brace
		if rowIndex < len(table.rows)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]}")

	jsonString = buf.String()

	return
}

/*
	Marshall gotables TableSet to JSON
*/
func (tableSet *TableSet) GetTableSetAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{"%s":`, tableSet.tableSetName))

	buf.WriteByte('[')
	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}

		var jsonTableString string
		jsonTableString, err = table.getTableAsJSON()
		if err != nil {
			return "", err
		}

		buf.WriteString(jsonTableString)

		if tableIndex < len(tableSet.tables)-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteString(`]}`)

	jsonString = buf.String()

	return
}

func (table *Table) getTableMetadataAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{"%s":`, table.tableName))
	buf.WriteByte('[')
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		buf.WriteByte(123) // Opening brace
		buf.WriteByte('"')
		buf.WriteString(table.colNames[colIndex])
		buf.WriteString(`":"`)
		buf.WriteString(table.colTypes[colIndex])
		buf.WriteString(`"}`)
		if colIndex < len(table.colNames)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]}")

	jsonString = buf.String()

	return
}

/*
	Marshall gotables TableSet metadata to JSON
*/
func (tableSet *TableSet) GetTableSetMetadataAsJSON() (jsonString string, err error) {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(`{"%s":`, tableSet.tableSetName))

	buf.WriteByte('[')
	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {

		var table *Table
		table, err = tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}

		var jsonTableString string
		jsonTableString, err = table.getTableMetadataAsJSON()
		if err != nil {
			return "", err
		}

		buf.WriteString(jsonTableString)

		if tableIndex < len(tableSet.tables)-1 {
			buf.WriteByte(',')
		}
	}

	buf.WriteString(`]}`)

	jsonString = buf.String()

	return
}

func newTableFromJSON(jsonMetadataString string, jsonString string) (table *Table, err error) {

	// Create empty table from metadata.

	dec := json.NewDecoder(strings.NewReader(jsonMetadataString))

	var token json.Token

	// Skip opening brace
	token, err = dec.Token()
	if err == io.EOF {
		return nil, fmt.Errorf("newTableFromJSON() unexpected EOF")
	}
	if err != nil {
		return nil, err
	}

	// Get table name
	token, err = dec.Token()
	if err == io.EOF {
		return nil, fmt.Errorf("newTableFromJSON() unexpected EOF")
	}
	if err != nil {
		return nil, err
	}
	switch token.(type) {
		case string:	// As expected
			table, err = NewTable(token.(string))
			if err != nil {
				return nil, err
			}
			err = table.SetStructShape(true)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("newTableFromJSON() expecting table name but found: %v", reflect.TypeOf(token))
	}

	var colNameNext bool = false
	var colName string
	var colTypeNext bool = false
	var colType string
	var prevDelim rune

	for {
		token, err = dec.Token()
		if err == io.EOF {
			return nil, fmt.Errorf("newTableFromJSON() unexpected EOF")
		}
		if err != nil {
			return nil, err
		}

		switch token.(type) {
			case json.Delim:
				delim := token.(json.Delim)
				switch delim {
					case 123: // Opening brace
						colNameNext = true
						prevDelim = 123	// Opening brace
					case 125:	// Closing brace
						if prevDelim == 125 {	// End of metadata object
							return table, nil
						}
						err = table.AppendCol(colName, colType)
						if err != nil {
							return nil, err
						}
						prevDelim = 125	// Closing brace
					case '[':
					case ']':
				}
			case string:
				if colNameNext {
					colName = token.(string)
					colNameNext = false
					colTypeNext = true
				} else if colTypeNext {
					colType = token.(string)
					colTypeNext = false
				} else {
					return nil, fmt.Errorf("newTableFromJSON() expecting colName or colType")
				}
			case bool:
			case float64:
			case json.Number:
			case nil:
			default:
				fmt.Printf("unknown json token type %T value %v\n", token, token)
		}
	}

	return table, nil
}
