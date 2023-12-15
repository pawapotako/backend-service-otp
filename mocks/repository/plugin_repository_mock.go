package mocks

import (
	"github.com/stretchr/testify/mock"
)

type pluginRepositoryMock struct {
	mock.Mock
}

func NewPluginRepositoryMock() *pluginRepositoryMock {
	return &pluginRepositoryMock{}
}

func (m *pluginRepositoryMock) EventProducer(topic string, event any) {
	m.Called(topic, event)
}
