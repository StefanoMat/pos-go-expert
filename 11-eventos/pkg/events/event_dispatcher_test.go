package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (t *TestEvent) GetName() string {
	return t.Name
}

func (t *TestEvent) GetPayload() interface{} {
	return t.Payload
}

func (t *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test", Payload: "payload"}
	suite.event2 = TestEvent{Name: "test2", Payload: "payload2"}
}
func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handler, s.dispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.dispatcher.handlers[s.event.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	assert.True(s.T(), s.dispatcher.Has(s.event.GetName(), &s.handler))
	assert.False(s.T(), s.dispatcher.Has(s.event2.GetName(), &s.handler))
}
func (s *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := s.dispatcher.Register(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	err = s.dispatcher.Register(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.dispatcher.handlers[s.event.GetName()]))

	err = s.dispatcher.Remove(s.event.GetName(), &s.handler)
	s.Nil(err)
	s.Equal(1, len(s.dispatcher.handlers[s.event.GetName()]))
	err = s.dispatcher.Remove(s.event.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(0, len(s.dispatcher.handlers[s.event.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &s.event)

	eventHandler2 := &MockHandler{}
	eventHandler2.On("Handle", &s.event)
	s.dispatcher.Register(s.event.GetName(), eventHandler)
	s.dispatcher.Register(s.event.GetName(), eventHandler2)

	s.dispatcher.Dispatch(&s.event)
	eventHandler.AssertExpectations(s.T())
	eventHandler2.AssertExpectations(s.T())
	eventHandler.AssertNumberOfCalls(s.T(), "Handle", 1)
	eventHandler2.AssertNumberOfCalls(s.T(), "Handle", 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
