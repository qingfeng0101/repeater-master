package handler

import "encoding/json"

type Payload struct {
	From        string           `json:"from"`
	To          []string         `json:"to" validate:"required"`
	Subject     string           `json:"subject"`
	Content     *json.RawMessage `json:"content" validate:"required"`
	ContentType string           `json:"content_type"`
}

type OldPayload struct {
	From        string `json:"from" form:"from"`
	To          string `json:"to" form:"to" validate:"required"`
	Subject     string `json:"subject" form:"subject"`
	Content     string `json:"content" form:"content" validate:"required"`
	ContentType string `json:"content_type" form:"content_type"`
}
