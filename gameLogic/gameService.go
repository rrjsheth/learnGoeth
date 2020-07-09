// Package gameLogic this file should only contain the interface
package gameLogic

//
// import (
//   "errors"
// )

// PokerGameService exported
type PokerGameService interface {
	JoinPokerGame(string, string, int) string
	InfoPokerGame(int) (string, int, int, error)
	// Raise(int, string, int) string
	// Call() ()
	Call(int, string, int, bool) (string, error)
	Bet(int, string, int, bool) (string, error)
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
