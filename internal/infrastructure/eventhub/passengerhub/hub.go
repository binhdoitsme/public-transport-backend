package passengerhub

import (
	"encoding/json"
	"log"
	"public-transport-backend/internal/common/slices"
	"public-transport-backend/internal/infrastructure/eventhub/eventhub"

	"github.com/google/uuid"
)

const topic string = "passengers"

type handlerFunc func(message interface{})

type PassengerEventHub struct {
	hub            *eventhub.EventHub
	handlersByType map[PassengerEventType][]string
	handlers       map[string]handlerFunc
}

func New(hub *eventhub.EventHub) *PassengerEventHub {
	passengerHub := &PassengerEventHub{
		hub:            hub,
		handlers:       make(map[string]handlerFunc),
		handlersByType: make(map[PassengerEventType][]string),
	}
	passengerHub.hub.Subscribe(topic, passengerHub.HandleEvent)
	return passengerHub
}

func (passengerHub *PassengerEventHub) handlersOfType(t PassengerEventType) []handlerFunc {
	result := make([]handlerFunc, 0)
	for _, handler := range passengerHub.handlersByType[t] {
		result = append(result, passengerHub.handlers[handler])
	}
	return result
}

func (passengerHub *PassengerEventHub) HandleEvent(message string) {
	messageData := &PassengerEvent{}
	err := json.Unmarshal([]byte(message), messageData)
	if err != nil {
		log.Println("error", err)
	}
	for _, handle := range passengerHub.handlersOfType(messageData.Type) {
		go handle(messageData)
	}
}

func (passengerHub *PassengerEventHub) Subscribe(t PassengerEventType, f handlerFunc) string {
	id := uuid.New().String()
	if _, exists := passengerHub.handlersByType[t]; !exists {
		passengerHub.handlersByType[t] = make([]string, 0)
	}
	passengerHub.handlers[id] = f
	passengerHub.handlersByType[t] = append(passengerHub.handlersByType[t], id)
	return id
}

func (passengerHub *PassengerEventHub) Unsubscribe(id string) {
	for t, ids := range passengerHub.handlersByType {
		passengerHub.handlersByType[t] = slices.RemoveValue(ids, id)
	}
}

func (passengerHub *PassengerEventHub) RequestApproval(id uint64) error {
	message := &PassengerEvent{
		Type: PassengerCreated,
		Data: struct {
			Id uint64 `json:"id"`
		}{id},
	}
	m, err := json.Marshal(message)
	if err != nil {
		return err
	}
	passengerHub.hub.Publish(topic, string(m))
	return nil
}
