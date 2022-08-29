package mediators

import (
	"amaterasu/src/domain"
	"errors"
)

var (
	RoomNotExistsInThisManager = errors.New("room not exists in manager")
	RoomAlreadyInThisManager   = errors.New("room not exists in manager")
)

type RoomMediator struct {
	rooms map[domain.RoomName]*domain.Room
}

func NewRoomMediator() *RoomMediator {
	return &RoomMediator{
		rooms: make(map[domain.RoomName]*domain.Room),
	}
}

func (r *RoomMediator) RegisterNewRoom(room *domain.Room) error {
	if r.RoomExists(room.GetName()) {
		return RoomAlreadyInThisManager
	}
	r.rooms[room.GetName()] = room
	return nil
}

func (r *RoomMediator) AddRoomListener(n domain.RoomName, l domain.Client) error {
	if !r.RoomExists(n) {
		return RoomNotExistsInThisManager
	}
	room := r.rooms[n]
	err := room.AddListener(l)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomMediator) RemoveRoomListener(n domain.RoomName, l domain.Client) error {
	if !r.RoomExists(n) {
		return RoomNotExistsInThisManager
	}
	room := r.rooms[n]
	err := room.RemoveListener(l)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomMediator) SendBroadcastMessage(n domain.RoomName, s domain.Client, message string) error {
	if !r.RoomExists(n) {
		return RoomNotExistsInThisManager
	}
	room := r.rooms[n]
	room.SendBroadcastMessage(s, message)
	return nil
}

func (r *RoomMediator) RoomExists(n domain.RoomName) bool {
	_, exists := r.rooms[n]
	return exists
}
