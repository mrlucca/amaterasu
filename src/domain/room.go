package domain

import "errors"

type RoomName string

var (
	ListenerAlreadyExists = errors.New("listener already exists in this room")
)

type Room struct {
	name      RoomName
	author    Client
	listeners map[NickName]Client
}

func CreateRoom(name RoomName, author Client) *Room {
	return &Room{
		name:      name,
		author:    author,
		listeners: make(map[NickName]Client),
	}
}

func (r *Room) SendBroadcastMessage(sender Client, msg string) {
	for _, client := range r.listeners {
		go client.SendMessageByRoom(sender, *r, msg)
	}
}

func (r *Room) AddListener(m Client) error {
	if _, ok := r.listeners[m.GetName()]; ok {
		return ListenerAlreadyExists
	}
	r.listeners[m.GetName()] = m
	return nil
}

func (r *Room) RemoveListener(m Client) {
	_, ok := r.listeners[m.GetName()]
	if !ok {
		return
	}
	delete(r.listeners, m.GetName())
}

func (r *Room) GetSize() int {
	return len(r.listeners)
}

func (r *Room) GetName() RoomName {
	return r.name
}
