// Code generated by genlib2. DO NOT EDIT.

package tensor

import (
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

var denseSetGetTests = []struct {
	of   Dtype
	data interface{}
	set  interface{}

	correct []interface{}
}{
	{Bool, []bool{true, false, true, false, true, false}, false, []interface{}{bool(true), bool(false), bool(true), bool(false), bool(true), bool(false)}},
	{Int, []int{0, 1, 2, 3, 4, 5}, 45, []interface{}{int(0), int(1), int(2), int(3), int(4), int(5)}},
	{Int8, []int8{0, 1, 2, 3, 4, 5}, 45, []interface{}{int8(0), int8(1), int8(2), int8(3), int8(4), int8(5)}},
	{Int16, []int16{0, 1, 2, 3, 4, 5}, 45, []interface{}{int16(0), int16(1), int16(2), int16(3), int16(4), int16(5)}},
	{Int32, []int32{0, 1, 2, 3, 4, 5}, 45, []interface{}{int32(0), int32(1), int32(2), int32(3), int32(4), int32(5)}},
	{Int64, []int64{0, 1, 2, 3, 4, 5}, 45, []interface{}{int64(0), int64(1), int64(2), int64(3), int64(4), int64(5)}},
	{Uint, []uint{0, 1, 2, 3, 4, 5}, 45, []interface{}{uint(0), uint(1), uint(2), uint(3), uint(4), uint(5)}},
	{Uint8, []uint8{0, 1, 2, 3, 4, 5}, 45, []interface{}{uint8(0), uint8(1), uint8(2), uint8(3), uint8(4), uint8(5)}},
	{Uint16, []uint16{0, 1, 2, 3, 4, 5}, 45, []interface{}{uint16(0), uint16(1), uint16(2), uint16(3), uint16(4), uint16(5)}},
	{Uint32, []uint32{0, 1, 2, 3, 4, 5}, 45, []interface{}{uint32(0), uint32(1), uint32(2), uint32(3), uint32(4), uint32(5)}},
	{Uint64, []uint64{0, 1, 2, 3, 4, 5}, 45, []interface{}{uint64(0), uint64(1), uint64(2), uint64(3), uint64(4), uint64(5)}},
	{Float32, []float32{0, 1, 2, 3, 4, 5}, 45, []interface{}{float32(0), float32(1), float32(2), float32(3), float32(4), float32(5)}},
	{Float64, []float64{0, 1, 2, 3, 4, 5}, 45, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}},
	{Complex64, []complex64{0, 1, 2, 3, 4, 5}, 45, []interface{}{complex64(0), complex64(1), complex64(2), complex64(3), complex64(4), complex64(5)}},
	{Complex128, []complex128{0, 1, 2, 3, 4, 5}, 45, []interface{}{complex128(0), complex128(1), complex128(2), complex128(3), complex128(4), complex128(5)}},
	{String, []string{"zero", "one", "two", "three", "four", "five"}, "HELLO WORLD", []interface{}{string("zero"), string("one"), string("two"), string("three"), string("four"), string("five")}},
}

func TestDense_setget(t *testing.T) {
	assert := assert.New(t)
	for _, gts := range denseSetGetTests {
		T := New(Of(gts.of), WithShape(len(gts.correct)))
		for i, v := range gts.correct {
			T.Set(i, v)
			got := T.Get(i)
			assert.Equal(v, got)
		}
	}
}

var denseMemsetTests = []struct {
	of    Dtype
	data  interface{}
	val   interface{}
	shape Shape

	correct interface{}
}{
	{Bool, []bool{true, false, true, false, true, false}, bool(false), Shape{2, 3}, []bool{false, false, false, false, false, false}},
	{Int, []int{0, 1, 2, 3, 4, 5}, int(45), Shape{2, 3}, []int{45, 45, 45, 45, 45, 45}},
	{Int8, []int8{0, 1, 2, 3, 4, 5}, int8(45), Shape{2, 3}, []int8{45, 45, 45, 45, 45, 45}},
	{Int16, []int16{0, 1, 2, 3, 4, 5}, int16(45), Shape{2, 3}, []int16{45, 45, 45, 45, 45, 45}},
	{Int32, []int32{0, 1, 2, 3, 4, 5}, int32(45), Shape{2, 3}, []int32{45, 45, 45, 45, 45, 45}},
	{Int64, []int64{0, 1, 2, 3, 4, 5}, int64(45), Shape{2, 3}, []int64{45, 45, 45, 45, 45, 45}},
	{Uint, []uint{0, 1, 2, 3, 4, 5}, uint(45), Shape{2, 3}, []uint{45, 45, 45, 45, 45, 45}},
	{Uint8, []uint8{0, 1, 2, 3, 4, 5}, uint8(45), Shape{2, 3}, []uint8{45, 45, 45, 45, 45, 45}},
	{Uint16, []uint16{0, 1, 2, 3, 4, 5}, uint16(45), Shape{2, 3}, []uint16{45, 45, 45, 45, 45, 45}},
	{Uint32, []uint32{0, 1, 2, 3, 4, 5}, uint32(45), Shape{2, 3}, []uint32{45, 45, 45, 45, 45, 45}},
	{Uint64, []uint64{0, 1, 2, 3, 4, 5}, uint64(45), Shape{2, 3}, []uint64{45, 45, 45, 45, 45, 45}},
	{Float32, []float32{0, 1, 2, 3, 4, 5}, float32(45), Shape{2, 3}, []float32{45, 45, 45, 45, 45, 45}},
	{Float64, []float64{0, 1, 2, 3, 4, 5}, float64(45), Shape{2, 3}, []float64{45, 45, 45, 45, 45, 45}},
	{Complex64, []complex64{0, 1, 2, 3, 4, 5}, complex64(45), Shape{2, 3}, []complex64{45, 45, 45, 45, 45, 45}},
	{Complex128, []complex128{0, 1, 2, 3, 4, 5}, complex128(45), Shape{2, 3}, []complex128{45, 45, 45, 45, 45, 45}},
	{String, []string{"zero", "one", "two", "three", "four", "five"}, string("HELLO WORLD"), Shape{2, 3}, []string{"HELLO WORLD", "HELLO WORLD", "HELLO WORLD", "HELLO WORLD", "HELLO WORLD", "HELLO WORLD"}},
}

func TestDense_memset(t *testing.T) {
	assert := assert.New(t)
	for _, mts := range denseMemsetTests {
		T := New(Of(mts.of), WithShape(mts.shape...))
		T.Memset(mts.val)
		assert.Equal(mts.correct, T.Data())

		T = New(Of(mts.of), WithShape(mts.shape...), WithBacking(mts.data))
		T2, _ := T.Slice(nil)
		T2.Memset(mts.val)
		assert.Equal(mts.correct, T2.Data())
	}
}

var denseZeroTests = []struct {
	of   Dtype
	data interface{}

	correct interface{}
}{
	{Bool, []bool{true, false, true, false, true, false}, []bool{false, false, false, false, false, false}},
	{Int, []int{0, 1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0, 0}},
	{Int8, []int8{0, 1, 2, 3, 4, 5}, []int8{0, 0, 0, 0, 0, 0}},
	{Int16, []int16{0, 1, 2, 3, 4, 5}, []int16{0, 0, 0, 0, 0, 0}},
	{Int32, []int32{0, 1, 2, 3, 4, 5}, []int32{0, 0, 0, 0, 0, 0}},
	{Int64, []int64{0, 1, 2, 3, 4, 5}, []int64{0, 0, 0, 0, 0, 0}},
	{Uint, []uint{0, 1, 2, 3, 4, 5}, []uint{0, 0, 0, 0, 0, 0}},
	{Uint8, []uint8{0, 1, 2, 3, 4, 5}, []uint8{0, 0, 0, 0, 0, 0}},
	{Uint16, []uint16{0, 1, 2, 3, 4, 5}, []uint16{0, 0, 0, 0, 0, 0}},
	{Uint32, []uint32{0, 1, 2, 3, 4, 5}, []uint32{0, 0, 0, 0, 0, 0}},
	{Uint64, []uint64{0, 1, 2, 3, 4, 5}, []uint64{0, 0, 0, 0, 0, 0}},
	{Float32, []float32{0, 1, 2, 3, 4, 5}, []float32{0, 0, 0, 0, 0, 0}},
	{Float64, []float64{0, 1, 2, 3, 4, 5}, []float64{0, 0, 0, 0, 0, 0}},
	{Complex64, []complex64{0, 1, 2, 3, 4, 5}, []complex64{0, 0, 0, 0, 0, 0}},
	{Complex128, []complex128{0, 1, 2, 3, 4, 5}, []complex128{0, 0, 0, 0, 0, 0}},
	{String, []string{"zero", "one", "two", "three", "four", "five"}, []string{"", "", "", "", "", ""}},
}

func TestDense_Zero(t *testing.T) {
	assert := assert.New(t)
	for _, mts := range denseZeroTests {

		typ := reflect.TypeOf(mts.data)
		val := reflect.ValueOf(mts.data)
		data := reflect.MakeSlice(typ, val.Len(), val.Cap())
		reflect.Copy(data, val)

		T := New(Of(mts.of), WithBacking(data.Interface()))
		T.Zero()
		assert.Equal(mts.correct, T.Data())

		T = New(Of(mts.of), WithBacking(mts.data))
		T2, _ := T.Slice(nil)
		T2.Zero()
		assert.Equal(mts.correct, T2.Data())
	}
}

func TestDense_Eq(t *testing.T) {
	eqFn := func(q *Dense) bool {
		a := q.Clone().(*Dense)
		if !q.Eq(a) {
			t.Error("Expected a clone to be exactly equal")
			return false
		}
		a.Zero()

		// Bools are excluded because the probability of having an array of all false is very high
		if q.Eq(a) && a.len() > 3 && a.Dtype() != Bool {
			t.Errorf("a %v", a.Data())
			t.Errorf("q %v", q.Data())
			t.Error("Expected *Dense to be not equal")
			return false
		}
		return true
	}
	if err := quick.Check(eqFn, &quick.Config{Rand: newRand(), MaxCount: quickchecks}); err != nil {
		t.Errorf("Failed to perform equality checks")
	}
}
