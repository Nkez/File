package models

import "github.com/google/uuid"

type Request struct {
	Method  string `json:"method"`
	URL     string `json:"url"`
	Headers Header `json:"headers"`
}

type Response struct {
	ID      uuid.UUID `json:"id"`
	Status  string    `json:"status"`
	Headers Header    `json:"headers"`
}

type Header struct {
	UserAgent   string `json:"user_agent"`
	ContentType string `json:"content_type"`
}

type ReqResMap struct {
	ID          uuid.UUID `json:"id"`
	ResponseMap Response  `json:"res"`
	RequestMap  Request   `json:"req"`
}

type Status struct {
	Answer string `json:"answer"`
}
