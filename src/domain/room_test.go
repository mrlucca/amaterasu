package domain_test

import (
	"amaterasu/src/domain"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func simpleClientFactoryFromNickName(name domain.NickName) domain.Client {
	c := new(net.Conn)
	client := domain.NewClient(name, *c)
	return *client
}

func simpleRoomFactoryFromClient(client domain.Client) *domain.Room {
	room := domain.CreateRoom(
		"test",
		client,
	)
	return room
}

func TestRoom_ErrorWhileAddEqualsListeners(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	err := room.AddListener(client)
	if !assert.Nil(t, err) {
		t.Error("first add is not safe error is: " + err.Error())
	}
	err = room.AddListener(client)
	if assert.Error(t, err) {
		assert.Equal(t, domain.ListenerAlreadyExists, err)
	} else {
		t.Error("no error handled")
	}
}

func TestRoom_GetSizeWithThreeDifferentClients(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	client2 := simpleClientFactoryFromNickName("test2")
	client3 := simpleClientFactoryFromNickName("test3")
	room := simpleRoomFactoryFromClient(client)
	room.AddListener(client)
	room.AddListener(client2)
	room.AddListener(client3)
	assert.Equal(t, 3, room.GetListenSize(), "room size is equals")
}

func TestRoom_GetSizeWithThreeEqualsClients(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	room.AddListener(client)
	room.AddListener(client)
	room.AddListener(client)
	assert.Equal(t, 1, room.GetListenSize(), "room size is equals")
}

func TestRoom_RemoveListenerWithInvalidClient(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	err := room.RemoveListener(client)
	if assert.Error(t, err) {
		assert.Equal(t, domain.ListenerNotAddedInThisRoom, err)
	}
}

func TestRoom_RemoveListenerWithValidClient(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	room.AddListener(client)
	assert.Equal(t, true, room.ListenerExists(client), "check if client listen exists")
	err := room.RemoveListener(client)
	assert.Nil(t, err)
}

func TestRoom_ListenerExistsWithoutListeners(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	assert.Equal(t, false, room.ListenerExists(client), "check if client listen exists")
}

func TestRoom_ListenerExistsWithListeners(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	anotherTest := simpleClientFactoryFromNickName("another_test")
	room := simpleRoomFactoryFromClient(client)
	room.AddListener(client)
	room.AddListener(anotherTest)
	assert.Equal(t, true, room.ListenerExists(client), "check if client listen exists")
}

func TestRoom_ListenerExistsWithRemovedListener(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	anotherTest := simpleClientFactoryFromNickName("another_test")
	room := simpleRoomFactoryFromClient(client)
	room.AddListener(client)
	room.AddListener(anotherTest)
	room.RemoveListener(client)
	assert.Equal(t, false, room.ListenerExists(client), "check if client listen exists")
}

func TestRoom_GetName(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	assert.Equal(t, domain.RoomName("test"), room.GetName(), "check if room name is correct")
}

func TestRoom_GetAuthor(t *testing.T) {
	client := simpleClientFactoryFromNickName("test")
	room := simpleRoomFactoryFromClient(client)
	assert.Equal(t, client, room.GetAuthor(), "check if room author is correct")
}
