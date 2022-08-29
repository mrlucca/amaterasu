package domain

const (
	LISTEN = "LISTEN"
	LOGIN  = "LOGIN"
	EXIT   = "EXIT"
	SEND   = "SEND"
	CREATE = "CREATE"
	LIST   = "LIST"
	CLOSE  = "CLOSE"
)

const (
	SendMessageInRoom = "broadcast"
	SendDirectMessage = "direct"
)

func LoginClient(client Client, args []string) NickName {
	return ""
}

func ListenClientInRoom(client Client, args []string) {

}

func ExitClientInRoom(client Client, args []string) {

}

func SendMessageFromClient(client Client, args []string) {
	messageType := args[0]
	switch messageType {

	}
}

func CreateRoomFromClient(client Client, args []string) {

}
