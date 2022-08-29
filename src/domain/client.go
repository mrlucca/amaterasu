package domain

import "net"

type NickName string

type Client struct {
	name NickName
	conn *net.Conn
}

func NewClient(name NickName, conn *net.Conn) *Client {
	return &Client{
		name: name,
		conn: conn,
	}
}

func (c *Client) Onmessage() {

}

func (c *Client) GetName() NickName {
	return c.name
}

func (c *Client) SendMessageByRoom(sender Client, room Room, msg string) {

}
