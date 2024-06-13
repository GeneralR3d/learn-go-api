package handlers

import (
	"encoding/json"
	"net/http"

	"learn-go-api/api"
	"learn-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"

)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {

	var params = api.CoinBalanceParams{}	//	decode the params using the custom method we made 

	var decoder *schema.Decoder = schema.NewDecoder()

	var err error // declaring beforehand

	err = decoder.Decode(&params, r.URL.Query())	//	decode fields and put them into the request struct

	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	//	instantiate database interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDataBase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}


	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)	//	tokenDetails is the reponse by the database
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	//	set value to reponse struct
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}