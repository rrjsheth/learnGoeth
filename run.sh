#!/bin/bash
PLAYERONEPUBLIC="0x438a6865772c275e35ba577ad5342366f7d3a7f2"
PLAYERONEPRIVATE="0x048e34540b752212d3060230fcc0dc8640bdd93078de78142a8071ef7ea6a71b"
PLAYERTWOPUBLIC="0xeb6f90ba15caea7cf48025dff785d5682c51fd0f"
PLAYERTWOPRIVATE="0xd21d6b5fcc6d67e02d1dc18a15daf744446f35d9f6b9dc44a48cf1965719b2a8"
# ganache-cli --accounts 2 --acctKeys ./player_keys.json
# rm ./player_keys.json
go run main.go $PLAYERONEPUBLIC $PLAYERONEPRIVATE $PLAYERTWOPUBLIC $PLAYERTWOPRIVATE


#curl -XPOST -d'{"gameNumber":0}' localhost:8080/infoGame
#curl -XPOST -d'{"transactionHash":"1","playerId":"31212","gameNumber":0}' localhost:8080/joinGame
