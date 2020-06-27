package main

import (
  "fmt"
  "os"
  "log"
  "strconv"
  "math/rand"
  "gamble"
  "github.com/notnil/joker/hand"
  "github.com/ethereum/go-ethereum/ethclient"
)

func main()  {
  args := os.Args[1:]

  if len(args) != 2 {
    log.Fatal("exactly 2 arguments should be inputted: found -- ", len(args))
  }

  fmt.Println("How many tokens would you like to bet?")
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
  deck := hand.NewDealer(rand.New(rand.NewSource(1021))).Deck()
	h1 := hand.New(deck.PopMulti(5))
	h2 := hand.New(deck.PopMulti(5))

	fmt.Println(h1)
	fmt.Println(h2)

	hands := hand.Sort(hand.SortingHigh, hand.DESC, h1, h2)
	fmt.Println("Winner is:", hands[0].Cards())
  if hands[0] == h1 {
    fmt.Println("winner is player 1")
  } else {
    fmt.Println("winner is player 2")
  }

  client, err := ethclient.Dial("http://127.0.0.1:8545")
  if err != nil {
    log.Fatalf("something went wrong", err)
  }
  gamble.TransferTokens(client, "31975964a567db9e9497077e59d990313085dc4723081c1064a4accbc7d9be94", "0xd156A9c65d5447A51a0567d6A2c7Df7B681b90b2", amount)
}
