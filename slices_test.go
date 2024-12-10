package utils

import (
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
