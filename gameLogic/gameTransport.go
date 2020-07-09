package gameLogic

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

//request and response
type infoPokerGameRequest struct {
	GameNumber int `json:"gameNumber"`
}

type infoPokerGameResponse struct {
	CasinoAccountNum string `json:"CasinoAccountNum"`
	GameNum          int    `json:"gameNumber"`
	BuyInAmount      int    `json:"buyInAmount"`
	Err              string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type joinPokerGameRequest struct {
	BuyInTransactionHash string `json:"transactionHash"`
	PlayerId             string `json:"playerId"`
	GameNumber           int    `json:"gameNumber"`
}
type joinPokerGameResponse struct {
	Hand string `json:"hand"`
	Err  string `json:"err,omitempty"`
}

type betRequest struct {
	GameNumber int    `json:"gameNumber"`
	PlayerId   string `json:"playerId"`
	BetAmount  int    `json:"betAmount"`
}

type betResponse struct {
	ResponseMessage string `json:"message"`
	Err             string `json:"err,omitempty"`
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

// TODO why do we need context here
func MakeJoinPokerGameRequest(p PokerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(joinPokerGameRequest)
		log.Println(request)
		hand, err := p.JoinPokerGame(req.BuyInTransactionHash, req.PlayerId, req.GameNumber)
		if err != nil {
			return joinPokerGameResponse{"", err.Error()}, nil
		}
		return joinPokerGameResponse{hand, ""}, nil
	}
}
func DecodeJoinPokerGameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request joinPokerGameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeBetRequest(p PokerService, isCall bool) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(betRequest)
		message, err := p.Bet(req.GameNumber, req.PlayerId, req.BetAmount, isCall)
		if err != nil {
			return betResponse{"", err.Error()}, nil
		}
		return betResponse{message, ""}, nil
	}
}

func DecodeBetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request betRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
