package utils

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func resetPasswordHashParams(t *testing.T) {
	require.Nil(t, os.Unsetenv(PasswordHashTimesName))
	require.Nil(t, os.Unsetenv(PasswordHashMemoryName))
	require.Nil(t, os.Unsetenv(PasswordHashThreadsName))
	require.Nil(t, os.Unsetenv(PasswordHashKeyLenName))
	require.Nil(t, os.Unsetenv(PasswordHashSaltLenName))
}

func Test_password_hash_and_compare(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashTimesName, "2"))
	require.Nil(t, os.Setenv(PasswordHashMemoryName, "32768"))
	require.Nil(t, os.Setenv(PasswordHashThreadsName, "6"))
	require.Nil(t, os.Setenv(PasswordHashKeyLenName, "64"))
	require.Nil(t, os.Setenv(PasswordHashSaltLenName, "32"))
	params, err := DefaultPasswordHashParams()
	require.Nil(t, err)
	require.Equal(
		t,
		PasswordHashParams{
			Times:   2,
			Memory:  32768,
			Threads: 6,
			KeyLen:  64,
			SaltLen: 32,
		},
		*params,
	)
	hash, err := HashPassword("test password", *params)
	require.Nil(t, err)
	eq, err := ComparePassword("test password", hash)
	require.Nil(t, err)
	require.True(t, eq)
}

func Test_password_hash_and_compare_returns_false_if_not_equal(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashTimesName, "2"))
	require.Nil(t, os.Setenv(PasswordHashMemoryName, "32768"))
	require.Nil(t, os.Setenv(PasswordHashThreadsName, "6"))
	require.Nil(t, os.Setenv(PasswordHashKeyLenName, "64"))
	require.Nil(t, os.Setenv(PasswordHashSaltLenName, "32"))
	params, err := DefaultPasswordHashParams()
	require.Nil(t, err)
	require.Equal(
		t,
		PasswordHashParams{
			Times:   2,
			Memory:  32768,
			Threads: 6,
			KeyLen:  64,
			SaltLen: 32,
		},
		*params,
	)
	hash, err := HashPassword("test password", *params)
	require.Nil(t, err)
	eq, err := ComparePassword("not password", hash)
	require.Nil(t, err)
	require.False(t, eq)
}

func Test_hash_time_returns_error_if_invalid_value(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashTimesName, "invalid"))
	_, err := DefaultPasswordHashParams()
	require.NotNil(t, err)
}

func Test_hash_memory_returns_error_if_invalid_value(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashMemoryName, "invalid"))
	_, err := DefaultPasswordHashParams()
	require.NotNil(t, err)
}

func Test_hash_threads_returns_error_if_invalid_value(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashThreadsName, "invalid"))
	_, err := DefaultPasswordHashParams()
	require.NotNil(t, err)
}

func Test_hash_key_len_returns_error_if_invalid_value(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashKeyLenName, "invalid"))
	_, err := DefaultPasswordHashParams()
	require.NotNil(t, err)
}

func Test_hash_salt_len_returns_error_if_invalid_value(t *testing.T) {
	defer resetPasswordHashParams(t)
	require.Nil(t, os.Setenv(PasswordHashSaltLenName, "invalid"))
	_, err := DefaultPasswordHashParams()
	require.NotNil(t, err)
}

func Test_HashPassword_handles_random_bytes_error(t *testing.T) {
	defer resetPasswordHashParams(t)
	params, err := DefaultPasswordHashParams()
	require.Nil(t, err)
	tmp := randomBytes
	defer func() { randomBytes = tmp }()
	randomBytes = func([]byte) (int, error) { return 0, errors.New("error") }
	_, err = HashPassword("test password", *params)
	require.NotNil(t, err)

}

func Test_ComparePassword_returns_error_if_invalid_hash_format(t *testing.T) {
	_, err := ComparePassword("password", "$$")
	require.ErrorIs(t, err, ErrInvalidHashFormat)
}

func Test_ComparePassword_returns_error_if_invalid_hash_algorithm(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$invalid$v=19$m=32768,t=2,p=6$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.ErrorIs(t, err, ErrInvalidHashAlgorithm)
}

func Test_ComparePassword_returns_error_if_version_wrong_format(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$m=1$m=32768,t=2,p=6$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.NotNil(t, err)
}

func Test_ComparePassword_returns_error_if_invalid_version(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=1$m=32768,t=2,p=6$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.ErrorIs(t, err, ErrInvalidHashVersion)
}

func Test_ComparePassword_returns_error_if_hash_params_empty(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.ErrorIs(t, err, ErrInvalidHashFormat)
}

func Test_ComparePassword_returns_error_if_invalid_hash_params_format(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$a=1,b=2,c=3$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.NotNil(t, err)
}

func Test_ComparePassword_returns_error_if_salt_empty(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$m=32768,t=2,p=6$$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.ErrorIs(t, err, ErrInvalidHashFormat)
}

func Test_ComparePassword_returns_error_if_invalid_hash_salt(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$m=32768,t=2,p=6$/*987$nwoK/4VQnDVgMT5nQvsfwtQJm/YVY/yPWnWzT7BKdR7JOoadEiodLyayWwGyyvWStDfNBFTDA/CL9FcTCX6jJQ",
	)
	require.NotNil(t, err)
}

func Test_ComparePassword_returns_error_if_password_hash_empty(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$m=32768,t=2,p=6$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$",
	)
	require.ErrorIs(t, err, ErrInvalidHashFormat)
}

func Test_ComparePassword_returns_error_if_invalid_password_hash(t *testing.T) {
	_, err := ComparePassword(
		"password",
		"$argon2id$v=19$m=32768,t=2,p=6$gor4v52ESeWIAflx1mg+Z1XPC31aM0r6SOICESoepmQ$/*987",
	)
	require.NotNil(t, err)
}
