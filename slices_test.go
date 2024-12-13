package utils

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sut struct {
	field1 int
	field2 string
}

func Test_Pluck(t *testing.T) {
	sut1 := []sut{sut{1, "one"}, sut{2, "two"}, sut{3, "three"}}
	assert.Equal(
		t, []int{1, 2, 3}, Pluck(sut1, func(f sut) int { return f.field1 }),
	)
}

func Test_JoinInteger(t *testing.T) {
	assert.Equal(t, "1,2,3", JoinInteger([]int{1, 2, 3}, ","))
}

func Test_SliceMapFunc(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r, err := SliceMapFunc(
		sut1, func(i int) (string, error) { return strconv.Itoa(i), nil },
	)
	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, r)
}

func Test_SliceMapFunc_returns_error(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r, err := SliceMapFunc(
		sut1, func(i int) (string, error) { return "", errors.New("test") },
	)
	assert.NotNil(t, err)
	assert.Equal(t, "test", err.Error())
	assert.Empty(t, r)
}

func Test_MapToType(t *testing.T) {
	var (
		sut1 interface{} = 1
		sut2 interface{} = 2
		sut              = []interface{}{sut1, sut2}
	)
	r, err := SliceMapFunc(sut, MapToType[int])
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2}, r)
}

func Test_MapToType_returns_error(t *testing.T) {
	var (
		sut1 interface{} = "1"
		sut2 interface{} = "2"
		sut              = []interface{}{sut1, sut2}
	)
	_, err := SliceMapFunc(sut, MapToType[int])
	assert.NotNil(t, err)
}

func Test_SliceFindFunc(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r := SliceFindFunc(sut1, func(i int) bool { return i == 2 })
	assert.Equal(t, 2, r)
}

func Test_SliceFindFunc_not_found(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r := SliceFindFunc(sut1, func(i int) bool { return i == 4 })
	assert.Zero(t, r)
}
