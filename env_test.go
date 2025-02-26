package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupEnvTest(tb testing.TB) {
	err := os.Setenv("TEST_ENV", "TEST")
	require.Nil(tb, err)
	err = os.Setenv("TEST_ENV_EMPTY", "")
	require.Nil(tb, err)
}

func Test_GetEnvWithDefault(t *testing.T) {
	setupEnvTest(t)
	got := GetEnvWithDefault("TEST_ENV", "defaultValue")
	require.Equal(t, "TEST", got)
	got = GetEnvWithDefault("TEST_ENV_EMPTY", "defaultValue")
	require.Empty(t, got)
	got = GetEnvWithDefault("NO_DEF", "defaultValue")
	require.Equal(t, "defaultValue", got)
}

func Test_GetEnvWithDefaultNE(t *testing.T) {
	setupEnvTest(t)
	got := GetEnvWithDefaultNE("TEST_ENV", "defaultValue")
	require.Equal(t, "TEST", got)
	got = GetEnvWithDefaultNE("TEST_ENV_EMPTY", "defaultValue")
	require.Equal(t, "defaultValue", got)
	got = GetEnvWithDefaultNE("NO_DEF", "defaultValue")
	require.Equal(t, "defaultValue", got)
}

func Test_MustGetEnv(t *testing.T) {
	setupEnvTest(t)
	got := MustGetEnv("TEST_ENV")
	require.Equal(t, "TEST", got)
	require.NoError(t, os.Unsetenv("TEST_ENV"))
	require.Panics(t, func() { MustGetEnv("TEST_ENV") })
}

func Test_MustGetEnvNE(t *testing.T) {
	setupEnvTest(t)
	got := MustGetEnvNE("TEST_ENV")
	require.Equal(t, "TEST", got)
	require.NoError(t, os.Setenv("TEST_ENV", ""))
	require.Panics(t, func() { MustGetEnvNE("TEST_ENV") })
}

func Test_GetEnvInt8(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     int8
		want    int8
		wantErr require.ValueAssertionFunc
	}{
		{"returns int8", "123", 0, 123, require.Nil},
		{"returns negative", "-123", 0, -123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "12345", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT8", tt.env))
				got, err := GetEnvInt8("TEST_ENV_INT8", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt16(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     int16
		want    int16
		wantErr require.ValueAssertionFunc
	}{
		{"returns int16", "123", 0, 123, require.Nil},
		{"returns negative", "-123", 0, -123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "1234567", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT16", tt.env))
				got, err := GetEnvInt16("TEST_ENV_INT16", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt32(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     int32
		want    int32
		wantErr require.ValueAssertionFunc
	}{
		{"returns int32", "123", 0, 123, require.Nil},
		{"returns negative", "-123", 0, -123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "12345678901", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT32", tt.env))
				got, err := GetEnvInt32("TEST_ENV_INT32", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt64(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     int64
		want    int64
		wantErr require.ValueAssertionFunc
	}{
		{"returns int64", "123", 0, 123, require.Nil},
		{"returns negative", "-123", 0, -123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{
			"overflow returns error", "12345678901234567890", 0, 0,
			require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT64", tt.env))
				got, err := GetEnvInt64("TEST_ENV_INT64", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint8(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     uint8
		want    uint8
		wantErr require.ValueAssertionFunc
	}{
		{"returns uint8", "123", 0, 123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"negative returns error", "-123", 0, 0, require.NotNil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "12345", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT8", tt.env))
				got, err := GetEnvUint8("TEST_ENV_UINT8", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint16(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     uint16
		want    uint16
		wantErr require.ValueAssertionFunc
	}{
		{"returns uint16", "123", 0, 123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"negative returns error", "-123", 0, 0, require.NotNil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "1234567", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT16", tt.env))
				got, err := GetEnvUint16("TEST_ENV_UINT16", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint32(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     uint32
		want    uint32
		wantErr require.ValueAssertionFunc
	}{
		{"returns uint32", "123", 0, 123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"negative returns error", "-123", 0, 0, require.NotNil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{"overflow returns error", "12345678901", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT32", tt.env))
				got, err := GetEnvUint32("TEST_ENV_UINT32", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint64(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     uint64
		want    uint64
		wantErr require.ValueAssertionFunc
	}{
		{"returns uint64", "123", 0, 123, require.Nil},
		{"empty returns default", "", 123, 123, require.Nil},
		{"negative returns error", "-123", 0, 0, require.NotNil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
		{"decimal number returns error", "123.45", 0, 0, require.NotNil},
		{
			"overflow returns error", "98765432109876543210", 0, 0,
			require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT64", tt.env))
				got, err := GetEnvUint64("TEST_ENV_UINT64", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvFloat32(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     float32
		want    float32
		wantErr require.ValueAssertionFunc
	}{
		{"returns float32", "123.4", 0, 123.4, require.Nil},
		{"returns negative", "-123.4", 0, -123.4, require.Nil},
		{"empty returns default", "", 123.4, 123.4, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_FLOAT32", tt.env))
				got, err := GetEnvFloat32("TEST_ENV_FLOAT32", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvFloat64(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     float64
		want    float64
		wantErr require.ValueAssertionFunc
	}{
		{"returns float64", "123.4", 0, 123.4, require.Nil},
		{"returns negative", "-123.4", 0, -123.4, require.Nil},
		{"empty returns default", "", 123.4, 123.4, require.Nil},
		{"invalid number returns error", "abc", 0, 0, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_FLOAT64", tt.env))
				got, err := GetEnvFloat64("TEST_ENV_FLOAT64", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvBool(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     bool
		want    bool
		wantErr require.ValueAssertionFunc
	}{
		{"returns true", "true", false, true, require.Nil},
		{"returns false", "false", true, false, require.Nil},
		{"empty returns default", "", true, true, require.Nil},
		{"invalid returns error", "trUe", false, false, require.NotNil},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_BOOL", tt.env))
				got, err := GetEnvBool("TEST_ENV_BOOL", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvCsv(t *testing.T) {
	tests := []struct {
		name string
		env  string
		def  []string
		want []string
	}{
		{
			"returns 1 values", "a", []string{}, []string{"a"},
		},
		{
			"returns 3 values", "a,b,c", []string{}, []string{"a", "b", "c"},
		},
		{
			"empty returns default", "", []string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			"trims empty values", ",a,,b, ,	,c,,", []string{},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_CSV", tt.env))
				got := GetEnvCsv("TEST_ENV_CSV", tt.def)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt8Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []int8
		want    []int8
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []int8{}, []int8{1}, require.Nil},
		{"returns 3 values", "1,2,3", []int8{}, []int8{1, 2, 3}, require.Nil},
		{
			"empty returns default", "", []int8{1, 2, 3},
			[]int8{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []int8{},
			[]int8{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []int8{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT8_CSV", tt.env))
				got, err := GetEnvInt8Csv("TEST_ENV_INT8_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt16Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []int16
		want    []int16
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []int16{}, []int16{1}, require.Nil},
		{"returns 3 values", "1,2,3", []int16{}, []int16{1, 2, 3}, require.Nil},
		{
			"empty returns default", "", []int16{1, 2, 3},
			[]int16{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []int16{},
			[]int16{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []int16{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT16_CSV", tt.env))
				got, err := GetEnvInt16Csv("TEST_ENV_INT16_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt32Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []int32
		want    []int32
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []int32{}, []int32{1}, require.Nil},
		{"returns 3 values", "1,2,3", []int32{}, []int32{1, 2, 3}, require.Nil},
		{
			"empty returns default", "", []int32{1, 2, 3},
			[]int32{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []int32{},
			[]int32{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []int32{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT32_CSV", tt.env))
				got, err := GetEnvInt32Csv("TEST_ENV_INT32_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvInt64Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []int64
		want    []int64
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []int64{}, []int64{1}, require.Nil},
		{"returns 3 values", "1,2,3", []int64{}, []int64{1, 2, 3}, require.Nil},
		{
			"empty returns default", "", []int64{1, 2, 3},
			[]int64{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []int64{},
			[]int64{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []int64{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_INT64_CSV", tt.env))
				got, err := GetEnvInt64Csv("TEST_ENV_INT64_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint8Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []uint8
		want    []uint8
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []uint8{}, []uint8{1}, require.Nil},
		{"returns 3 values", "1,2,3", []uint8{}, []uint8{1, 2, 3}, require.Nil},
		{
			"empty returns default", "", []uint8{1, 2, 3},
			[]uint8{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []uint8{},
			[]uint8{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []uint8{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT8_CSV", tt.env))
				got, err := GetEnvUint8Csv("TEST_ENV_UINT8_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint16Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []uint16
		want    []uint16
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []uint16{}, []uint16{1}, require.Nil},
		{
			"returns 3 values", "1,2,3", []uint16{}, []uint16{1, 2, 3},
			require.Nil,
		},
		{
			"empty returns default", "", []uint16{1, 2, 3},
			[]uint16{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []uint16{},
			[]uint16{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []uint16{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT16_CSV", tt.env))
				got, err := GetEnvUint16Csv("TEST_ENV_UINT16_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint32Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []uint32
		want    []uint32
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []uint32{}, []uint32{1}, require.Nil},
		{
			"returns 3 values", "1,2,3", []uint32{}, []uint32{1, 2, 3},
			require.Nil,
		},
		{
			"empty returns default", "", []uint32{1, 2, 3},
			[]uint32{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []uint32{},
			[]uint32{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []uint32{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT32_CSV", tt.env))
				got, err := GetEnvUint32Csv("TEST_ENV_UINT32_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvUint64Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []uint64
		want    []uint64
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1", []uint64{}, []uint64{1}, require.Nil},
		{
			"returns 3 values", "1,2,3", []uint64{}, []uint64{1, 2, 3},
			require.Nil,
		},
		{
			"empty returns default", "", []uint64{1, 2, 3},
			[]uint64{1, 2, 3}, require.Nil,
		},
		{
			"trims empty values", ",1,,2, ,	,3,,", []uint64{},
			[]uint64{1, 2, 3}, require.Nil,
		},
		{
			"invalid number returns error", "1,abc", []uint64{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_UINT64_CSV", tt.env))
				got, err := GetEnvUint64Csv("TEST_ENV_UINT64_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvFloat32Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []float32
		want    []float32
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1.2", []float32{}, []float32{1.2}, require.Nil},
		{
			"returns 3 values", "1.2,2.3,3.4", []float32{},
			[]float32{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"empty returns default", "", []float32{1.2, 2.3, 3.4},
			[]float32{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"trims empty values", ",1.2,,2.3, ,	,3.4,,", []float32{},
			[]float32{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"invalid number returns error", "1.2,abc", []float32{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_FLOAT32_CSV", tt.env))
				got, err := GetEnvFloat32Csv("TEST_ENV_FLOAT32_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvFloat64Csv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []float64
		want    []float64
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "1.2", []float64{}, []float64{1.2}, require.Nil},
		{
			"returns 3 values", "1.2,2.3,3.4", []float64{},
			[]float64{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"empty returns default", "", []float64{1.2, 2.3, 3.4},
			[]float64{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"trims empty values", ",1.2,,2.3, ,	,3.4,,", []float64{},
			[]float64{1.2, 2.3, 3.4}, require.Nil,
		},
		{
			"invalid number returns error", "1.2,abc", []float64{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_FLOAT64_CSV", tt.env))
				got, err := GetEnvFloat64Csv("TEST_ENV_FLOAT64_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}

func Test_GetEnvBoolCsv(t *testing.T) {
	tests := []struct {
		name    string
		env     string
		def     []bool
		want    []bool
		wantErr require.ValueAssertionFunc
	}{
		{"returns 1 values", "true", []bool{}, []bool{true}, require.Nil},
		{
			"returns 3 values", "t,f,t", []bool{},
			[]bool{true, false, true}, require.Nil,
		},
		{
			"empty returns default", "", []bool{true, false, true},
			[]bool{true, false, true}, require.Nil,
		},
		{
			"trims empty values", ",1,,0, ,	,1,,", []bool{},
			[]bool{true, false, true}, require.Nil,
		},
		{
			"invalid returns error", "true,TrUe", []bool{},
			nil, require.NotNil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				require.Nil(t, os.Setenv("TEST_ENV_BOOL_CSV", tt.env))
				got, err := GetEnvBoolCsv("TEST_ENV_BOOL_CSV", tt.def)
				tt.wantErr(t, err)
				require.Equal(t, tt.want, got)
			},
		)
	}
}
