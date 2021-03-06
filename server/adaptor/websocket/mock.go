package websocket

import (
	"github.com/stretchr/testify/mock"
)

// Mock mocks a websocket adaptor
type Mock struct {
	mock.Mock

	Token     string
	ReadChan  chan interface{}
	WriteChan chan interface{}
}

// NewMock create a new websocket adaptor mock, with helper read/write
// chans already created
func NewMock(accessToken string, writeChanBuffer uint) *Mock {
	return &Mock{
		Token:     accessToken,
		ReadChan:  make(chan interface{}),
		WriteChan: make(chan interface{}, writeChanBuffer),
	}
}

// AccessToken mocks token getter
func (m *Mock) AccessToken() string {
	args := m.Called()
	return args.String(0)
}

// OnAccessToken is a helper method to setup an "AccessToken()" handler
// with the mock accessToken
func (m *Mock) OnAccessToken() *mock.Call {
	return m.On("AccessToken").Return(m.Token)
}

// Read mocks read channel getter
func (m *Mock) Read() <-chan interface{} {
	args := m.Called()
	return args.Get(0).(<-chan interface{})
}

// OnRead is a helper method to setup a "Read()" handler
// with the mock readChan
func (m *Mock) OnRead() *mock.Call {
	var readOnlyChan <-chan interface{} = m.ReadChan
	return m.On("Read").Return(readOnlyChan)
}

// Write mocks write channel getter
func (m *Mock) Write() chan<- interface{} {
	args := m.Called()
	return args.Get(0).(chan<- interface{})
}

// OnWrite is a helper method to setup a "Write()" handler
// with the mock writeChan
func (m *Mock) OnWrite() *mock.Call {
	var writeOnlyChan chan<- interface{} = m.WriteChan
	return m.On("Write").Return(writeOnlyChan)
}
