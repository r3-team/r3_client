package types

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

// generic request types
type Request struct {
	Ressource string          `json:"ressource"`
	Action    string          `json:"action"`
	Payload   json.RawMessage `json:"payload"`
}
type RequestTransaction struct {
	TransactionNr uint64    `json:"transactionNr"`
	Requests      []Request `json:"requests"`
}
type Response struct {
	Payload json.RawMessage `json:"payload"`
}
type ResponseTransaction struct {
	TransactionNr uint64     `json:"transactionNr"`
	Responses     []Response `json:"responses"`
	Error         string     `json:"error"`
}

// unrequested messages
type UnreqResponse struct {
	Ressource string          `json:"ressource"`
	Result    string          `json:"result"`
	Payload   json.RawMessage `json:"payload"`
}
type UnreqResponseTransaction struct {
	TransactionNr uint64          `json:"transactionNr"`
	Responses     []UnreqResponse `json:"responses"`
	Error         string          `json:"error"`
}

// payloads
type UnreqResponsePayloadFileRequested struct {
	AttributeId uuid.UUID `json:"attributeId"`
	ChooseApp   bool      `json:"chooseApp"`
	FileId      uuid.UUID `json:"fileId"`
	FileHash    string    `json:"fileHash"`
	FileName    string    `json:"fileName"`
}
type UnreqResponsePayloadKeystrokesDo struct {
	Strokes string `json:"strokes"`
}
type RequestPayloadLogin struct {
	LoginId    int64  `json:"loginId"`
	TokenFixed string `json:"tokenFixed"`
}
type RequestPayloadClientEventExec struct {
	Id        uuid.UUID     `json:"id"`
	Arguments []interface{} `json:"arguments"`
}
type ResponsePayloadLogin struct {
	LanguageCode string `json:"languageCode"`
	Token        string `json:"token"`
}
