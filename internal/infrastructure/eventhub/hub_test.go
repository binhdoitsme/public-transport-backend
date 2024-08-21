package eventhub_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"public-transport-backend/internal/infrastructure/eventhub" // Update with the correct import path
)

func TestEventHub_AddTopic(t *testing.T) {
	hub := eventhub.New()

	// Ensure that adding a topic works as expected.
	hub.AddTopic("topic1")
	assert.Contains(t, hub.Topics(), "topic1")
}

func TestEventHub_Subscribe(t *testing.T) {
	hub := eventhub.New()

	// Start the hub to handle new topics
	hub.Start()
	defer hub.Stop()

	// Subscribe to a topic
	handlerInvoked := false
	handler := func(msg string) {
		handlerInvoked = true
	}

	_, err := hub.Subscribe("topic1", handler)
	assert.NoError(t, err)

	// Publish a message to the topic
	hub.Publish("topic1", "test message")

	// Wait for a short duration to allow the handler to be invoked
	time.Sleep(100 * time.Millisecond)

	assert.True(t, handlerInvoked, "Handler should have been invoked")
}

func TestEventHub_Publish(t *testing.T) {
	hub := eventhub.New()

	// Start the hub to handle new topics
	hub.Start()
	defer hub.Stop()

	// Subscribe to a topic with a handler that captures the message
	var capturedMsg string
	handler := func(msg string) {
		capturedMsg = msg
	}

	_, err := hub.Subscribe("topic1", handler)
	assert.NoError(t, err)

	// Publish a message to the topic
	hub.Publish("topic1", "Hello, EventHub!")

	// Wait for the handler to be invoked
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, "Hello, EventHub!", capturedMsg, "Handler should have received the published message")
}

func TestEventHub_MultipleHandlers(t *testing.T) {
	hub := eventhub.New()

	// Start the hub to handle new topics
	hub.Start()
	defer hub.Stop()

	// Use a WaitGroup to synchronize handler invocations
	var wg sync.WaitGroup
	wg.Add(2)

	// Subscribe two handlers to the same topic
	handler1 := func(msg string) {
		defer wg.Done()
	}

	handler2 := func(msg string) {
		defer wg.Done()
	}

	_, err := hub.Subscribe("topic1", handler1)
	assert.NoError(t, err)
	_, err = hub.Subscribe("topic1", handler2)
	assert.NoError(t, err)

	// Publish a message to the topic
	hub.Publish("topic1", "test message")

	// Wait for both handlers to be invoked
	done := waitTimeout(&wg, time.Second)

	// Check that both handlers were invoked within the timeout
	assert.False(t, done, "Both handlers should have been invoked within the timeout")
}

func TestEventHub_Stop(t *testing.T) {
	hub := eventhub.New()

	// Start the hub to handle new topics
	hub.Start()

	// Subscribe a handler that will not be invoked after Stop is called
	handlerInvoked := false
	handler := func(msg string) {
		handlerInvoked = true
	}

	_, err := hub.Subscribe("topic1", handler)
	assert.NoError(t, err)

	// Stop the hub
	hub.Stop()

	// Publish a message after stopping the hub
	hub.Publish("topic1", "test message")

	// Wait for a short duration to ensure handler is not invoked
	time.Sleep(100 * time.Millisecond)

	assert.False(t, handlerInvoked, "Handler should not have been invoked after hub was stopped")
}

// waitTimeout waits for the wait group to finish or times out
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
