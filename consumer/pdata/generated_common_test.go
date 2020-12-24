// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run cmd/pdatagen/main.go".

package pdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	otlpcommon "go.opentelemetry.io/collector/data/opentelemetry-proto-gen/common/v1"
)

func TestInstrumentationLibrary_CopyTo(t *testing.T) {
	ms := NewInstrumentationLibrary()
	generateTestInstrumentationLibrary().CopyTo(ms)
	assert.EqualValues(t, generateTestInstrumentationLibrary(), ms)
}

func TestInstrumentationLibrary_Name(t *testing.T) {
	ms := NewInstrumentationLibrary()
	assert.EqualValues(t, "", ms.Name())
	testValName := "test_name"
	ms.SetName(testValName)
	assert.EqualValues(t, testValName, ms.Name())
}

func TestInstrumentationLibrary_Version(t *testing.T) {
	ms := NewInstrumentationLibrary()
	assert.EqualValues(t, "", ms.Version())
	testValVersion := "test_version"
	ms.SetVersion(testValVersion)
	assert.EqualValues(t, testValVersion, ms.Version())
}

func TestAnyValueArray(t *testing.T) {
	es := NewAnyValueArray()
	assert.EqualValues(t, 0, es.Len())
	es = newAnyValueArray(&[]otlpcommon.AnyValue{})
	assert.EqualValues(t, 0, es.Len())

	es.Resize(7)
	emptyVal := NewAttributeValue()
	testVal := generateTestAttributeValue()
	assert.EqualValues(t, 7, es.Len())
	for i := 0; i < es.Len(); i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
		fillTestAttributeValue(es.At(i))
		assert.EqualValues(t, testVal, es.At(i))
	}
}

func TestAnyValueArray_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestAnyValueArray()
	dest := NewAnyValueArray()
	src := generateTestAnyValueArray()
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestAnyValueArray(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.EqualValues(t, generateTestAnyValueArray(), dest)
	assert.EqualValues(t, 0, src.Len())
	assert.EqualValues(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestAnyValueArray().MoveAndAppendTo(dest)
	assert.EqualValues(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i))
		assert.EqualValues(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestAnyValueArray_CopyTo(t *testing.T) {
	dest := NewAnyValueArray()
	// Test CopyTo to empty
	NewAnyValueArray().CopyTo(dest)
	assert.EqualValues(t, NewAnyValueArray(), dest)

	// Test CopyTo larger slice
	generateTestAnyValueArray().CopyTo(dest)
	assert.EqualValues(t, generateTestAnyValueArray(), dest)

	// Test CopyTo same size slice
	generateTestAnyValueArray().CopyTo(dest)
	assert.EqualValues(t, generateTestAnyValueArray(), dest)
}

func TestAnyValueArray_Resize(t *testing.T) {
	es := generateTestAnyValueArray()
	emptyVal := NewAttributeValue()
	// Test Resize less elements.
	const resizeSmallLen = 4
	expectedEs := make(map[*otlpcommon.AnyValue]bool, resizeSmallLen)
	for i := 0; i < resizeSmallLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, resizeSmallLen, len(expectedEs))
	es.Resize(resizeSmallLen)
	assert.Equal(t, resizeSmallLen, es.Len())
	foundEs := make(map[*otlpcommon.AnyValue]bool, resizeSmallLen)
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)

	// Test Resize more elements.
	const resizeLargeLen = 7
	oldLen := es.Len()
	expectedEs = make(map[*otlpcommon.AnyValue]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).orig] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.Resize(resizeLargeLen)
	assert.Equal(t, resizeLargeLen, es.Len())
	foundEs = make(map[*otlpcommon.AnyValue]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).orig] = true
	}
	assert.EqualValues(t, expectedEs, foundEs)
	for i := oldLen; i < resizeLargeLen; i++ {
		assert.EqualValues(t, emptyVal, es.At(i))
	}

	// Test Resize 0 elements.
	es.Resize(0)
	assert.Equal(t, 0, es.Len())
}

func TestAnyValueArray_Append(t *testing.T) {
	es := generateTestAnyValueArray()

	emptyVal := NewAttributeValue()
	es.Append(emptyVal)
	assert.EqualValues(t, *(es.At(7)).orig, *emptyVal.orig)

	emptyVal2 := NewAttributeValue()
	es.Append(emptyVal2)
	assert.EqualValues(t, *(es.At(8)).orig, *emptyVal2.orig)

	assert.Equal(t, 9, es.Len())
}

func generateTestInstrumentationLibrary() InstrumentationLibrary {
	tv := NewInstrumentationLibrary()
	fillTestInstrumentationLibrary(tv)
	return tv
}

func fillTestInstrumentationLibrary(tv InstrumentationLibrary) {
	tv.SetName("test_name")
	tv.SetVersion("test_version")
}

func generateTestAnyValueArray() AnyValueArray {
	tv := NewAnyValueArray()
	fillTestAnyValueArray(tv)
	return tv
}

func fillTestAnyValueArray(tv AnyValueArray) {
	tv.Resize(7)
	for i := 0; i < tv.Len(); i++ {
		fillTestAttributeValue(tv.At(i))
	}
}
