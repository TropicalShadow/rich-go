package ipc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

var Socket net.Conn

func CloseSocket() error {
	if Socket != nil {
		_ = Socket.Close()
		Socket = nil
	}
	return nil
}

func parseBaseResponse(rawResponse []byte) (*BaseResponse, error) {
	var baseResp BaseResponse
	err := json.Unmarshal(rawResponse, &baseResp)
	if err != nil {
		fmt.Println(string(rawResponse))
		return nil, fmt.Errorf("failed to parse base response: %w", err)
	}
	return &baseResp, nil
}

func ParseResponseData[T any](rawResponse []byte) (*T, error) {
	baseResp, err := parseBaseResponse(rawResponse)
	if err != nil {
		return nil, err
	}

	var result T

	switch baseResp.CMD {
	case "DISPATCH":
		var handshakeData HandShakeDataResponse
		if err := json.Unmarshal(baseResp.Data, &handshakeData); err != nil {
			return nil, fmt.Errorf("failed to parse handshake data: %w", err)
		}
		result = any(handshakeData).(T)
	case "SET_ACTIVITY":
		var activityData ResponseActivity
		if err := json.Unmarshal(baseResp.Data, &activityData); err != nil {
			return nil, fmt.Errorf("failed to parse activity data: %w", err)
		}
		result = any(activityData).(T)
	case "ERROR":
		var errorData ErrorDataResponse
		if err := json.Unmarshal(baseResp.Data, &errorData); err != nil {
			return nil, fmt.Errorf("failed to parse error data: %w", err)
		}
		result = any(errorData).(T)
	default:
		result = any(baseResp.Data).(T)
	}

	return &result, nil
}

func Read() []byte {
	var buffer bytes.Buffer
	buf := make([]byte, 1024) // Consider making this dynamic? SET_ACTIVITY can return a wide range of sizes

	for {
		payloadLength, err := Socket.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		if payloadLength > 8 {
			buffer.Write(buf[8:payloadLength])
		}

		if payloadLength < len(buf) {
			break
		}
	}

	return buffer.Bytes()
}

// Send opcode and payload to the unix Socket
func Send(opcode int, payload string) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(len(payload)))
	if err != nil {
		return nil, err
	}

	buf.Write([]byte(payload))
	_, err = Socket.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return Read(), nil
}
