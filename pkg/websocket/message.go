package websocket

const (
	TextMessage = 1
	PingMessage = 9
)

type Message struct {
	messageType    int
	messageContent string
}

func (message Message) Type() int {
	return message.messageType
}

func (message Message) Content() string {
	return message.messageContent
}
