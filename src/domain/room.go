package domain

import "errors"

type RoomName string

var (
	ListenerAlreadyExists      = errors.New("listener already exists in this room")
	ListenerNotAddedInThisRoom = errors.New("listener not added in this room")
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
		go client.SendMessageFromRoom(sender, *r, msg)
	}
}

func (r *Room) AddListener(l Client) error {
	if r.ListenerExists(l) {
		return ListenerAlreadyExists
	}
	r.listeners[l.GetName()] = l
	return nil
}

func (r *Room) RemoveListener(l Client) error {
	if !r.ListenerExists(l) {
		return ListenerNotAddedInThisRoom
	}
	delete(r.listeners, l.GetName())
	return nil
}
func (r *Room) ListenerExists(l Client) bool {
	_, ok := r.listeners[l.GetName()]
	return ok
}

func (r *Room) GetListenSize() int {
	return len(r.listeners)
}

func (r *Room) GetName() RoomName {
	return r.name
}

func (r *Room) GetAuthor() Client {
	return r.author
}
