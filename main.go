package main

import (
	"log"
	"net/http"
	"os"
	gL "poker/gameLogic"

	"github.com/ethereum/go-ethereum/ethclient"
	Logger "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := Logger.NewLogfmtLogger(os.Stderr)
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("exactly 1 argument -- the public key of casino account -- should be inputted: found -- ", len(args))
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("something went wrong", err)
	}

	gL.StartGame(client, args[0])

	pokerService := gL.PokerService{}

	infoGameHandler := httptransport.NewServer(
		gL.MakeInfoPokerGameEndpoint(pokerService),
		gL.DecodeInfoPokerGameRequest,
		gL.EncodeResponse,
	)

	joinGameHandler := httptransport.NewServer(
		gL.MakeJoinPokerGameRequest(pokerService),
		gL.DecodeJoinPokerGameRequest,
		gL.EncodeResponse,
	)

	betHandler := httptransport.NewServer(
		gL.MakeBetRequest(pokerService, false),
		gL.DecodeBetRequest,
		gL.EncodeResponse,
	)

	callHandler := httptransport.NewServer(
		gL.MakeBetRequest(pokerService, true),
		gL.DecodeBetRequest,
		gL.EncodeResponse,
	)

	http.Handle("/infoGame", infoGameHandler)
	http.Handle("/joinGame", joinGameHandler)
	http.Handle("/bet", betHandler)
	http.Handle("/call", callHandler)

	// create a handler with some middleware that will handle the further logic part
	// will call teh call method but also the winner logic
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
