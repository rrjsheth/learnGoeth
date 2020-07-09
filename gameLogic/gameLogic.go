// this file should only contain the business logic of the service
package gameLogic

import (
	//  "fmt"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
	// "poker/gamble"
	// "poker/balance"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/notnil/joker/hand"
)

type accountKeys struct {
	public  string
	private string
}

type GameMetaData struct {
	casinoAccountNumber string
	players             [](*PlayerMetaData)
	gameNumber          int
	totalPotMoney       int
	buyInAmount         int
	turn                string
	latestBetValue      int
}

type PlayerMetaData struct {
	playerID    string
	moneyLeft   int
	currentBet  int
	currentHand *(hand.Hand)
}

type PokerService struct{}

var OnGoingGames []GameMetaData
var globalEthClient *ethclient.Client

func getGame(gameNumber int) *GameMetaData {
	for i, game := range OnGoingGames {
		if gameNumber == game.gameNumber {
			// log.Println("check if game is same as ongoinggame[0]", &game == &OnGoingGames[i])
			return &OnGoingGames[i]
		}
	}
	return nil
}

func getPlayer(gameNumber int, playerID string) *PlayerMetaData {
	game := getGame(gameNumber)
	for _, player := range game.players {
		if playerID == player.playerID {
			return player
		}
	}
	return nil
}

func (s PokerService) InfoPokerGame(gameNumber int) (string, int, int, error) {
	game := getGame(gameNumber)
	if game == nil {
		return "", -1, 0, errors.New("Did not find gameNumber" + strconv.Itoa(gameNumber))
	}
	return game.casinoAccountNumber, game.gameNumber, game.buyInAmount, nil
}

// check if the user has paid his buy in to get into the game
// what is the player's id
// need to return the hand to start playing, the other has to be errors to return
func (s PokerService) JoinPokerGame(buyInTransactionHash string, playerID string, gameNumber int) (string, error) {
	game := getGame(gameNumber)
	if game == nil {
		log.Println("did nto find that game; number provided does not exist", gameNumber)
		return "", errors.New("did nto find that game; number provided does not exist" + strconv.Itoa(gameNumber))
	}

	if player := getPlayer(gameNumber, playerID); player != nil {
		return "", errors.New("You are already in the game; right now buying in again isnt allowed")
	}

	// come back and learn exactly how the rand system and the poker deck system is working
	// use rand.seed in order to get random seed
	// rand.Seed(time.Now().UnixNano())
	deck := hand.NewDealer(rand.New(rand.NewSource(time.Now().UnixNano()))).Deck()
	dealtHand := hand.New(deck.PopMulti(5))
	log.Println(dealtHand, buyInTransactionHash)

	// TODO check if using signatures is better than doing these checks
	// amountTransferred, senderAddress, receiverAddress := gamble.GetTransactionInfo(globalEthClient, buyInTransactionHash)
	// if amountTransferred < game.buyInAmount {
	//   fmt.Println("did not pay enough money to get into game")
	//   return "", errors.New("did not pay enough money to get into game")
	// } else if senderAddress == playerID {
	//   fmt.Println("player id did not match id of payer of buyin")
	//   return "", errors.New("player id did not match id of payer of buyin")
	// } else if receiverAddress == game.casinoAccountNumber {
	//   fmt.Println("did not send money to the right account")
	//   return "", errors.New("did not send money to the right account")
	// }
	// game.players = append(game.players, &PlayerMetaData{playerID, amountTransferred, 0})
	game.players = append(game.players, &PlayerMetaData{playerID, game.buyInAmount, 0, dealtHand})
	log.Println(OnGoingGames)
	log.Println(&OnGoingGames[0] == game)
	returnMsg := "you have successfully joined the game. Here is your first hand: " + dealtHand.String()
	return returnMsg, nil
}

// func (s PokerService) Raise(gameNumber int, playerID string, raiseAmount int) string {
//   game := getGame(gameNumber)
//   player := getPlayer(gameNumber, playerID)
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

func (s PokerService) Bet(gameNumber int, playerID string, betAmount int, isCall bool) (string, error) {
	game := getGame(gameNumber)
	if game == nil {
		return "", errors.New("Did not find gameNumber " + strconv.Itoa(gameNumber))
	}
	player := getPlayer(gameNumber, playerID)
	if player == nil {
		return "", errors.New("Did not find playerID " + playerID)
	}

	fmt.Printf("Player 1 bets %d tokens\n", betAmount)
	player.moneyLeft = player.moneyLeft - betAmount
	player.currentBet = player.currentBet + betAmount
	game.totalPotMoney = game.totalPotMoney + betAmount
	game.latestBetValue = player.currentBet

	log.Printf("%+v\n%v", game, *(game.players[0]))
	log.Println("Pot now has ", game.totalPotMoney)
	log.Println("on going games", OnGoingGames)
	if isCall {
		defer s.winLogic(gameNumber)
		return "successfully called", nil
	}
	return "successfully betted", nil
}

func (s PokerService) winLogic(gameNumber int) {
	game := getGame(gameNumber)
	allHands := make([](*hand.Hand), len(game.players))
	handOwner := make(map[string]string) // make map here to find the player with that hand
	log.Println(game.players, len(game.players))
	for index, player := range game.players {
		allHands[index] = player.currentHand
		handOwner[player.currentHand.String()] = player.playerID
	}
	// come back and make this a little better
	// put them all into a map or slice
	log.Println("length of hands ", len(allHands))
	log.Println(allHands)
	// TODO winner logic is defined here
	sortedHands := hand.Sort(hand.SortingHigh, hand.DESC, allHands...)
	fmt.Println("Winner is:", sortedHands[0].Cards(), handOwner[sortedHands[0].String()])

	// put transfer money logic here

}

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
