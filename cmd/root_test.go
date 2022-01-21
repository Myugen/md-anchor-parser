package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestExecute(t *testing.T) {
	assertt := assert.New(t)
	mockedReader := new(MyMockedObject)
	mockedReader.On("Read").Return()

	cmd := NewRootCmd(mockedReader)
	err := cmd.Execute()

	assertt.Equal(nil, err)
	mockedReader.AssertCalled(t, "Read")
}

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) Read() {
	m.Called()
}
