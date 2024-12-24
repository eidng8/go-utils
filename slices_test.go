package utils

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type sut struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

func Test_CloneDeepJsonable_primitive(t *testing.T) {
	sut1 := []int{1, 2, 3}
	actual, err := CloneDeepJsonable(&sut1)
	require.Nil(t, err)
	require.Equal(t, &sut1, actual)
}

func Test_CloneDeepJsonable_struct(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	actual, err := CloneDeepJsonable(&sut1)
	require.Nil(t, err)
	require.Equal(t, &sut1, actual)
}

func Test_CloneDeepJsonable_nil_returns_nil(t *testing.T) {
	actual, err := CloneDeepJsonable[*any, any](nil)
	require.Nil(t, err)
	require.Nil(t, actual)
}

// type jsonMockMarshalError struct {
// }
//
// func (j jsonMockMarshalError) BorrowIterator(data []byte) *jsoniter.Iterator {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) ReturnIterator(iter *jsoniter.Iterator) {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) BorrowStream(writer io.Writer) *jsoniter.Stream {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) ReturnStream(stream *jsoniter.Stream) {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) MarshalToString(v interface{}) (string, error) {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) Marshal(v interface{}) ([]byte, error) {
//     return nil, errors.New("test")
// }
//
// func (j jsonMockMarshalError) MarshalIndent(
//     v interface{}, prefix, indent string,
// ) (
//     []byte, error,
// ) {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) UnmarshalFromString(
//     str string, v interface{},
// ) error {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) Unmarshal(data []byte, v interface{}) error {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) Get(
//     data []byte, path ...interface{},
// ) jsoniter.Any {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) NewEncoder(writer io.Writer) *jsoniter.Encoder {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) NewDecoder(reader io.Reader) *jsoniter.Decoder {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) Valid(data []byte) bool {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) RegisterExtension(extension jsoniter.Extension) {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) DecoderOf(typ reflect2.Type) jsoniter.ValDecoder {
//     // TODO implement me
//     panic("implement me")
// }
//
// func (j jsonMockMarshalError) EncoderOf(typ reflect2.Type) jsoniter.ValEncoder {
//     // TODO implement me
//     panic("implement me")
// }
//
// var _ jsoniter.API = (*jsonMockMarshalError)(nil)

func Test_CloneDeepJsonable_marshal_throws_error(t *testing.T) {
	sut1 := []int{1, 2, 3}
	oj := Jsoniter
	defer func() { Jsoniter = oj }()
	Jsoniter = NewMockAPI(t)
	Jsoniter.(*MockAPI).EXPECT().Marshal(&sut1).Return(nil, errors.New("test"))
	actual, err := CloneDeepJsonable(&sut1)
	require.NotNil(t, err)
	require.Nil(t, actual)
}

func Test_CloneDeepJsonable_unmarshal_throws_error(t *testing.T) {
	var tv []int
	sut1 := []int{1, 2, 3}
	oj := Jsoniter
	defer func() { Jsoniter = oj }()
	Jsoniter = NewMockAPI(t)
	Jsoniter.(*MockAPI).EXPECT().Marshal(&sut1).Once().
		Return([]byte("[1,2,3]"), nil)
	Jsoniter.(*MockAPI).EXPECT().Unmarshal([]byte("[1,2,3]"), &tv).
		Return(errors.New("test"))
	actual, err := CloneDeepJsonable(&sut1)
	require.NotNil(t, err)
	require.Nil(t, actual)
}

func Test_Filter(t *testing.T) {
	sut1 := []int{1, 2, 3}
	actual := FilterFunc(
		sut1, func(v, idx int, a []int) bool {
			require.Equal(t, sut1[idx], v)
			require.Equal(t, sut1, a)
			return v > 1
		},
	)
	require.Equal(t, []int{2, 3}, actual)
}

func Test_Filter_not_found_returns_nil(t *testing.T) {
	sut1 := []int{1, 2, 3}
	actual := FilterFunc(
		sut1, func(v, idx int, a []int) bool {
			require.Equal(t, sut1[idx], v)
			require.Equal(t, sut1, a)
			return v > 3
		},
	)
	require.Nil(t, actual)
}

func Test_ContainsAny(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{2, 3}
	require.True(t, ContainsAny(sut1, sut2))
}

func Test_ContainsAny_not_found(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{4, 5}
	require.False(t, ContainsAny(sut1, sut2))
}

func Test_Intersect(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{2, 3}
	actual := Intersect(sut1, sut2)
	require.Equal(t, []int{2, 3}, actual)
}

func Test_Intersect_not_found_returns_nil(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{4, 5}
	actual := Intersect(sut1, sut2)
	require.Nil(t, actual)
}

func Test_IntersectFunc_with_primitive_array(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{2, 3}
	actual := IntersectFunc(
		sut1, sut2, func(v, idx int, a []int) int {
			require.Equal(t, a[idx], v)
			return v
		},
	)
	require.Equal(t, []int{2, 3}, actual)
}

func Test_IntersectFunc_with_struct_array(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []sut{{2, "two"}, {4, "four"}}
	actual := IntersectFunc(
		sut1, sut2, func(v sut, idx int, a []sut) int {
			require.Equal(t, a[idx], v)
			return v.Field1
		},
	)
	require.Equal(t, []sut{{2, "two"}}, actual)
}

func Test_IntersectFunc_with_struct_ptr_array(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []sut{{2, "two"}, {4, "four"}}
	sut3 := []*sut{&sut1[0], &sut1[1], &sut1[2]}
	sut4 := []*sut{&sut2[0], &sut2[1]}
	actual := IntersectFunc(
		sut3, sut4, func(v *sut, idx int, a []*sut) int {
			require.Equal(t, a[idx], v)
			return v.Field1
		},
	)
	require.Equal(t, []*sut{&sut1[1]}, actual)
}

func Test_IntersectFunc_not_found_returns_nil(t *testing.T) {
	sut1 := []int{1, 2, 3}
	sut2 := []int{4, 5}
	actual := IntersectFunc(
		sut1, sut2, func(v, idx int, a []int) int {
			require.Equal(t, a[idx], v)
			return v
		},
	)
	require.Nil(t, actual)
}

func Test_JoinInteger(t *testing.T) {
	require.Equal(t, "1,2,3", JoinInteger([]int{1, 2, 3}, ","))
}

func Test_MapToType(t *testing.T) {
	var sut interface{} = "abc"
	r, err := MapToType[string](sut)
	require.Nil(t, err)
	require.Equal(t, "abc", r)
}

func Test_MapToType_returns_error(t *testing.T) {
	var sut interface{} = "abc"
	_, err := MapToType[int](sut)
	require.NotNil(t, err)
}

func Test_SliceFindFunc(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r := SliceFindFunc(
		sut1, func(v, idx int, a []int) bool {
			require.Equal(t, sut1[idx], v)
			return 2 == v
		},
	)
	require.Equal(t, 2, r)
}

func Test_SliceFindFunc_struct(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	r := SliceFindFunc(
		sut1, func(v sut, idx int, a []sut) bool {
			require.Equal(t, sut1[idx], v)
			return 2 == v.Field1
		},
	)
	require.Equal(t, sut{2, "two"}, r)
}

func Test_SliceFindFunc_struct_ptr(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []*sut{&sut1[0], &sut1[1], &sut1[2]}
	r := SliceFindFunc(
		sut2, func(v *sut, idx int, a []*sut) bool {
			require.Equal(t, &sut1[idx], v)
			return 2 == v.Field1
		},
	)
	require.Equal(t, &sut1[1], r)
}

func Test_SliceFindFunc_not_found_returns_zero(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r := SliceFindFunc(
		sut1, func(v, idx int, a []int) bool {
			require.Equal(t, sut1[idx], v)
			return 4 == v
		},
	)
	require.Zero(t, r)
}

func Test_Pluck(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	actual := Pluck(
		sut1, func(v sut, idx int, a []sut) int {
			require.Equal(t, sut1[idx], v)
			return v.Field1
		},
	)
	require.Equal(t, []int{1, 2, 3}, actual)
}

func Test_Pluck_ptr(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []*sut{&sut1[0], &sut1[1], &sut1[2]}
	actual := Pluck(
		sut2, func(v *sut, idx int, a []*sut) int {
			require.Equal(t, &sut1[idx], v)
			return v.Field1
		},
	)
	require.Equal(t, []int{1, 2, 3}, actual)
}

func Test_SliceMapFunc(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r, err := SliceMapFunc[[]int, []string, int, string](
		sut1, func(v, idx int, a []int) (string, error) {
			require.Equal(t, sut1[idx], v)
			return strconv.Itoa(v), nil
		},
	)
	require.Nil(t, err)
	require.Equal(t, []string{"1", "2", "3"}, r)
}

func Test_SliceMapFunc_to_struct_ptr(t *testing.T) {
	sut1 := []int{1, 2, 3}
	exp := []sut{{1, "test"}, {2, "test"}, {3, "test"}}
	expected := []*sut{&exp[0], &exp[1], &exp[2]}
	r, err := SliceMapFunc[[]int, []*sut, int, *sut](
		sut1, func(v, idx int, a []int) (*sut, error) {
			require.Equal(t, sut1[idx], v)
			return &sut{v, "test"}, nil
		},
	)
	require.Nil(t, err)
	require.Equal(t, expected, r)
}

func Test_SliceMapFunc_from_struct_ptr(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []*sut{&sut1[0], &sut1[1], &sut1[2]}
	r, err := SliceMapFunc[[]*sut, []string, *sut, string](
		sut2, func(v *sut, idx int, a []*sut) (string, error) {
			require.Equal(t, &sut1[idx], v)
			return v.Field2, nil
		},
	)
	require.Nil(t, err)
	require.Equal(t, []string{"one", "two", "three"}, r)
}

func Test_SliceMapFunc_returns_error(t *testing.T) {
	sut1 := []int{1, 2, 3}
	r, err := SliceMapFunc[[]int, []string, int, string](
		sut1, func(v, idx int, a []int) (string, error) {
			require.Equal(t, sut1[idx], v)
			return "", errors.New("test")
		},
	)
	require.NotNil(t, err)
	require.Equal(t, "test", err.Error())
	require.Empty(t, r)
}

func Test_ApplyFunc_primitive(t *testing.T) {
	sut1 := []int{1, 2, 3}
	actual := ApplyFunc(
		sut1, func(v int, idx int, a []int) {
			require.Equal(t, sut1, a)
			require.Equal(t, sut1[idx], v)
			a[idx]++
		},
	)
	require.Equal(t, []int{2, 3, 4}, actual)
	require.Equal(t, []int{2, 3, 4}, sut1)
}

func Test_ApplyFunc_struct_pass_by_value_make_no_change(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	actual := ApplyFunc(
		sut1, func(v sut, idx int, a []sut) {
			require.Equal(t, sut1, a)
			require.Equal(t, sut1[idx], v)
			v.Field1++
		},
	)
	require.Equal(
		t, []sut{{1, "one"}, {2, "two"}, {3, "three"}}, actual,
	)
	require.Equal(
		t, []sut{{1, "one"}, {2, "two"}, {3, "three"}}, sut1,
	)
}

func Test_ApplyFunc_struct_ptr(t *testing.T) {
	sut1 := []sut{{1, "one"}, {2, "two"}, {3, "three"}}
	sut2 := []*sut{&sut1[0], &sut1[1], &sut1[2]}
	actual := ApplyFunc(
		sut2, func(v *sut, idx int, a []*sut) {
			require.Equal(t, sut2, a)
			require.Equal(t, sut2[idx], v)
			v.Field1++
		},
	)
	require.Equal(
		t, []*sut{{2, "one"}, {3, "two"}, {4, "three"}}, actual,
	)
	require.Equal(
		t, []sut{{2, "one"}, {3, "two"}, {4, "three"}}, sut1,
	)
}

func Test_ApplyFunc_empty_slice(t *testing.T) {
	var sut1 []int
	ApplyFunc(
		sut1, func(v int, _ int, _ []int) {
			require.Fail(t, "should not be called")
		},
	)
	require.Empty(t, sut1)
}
