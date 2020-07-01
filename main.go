package main

import (
  "os"
  "log"
  gL "poker/gameLogic"
  "github.com/ethereum/go-ethereum/ethclient"
)

func main()  {
  args := os.Args[1:]

  if len(args) != 4 {
    log.Fatal("exactly 4 arguments should be inputted: found -- ", len(args))
  }

  client, err := ethclient.Dial("http://127.0.0.1:8545")
  if err != nil {
    log.Fatalf("something went wrong", err)
  }

  gL.StartGame(client, args)
}
