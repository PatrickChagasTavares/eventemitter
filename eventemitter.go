package eventemitter

import "fmt"

var (
	listEvents     map[string][]func(data interface{}) = make(map[string][]func(data interface{}))
	listEventsOnce map[string][]func(data interface{}) = make(map[string][]func(data interface{}))
)

func Listen(name string, f func(data interface{})) {
	listen(name, f, false)
}

func ListenOnce(name string, f func(data interface{})) {
	listen(name, f, true)
}

func Emit(name string, data interface{}) {
	event, ok := listEvents[name]
	if !ok {
		event, ok = listEventsOnce[name]
		if !ok {
			return
		}

		defer RemoveEvent(name)
	}

	for _, f := range event {
		run(f, data)
	}
}

func Reset() {
	listEvents = make(map[string][]func(data interface{}))
	listEventsOnce = make(map[string][]func(data interface{}))
}

func RemoveEvent(name string) {
	delete(listEvents, name)
	delete(listEventsOnce, name)
}

func listen(name string, f func(data interface{}), once bool) {
	list := listEvents
	if once {
		list = listEventsOnce
	}

	event, ok := list[name]

	if ok {
		event = append(event, f)
		list[name] = event
		return
	}
	list[name] = []func(data interface{}){f}
}

func run(f func(data interface{}), data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			return
		}
	}()

	f(data)
}
