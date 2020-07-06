// this file should only contain the business logic of the service
package gameLogic

import (
//  "fmt"
  "log"
  "math/rand"
  "time"
  "errors"
  "strconv"
  // "poker/gamble"
  // "poker/balance"
  "github.com/notnil/joker/hand"
  "github.com/ethereum/go-ethereum/ethclient"

)
type accountKeys struct {
  public string
  private string
}

type GameMetaData struct {
  casinoAccountNumber string
  players [](*PlayerMetaData)
  gameNumber int
  totalPotMoney int
  buyInAmount int
  turn string
  latestBetValue int
}

type PlayerMetaData struct {
  playerId string
  moneyLeft int
  currentBet int
}

type PokerService struct {}

var OnGoingGames []GameMetaData
var globalEthClient *ethclient.Client

func getGame(gameNumber int) *GameMetaData {
  for _, game := range OnGoingGames {
    if gameNumber == game.gameNumber {
      return &game
    }
  }
    return nil
}

// func getPlayer(gameNumber int, playerId string) *PlayerMetaData {
//   game := getGame(gameNumber)
//   for _, player := range game.players {
//     if playerId == player.playerId {
//       return &player
//     }
//   }
//     return nil
// }

func (s PokerService) InfoPokerGame(gameNumber int) (string, int, int, error) {
  game := getGame(gameNumber)
  if game == nil {
    return "", -1, 0, errors.New("Did not find gameNumber" + strconv.Itoa(gameNumber) )
  }
  return game.casinoAccountNumber, game.gameNumber, game.buyInAmount, nil
}

// check if the user has paid his buy in to get into the game
// what is the player's id
// need to return the hand to start playing, the other has to be errors to return
func (s PokerService) JoinPokerGame(buyInTransactionHash string, playerId string, gameNumber int) (string, error) {
  game := getGame(gameNumber)
  if game == nil {
    log.Println("did nto find that game; number provided does not exist", gameNumber)
    return "", errors.New("did nto find that game; number provided does not exist" + strconv.Itoa(gameNumber))
  }

  // come back and learn exactly how the rand system and the poker deck system is working
  // use rand.seed in order to get random seed
  // rand.Seed(time.Now().UnixNano())
  deck := hand.NewDealer(rand.New(rand.NewSource(time.Now().UnixNano()))).Deck()
	dealtHand := hand.New(deck.PopMulti(5))
	log.Println(dealtHand, buyInTransactionHash)

  // TODO check if using signatures is better than doing these checks
  amountTransferred, senderAddress, receiverAddress := gamble.GetTransactionInfo(globalEthClient, buyInTransactionHash)
  if amountTransferred < game.buyInAmount {
    fmt.Println("did not pay enough money to get into game")
    return "", errors.New("did not pay enough money to get into game")
  } else if senderAddress == playerId {
    fmt.Println("player id did not match id of payer of buyin")
    return "", errors.New("player id did not match id of payer of buyin")
  } else if receiverAddress == game.casinoAccountNumber {
    fmt.Println("did not send money to the right account")
    return "", errors.New("did not send money to the right account")
  }

  // TODO put actual amout that user send to casino account
  game.players = append(game.players, &PlayerMetaData{playerId, amountTransferred, 0})
  return dealtHand.String(), nil
}

// func (s PokerService) Raise(gameNumber int, playerId string, raiseAmount int) string {
//   game := getGame(gameNumber)
//   player := getPlayer(gameNumber, playerId)
//   if game == nil {
//     return "", -1, 0, errors.New("Did not find gameNumber" + strconv.Itoa(gameNumber) )
//   }
//   newLatestBetAmount := player.currentBet + raiseAmount
//   if player.moneyLeft < newLatestBetAmount {
//     return "You do not have enough money to raise " + strconv.Itoa(raiseAmount)
//   }
//   if newLatestBetAmount < game.latestBetValue {
//     return fmt.Sprintf("you need to raise at least %d\nwe received request for only %d",game.latestBetValue-player.currentBet , raiseAmount)
//   }
//
//   // TODO raise the actual amount
//   return "successfully raised"
// }
// func (s PokerService) Bet() string{
      // fmt.Println("Player 1: How many tokens would you like to bet?")
      // var amount string
      // _, err := fmt.Scanln(&amount)
      // if err != nil {
      //   log.Fatal("was not able to get input from user")
      // }
      // amountInt, err := strconv.Atoi(amount)
      // if err != nil {
      //   log.Fatal("did not receive valid input format from user -- please enter integer value")
      // }
      // fmt.Println("Pot now has ", amountInt*2)
//   return "successfully betted"
// }
// func (s PokerService) Fold() string{
//   return "successfully folded"
// }
func StartGame(client *ethclient.Client, casinoAccountNumber string) {
  globalEthClient = client
  OnGoingGames = append(OnGoingGames, GameMetaData{casinoAccountNumber, nil, 0, 0, 1000, "-1", 0}) // fill out this struct
  // playerOneKeys := accountKeys{keys[0], keys[1][2:]}
  // playerTwoKeys := accountKeys{keys[2], keys[3][2:]}

  // TODO winner logic is defined here
	// sortedHands := hand.Sort(hand.SortingHigh, hand.DESC, h1, h2)
	// fmt.Println("Winner is:", sortedHands[0].Cards())

  // var transactionId string
  // if sortedHands[0] == h1 {
  //   fmt.Println("Winner is player 1")
  //   transactionId = gamble.TransferTokens(client, playerTwoKeys.private, playerOneKeys.public, amountInt)
  // } else {
  //   fmt.Println("Winner is player 2")
  //   transactionId = gamble.TransferTokens(client, playerOneKeys.private, playerTwoKeys.public, amountInt)
  // }
  //
  // fmt.Println("balance of player 1", balance.Balance(client, playerOneKeys.public))
  // fmt.Println("balance of player 2", balance.Balance(client, playerTwoKeys.public))
  // gamble.GetTransactionInfo(globalEthClient, transactionId)
}

// expose api to bet and to get balance
