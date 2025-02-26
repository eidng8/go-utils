package testing

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-utils"
)

type MockUuidConfig struct {
	GetPanics                   bool
	NewReturnsError             bool
	MarshalBinaryReturnsError   bool
	MarshalTextReturnsError     bool
	UnmarshalBinaryReturnsError bool
	UnmarshalTextReturnsError   bool
	VersionPanics               bool
}

type MockUuid struct {
	id  utils.UUID
	cfg MockUuidConfig
}

func NewUuidMock(cfg MockUuidConfig) MockUuid {
	m := MockUuid{}
	m.id = &utils.Uuid{}
	m.cfg = cfg
	return m
}

func (m *MockUuid) Get() uuid.UUID {
	if m.cfg.GetPanics {
		panic("Get error")
	}
	return m.id.Get()
}

func (m *MockUuid) New() error {
	if m.cfg.NewReturnsError {
		return assert.AnError
	}
	return m.id.New()
}

func (m *MockUuid) MarshalBinary() ([]byte, error) {
	if m.cfg.MarshalBinaryReturnsError {
		return nil, assert.AnError
	}
	return m.id.MarshalBinary()
}

func (m *MockUuid) MarshalText() ([]byte, error) {
	if m.cfg.MarshalTextReturnsError {
		return nil, assert.AnError
	}
	return m.id.MarshalText()
}

func (m *MockUuid) UnmarshalBinary(data []byte) error {
	if m.cfg.UnmarshalBinaryReturnsError {
		return assert.AnError
	}
	return m.id.UnmarshalBinary(data)
}

func (m *MockUuid) UnmarshalText(data []byte) error {
	if m.cfg.UnmarshalTextReturnsError {
		return assert.AnError
	}
	return m.id.UnmarshalText(data)
}

func (m *MockUuid) Version() uuid.Version {
	if m.cfg.VersionPanics {
		panic("Version error")
	}
	return m.id.Version()
}

var _ utils.UUID = &MockUuid{}
