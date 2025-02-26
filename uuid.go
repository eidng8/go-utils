package utils

import "github.com/google/uuid"

var (
	uuidV7 = uuid.NewV7
	uuidV6 = uuid.NewV6
	uuidV4 = uuid.NewRandom

	_ UUID = &Uuid{}
)

type UUID interface {
	Get() uuid.UUID
	New() error
	MarshalBinary() ([]byte, error)
	MarshalText() ([]byte, error)
	UnmarshalBinary(data []byte) error
	UnmarshalText(data []byte) error
	Version() uuid.Version
}

type Uuid struct {
	id uuid.UUID
}

func (u *Uuid) Get() uuid.UUID {
	return u.id
}

func (u *Uuid) New() (err error) {
	u.id, err = NewUuid()
	return
}

func (u *Uuid) MarshalBinary() ([]byte, error) {
	return u.id.MarshalBinary()
}

func (u *Uuid) MarshalText() ([]byte, error) {
	return u.id.MarshalText()
}

func (u *Uuid) UnmarshalBinary(data []byte) error {
	return u.id.UnmarshalBinary(data)
}

func (u *Uuid) UnmarshalText(data []byte) error {
	return u.id.UnmarshalText(data)
}

func (u *Uuid) Version() uuid.Version {
	return u.id.Version()
}

// NewUuid returns a new UUID. It tries to generate a UUID using V7, V6, and
// finally falls back to V4. If all methods fail, it returns an error.
func NewUuid() (uuid.UUID, error) {
	id, err := uuidV7()
	if nil == err {
		return id, nil
	}
	id, err = uuidV6()
	if nil == err {
		return id, nil
	}
	id, err = uuidV4()
	if nil == err {
		return id, nil
	}
	return uuid.Nil, err
}
