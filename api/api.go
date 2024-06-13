package api

import (
	"encoding/json"
	"net/http"
)

//	coin balance params
type CoinBalanceParams struct{
	Username string
}

// coin balance response
type CoinBalanceResponse struct{

	// success code
	Code int

	// account balance
	Balance int64
}

//	error response
type Error struct{
	
	// error code
	Code int

	// error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-type", "application/json")	// set the header type, meaning we want to return a json
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)	//converts the response(which is an error) into a json
}

//	simply just 2 wrapper functions of the private writeError() function. The error codes come from http package
var (
	//	specific error message from error type native to GO, not our custom error struct
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	//	generic error message
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An unexpected error has occured", http.StatusInternalServerError)
	}
)


