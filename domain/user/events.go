package domain

import (
	"github.com/jperdior/recordari-kit/application/event"
)

//TODO: find a way to have this in a shared package

const UserRegisteredType event.Type = "user_registered"

type UserRegisteredEvent struct {
	event.BaseEvent
	email string
	roles []string
}

func (e UserRegisteredEvent) Email() string {
	return e.email
}

func (e UserRegisteredEvent) Roles() []string {
	return e.roles
}

func (e UserRegisteredEvent) Type() event.Type {
	return UserRegisteredType
}

func (e UserRegisteredEvent) ToDTO() event.EventDTO {
	return UserRegisteredEventDto{
		BaseEventDTO: e.BaseEvent.ToDTO(),
		Email:        e.email,
		Roles:        e.roles,
	}
}

type UserRegisteredEventDto struct {
	event.BaseEventDTO
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

func (e UserRegisteredEventDto) ToEvent() UserRegisteredEvent {
	return UserRegisteredEvent{
		BaseEvent: e.BaseEventDTO.ToBaseEvent(),
		email:     e.Email,
		roles:     e.Roles,
	}
}
