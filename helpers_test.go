package gotables

/*
	helpers.go
*/

import (
	"bytes"
//	"fmt"
//	"os"
//	"runtime/debug"
	"testing"
)

/*
Copyright (c) 2017 Malcolm Gorman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

//	-----------------------------------------------------------
//	TestSet<type>() functions for each of 17 types.
//	-----------------------------------------------------------

//	Test Set and Get table cell in colName at rowIndex to newValue string
func TestSetAndGetString(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected string
	}{

		{ "ABC" },
		{ "abc" },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetString(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetString(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetString() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue bool
func TestSetAndGetBool(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected bool
	}{

		{ false },
		{ true },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetBool(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetBool(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetBool() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int
func TestSetAndGetInt(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int
	}{

		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int8
func TestSetAndGetInt8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int8
	}{

		{ -128 },
		{ 127 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt8(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt8(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int16
func TestSetAndGetInt16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int16
	}{

		{ -32768 },
		{ 32767 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt16(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt16(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int32
func TestSetAndGetInt32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int32
	}{

		{ -2147483648 },
		{ 2147483647 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue int64
func TestSetAndGetInt64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int64
	}{

		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetInt64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint
func TestSetAndGetUint(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint
	}{

		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue byte
func TestSetAndGetByte(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected byte
	}{

		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByte(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetByte(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetByte() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint8
func TestSetAndGetUint8(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint8
	}{

		{ 0 },
		{ 255 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint8(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint8() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint16
func TestSetAndGetUint16(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint16
	}{

		{ 0 },
		{ 65535 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint16(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint16(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint16() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint32
func TestSetAndGetUint32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint32
	}{

		{ 0 },
		{ 4294967295 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue uint64
func TestSetAndGetUint64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint64
	}{

		{ 0 },
		{ 18446744073709551615 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float32
func TestSetAndGetFloat32(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float32
	}{

		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat32(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetFloat32(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetFloat32() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue float64
func TestSetAndGetFloat64(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float64
	}{

		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetFloat64(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetFloat64(colName, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetFloat64() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []byte
func TestSetAndGetByteSlice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []byte
	}{

		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetByteSlice(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetByteSlice(colName, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting GetByteSlice() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colName at rowIndex to newValue []uint8
func TestSetAndGetUint8Slice(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []uint8
	}{

		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const rowIndex = 0

	for _, test := range tests {

		err = table.SetUint8Slice(colName, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint8Slice(colName, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting GetUint8Slice() value %v, not %v", test.expected, value)
		}
	}
}

//	--------------------------------------------------------------------
//	TestSet<type>ByColIndex() functions for each of 17 types.
//	--------------------------------------------------------------------

//	Test Set and Get table cell in colIndex at rowIndex to newValue string
func TestSetAndGetStringByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "StringValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "string")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected string
	}{

		{ "ABC" },
		{ "abc" },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetStringByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetStringByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetStringByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue bool
func TestSetAndGetBoolByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "BoolValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "bool")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected bool
	}{

		{ false },
		{ true },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetBoolByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetBoolByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetBoolByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int
func TestSetAndGetIntByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "IntValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int
	}{

		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetIntByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetIntByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetIntByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int8
func TestSetAndGetInt8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int8
	}{

		{ -128 },
		{ 127 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt8ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt8ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int16
func TestSetAndGetInt16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int16
	}{

		{ -32768 },
		{ 32767 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt16ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt16ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int32
func TestSetAndGetInt32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int32
	}{

		{ -2147483648 },
		{ 2147483647 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt32ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue int64
func TestSetAndGetInt64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Int64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "int64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected int64
	}{

		{ -9223372036854775808 },
		{ 9223372036854775807 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetInt64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetInt64ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetInt64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint
func TestSetAndGetUintByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "UintValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint
	}{

		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUintByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUintByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUintByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue byte
func TestSetAndGetByteByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected byte
	}{

		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetByteByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetByteByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint8
func TestSetAndGetUint8ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint8
	}{

		{ 0 },
		{ 255 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint8ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint8ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint16
func TestSetAndGetUint16ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint16Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint16")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint16
	}{

		{ 0 },
		{ 65535 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint16ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint16ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint16ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint32
func TestSetAndGetUint32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint32
	}{

		{ 0 },
		{ 4294967295 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint32ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue uint64
func TestSetAndGetUint64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "uint64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected uint64
	}{

		{ 0 },
		{ 18446744073709551615 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint64ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetUint64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float32
func TestSetAndGetFloat32ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float32Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float32")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float32
	}{

		{ 1.401298464324817e-45 },
		{ 3.4028234663852886e+38 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat32ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetFloat32ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetFloat32ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue float64
func TestSetAndGetFloat64ByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Float64Value"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "float64")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected float64
	}{

		{ 5e-324 },
		{ 1.7976931348623157e+308 },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetFloat64ByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetFloat64ByColIndex(colIndex, rowIndex)

		if value != test.expected {

			t.Errorf("expecting GetFloat64ByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []byte
func TestSetAndGetByteSliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "ByteSliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]byte")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []byte
	}{

		{ []byte{ 0 } },
		{ []byte{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetByteSliceByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetByteSliceByColIndex(colIndex, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting GetByteSliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}

//	Test Set and Get table cell in colIndex at rowIndex to newValue []uint8
func TestSetAndGetUint8SliceByColIndex(t *testing.T) {

	// See: TestSet<type>() functions

	const colName string = "Uint8SliceValue"

    table, err := NewTable("SetAndGet")
    if err != nil { t.Error(err) }

	err = table.AppendCol(colName, "[]uint8")
    if err != nil { t.Error(err) }

	err = table.AppendRow()
    if err != nil { t.Error(err) }

	var tests = []struct {
		expected []uint8
	}{

		{ []uint8{ 0 } },
		{ []uint8{ 255 } },
	}

	const colIndex = 0
	const rowIndex = 0

	for _, test := range tests {
		err = table.SetUint8SliceByColIndex(colIndex, rowIndex, test.expected)
	    if err != nil { t.Error(err) }

		value, _ := table.GetUint8SliceByColIndex(colIndex, rowIndex)

		if !bytes.Equal(value, test.expected) {

			t.Errorf("expecting GetUint8SliceByColIndex() value %v, not %v", test.expected, value)
		}
	}
}
