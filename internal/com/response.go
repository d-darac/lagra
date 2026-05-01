package com

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    *ErrorCode `json:"code,omitempty"`
	Message string     `json:"message"`
	Param   *string    `json:"param,omitempty"`
	Type    ErrorType  `json:"type"`
}

type ErrorListResponse struct {
	Errors []ErrorResponse `json:"errors"`
}

type ListResponse struct {
	Url     string `json:"url"`
	Data    []any  `json:"data"`
	HasMore bool   `json:"has_more"`
}

func NewListResponse(url string) *ListResponse {
	return &ListResponse{
		Data: make([]any, 0),
		Url:  url,
	}
}

func (e ErrorResponse) Respond(w http.ResponseWriter, statCode int, err error) {
	if err != nil {
		log.Println(err)
	}
	if statCode >= http.StatusInternalServerError {
		log.Printf("responding with status %d", statCode)
	}
	RespondJSON(w, statCode, struct {
		Error ErrorResponse `json:"error"`
	}{Error: e})
}

func (e ErrorListResponse) Respond(w http.ResponseWriter, statCode int, err error) {
	if err != nil {
		log.Println(err)
	}
	if statCode >= http.StatusInternalServerError {
		log.Printf("responding with status %d", statCode)
	}
	RespondJSON(w, statCode, e)
}

func RespondError(w http.ResponseWriter, err error) {
	appErr := ErrorToAppError(err)
	RespondJSON(w, appErr.Status, struct {
		Error ErrorResponse `json:"error"`
	}{
		Error: ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Param:   appErr.Param,
			Type:    appErr.Type,
		},
	})
}

func RespondErrorList(w http.ResponseWriter, errs []*AppError) {
	errListRes := ErrorListResponse{}
	for _, err := range errs {
		errListRes.Errors = append(errListRes.Errors, ErrorResponse{
			Code:    err.Code,
			Message: err.Message,
			Param:   err.Param,
			Type:    err.Type,
		})

	}
	RespondJSON(w, http.StatusBadRequest, errListRes)
}

func RespondJSON(w http.ResponseWriter, statCode int, payload any) {
	setDefaultHeaders(w)
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		RespondError(w, err)
		return
	}
	w.WriteHeader(statCode)
	if statCode != http.StatusNoContent {
		w.Write([]byte(data)) // #nosec G104
	}
}

func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
