package main

import (
  "os"
  "log"
  gL "poker/gameLogic"
  Logger "github.com/go-kit/kit/log"
  "github.com/ethereum/go-ethereum/ethclient"
  "net/http"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main()  {
  logger := Logger.NewLogfmtLogger(os.Stderr)
  args := os.Args[1:]

  if len(args) != 4 {
    log.Fatal("exactly 4 arguments should be inputted: found -- ", len(args))
  }

  client, err := ethclient.Dial("http://127.0.0.1:8545")
  if err != nil {
    log.Fatalf("something went wrong", err)
  }

  gL.StartGame(client, args)

  pokerService := gL.PokerService{}

  infoGameHandler := httptransport.NewServer(
		gL.MakeInfoPokerGameEndpoint(pokerService),
		gL.DecodeInfoPokerGameRequest,
		gL.EncodeResponse,
	)

  http.Handle("/infoGame", infoGameHandler)
  logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
