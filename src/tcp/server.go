package tcp

import (
	"amaterasu/src/domain"
	"amaterasu/src/mediators"
	"log"
	"net"
)

type Server struct {
	clientMediator *mediators.ClientMediator
	roomMediator   *mediators.RoomMediator
}

func NewServer() *Server {
	clientMediator := mediators.NewClientMediator()
	roomMediator := mediators.NewRoomMediator()
	return &Server{
		clientMediator: clientMediator,
		roomMediator:   roomMediator,
	}
}

func (s Server) Run() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		c := domain.NewClient(conn, s.clientMediator, s.roomMediator)
		go c.OnMessageHandler()
	}
}
