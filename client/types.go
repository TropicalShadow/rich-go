package client

import "github.com/hugolgst/rich-go/ipc"

type Handshake struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

type Args struct {
	Pid      int                  `json:"pid"`
	Activity *ipc.PayloadActivity `json:"activity"`
}
