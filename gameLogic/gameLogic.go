package gameLogic

import (
  "fmt"
  "log"
  "strconv"
  "math/rand"
  "time"

  "poker/gamble"
  "poker/balance"
  "github.com/notnil/joker/hand"
  "github.com/ethereum/go-ethereum/ethclient"

)
type accountKeys struct {
  public string
  private string
}
func StartGame(client *ethclient.Client, keys []string) {
  playerOneKeys := accountKeys{keys[0], keys[1][2:]}
  playerTwoKeys := accountKeys{keys[2], keys[3][2:]}
  fmt.Println(playerOneKeys)
  fmt.Println(playerTwoKeys)

  fmt.Println("Player 1: How many tokens would you like to bet?")
  var amount string
  _, err := fmt.Scanln(&amount)
  if err != nil {
    log.Fatal("was not able to get input from user")
  }
  amountInt, err := strconv.Atoi(amount)
  if err != nil {
    log.Fatal("did not receive valid input format from user -- please enter integer value")
  }
  fmt.Println("Pot now has ", amountInt*2)

  // come back and learn exactly how the rand system and the poker deck system is working
  // use rand.seed in order to get random seed
  // rand.Seed(time.Now().UnixNano())
  deck := hand.NewDealer(rand.New(rand.NewSource(time.Now().UnixNano()))).Deck()
	h1 := hand.New(deck.PopMulti(5))
	h2 := hand.New(deck.PopMulti(5))

	fmt.Println(h1)
	fmt.Println(h2)

  
	sortedHands := hand.Sort(hand.SortingHigh, hand.DESC, h1, h2)
	fmt.Println("Winner is:", sortedHands[0].Cards())
  if sortedHands[0] == h1 {
    fmt.Println("Winner is player 1")
    gamble.TransferTokens(client, playerTwoKeys.private, playerOneKeys.public, amountInt)
  } else {
    fmt.Println("Winner is player 2")
    gamble.TransferTokens(client, playerOneKeys.private, playerTwoKeys.public, amountInt)
  }

  fmt.Println("balance of player 1", balance.Balance(client, playerOneKeys.public))
  fmt.Println("balance of player 2", balance.Balance(client, playerTwoKeys.public))
}

// expose api to bet and to get balance
