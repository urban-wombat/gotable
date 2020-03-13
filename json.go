package gotables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
)

type circRefMap map[*Table]struct{}
var empty struct{}

var replaceSpaces *regexp.Regexp = regexp.MustCompile(` `)

// const metadataTableNamePrefix string = "metadata::"
// const dataTableNamePrefix string     = "data::"

// /*
// 	Marshal json from the rows of data in this table.
// 
// 	A *gotables.Table is composed of metadata and data:-
// 		1. Metadata:-
// 			* Table name
// 			* Column names
// 			* Column types
// 		2. Data:
// 			* Rows of data
// 
// 	To generate json metadata and data:-
// 		1. Meta: call method table.GetTableMetadataAsJSON()
// 		2. Data: call method table.GetTableDataAsJSON()
// */
// func (table *Table) GetTableDataAsJSON() (jsonDataString string, err error) {
// where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
// 
// 	if table == nil {
// 		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
// 	}
// 
// 	var refMap circRefMap = map[*Table]struct{}{}
// 	refMap[table] = empty	// Put this top-level table into the map.
// 
// 	var buf bytes.Buffer
// 
// 	buf.WriteByte(123)	// Opening brace outermost
// 
// where()
// 	buf.WriteString(fmt.Sprintf(`"tableName":%q,`, table.Name()))
// 
// 	err = getTableAsJSON_recursive(table, &buf, refMap)
// 	if err != nil {
// 		return "", err
// 	}
// 
// 	buf.WriteByte(125)	// Closing brace outermost
// 
// 	jsonDataString = buf.String()
// 
// 	return
// }

//	func (table *Table) GetTableAsJSONIndent(prefix string, indent string) (jsonStringIndented string, err error) {
//	where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
//	
//		jsonString, err := table.GetTableAsJSON()
//		if err != nil {
//			return "", err
//		}
//	
//		var buf bytes.Buffer
//		err = json.Indent(&buf, []byte(jsonString), "", "\t")
//		if err != nil {
//			return "", err
//		}
//		jsonStringIndented = buf.String()
//	
//		return
//	}

/*
	Marshal json from the metadata in this table.

	A *gotables.Table is composed of metadata and data:-
		1. Metadata:-
			* Table name
			* Column names
			* Column types
		2. Data:
			* Rows of data

	To generate json metadata and data:-
		1. Meta: call method table.GetTableMetadataAsJSON()
		2. Data: call method table.GetTableDataAsJSON()

	Note: The table must have at least 1 col defined (zero rows are okay).
*/
func (table *Table) GetTableMetadataAsJSON() (jsonMetadataString string, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	if table.ColCount() == 0 {
		// return "", fmt.Errorf("%s: in table [%s]: cannot marshal json metadata from a table with zero columns", UtilFuncName(), table.Name())
		return "[]", nil
	}

	var buf bytes.Buffer

//	buf.WriteString(fmt.Sprintf(`"%s%s":[`, metadataTableNamePrefix, table.tableName))
	buf.WriteString(`"metadata":[`)	// Open array of metadata.
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		buf.WriteByte(123) // Opening brace around heading element (name: type)
		buf.WriteByte('"')
		buf.WriteString(table.colNames[colIndex])
		buf.WriteString(`":"`)
		buf.WriteString(table.colTypes[colIndex])
		buf.WriteByte('"')
		buf.WriteByte(125) // Closing brace around heading element (name: type)
		if colIndex < len(table.colNames)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')	// Close array of metadata.

	jsonMetadataString = buf.String()

where(jsonMetadataString)

	return
}

func (table *Table) GetTableMetadataAsJSONIndent(prefix string, indent string) (jsonDataString string, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	jsonString, err := table.GetTableMetadataAsJSON()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonString), "", "\t")
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

//	func newTableFromJSONMetadata(jsonMetadataString string) (table *Table, err error) {
//	where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
//	
//		if jsonMetadataString == "" {
//			return nil, fmt.Errorf("%s: jsonMetadataString is empty", UtilFuncName())
//		}
//	
//		// Create empty table from metadata.
//		/* Note: To preserve column order, we do NOT use JSON marshalling into a map,
//		   because iterating over a map returns values in random order.
//		   Instead, we use the json decoder. (The data rows (later in this function)
//		   ARE decoded using a map.)
//		   Actually, the jsonMetadataString is an array, so it probably WOULD work.
//		   TODO: Use a map to decode jsonMetadataString
//		*/
//	
//		dec := json.NewDecoder(strings.NewReader(jsonMetadataString))
//		var token json.Token
//	
//		// Skip opening brace
//		token, err = dec.Token()
//		if err == io.EOF {
//			return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//		}
//		if err != nil {
//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//		}
//	
//		// Get table name
//		token, err = dec.Token()
//		if err == io.EOF {
//			return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//		}
//		if err != nil {
//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//		}
//	
//		// Get the table name.
//		var metadataTableName string
//		var tableName string
//		switch token.(type) {
//		case string: // As expected
//			metadataTableName = token.(string)
//	
//	//		// Strip off metadataTableNamePrefix, leaving the table name.
//	//		tableName = metadataTableName[len(metadataTableNamePrefix):]
//	
//			table, err = NewTable(tableName)
//			if err != nil {
//				return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//			}
//		default:
//			return nil, fmt.Errorf("%s ERROR %s: expecting table name but found: %v", UtilFuncSource(), UtilFuncName(), reflect.TypeOf(token))
//		}
//	
//		// Simple parsing flags and values.
//		var colNameNext bool = false
//		var colName string
//		var colTypeNext bool = false
//		var colType string
//		var prevDelim rune
//	
//	Loop:
//		for {
//			token, err = dec.Token()
//			if err == io.EOF {
//				return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//			}
//			if err != nil {
//				return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//			}
//	
//			switch token.(type) {
//			case json.Delim:
//				delim := token.(json.Delim)
//				switch delim {
//				case 123: // Opening brace
//					colNameNext = true
//					prevDelim = 123 // Opening brace
//				case 125: // Closing brace
//					if prevDelim == 125 { // Closing brace: end of JSON metadata object
//						// Table metadata is now completely initialised. Now do the rows of data.
//						//							return table, nil
//						break Loop
//					}
//					// We now have a colName-plus-colType pair. Add this col to table.
//					err = table.AppendCol(colName, colType)
//					if err != nil {
//						return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//					}
//					prevDelim = 125 // Closing brace: end of col
//				case '[': // Ignore slice signifiers in type names
//				case ']': // Ignore slice signifiers in type names
//				}
//			case string:
//				if colNameNext {
//					colName = token.(string)
//					colNameNext = false
//					colTypeNext = true
//				} else if colTypeNext {
//					colType = token.(string)
//					colTypeNext = false
//				} else {
//					return nil, fmt.Errorf("newTableFromJSON(): expecting colName or colType")
//				}
//			case bool:
//				return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//			case float64:
//				return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//			case json.Number:
//				return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//			case nil:
//				return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//			default:
//				fmt.Printf("unknown json token type %T value %v\n", token, token)
//			}
//		}
//	
//		return table, nil
//	}

func newTableFromJSONData(metadataTable *Table, jsonDataString string) (table *Table, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
where(fmt.Sprintf("jsonDataString = %s", jsonDataString))
	// Strictly speaking, this doesn't create a new table, but the naming is more consistent with
	// newTableFromJSONMetadata() which it goes with.

	// Append rows of table data from JSON.

	/*
	   Note: Here we use a map for rows of data now that we have already preserved col order.
	   Unmarshal does all the parsing for us.
	*/

	// newTableFromJSONMetadata() has already created the table and populated it with
	// metadata: col names, col types. Here we will populate it with data rows.
	table = metadataTable	// Use as input the output table from newTableFromJSONMetadata()
	metadataTableName := table.Name()

	var unmarshalled interface{}
where("***CALLING** json.Unmarshal() ...")
	err = json.Unmarshal([]byte(jsonDataString), &unmarshalled)
	if err != nil {
		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
	}

	var tableMap map[string]interface{} = unmarshalled.(map[string]interface{})
where(fmt.Sprintf("tableMap (UNMARSHALLED) = %v", tableMap))
where()

	// Check that this JSON data (rows) object table name matches the JSON metadata object table name.
	// (Could have simply used metadataTableName as the key to a lookup.)
	var dataTableName string
i := 1
	for dataTableName, _ = range tableMap {
		// There should be only one key, and it should be the table name.
fmt.Printf("tableName[%d] = %s\n", i, dataTableName)
i++
	}

//	// Strip off metadataTableNamePrefix, leaving the table name.
//	dataTableName = dataTableName[len(dataTableNamePrefix):]
where(dataTableName)

	if dataTableName != metadataTableName {
		return nil, fmt.Errorf("newTableFromJSON(): unexpected JSON metadataTableName %q != JSON dataTableName %q",
			metadataTableName, dataTableName)
	}

//where(tableMap)
where(fmt.Sprintf("dataTableName = %s", dataTableName))
where(fmt.Sprintf("tableMap = %v", tableMap))
// DOING: map containing 2 maps?
//	var tableInterface []
fmt.Printf("tableMap type = %T\n", tableMap)
//	var something map[string]interface{} = tableMap[dataTableName].([]interface{})
//fmt.Printf("something type = %T\n", something)
//fmt.Printf("something = %v\n", something)
os.Exit(4)
/*
//	var rowsInterface []interface{} = tableMap[dataTableName].([]interface{})
	var rowsInterface []interface{} = something[dataTableName].([]interface{})
//where(rowsInterface)

	// Loop through the JSON data rows.
	for rowIndex, row := range rowsInterface {
		table.AppendRow()
		var rowMap map[string]interface{} = row.(map[string]interface{})
		for colName, val := range rowMap {
			var colIndex = table.colNamesMap[colName]
			var colType string = table.colTypes[colIndex]
//where(fmt.Sprintf("coltype: %q", colType))
//where(fmt.Sprintf("val type: %T", val))
//where(fmt.Sprintf("val value: %v", val))
//where()
			switch val.(type) {
			case string:
				err = table.SetString(colName, rowIndex, val.(string))
			case float64:	// All JSON number values are stored as float64
				switch colType {	// We need to convert them back to gotables numeric types
				case "int":
					err = table.SetInt(colName, rowIndex, int(val.(float64)))
				case "uint":
					err = table.SetUint(colName, rowIndex, uint(val.(float64)))
				case "byte":
					err = table.SetByte(colName, rowIndex, byte(val.(float64)))
				case "int8":
					err = table.SetInt8(colName, rowIndex, int8(val.(float64)))
				case "int16":
					err = table.SetInt16(colName, rowIndex, int16(val.(float64)))
				case "int32":
					err = table.SetInt32(colName, rowIndex, int32(val.(float64)))
				case "int64":
					err = table.SetInt64(colName, rowIndex, int64(val.(float64)))
				case "uint8":
					err = table.SetUint8(colName, rowIndex, uint8(val.(float64)))
				case "uint16":
					err = table.SetUint16(colName, rowIndex, uint16(val.(float64)))
				case "uint32":
					err = table.SetUint32(colName, rowIndex, uint32(val.(float64)))
				case "uint64":
					err = table.SetUint64(colName, rowIndex, uint64(val.(float64)))
				case "float32":
					err = table.SetFloat32(colName, rowIndex, float32(val.(float64)))
				case "float64":
					err = table.SetFloat64(colName, rowIndex, float64(val.(float64)))
				}
				if err != nil {
					err := fmt.Errorf("could not convert JSON float64 to gotables %s", colType)
					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
				}
			case bool:
				err = table.SetBool(colName, rowIndex, val.(bool))
			case []interface{}: // This cell is a slice
				var interfaceSlice []interface{} = val.([]interface{})
				var byteSlice []byte = []byte{}
				for _, sliceVal := range interfaceSlice {
					byteSlice = append(byteSlice, byte(sliceVal.(float64)))
				}
				err = table.SetByteSlice(colName, rowIndex, byteSlice)
			case map[string]interface{}:	// This cell is a table
// TODO We need to somehow parse this into a table!
				err = table.SetTable(colName, rowIndex, val.(*Table))
			case nil:
				// TODO: This may break nested tables.
				return nil, fmt.Errorf("newTableFromJSON(): unexpected nil value")
			default:
				return nil, fmt.Errorf("%s ERROR %s: unexpected value of type: %v", UtilFuncSource(), UtilFuncName(), reflect.TypeOf(val))
			}

			// Single error handler for all the table.Set...() calls.
			if err != nil {
				return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
			}
		}
	}
*/

	return table, nil
}

//	//	/*
//	//		Unmarshal a document of JSON metadata and a document of JSON data to a *gotables.Table
//	//	
//	//		Two JSON documents are required:-
//	//			1. JSON metadata which contains the tableName, colNames and colTypes.
//	//			2. JSON data which contains zero or more rows of data that map to the metadata.
//	//	
//	//		The two documents must match: the metadata must match the corresponding data.
//	//	*/
//	//	func NewTableFromJSON(jsonMetadataString string, jsonDataString string) (table *Table, err error) {
//	//	
//	//	//	if jsonMetadataString == "" {
//	//	//		return nil, fmt.Errorf("newTableFromJSON(): jsonMetadataString is empty")
//	//	//	}
//	//	//
//	//	//	if jsonDataString == "" {
//	//	//		return nil, fmt.Errorf("newTableFromJSON(): jsonDataString is empty")
//	//	//	}
//	//	//
//	//	//	// Create empty table from metadata.
//	//	//	/* Note: To preserve column order, we do NOT use JSON marshalling into a map,
//	//	//	   because iterating over a map returns values in random order.
//	//	//	   Instead, we use the json decoder. (The data rows (later in this function)
//	//	//	   ARE decoded using a map.)
//	//	//	   Actually, the jsonMetadataString is an array, so it probably WOULD work.
//	//	//	   TODO: Use a map to decode jsonMetadataString
//	//	//	*/
//	//	//
//	//	//	dec := json.NewDecoder(strings.NewReader(jsonMetadataString))
//	//	//	var token json.Token
//	//	//
//	//	//	// Skip opening brace
//	//	//	token, err = dec.Token()
//	//	//	if err == io.EOF {
//	//	//		return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//	//	//	}
//	//	//	if err != nil {
//	//	//		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//	}
//	//	//
//	//	//	// Get table name
//	//	//	token, err = dec.Token()
//	//	//	if err == io.EOF {
//	//	//		return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//	//	//	}
//	//	//	if err != nil {
//	//	//		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//	}
//	//	//
//	//	//	// Get the table name.
//	//	//	var metadataTableName string
//	//	//	switch token.(type) {
//	//	//	case string: // As expected
//	//	//		metadataTableName = token.(string)
//	//	//		table, err = NewTable(metadataTableName)
//	//	//		if err != nil {
//	//	//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//		}
//	//	//	default:
//	//	//		return nil, fmt.Errorf("%s ERROR %s: expecting table name but found: %v", UtilFuncSource(), UtilFuncName(), reflect.TypeOf(token))
//	//	//	}
//	//	//
//	//	//	// Simple parsing flags and values.
//	//	//	var colNameNext bool = false
//	//	//	var colName string
//	//	//	var colTypeNext bool = false
//	//	//	var colType string
//	//	//	var prevDelim rune
//	//	//
//	//	//Loop:
//	//	//	for {
//	//	//		token, err = dec.Token()
//	//	//		if err == io.EOF {
//	//	//			return nil, fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
//	//	//		}
//	//	//		if err != nil {
//	//	//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//		}
//	//	//
//	//	//		switch token.(type) {
//	//	//		case json.Delim:
//	//	//			delim := token.(json.Delim)
//	//	//			switch delim {
//	//	//			case 123: // Opening brace
//	//	//				colNameNext = true
//	//	//				prevDelim = 123 // Opening brace
//	//	//			case 125: // Closing brace
//	//	//				if prevDelim == 125 { // Closing brace: end of JSON metadata object
//	//	//					// Table metadata is now completely initialised. Now do the rows of data.
//	//	//					//							return table, nil
//	//	//					break Loop
//	//	//				}
//	//	//				// We now have a colName-plus-colType pair. Add this col to table.
//	//	//				err = table.AppendCol(colName, colType)
//	//	//				if err != nil {
//	//	//					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//				}
//	//	//				prevDelim = 125 // Closing brace: end of col
//	//	//			case '[': // Ignore slice signifiers in type names
//	//	//			case ']': // Ignore slice signifiers in type names
//	//	//			}
//	//	//		case string:
//	//	//			if colNameNext {
//	//	//				colName = token.(string)
//	//	//				colNameNext = false
//	//	//				colTypeNext = true
//	//	//			} else if colTypeNext {
//	//	//				colType = token.(string)
//	//	//				colTypeNext = false
//	//	//			} else {
//	//	//				return nil, fmt.Errorf("newTableFromJSON(): expecting colName or colType")
//	//	//			}
//	//	//		case bool:
//	//	//			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//	//	//		case float64:
//	//	//			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//	//	//		case json.Number:
//	//	//			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//	//	//		case nil:
//	//	//			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
//	//	//		default:
//	//	//			fmt.Printf("unknown json token type %T value %v\n", token, token)
//	//	//		}
//	//	//	}
//	//	
//	//		var metadataTable *Table
//	//		metadataTable, err = newTableFromJSONMetadata(jsonMetadataString)
//	//		if err != nil {
//	//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//		}
//	//	
//	//	//	metadataTableName := table.Name()
//	//	//
//	//	//	// Append rows of table data from JSON.
//	//	//	/*
//	//	//	   Note: Here we use a map for rows of data now that we have already preserved col order.
//	//	//	   Unmarshal does all the parsing for us.
//	//	//	*/
//	//	//
//	//	//	var unmarshalled interface{}
//	//	//	err = json.Unmarshal([]byte(jsonDataString), &unmarshalled)
//	//	//	if err != nil {
//	//	//		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//	}
//	//	//
//	//	//	var tableMap map[string]interface{} = unmarshalled.(map[string]interface{})
//	//	//
//	//	//	// Check that this JSON data (rows) object table name matches the JSON metadata object table name.
//	//	//	// (Could have simply used metadataTableName as the key to a lookup.)
//	//	//	var dataTableName string
//	//	//	for dataTableName, _ = range tableMap {
//	//	//		// There should be only one key, and it should be the table name.
//	//	//	}
//	//	//	if dataTableName != metadataTableName {
//	//	//		return nil, fmt.Errorf("newTableFromJSON(): unexpected JSON metadataTableName %q != JSON dataTableName %q",
//	//	//			metadataTableName, dataTableName)
//	//	//	}
//	//	//
//	//	//	var rowsInterface []interface{} = tableMap[dataTableName].([]interface{})
//	//	//where(rowsInterface)
//	//	//
//	//	//	// Loop through the JSON data rows.
//	//	//	for rowIndex, row := range rowsInterface {
//	//	//		table.AppendRow()
//	//	//		var rowMap map[string]interface{} = row.(map[string]interface{})
//	//	//		for colName, val := range rowMap {
//	//	//			var colIndex = table.colNamesMap[colName]
//	//	//			var colType string = table.colTypes[colIndex]
//	//	//where(fmt.Sprintf("coltype: %q", colType))
//	//	//where(fmt.Sprintf("val type: %T", val))
//	//	//where(fmt.Sprintf("val value: %v", val))
//	//	//where()
//	//	//			switch val.(type) {
//	//	//			case string:
//	//	//				err = table.SetString(colName, rowIndex, val.(string))
//	//	//			case float64:	// All JSON number values are stored as float64
//	//	//				switch colType {	// We need to convert them back to gotables numeric types
//	//	//				case "int":
//	//	//					err = table.SetInt(colName, rowIndex, int(val.(float64)))
//	//	//				case "uint":
//	//	//					err = table.SetUint(colName, rowIndex, uint(val.(float64)))
//	//	//				case "byte":
//	//	//					err = table.SetByte(colName, rowIndex, byte(val.(float64)))
//	//	//				case "int8":
//	//	//					err = table.SetInt8(colName, rowIndex, int8(val.(float64)))
//	//	//				case "int16":
//	//	//					err = table.SetInt16(colName, rowIndex, int16(val.(float64)))
//	//	//				case "int32":
//	//	//					err = table.SetInt32(colName, rowIndex, int32(val.(float64)))
//	//	//				case "int64":
//	//	//					err = table.SetInt64(colName, rowIndex, int64(val.(float64)))
//	//	//				case "uint8":
//	//	//					err = table.SetUint8(colName, rowIndex, uint8(val.(float64)))
//	//	//				case "uint16":
//	//	//					err = table.SetUint16(colName, rowIndex, uint16(val.(float64)))
//	//	//				case "uint32":
//	//	//					err = table.SetUint32(colName, rowIndex, uint32(val.(float64)))
//	//	//				case "uint64":
//	//	//					err = table.SetUint64(colName, rowIndex, uint64(val.(float64)))
//	//	//				case "float32":
//	//	//					err = table.SetFloat32(colName, rowIndex, float32(val.(float64)))
//	//	//				case "float64":
//	//	//					err = table.SetFloat64(colName, rowIndex, float64(val.(float64)))
//	//	//				}
//	//	//				if err != nil {
//	//	//					err := fmt.Errorf("could not convert JSON float64 to gotables %s", colType)
//	//	//					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//				}
//	//	//			case bool:
//	//	//				err = table.SetBool(colName, rowIndex, val.(bool))
//	//	//			case []interface{}: // This cell is a slice
//	//	//				var interfaceSlice []interface{} = val.([]interface{})
//	//	//				var byteSlice []byte = []byte{}
//	//	//				for _, sliceVal := range interfaceSlice {
//	//	//					byteSlice = append(byteSlice, byte(sliceVal.(float64)))
//	//	//				}
//	//	//				err = table.SetByteSlice(colName, rowIndex, byteSlice)
//	//	//			case map[string]interface{}:	// This cell is a table
//	//	//// TODO We need to somehow parse this into a table!
//	//	//				err = table.SetTable(colName, rowIndex, val.(*Table))
//	//	//			case nil:
//	//	//				// TODO: This may break nested tables.
//	//	//				return nil, fmt.Errorf("newTableFromJSON(): unexpected nil value")
//	//	//			default:
//	//	//				return nil, fmt.Errorf("%s ERROR %s: unexpected value of type: %v", UtilFuncSource(), UtilFuncName(), reflect.TypeOf(val))
//	//	//			}
//	//	//
//	//	//			// Single error handler for all the table.Set...() calls.
//	//	//			if err != nil {
//	//	//				return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//	//			}
//	//	//		}
//	//	//	}
//	//	
//	//		table, err = newTableFromJSONData(metadataTable, jsonDataString)
//	//		if err != nil {
//	//			return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	//		}
//	//	
//	//		return table, nil
//	//	}

///*
//	Unmarshal a document of JSON metadata and a document of JSON data to a *gotables.Table
//
//	Two JSON documents are required:-
//		1. JSON metadata which contains the tableName, colNames and colTypes.
//		2. JSON data which contains zero or more rows of data that map to the metadata.
//
//	The two documents must match: the metadata must match the corresponding data.
//*/
//func NewTableFromJSON(jsonMetadataString string, jsonDataString string) (table *Table, err error) {
//where("inside NewTableFromJSON")
//
//	var metadataTable *Table
//where("***CALLING** newTableFromJSONMetadata()")
//	metadataTable, err = newTableFromJSONMetadata(jsonMetadataString)
//	if err != nil {
//		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	}
//where(metadataTable)
//
//where("***CALLING** newTableFromJSONData(metadataTable, jsonDataString) ...")
//	table, err = newTableFromJSONData(metadataTable, jsonDataString)
//	if err != nil {
//		return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
//	}
//
//	return table, nil
//}

///*
//	Unmarshal JSON documents to a *gotables.TableSet
//*/
//func NewTableSetFromJSON(jsonStrings []string) (tableSet *TableSet, err error) {
//
//	if jsonStrings == nil {
//		return nil, fmt.Errorf("jsonStrings == nil")
//	}
//
//	tableSet, err = NewTableSet("")
//	if err != nil {
//		return nil, err
//	}
//
//	for tableIndex := 0; tableIndex < len(jsonStrings); tableIndex++ {
//		table, err := NewTableFromJSON(jsonStrings[tableIndex])
//		if err != nil {
//			return nil, err
//		}
//
//		err = tableSet.AppendTable(table)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return
//}

func (table *Table) GetTableAsJSON() (json string, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var refMap circRefMap = map[*Table]struct{}{}
	var buf bytes.Buffer

	buf.WriteByte(123)	// Opening brace outermost

where("***CALLING** getTableAsJSON_recursive()")
	err = getTableAsJSON_recursive(table, &buf, refMap)
	if err != nil {
		return "", err
	}

	buf.WriteByte(125)	// Closing brace outermost

	json = buf.String()

	return
}

func getTableAsJSON_recursive(table *Table, buf *bytes.Buffer, refMap circRefMap) (err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if table == nil {
		return fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	// Add this table to the circular reference map.
	refMap[table] = empty

	buf.WriteString(fmt.Sprintf(`"tableName":%q,`, table.Name()))

//	buf.WriteString(fmt.Sprintf(`"%s%s":[`, metadataTableNamePrefix, table.tableName))
	buf.WriteString(`"metadata":[`)
	for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
		buf.WriteByte(123) // Opening brace around heading element (name: type)
		buf.WriteByte('"')
		buf.WriteString(table.colNames[colIndex])
		buf.WriteString(`":"`)
		buf.WriteString(table.colTypes[colIndex])
		buf.WriteByte('"')
		buf.WriteByte(125) // Closing brace around heading element (name: type)
		if colIndex < len(table.colNames)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')
	buf.WriteByte(',')	// Between metadata and data.

	// Get data

//	buf.WriteString(fmt.Sprintf(`"%s%s":[`, dataTableNamePrefix, table.Name()))	// Begin array of rows.
	buf.WriteString(`"data":[`)
	for rowIndex := 0; rowIndex < len(table.rows); rowIndex++ {
		buf.WriteByte('[')	// Begin array of column cells.
		for colIndex := 0; colIndex < len(table.colNames); colIndex++ {
			buf.WriteByte(123) // Opening brace
			buf.WriteString(fmt.Sprintf("%q:", table.colNames[colIndex]))
			var val interface{}
			val, err = table.GetValByColIndex(colIndex, rowIndex)
			if err != nil {
				return err
			}

			switch val.(type) {

			case string:
				buf.WriteString(fmt.Sprintf("%q", val.(string)))

			case bool, int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
				var valStr string
				valStr, err = table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				buf.WriteString(valStr)

			case []byte:
				var valStr string
				valStr, err := table.GetValAsStringByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
				// Insert comma delimiters between slice elements.
				//				valStr = strings.ReplaceAll(valStr, " ", ",")	// New in Go 1.11?
				valStr = replaceSpaces.ReplaceAllString(valStr, ",")
				buf.WriteString(valStr)

			case *Table:

				var nestedTable *Table
				nestedTable, err = table.GetTableByColIndex(colIndex, rowIndex)
				if err != nil {
					return err
				}
//fmt.Printf("#3 getTableAsJSON_recursive(nestedTable) = %p\n", nestedTable)

				_, exists := refMap[nestedTable]
				if exists {
					err = fmt.Errorf("%s: circular reference: a reference to table [%s] already exists",
						UtilFuncName(), nestedTable.Name())
					return
				}

				isNilTable, err := nestedTable.IsNilTable()
				if err != nil {
					return err
				}
				if isNilTable {
					buf.WriteString("null")
				} else {
					buf.WriteByte(123)	// Begin nested table.
					err = getTableAsJSON_recursive(nestedTable, buf, refMap)
					if err != nil {
						return err
					}
					buf.WriteByte(125)	// End nested table.
				}

			default:
				buf.WriteString(`"TYPE UNKNOWN"`)
			}

			buf.WriteByte(125) // Closing brace
			if colIndex < len(table.colNames)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')	// End array of column cells.
		if rowIndex < len(table.rows)-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteByte(']')	// End array of rows.

	return
}

func (table *Table) GetTableAsJSONIndent(prefix string, indent string) (jsonStringIndented string, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	jsonString, err := table.GetTableAsJSON()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, []byte(jsonString), "", "\t")
	if err != nil {
		return "", err
	}
	jsonStringIndented = buf.String()

	return
}

func checkJsonDecodeError(checkErr error) (err error) {
	if checkErr == io.EOF {
		return fmt.Errorf("%s ERROR %s: unexpected EOF", UtilFuncSource(), UtilFuncName())
	}

	if checkErr != nil {
		return fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
	}

	return nil
}

// func newTableFromJSON_recursive(dec *json.Decoder, tableIn *Table) (table *Table, err error) {
// where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
// 
// 	/*
// 		Create empty table from metadata.
// 		Note: To preserve column order, we do NOT use JSON marshalling into a map,
// 		because iterating over a map returns values in random order.
// 		Instead, we use the json decoder.
// 	*/
// 
// 	var token json.Token
// 
// 	// Skip overall object opening brace
// where("Skip overall object opening brace")
// 	token, err = dec.Token()
// 	if err = checkJsonDecodeError(err); err != nil {
// 		return nil, err
// 	}
// where(fmt.Sprintf("TOKEN: %s", token))
// 
// 	// Get table name
// 	token, err = dec.Token()
// 	if err = checkJsonDecodeError(err); err != nil {
// 		return nil, err
// 	}
// 	var tableName string = token.(string)
// where(fmt.Sprintf("TOKEN: %s", token))
// 	tableIn, err = NewTable(tableName)
// 	if err != nil {
// 		return
// 	}
// where(tableIn)
// 
// 	// Skip metadata opening brace
// where("Skip metadata opening brace")
// 	token, err = dec.Token()
// 	if err = checkJsonDecodeError(err); err != nil {
// 		return nil, err
// 	}
// where(fmt.Sprintf("TOKEN: %s", token))
// 
// 	// Skip metadata name
// where("Skip metadata name")
// 	token, err = dec.Token()
// 	if err = checkJsonDecodeError(err); err != nil {
// 		return nil, err
// 	}
// where(fmt.Sprintf("TOKEN: %s", token))
// 
// 	// Skip opening colnames/types array square bracket
// where("Skip opening colnames/types array square bracket")
// 	token, err = dec.Token()
// 	if err = checkJsonDecodeError(err); err != nil {
// 		return nil, err
// 	}
// where(fmt.Sprintf("TOKEN: %s", token))
// 
// 	// Simple parsing flags and values.
// 	var colNameNext bool = false
// 	var colName string
// 	var colTypeNext bool = false
// 	var colType string
// 	var prevDelim rune
// 
// repeatLimit := 4
// i := 1
// where(fmt.Sprintf("for i=%d to %d ...", i, repeatLimit))
// Loop:
// 	for {
// 		token, err = dec.Token()
// where(fmt.Sprintf("i=%d TOKEN: %s", i, token))
// 		if err = checkJsonDecodeError(err); err != nil {
// 			return nil, err
// 		}
// 
// 		switch token.(type) {
// 		case json.Delim:
// where("case json.Delim:")
// 			delim := token.(json.Delim)
// 			switch delim {
// 			case 123: // Opening brace
// where("	case 123: // Opening brace")
// 				colNameNext = true
// 				prevDelim = 123	// Opening brace
// 			case 125: // Closing brace
// where("	case 125: // Closing brace")
// 				if prevDelim == 125	{	// 2 closing braces in a row.
// 					// Closing brace: end of JSON metadata object
// where("	prevDelim == 125 // Closing brace: end of JSON metadata object")
// 					// Table metadata is now completely initialised. Now do the rows of data.
// 					//							return tableIn, nil
// where("		break Loop")
// 					break Loop
// 				}
// 				// We now have a colName-plus-colType pair. Add this col to tableIn.
// where(fmt.Sprintf("	tableIn.AppendCol(colName=%q, colType=%q)", colName, colType))
// 				err = tableIn.AppendCol(colName, colType)
// 				if err != nil {
// 					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
// 				}
// 				prevDelim = 125 // Closing brace: end of col
// 			case '[': // Ignore slice signifiers in type names
// where(fmt.Sprintf("	case '[' // Ignore slice signifiers in type names"))
// 			case ']': // Ignore slice signifiers in type names
// where(fmt.Sprintf("	case ']' // Ignore slice signifiers in type names"))
// 			}
// 		case string:
// where("case string:")
// 			if colNameNext {
// where("colNameNext")
// 				colName = token.(string)
// 				colNameNext = false
// 				colTypeNext = true
// 			} else if colTypeNext {
// where("colTypeNext")
// 				colType = token.(string)
// 				colTypeNext = false
// 			} else {
// 				return nil, fmt.Errorf("newTableFromJSON(): expecting colName or colType")
// 			}
// 		case bool:
// where("case bool:")
// 			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
// 		case float64:
// where("case float64:")
// 			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
// 		case json.Number:
// where("case json.Number:")
// 			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
// 		case nil:
// where("case nil:")
// 			return nil, fmt.Errorf("newTableFromJSON(): unexpected value of type: %v", reflect.TypeOf(token))
// 		default:
// where("default:")
// 			return nil, fmt.Errorf("unknown json token type %T value %v\n", token, token)
// 		}
// where(tableIn)
// if i > repeatLimit {
// where(fmt.Sprintf("i=%d > repeatLimit=%d", i, repeatLimit))
// os.Exit(i)
// i++
// }
// 	}
// 
// where(fmt.Sprintf("return nil i=%d", i))
// 	return nil, nil
// }

func newTableFromJSON_recursive(m map[string]interface{}) (table *Table, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	var exists bool

	/*
	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return nil, err
	}
	*/

where()
	/*
		We don't know the order map values will be returned if we iterate of the map:
		(1) tableName
		(2) metadata
		(3) data (if any)
		So we retrieve each of the 3 (possibly 2) top-level map values individually.
	*/

where()
	// (1) Retrieve and process table name.
	var tableName string
	tableName, exists = m["tableName"].(string)
	if !exists {
		return nil, fmt.Errorf("JSON is missing table name")
	}
	table, err = NewTable(tableName)
	if err != nil {
		return nil, err
	}
/*
err = table.SetStructShape(true)
if err != nil {
	return nil, err
}
*/

where()
	// (2) Retrieve and process metadata.
	var metadata []interface{}
//	metadata, exists = m[fmt.Sprintf("metadata::%s", tableName)].([]interface{})
	metadata, exists = m["metadata"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("JSON is missing table metadata")
	}
	// Loop through the array of metadata.
	for _, colNameAndType := range metadata {
		var colName string
		var colType string
		var val interface{}
		for colName, val = range colNameAndType.(map[string]interface{}) {
			// There's only one map element here: colName and colType.
		}
		colType, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("expecting col type value from JSON string value but got type %T: %v", val, val)
		}

where()
		err = table.AppendCol(colName, colType)
		if err != nil {
			return nil, err
		}
	}
where(fmt.Sprintf("\n%v\n", table))
where()
	// (3) Retrieve and process data (if any).
	var data []interface{}
//	data, exists = m[fmt.Sprintf("data::%s", tableName)].([]interface{})
	data, exists = m["data"].([]interface{})
	if !exists {
		// Zero rows in this table. That's okay.
		return table, nil
	}
where(data)
	// Loop through the array of rows.
	for rowIndex, val := range data {
		where(fmt.Sprintf("row [%d] %v", rowIndex, val))
		err = table.AppendRow()
		if err != nil {
			return nil, err
		}

where()
		var row []interface{} = val.([]interface{})
		for colIndex, val := range row {
			where(fmt.Sprintf("\t\tcol [%d] %v", colIndex, val))
			var cell interface{}
			for _, cell = range val.(map[string]interface{}) {
				// There's only one map element here: colName and colType.
				where(fmt.Sprintf("\t\t\tcol=%d row=%d celltype=%T cell=%v", colIndex, rowIndex, cell, cell))

where()
				var colType string = table.colTypes[colIndex]
				switch cell.(type) {
				case string:
					err = table.SetStringByColIndex(colIndex, rowIndex, cell.(string))
				case float64:	// All JSON number values are stored as float64
					switch colType {	// We need to convert them back to gotables numeric types
					case "int":
						err = table.SetIntByColIndex(colIndex, rowIndex, int(cell.(float64)))
					case "uint":
						err = table.SetUintByColIndex(colIndex, rowIndex, uint(cell.(float64)))
					case "byte":
						err = table.SetByteByColIndex(colIndex, rowIndex, byte(cell.(float64)))
					case "int8":
						err = table.SetInt8ByColIndex(colIndex, rowIndex, int8(cell.(float64)))
					case "int16":
						err = table.SetInt16ByColIndex(colIndex, rowIndex, int16(cell.(float64)))
					case "int32":
						err = table.SetInt32ByColIndex(colIndex, rowIndex, int32(cell.(float64)))
					case "int64":
						err = table.SetInt64ByColIndex(colIndex, rowIndex, int64(cell.(float64)))
					case "uint8":
						err = table.SetUint8ByColIndex(colIndex, rowIndex, uint8(cell.(float64)))
					case "uint16":
						err = table.SetUint16ByColIndex(colIndex, rowIndex, uint16(cell.(float64)))
					case "uint32":
						err = table.SetUint32ByColIndex(colIndex, rowIndex, uint32(cell.(float64)))
					case "uint64":
						err = table.SetUint64ByColIndex(colIndex, rowIndex, uint64(cell.(float64)))
					case "float32":
						err = table.SetFloat32ByColIndex(colIndex, rowIndex, float32(cell.(float64)))
					case "float64":
						err = table.SetFloat64ByColIndex(colIndex, rowIndex, float64(cell.(float64)))
					}
					// Single error handler for all the calls in this switch statement.
					if err != nil {
						err := fmt.Errorf("could not convert JSON float64 to gotables %s", colType)
						return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
					}
				case bool:
					err = table.SetBoolByColIndex(colIndex, rowIndex, cell.(bool))
				case []interface{}: // This cell is a slice
					var interfaceSlice []interface{} = cell.([]interface{})
					var byteSlice []byte = []byte{}
					for _, sliceVal := range interfaceSlice {
						byteSlice = append(byteSlice, byte(sliceVal.(float64)))
					}
					err = table.SetByteSliceByColIndex(colIndex, rowIndex, byteSlice)
					if err != nil {
						return nil, err
					}
				case map[string]interface{}:	// This cell is a table.
					switch colType {
						case "*Table", "*gotables.Table":
						tableNested, err := newTableFromJSON_recursive(cell.(map[string]interface{}))
						if err != nil {
							return nil, err
						}
						err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
						if err != nil {
							return nil, err
						}
						default:
							return nil, fmt.Errorf("newTableFromJSON_recursive(): unexpected cell value at [%s].(%d,%d)",
								tableName, colIndex, rowIndex)
					}
				case nil:	// This cell is a nil table.
					switch colType {
						case "*Table", "*gotables.Table":
							var tableNested *Table = NewNilTable()
							err = table.SetTableByColIndex(colIndex, rowIndex, tableNested)
							if err != nil {
								return nil, err
							}
						default:
							return nil, fmt.Errorf("newTableFromJSON_recursive(): unexpected nil value at [%s].(%d,%d)",
								tableName, colIndex, rowIndex)
					}
				default:
					return nil, fmt.Errorf("%s ERROR %s: unexpected value of type: %v",
						UtilFuncSource(), UtilFuncName(), reflect.TypeOf(val))
				}
				// Single error handler for all the calls in this switch statement.
				if err != nil {
					return nil, fmt.Errorf("%s ERROR %s: %v", UtilFuncSource(), UtilFuncName(), err)
				}
			}
		}
	}
where(table)

where()
	return
}

///*
//	Marshal gotables TableSet to JSON
//*/
//func (tableSet *TableSet) GetTableSetAsJSON() (jsonStrings []string, err error) {
//where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))
//
//	if tableSet == nil {
//		return nil, fmt.Errorf("%s %s tableSet is <nil>", UtilFuncSource(), UtilFuncName())
//	}
//
//	for tableIndex := 0; tableIndex < len(tableSet.tables); tableIndex++ {
//
//		var table *Table
//		table, err = tableSet.TableByTableIndex(tableIndex)
//		if err != nil {
//			return nil, err
//		}
//
//		var jsonString string
//		jsonString, err = table.GetTableAsJSON()
//		if err != nil {
//			return nil, err
//		}
//		jsonStrings = append(jsonStrings, jsonString)
//	}
//
//	return
//}

/*
	Marshal gotables TableSet to JSON
*/
func (tableSet *TableSet) GetTableSetAsJSON() (jsonTableSet string, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	if tableSet == nil {
		return "", fmt.Errorf("%s ERROR: table.%s: table is <nil>", UtilFuncSource(), UtilFuncName())
	}

	var buf bytes.Buffer

	buf.WriteByte(123)	// Opening brace outermost
	buf.WriteString(fmt.Sprintf(`"tableSetName":%q,`, tableSet.Name()))
//	buf.WriteByte('[')	// Opening array of tables
	buf.WriteString(`"tables":[`)	// Opening array of tables

	var tableCount int = tableSet.TableCount()
	for tableIndex := 0; tableIndex < tableCount; tableIndex++ {
		table, err := tableSet.TableByTableIndex(tableIndex)
		if err != nil {
			return "", err
		}

where("***CALLING** getTableAsJSON()")
		var jsonTable string
		jsonTable, err = table.GetTableAsJSON()
		if err != nil {
			return "", err
		}

		buf.WriteString(jsonTable)

		if tableIndex < tableCount-1 {
			buf.WriteByte(',')	// Delimiter between tables
		}
	}

	buf.WriteByte(']')	// Closing array of tables
	buf.WriteByte(125)	// Closing brace outermost

	jsonTableSet = buf.String()

	return
}

/*
	Unmarshal JSON documents to a *gotables.TableSet
*/
func NewTableSetFromJSON(jsonTableSet string) (tableSet *TableSet, err error) {

	if jsonTableSet == "" {
		return nil, fmt.Errorf("%s: jsonTableSet is empty", UtilFuncName())
	}

	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonTableSet), &m)
	if err != nil {
		return nil, err
	}

	// (1) Retrieve and process TableSet name.
	var tableSetName string
	var exists bool
	tableSetName, exists = m["tableSetName"].(string)
	if !exists {
		return nil, fmt.Errorf("JSON is missing tableSet name")
	}

	tableSet, err = NewTableSet(tableSetName)
	if err != nil {
		return nil, err
	}

	// (2) Retrieve and process tables.
	var tablesMap []interface{}
	tablesMap, exists = m["tables"].([]interface{})
	if !exists {
		return nil, fmt.Errorf("JSON is missing tables")
	}

	var tableMap map[string]interface{}
	var tableMapInterface interface{}

	// Loop through the array of tables.
	for _, tableMapInterface = range tablesMap {

		tableMap = tableMapInterface.(map[string]interface{})

		var table *Table
		table, err = newTableFromJSON_recursive(tableMap)
		if err != nil {
			return nil, err
		}

		err = tableSet.Append(table)
		if err != nil {
			return nil, err
		}
	}

	return
}

func NewTableFromJSON(jsonString string) (table *Table, err error) {
where(fmt.Sprintf("***INSIDE*** %s", UtilFuncName()))

	// This is similar to NewTableFromString which first gets a TableSet.

	if jsonString == "" {
		return nil, fmt.Errorf("%s: jsonString is empty", UtilFuncName())
	}

	tableSet, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	tableCount := tableSet.TableCount()
	if tableCount != 1 {
		return nil, fmt.Errorf("%s: expecting a JSON string containing 1 table but found %d table%s",
			 UtilFuncName(), tableCount, plural(tableCount))
	}

	table, err = tableSet.TableByTableIndex(0)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	return table, nil
}

func NewTableFromJSONByTableName(jsonString string, tableName string) (table *Table, err error) {
	tableSet, err := NewTableSetFromJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	table, err = tableSet.Table(tableName)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", UtilFuncName(), err)
	}

	return
}
