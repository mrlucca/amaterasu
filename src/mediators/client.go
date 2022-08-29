package mediators

import (
	"amaterasu/src/domain"
	"errors"
)

var (
	ClientNotExistsInThisManager = errors.New("client not exists in this manager")
)

type ClientMediator struct {
	clients map[domain.NickName]domain.Client
}

func NewClientMediator() *ClientMediator {
	return &ClientMediator{
		clients: make(map[domain.NickName]domain.Client),
	}
}

func (m *ClientMediator) RegisterClient(client domain.Client) error {
	m.clients[client.GetName()] = client
	return nil
}

func (m *ClientMediator) RemoveClient(client domain.Client) error {
	delete(m.clients, client.GetName())
	return nil
}

func (m *ClientMediator) SendDirectMessage(s domain.Client, n domain.NickName, message string) error {
	client, clientExists := m.clients[n]
	if !clientExists {
		return ClientNotExistsInThisManager
	}
	client.SendMessageFromAnotherClient(s, message)
	return nil
}