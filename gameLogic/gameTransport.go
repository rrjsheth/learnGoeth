package gameLogic

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

//request and response
type infoPokerGameRequest struct {
  GameNumber int `json:"gameNumber"`
}

type infoPokerGameResponse struct {
  CasinoAccountNum string `json:"CasinoAccountNum"`
  GameNum int `json:"gameNumber"`
  BuyInAmount int `json:"buyInAmount"`
  Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}


//Endpoints
func MakeInfoPokerGameEndpoint(p PokerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(infoPokerGameRequest)
		acntNum, gameNum, buyIn, err := p.InfoPokerGame(req.GameNumber)
		if err != nil {
			return infoPokerGameResponse{acntNum, gameNum, buyIn, err.Error()}, nil
		}
		return infoPokerGameResponse{acntNum, gameNum, buyIn, ""}, nil
	}
}

func DecodeInfoPokerGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request infoPokerGameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
