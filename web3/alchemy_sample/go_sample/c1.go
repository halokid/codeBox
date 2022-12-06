package main

import (
  "context"
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
  client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/Ocf_EETGbdsBqKJ3f-DNC4YfxK7F1JB9")
  if err != nil {
    log.Fatal(err)
  }

  // Get the balance of an account
  account := common.HexToAddress("0x9851099eF06BAD791045281c193847CF32da70E8")
  balance, err := client.BalanceAt(context.Background(), account, nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Account balance:", balance) // 25893180161173005034

  // Get the latest known block
  block, err := client.BlockByNumber(context.Background(), nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Latest block:", block.Number().Uint64())
}



