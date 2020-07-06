// this file should only contain the interface
package gameLogic
// 
// import (
//   "errors"
// )

type PokerGameService interface {
  // JoinPokerGame(string, string, string) string
  InfoPokerGame(int) (string, int, int, error)
  // Raise(int, string, int) string
  // Call() ()
  // Bet() string
  // Fold() string
}

// type joinPokerGameRequest struct {
// 	S string `json:"s"`
// }
//
// type joinPokerGameResponse struct {
// 	V   string `json:"v"`
// 	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
// }
