package utils

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

// GetEnvWithDefault returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found or is empty.
func GetEnvWithDefault(key, defaultValue string) string {
	return GetEnv(key, defaultValue, false)
}

// GetEnvWithDefaultNE returns the value of the environment variable named by
// the key. It is guaranteed to return the default value if the environment
// variable is not found. If the environment variable is found, it is guaranteed
// to return the default value if the environment variable is empty.
func GetEnvWithDefaultNE(key, defaultValue string) string {
	return GetEnv(key, defaultValue, true)
}

// GetEnv returns the value of the environment variable named by
// the key. It is guaranteed to return the `defaultValue` if the environment
// variable is not found. If the environment variable is found, it is guaranteed
// to return the `defaultValue` if the environment variable is empty and
// `nonEmpty` is `true`.
func GetEnv(key, defaultValue string, nonEmpty bool) string {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	if nonEmpty && "" == val {
		return defaultValue
	}
	return val
}

func GetEnvCsv(key string, defaultValue []string) []string {
	val := GetEnvWithDefault(key, "")
	if "" == val {
		return defaultValue
	}
	return slices.DeleteFunc(
		strings.Split(val, ","),
		func(s string) bool { return "" == strings.TrimSpace(s) },
	)
}

func GetEnvInt(key string, defaultValue int64, bitSize int) (int64, error) {
	val := GetEnvWithDefault(key, "")
	if "" == val {
		return defaultValue, nil
	}
	ret, err := strconv.ParseInt(val, 10, bitSize)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetEnvInt8(key string, defaultValue int8) (int8, error) {
	ret, err := GetEnvInt(key, int64(defaultValue), 8)
	return int8(ret), err
}

func GetEnvInt16(key string, defaultValue int16) (int16, error) {
	ret, err := GetEnvInt(key, int64(defaultValue), 16)
	return int16(ret), err
}

func GetEnvInt32(key string, defaultValue int32) (int32, error) {
	ret, err := GetEnvInt(key, int64(defaultValue), 32)
	return int32(ret), err
}

func GetEnvInt64(key string, defaultValue int64) (int64, error) {
	return GetEnvInt(key, int64(defaultValue), 64)
}

func GetEnvUint(key string, defaultValue uint64, bitSize int) (uint64, error) {
	val := GetEnvWithDefault(key, "")
	if "" == val {
		return defaultValue, nil
	}
	ret, err := strconv.ParseUint(val, 10, bitSize)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetEnvUint8(key string, defaultValue uint8) (uint8, error) {
	ret, err := GetEnvUint(key, uint64(defaultValue), 8)
	return uint8(ret), err
}

func GetEnvUint16(key string, defaultValue uint16) (uint16, error) {
	ret, err := GetEnvUint(key, uint64(defaultValue), 16)
	return uint16(ret), err
}

func GetEnvUint32(key string, defaultValue uint32) (uint32, error) {
	ret, err := GetEnvUint(key, uint64(defaultValue), 32)
	return uint32(ret), err
}

func GetEnvUint64(key string, defaultValue uint64) (uint64, error) {
	return GetEnvUint(key, uint64(defaultValue), 64)
}

func GetEnvFloat(key string, defaultValue float64, bitSize int) (
	float64, error,
) {
	val := GetEnvWithDefault(key, "")
	if "" == val {
		return defaultValue, nil
	}
	ret, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func GetEnvFloat32(key string, defaultValue float32) (float32, error) {
	ret, err := GetEnvFloat(key, float64(defaultValue), 32)
	return float32(ret), err
}

func GetEnvFloat64(key string, defaultValue float64) (float64, error) {
	return GetEnvFloat(key, defaultValue, 64)
}

func GetEnvBool(key string, defaultValue bool) (bool, error) {
	val := GetEnvWithDefault(key, "")
	if "" == val {
		return defaultValue, nil
	}
	ret, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func GetEnvIntCsv(key string, defaultValue []int64, bitSize int) (
	[]int64, error,
) {
	ss := GetEnvCsv(key, nil)
	if nil == ss {
		return defaultValue, nil
	}
	vs, err := SliceMapFunc[[]string, []int64](
		ss, func(s string) (int64, error) {
			return strconv.ParseInt(s, 10, bitSize)
		},
	)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func GetEnvInt8Csv(key string, defaultValue []int8) ([]int8, error) {
	dv, _ := SliceMapFunc[[]int8, []int64](
		defaultValue, func(i int8) (int64, error) { return int64(i), nil },
	)
	vs, err := GetEnvIntCsv(key, dv, 8)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]int64, []int8](
		vs, func(i int64) (int8, error) { return int8(i), nil },
	)
}

func GetEnvInt16Csv(key string, defaultValue []int16) ([]int16, error) {
	dv, _ := SliceMapFunc[[]int16, []int64](
		defaultValue, func(i int16) (int64, error) { return int64(i), nil },
	)
	vs, err := GetEnvIntCsv(key, dv, 16)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]int64, []int16](
		vs, func(i int64) (int16, error) { return int16(i), nil },
	)
}

func GetEnvInt32Csv(key string, defaultValue []int32) ([]int32, error) {
	dv, _ := SliceMapFunc[[]int32, []int64](
		defaultValue, func(i int32) (int64, error) { return int64(i), nil },
	)
	vs, err := GetEnvIntCsv(key, dv, 32)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]int64, []int32](
		vs, func(i int64) (int32, error) { return int32(i), nil },
	)
}

func GetEnvInt64Csv(key string, defaultValue []int64) ([]int64, error) {
	return GetEnvIntCsv(key, defaultValue, 64)
}

func GetEnvUintCsv(key string, defaultValue []uint64, bitSize int) (
	[]uint64, error,
) {
	ss := GetEnvCsv(key, nil)
	if nil == ss {
		return defaultValue, nil
	}
	vs, err := SliceMapFunc[[]string, []uint64](
		ss, func(s string) (uint64, error) {
			return strconv.ParseUint(s, 10, bitSize)
		},
	)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func GetEnvUint8Csv(key string, defaultValue []uint8) ([]uint8, error) {
	dv, _ := SliceMapFunc[[]uint8, []uint64](
		defaultValue, func(i uint8) (uint64, error) { return uint64(i), nil },
	)
	vs, err := GetEnvUintCsv(key, dv, 8)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]uint64, []uint8](
		vs, func(i uint64) (uint8, error) { return uint8(i), nil },
	)
}

func GetEnvUint16Csv(key string, defaultValue []uint16) ([]uint16, error) {
	dv, _ := SliceMapFunc[[]uint16, []uint64](
		defaultValue, func(i uint16) (uint64, error) { return uint64(i), nil },
	)
	vs, err := GetEnvUintCsv(key, dv, 16)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]uint64, []uint16](
		vs, func(i uint64) (uint16, error) { return uint16(i), nil },
	)
}

func GetEnvUint32Csv(key string, defaultValue []uint32) ([]uint32, error) {
	dv, _ := SliceMapFunc[[]uint32, []uint64](
		defaultValue, func(i uint32) (uint64, error) { return uint64(i), nil },
	)
	vs, err := GetEnvUintCsv(key, dv, 32)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]uint64, []uint32](
		vs, func(i uint64) (uint32, error) { return uint32(i), nil },
	)
}

func GetEnvUint64Csv(key string, defaultValue []uint64) ([]uint64, error) {
	return GetEnvUintCsv(key, defaultValue, 64)
}

func GetEnvFloatCsv(key string, defaultValue []float64, bitSize int) (
	[]float64, error,
) {
	ss := GetEnvCsv(key, nil)
	if nil == ss {
		return defaultValue, nil
	}
	vs, err := SliceMapFunc[[]string, []float64](
		ss, func(s string) (float64, error) {
			return strconv.ParseFloat(
				s, bitSize,
			)
		},
	)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func GetEnvFloat32Csv(key string, defaultValue []float32) ([]float32, error) {
	dv, _ := SliceMapFunc[[]float32, []float64](
		defaultValue,
		func(f float32) (float64, error) { return float64(f), nil },
	)
	vs, err := GetEnvFloatCsv(key, dv, 32)
	if err != nil {
		return nil, err
	}
	return SliceMapFunc[[]float64, []float32](
		vs, func(f float64) (float32, error) { return float32(f), nil },
	)
}

func GetEnvFloat64Csv(key string, defaultValue []float64) ([]float64, error) {
	return GetEnvFloatCsv(key, defaultValue, 64)
}

func GetEnvBoolCsv(key string, defaultValue []bool) ([]bool, error) {
	ss := GetEnvCsv(key, nil)
	if nil == ss {
		return defaultValue, nil
	}
	vs, err := SliceMapFunc[[]string, []bool](
		ss,
		func(s string) (bool, error) { return strconv.ParseBool(s) },
	)
	if err != nil {
		return nil, err
	}
	return vs, nil
}
