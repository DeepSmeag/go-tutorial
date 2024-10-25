package handlers

import (
	"encoding/json"
	"goapi/api"
	"goapi/internal/tools"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams = api.CoinBalanceParams{}

	var decoder *schema.Decoder = schema.NewDecoder()
	var err error = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error()
		api.InternalErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
	var response = api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: tokenDetails.Coins,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
