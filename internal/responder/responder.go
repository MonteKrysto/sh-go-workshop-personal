package responder

import (
	"encoding/json"
	"net/http"
)

// Responder is responsible for returning uniform responses to the client

var (
	ResourceCreated       = "Resource posted successfully"
	ResourceFetched       = "Resource fetched successfully"
	ResourcesFetched      = "Resources fetched successfully"
	ResourceUpdated       = "Resource updated successfully"
	ResourceDeleted       = "Resource deleted successfully"
	ResourceNotFound      = "Resource not found"
	ResourceNotCreated    = "Resource not created"
	ResourceUnprocessable = "Resource not processable"
)

type ResponseBody struct {
	Code    int  `json:"status"`
	Success bool `json:"success"`
	//Message string      `json:"message"`
	Data interface{} `json:"data"`
}

type ErrorResponseBody struct {
	Code    int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var response ResponseBody
var errorResponse ErrorResponseBody

func JsonResponse(w http.ResponseWriter, code int, status bool, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response.Success = status
	//response.Message = message
	response.Data = data
	response.Code = code

	json.NewEncoder(w).Encode(response)
}

func JsonErrorResponse(w http.ResponseWriter, code int, status bool, message string) {
	w.Header().Set("Content-Type", "application/json")
	errorResponse.Success = status
	errorResponse.Message = message
	errorResponse.Data = nil
	errorResponse.Code = code

	json.NewEncoder(w).Encode(errorResponse)
}
