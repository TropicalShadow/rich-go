package client

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tropicalshadow/rich-go/ipc"
	"os"
	"strings"
)

type Client struct {
	Logged bool `json:"-"`

	HandshakeData   *ipc.HandShakeDataResponse `json:"handshake_config"`
	CurrentActivity *Activity                  `json:"current_activity"`
}

func NewClient() *Client {
	return &Client{Logged: false}
}

func (c *Client) User() *ipc.User {
	return c.HandshakeData.User
}

// Login connects to Discord IPC
func (c *Client) Login(clientId string) error {
	return c.LoginWithPipe(clientId, "0")
}

// LoginWithPipe connects to Discord IPC with a specific pipe id
func (c *Client) LoginWithPipe(clientId string, pipe string) error {
	if !c.Logged {
		payload, err := json.Marshal(Handshake{"1", clientId})
		if err != nil {
			return err
		}

		err = ipc.OpenSocket(pipe)
		if err != nil {
			return err
		}

		response, err := ipc.Send(0, string(payload))
		if err != nil {
			return err
		}

		if strings.Contains(string(response), "Invalid Client ID") {
			return errors.New("invalid client id")
		}

		handshake, err := ipc.ParseResponseData[ipc.HandShakeDataResponse](response)
		if err != nil {
			return err
		}

		if handshake == nil {
			return errors.New("handshake is nil")
		}

		c.HandshakeData = handshake
	}
	c.Logged = true

	return nil
}

func (c *Client) Logout() error {
	c.HandshakeData = nil
	c.Logged = false

	return ipc.CloseSocket()
}

// IsLogged returns whether the client is logged in
func (c *Client) IsLogged() bool {
	return c.Logged && ipc.Socket != nil
}

func (c *Client) ClearActivity() (*Activity, error) {
	c.CurrentActivity = nil
	if !c.Logged {
		return nil, errors.New("client is not logged in")
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			nil,
		},
		getNonce(),
	})

	if err != nil {
		return nil, err
	}

	data, err := ipc.Send(1, string(payload))
	if err != nil {
		_ = c.Logout()
		return nil, err
	}

	parsedResponse, err := ipc.ParseResponseData[ipc.ResponseActivity](data)
	if err != nil {
		return nil, err
	}

	if parsedResponse == nil {
		return nil, errors.New("parsedResponse is nil")
	}

	return fromPayload(parsedResponse), nil
}

func (c *Client) SetActivity(activity Activity) (*Activity, error) {
	if !c.Logged {
		return nil, errors.New("client is not logged in")
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			activity.toPayload(),
		},
		getNonce(),
	})

	if err != nil {
		return nil, err
	}

	data, err := ipc.Send(1, string(payload))
	if err != nil {
		_ = c.Logout()
		return nil, err
	}

	parsedResponse, err := ipc.ParseResponseData[ipc.ResponseActivity](data)
	if err != nil {
		return nil, err
	}

	if parsedResponse == nil {
		return nil, errors.New("parsedResponse is nil")
	}

	c.CurrentActivity = fromPayload(parsedResponse)

	return c.CurrentActivity, nil
}

func getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
