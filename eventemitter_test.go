package eventemitter_test

import (
	"testing"

	"github.com/PatrickChagastavares/eventemitter"
	"github.com/stretchr/testify/require"
)

func TestListen(t *testing.T) {
	t.Cleanup(eventemitter.Reset)

	eventName := "TestListen"

	actionCalled := false
	action := func(data interface{}) {
		actionCalled = true
	}

	eventemitter.Listen(eventName, action)

	require.False(t, actionCalled)
	eventemitter.Emit("TestListen", nil)
	require.True(t, actionCalled)
}

func TestMultipleEvents(t *testing.T) {
	t.Cleanup(eventemitter.Reset)

	eventNameA := "TestMultipleEventsA"

	actionCalled1 := false
	action1 := func(data interface{}) {
		actionCalled1 = true
	}

	eventNameB := "TestMultipleEventsB"

	actionCalled2 := false
	action2 := func(data interface{}) {
		actionCalled2 = true
	}

	eventemitter.Listen(eventNameA, action1)
	eventemitter.Listen(eventNameB, action2)

	require.False(t, actionCalled1)
	require.False(t, actionCalled2)

	eventemitter.Emit("TestMultipleEventsA", nil)
	require.True(t, actionCalled1)
	require.False(t, actionCalled2)

	eventemitter.Emit("TestMultipleEventsB", nil)
	require.True(t, actionCalled1)
	require.True(t, actionCalled2)
}

func TestMultipleListenersToSameEvent(t *testing.T) {
	t.Cleanup(eventemitter.Reset)

	eventNameA := "TestMultipleListenersToSameEvent"

	actionCalled1 := false
	action1 := func(data interface{}) {
		actionCalled1 = true
	}

	actionCalled2 := false
	action2 := func(data interface{}) {
		actionCalled2 = true
	}

	eventemitter.Listen(eventNameA, action1)
	eventemitter.Listen(eventNameA, action2)

	require.False(t, actionCalled1)
	require.False(t, actionCalled2)

	eventemitter.Emit("TestMultipleListenersToSameEvent", nil)
	require.True(t, actionCalled1)
	require.True(t, actionCalled2)
}

func TestEmitWithData(t *testing.T) {
	t.Cleanup(eventemitter.Reset)
	eventName := "TestEmitWithData"

	actionCalled := false
	var actionData interface{}
	action := func(data interface{}) {
		actionCalled = true
		actionData = data
	}

	eventemitter.Listen(eventName, action)

	require.False(t, actionCalled)
	eventemitter.Emit("TestEmitWithData", "oi")
	require.True(t, actionCalled)
	require.Equal(t, "oi", actionData)
}

func TestRemoveEvent(t *testing.T) {
	t.Cleanup(eventemitter.Reset)
	eventName := "TestRemoveEvent"

	actionCalled := false
	action := func(data interface{}) {
		actionCalled = true
	}

	eventemitter.Listen(eventName, action)
	eventemitter.RemoveEvent(eventName)

	require.False(t, actionCalled)
	eventemitter.Emit("TestRemoveEvent", nil)
	require.False(t, actionCalled)
}

func TestKeepRunningOnPanic(t *testing.T) {
	t.Cleanup(eventemitter.Reset)
	eventName := "TestKeepRunningOnPanic"

	action1Called := false
	action1 := func(data interface{}) {
		action1Called = true
	}

	actionPanic := func(data interface{}) {
		panic("ahhh")
	}

	action2Called := false
	action2 := func(data interface{}) {
		action2Called = true
	}

	eventemitter.Listen(eventName, action1)
	eventemitter.Listen(eventName, actionPanic)
	eventemitter.Listen(eventName, action2)

	require.NotPanics(t, func() {
		eventemitter.Emit("TestKeepRunningOnPanic", nil)
	})

	require.True(t, action1Called)
	require.True(t, action2Called)
}

func TestListenOnce(t *testing.T) {
	t.Cleanup(eventemitter.Reset)

	eventName := "TestListenOnce"

	actionCalledTimes := 0
	action := func(data interface{}) {
		actionCalledTimes++
	}

	eventemitter.ListenOnce(eventName, action)

	require.Equal(t, 0, actionCalledTimes)

	eventemitter.Emit("TestListenOnce", nil)
	require.Equal(t, 1, actionCalledTimes)

	eventemitter.Emit("TestListenOnce", nil)
	require.Equal(t, 1, actionCalledTimes)
}

func TestReset(t *testing.T) {
	t.Cleanup(eventemitter.Reset)

	eventName1 := "TestReset1"

	action1Called := false
	action1 := func(data interface{}) {
		action1Called = true
	}

	eventName2 := "TestReset2"

	action2Called := false
	action2 := func(data interface{}) {
		action2Called = true
	}

	eventemitter.Listen(eventName1, action1)
	eventemitter.Listen(eventName2, action2)

	eventemitter.Reset()

	eventemitter.Emit("TestReset1", nil)
	eventemitter.Emit("TestReset2", nil)

	require.False(t, action1Called)
	require.False(t, action2Called)
}
