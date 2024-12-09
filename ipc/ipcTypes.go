package ipc

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	CMD   string          `json:"cmd"`
	EVT   string          `json:"evt"`
	Data  json.RawMessage `json:"data"` // Raw JSON to allow further processing
	NONCE string          `json:"nonce"`
}

type HandShakeResponse struct {
	CMD   string                `json:"cmd"`
	Data  HandShakeDataResponse `json:"data"`
	EVT   string                `json:"evt"`
	NONCE string                `json:"nonce"`
}

type ErrorDataResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorDataResponse) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

type HandShakeDataResponse struct {
	User   *User   `json:"user"`
	Config *Config `json:"config"`
	V      int     `json:"v"`
}

type Config struct {
	CDN_HOST    string `json:"cdn_host"`
	API_HOST    string `json:"api_endpoint"`
	ENVIRONMENT string `json:"environment"`
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
	Bot           bool   `json:"bot"`
}

type ResponseActivity struct {
	Details       string             `json:"details,omitempty"`
	State         string             `json:"state,omitempty"`
	Assets        PayloadAssets      `json:"assets,omitempty"`
	Party         *PayloadParty      `json:"party,omitempty"`
	Timestamps    *PayloadTimestamps `json:"timestamps,omitempty"`
	Secrets       *PayloadSecrets    `json:"secrets,omitempty"`
	Buttons       *[]string          `json:"buttons,omitempty"`
	Name          string             `json:"name,omitempty"`
	Type          int                `json:"type,omitempty"`
	ApplicationID string             `json:"application_id,omitempty"`
}

type PayloadActivity struct {
	Details    string             `json:"details,omitempty"`
	State      string             `json:"state,omitempty"`
	Assets     PayloadAssets      `json:"assets,omitempty"`
	Party      *PayloadParty      `json:"party,omitempty"`
	Timestamps *PayloadTimestamps `json:"timestamps,omitempty"`
	Secrets    *PayloadSecrets    `json:"secrets,omitempty"`
	Buttons    []*PayloadButton   `json:"buttons,omitempty"`
}

type PayloadAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type PayloadParty struct {
	ID   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type PayloadTimestamps struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type PayloadSecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

type PayloadButton struct {
	Label string `json:"label,omitempty"`
	Url   string `json:"url,omitempty"`
}
