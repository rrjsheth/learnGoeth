package gamble

import (
  "context"
  "crypto/ecdsa"
  "fmt"
  "log"
  "math/big"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/ethclient"
)

func GetTransactionInfo(client *ethclient.Client, transactionHash string) ( amountTransferred *big.Int, senderAddress *common.Address, receiverAddress *common.Address) {
  transactionDetails, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(transactionHash))
  if err != nil {
    log.Fatal("error wihle trying to get transaction", err)
  }

  amountTransferred = transactionDetails.Value()
  receiverAddress = transactionDetails.To()
  if msg, err := transactionDetails.AsMessage(types.HomesteadSigner{}); err == nil {
    senderAddress = msg.From()
  }
  fmt.Println("transaction pending?", isPending)
  return
}

func TransferTokens(client *ethclient.Client, privateSourceAddr string, publicDestAddr string, betAmount int) string {
  privateKey, err := crypto.HexToECDSA(privateSourceAddr)
  if err != nil {
    log.Fatal(err)
  }

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // called type assertion

  if !ok {
    log.Fatal("something wrong")
  }

  fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
  nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
  if err != nil {
    log.Fatal("nonce error")
  }

  value := big.NewInt(5e18)
  gasLimit := uint64(21000)
  gasPrice, err := client.SuggestGasPrice(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  toAddress := common.HexToAddress(publicDestAddr)
  var data []byte
  tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
  signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(signedTx)
  err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %s wei to %s: %s\n", value.String(), toAddress.Hex(), signedTx.Hash().Hex())
  fmt.Println("hash value", signedTx.Hash())
  return signedTx.Hash().Hex()
}
