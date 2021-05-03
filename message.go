package message

type ReplyMessage struct {
	MessageType string // "0"
	State       bool
	Err         string
}

type LoginMessage struct {
	MessageType string // "1"
	Username    string
	Password    string
}

type SendMessage struct {
	MessageType string // "2"
	Message     string
	Sendername  string
}

type FileNotifyMessage struct {
	MessageType string // "3"
	Sendername  string
	Filename    string
	Length      string
}

type PullRequestMessage struct {
	MessageType string // "4"
	Sendername  string
	Filename    string
	Offset      string
}

type DataMessage struct {
	MessageType string // "6"
	Sendername  string
	Filename    string
	Length      string
	Offset      string
	Data        []byte
}

type ChatRequest struct {
	MessageType string // "7"
	Mode        string
	ID          string
}

type LogoutRequest struct {
	MessageType string // "8"
}

//type FileData struct {
//	MessageType string //3:请求发送 4:同意接收 5:拒绝接收 6:发送数据
//	Sendername  string
//	Filename    string
//	Offset      string
//	Data        []byte
//}
