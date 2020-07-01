package balance

import (
  "fmt"
  "math/big"
  "context"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"

)
func Balance(client *ethclient.Client, publicAddr string) *big.Int {
  account := common.HexToAddress(publicAddr)
  balance, err := client.BalanceAt(context.Background(), account, nil)
  if err != nil {
    fmt.Println(err)
  }
  return balance;
}


// expose api to bet and to get balance
//
