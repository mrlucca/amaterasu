package domain_test

import (
	"amaterasu/src/domain"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestRoom_ErrorWhileAddEqualsListeners(t *testing.T) {
	c := new(net.Conn)
	client := domain.NewClient("test", c)
	room := domain.CreateRoom(
		"test",
		*client,
	)
	err := room.AddListener(*client)
	if !assert.Nil(t, err) {
		t.Error("first add is not safe error is: " + err.Error())
	}
	err = room.AddListener(*client)
	if assert.Error(t, err) {
		assert.Equal(t, domain.ListenerAlreadyExists, err)
	} else {
		t.Error("no error handled")
	}
}

func TestRoom_GetSizeWithThreeDifferentClients(t *testing.T) {
	c := new(net.Conn)
	client := domain.NewClient("test", c)
	client2 := domain.NewClient("test2", c)
	client3 := domain.NewClient("test3", c)
	room := domain.CreateRoom(
		"test",
		*client,
	)
	room.AddListener(*client)
	room.AddListener(*client2)
	room.AddListener(*client3)
	assert.Equal(t, 3, room.GetSize(), "room size is equals")
}

func TestRoom_GetSizeWithThreeEqualsClients(t *testing.T) {
	c := new(net.Conn)
	client := domain.NewClient("test", c)
	room := domain.CreateRoom(
		"test",
		*client,
	)
	room.AddListener(*client)
	room.AddListener(*client)
	room.AddListener(*client)
	assert.Equal(t, 1, room.GetSize(), "room size is equals")
}
