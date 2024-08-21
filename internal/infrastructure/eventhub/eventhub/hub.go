package eventhub

import (
	"github.com/google/uuid"
)

type EventHub struct {
	topics        map[string]chan string
	topicHandlers map[string][]string
	handlers      map[string]func(string)
	newTopics     chan string
	isActive      bool
}

func New() *EventHub {
	return &EventHub{
		topics:        make(map[string]chan string),
		topicHandlers: make(map[string][]string),
		handlers:      make(map[string]func(string)),
		newTopics:     make(chan string),
	}
}

func (hub *EventHub) Topics() []string {
	topics := make([]string, 0)
	for t := range hub.topics {
		topics = append(topics, t)
	}
	return topics
}

func (hub *EventHub) AddTopic(name string) {
	if _, exists := hub.topics[name]; exists {
		return
	}

	hub.topics[name] = make(chan string)
	go func() { hub.newTopics <- name }()
}

func (hub *EventHub) Subscribe(topic string, handler func(string)) (string, error) {
	hub.AddTopic(topic)
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if _, exists := hub.topicHandlers[topic]; !exists {
		hub.topicHandlers[topic] = make([]string, 0)
	}
	idString := id.String()
	hub.topicHandlers[topic] = append(hub.topicHandlers[topic], idString)
	hub.handlers[idString] = handler

	return idString, nil
}

// func (hub *EventHub) Unsubscribe(topic string, id string) {

// }

func (hub *EventHub) Publish(topic string, message string) {
	channel, exists := hub.topics[topic]
	if !hub.isActive || !exists {
		return
	}

	channel <- message
}

func (hub *EventHub) Start() {
	hub.isActive = true
	go func() {
		for {
			topic := <-hub.newTopics
			if topic == "<terminate>" {
				break
			}
			channel := hub.topics[topic]
			go func() {
				for {
					msg := <-channel
					if msg == "<terminate>" {
						break
					}
					handlers := hub.topicHandlers[topic]
					for _, h := range handlers {
						go hub.handlers[h](msg)
					}
				}
			}()
		}
	}()
}

func (hub *EventHub) Stop() {
	hub.isActive = false
	for _, channel := range hub.topics {
		channel <- "<terminate>"
	}
	hub.newTopics <- "<terminate>"
}
