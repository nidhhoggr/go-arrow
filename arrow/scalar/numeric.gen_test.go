// Code generated by scalar/numeric.gen_test.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scalar_test

import (
	"testing"

	"github.com/nidhhoggr/go-arrow/arrow"
	"github.com/nidhhoggr/go-arrow/arrow/scalar"
	"github.com/stretchr/testify/assert"
)

func TestBasicInt8Scalars(t *testing.T) {
	value := int8(1)

	scalarVal := scalar.NewInt8Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Int8
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := int8(2)
	scalarOther := scalar.NewInt8Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Int8)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarInt8(t *testing.T) {
	three := scalar.MakeScalar(int8(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewInt8Scalar(3), three)

	assertMakeScalar(t, scalar.NewInt8Scalar(3), int8(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Int8, "3", scalar.NewInt8Scalar(3))
}

func TestBasicInt16Scalars(t *testing.T) {
	value := int16(1)

	scalarVal := scalar.NewInt16Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Int16
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := int16(2)
	scalarOther := scalar.NewInt16Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Int16)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarInt16(t *testing.T) {
	three := scalar.MakeScalar(int16(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewInt16Scalar(3), three)

	assertMakeScalar(t, scalar.NewInt16Scalar(3), int16(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Int16, "3", scalar.NewInt16Scalar(3))
}

func TestBasicInt32Scalars(t *testing.T) {
	value := int32(1)

	scalarVal := scalar.NewInt32Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Int32
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := int32(2)
	scalarOther := scalar.NewInt32Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Int32)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarInt32(t *testing.T) {
	three := scalar.MakeScalar(int32(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewInt32Scalar(3), three)

	assertMakeScalar(t, scalar.NewInt32Scalar(3), int32(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Int32, "3", scalar.NewInt32Scalar(3))
}

func TestBasicInt64Scalars(t *testing.T) {
	value := int64(1)

	scalarVal := scalar.NewInt64Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Int64
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := int64(2)
	scalarOther := scalar.NewInt64Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Int64)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarInt64(t *testing.T) {
	three := scalar.MakeScalar(int64(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewInt64Scalar(3), three)

	assertMakeScalar(t, scalar.NewInt64Scalar(3), int64(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Int64, "3", scalar.NewInt64Scalar(3))
}

func TestBasicUint8Scalars(t *testing.T) {
	value := uint8(1)

	scalarVal := scalar.NewUint8Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Uint8
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := uint8(2)
	scalarOther := scalar.NewUint8Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Uint8)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarUint8(t *testing.T) {
	three := scalar.MakeScalar(uint8(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewUint8Scalar(3), three)

	assertMakeScalar(t, scalar.NewUint8Scalar(3), uint8(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Uint8, "3", scalar.NewUint8Scalar(3))
}

func TestBasicUint16Scalars(t *testing.T) {
	value := uint16(1)

	scalarVal := scalar.NewUint16Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Uint16
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := uint16(2)
	scalarOther := scalar.NewUint16Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Uint16)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarUint16(t *testing.T) {
	three := scalar.MakeScalar(uint16(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewUint16Scalar(3), three)

	assertMakeScalar(t, scalar.NewUint16Scalar(3), uint16(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Uint16, "3", scalar.NewUint16Scalar(3))
}

func TestBasicUint32Scalars(t *testing.T) {
	value := uint32(1)

	scalarVal := scalar.NewUint32Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Uint32
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := uint32(2)
	scalarOther := scalar.NewUint32Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Uint32)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarUint32(t *testing.T) {
	three := scalar.MakeScalar(uint32(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewUint32Scalar(3), three)

	assertMakeScalar(t, scalar.NewUint32Scalar(3), uint32(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Uint32, "3", scalar.NewUint32Scalar(3))
}

func TestBasicUint64Scalars(t *testing.T) {
	value := uint64(1)

	scalarVal := scalar.NewUint64Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Uint64
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := uint64(2)
	scalarOther := scalar.NewUint64Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Uint64)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarUint64(t *testing.T) {
	three := scalar.MakeScalar(uint64(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewUint64Scalar(3), three)

	assertMakeScalar(t, scalar.NewUint64Scalar(3), uint64(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Uint64, "3", scalar.NewUint64Scalar(3))
}

func TestBasicFloat32Scalars(t *testing.T) {
	value := float32(1)

	scalarVal := scalar.NewFloat32Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Float32
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := float32(2)
	scalarOther := scalar.NewFloat32Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Float32)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarFloat32(t *testing.T) {
	three := scalar.MakeScalar(float32(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewFloat32Scalar(3), three)

	assertMakeScalar(t, scalar.NewFloat32Scalar(3), float32(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Float32, "3", scalar.NewFloat32Scalar(3))
}

func TestBasicFloat64Scalars(t *testing.T) {
	value := float64(1)

	scalarVal := scalar.NewFloat64Scalar(value)
	assert.Equal(t, value, scalarVal.Value)
	assert.True(t, scalarVal.IsValid())
	assert.NoError(t, scalarVal.ValidateFull())

	expectedType := arrow.PrimitiveTypes.Float64
	assert.True(t, arrow.TypeEqual(scalarVal.DataType(), expectedType))

	other := float64(2)
	scalarOther := scalar.NewFloat64Scalar(other)
	assert.NotEqual(t, scalarVal, scalarOther)
	assert.False(t, scalar.Equals(scalarVal, scalarOther))

	scalarVal.Value = other
	assert.Equal(t, other, scalarVal.Value)
	assert.Equal(t, scalarVal, scalarOther)
	assert.True(t, scalar.Equals(scalarVal, scalarOther))

	nullVal := scalar.MakeNullScalar(arrow.PrimitiveTypes.Float64)
	assert.False(t, nullVal.IsValid())
	assert.NoError(t, nullVal.ValidateFull())
}

func TestMakeScalarFloat64(t *testing.T) {
	three := scalar.MakeScalar(float64(3))
	assert.NoError(t, three.ValidateFull())
	assert.Equal(t, scalar.NewFloat64Scalar(3), three)

	assertMakeScalar(t, scalar.NewFloat64Scalar(3), float64(3))
	assertParseScalar(t, arrow.PrimitiveTypes.Float64, "3", scalar.NewFloat64Scalar(3))
}
