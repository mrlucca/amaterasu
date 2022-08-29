package domain

import (
	"bufio"
	"net"
	"strings"
)

type NickName string

type Client struct {
	name NickName
	conn net.Conn
}

func NewClient(name NickName, conn net.Conn) *Client {
	return &Client{
		name: name,
		conn: conn,
	}
}

func (c *Client) OnMessage() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.ToUpper(strings.TrimSpace(args[0]))
		args = args[1:]
		switch cmd {
		case LOGIN:
			c.name = LoginClient(*c, args)
		case LISTEN:
			ListenClientInRoom(*c, args)
		case EXIT:
			ExitClientInRoom(*c, args)
		case SEND:
			SendMessageFromClient(*c, args)
		case CREATE:
			CreateRoomFromClient(*c, args)
		}
	}

}

func (c *Client) GetName() NickName {
	return c.name
}

func (c *Client) SendMessageFromRoom(sender Client, room Room, msg string) {

}

func (c *Client) SendMessageFromAnotherClient(sender Client, msg string) {

}
