package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Dbmock struct {
	mock.Mock
}

func (m *Dbmock) Delete(id string) error {
	result := m.Called(id)

	if result.Get(0) != nil {
		return result.Error(0)
	}

	return nil
}

func (m *Dbmock) Update(id string, value map[string]string) error {
	result := m.Called(id, value)

	if result.Get(0) != nil {
		return result.Error(0)
	}

	return nil
}

func TestUpdateProductFailed(t *testing.T) {
	dbmock := &Dbmock{}

	dbmock.On("Delete", "1").Return(errors.New("something"))
	pdservice := NewProductService(dbmock)

	err := pdservice.Delete("1")
	assert.Nil(t, err)
}
