package main

import (
  "fmt"
  "crypto/sha256"
  "time"
)

//区块链的数据结构
type Block struct {
  Index                 int64   `json:"index"`              //区块索引
  PreviousHash          string  `json:"previousHash"`       //前块哈希
  Timestamp             int64   `json:"timestamp"`          //时间戳
  Data                  string  `json:"data"`               //区块数据
  Hash                  string `json:"hash"`                //区块哈希
}


//区块链的 Hash值 的计算方法
func calculateHashForBlock(b *Block) string {
  return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d%s%d%s",
                                          b.Index, b.PreviousHash,
                                          b.Timestamp, b.Data))))
}


//生成一个新的区块
func generateNextBlock(data string) (nb *Block) {
  var previousBlock = getLastBlock()
  nb = &Block {
      Data:                   data,
      PreviousHash:           previousBlock.Hash,
      Index:                  previousBlock.Index + 1,
      Timestamp:              time.Now().Unix(),
  }

  nb.Hash = calculateHashForBlock(nb)
  return
}


//定义创世块的数据
var genesisBlock = &Block {
  Index:                  0,
  PreviousHash:           "0",
  Timestamp:              1465154705,
  Data:                   "my genesis block!",
  Hash:                   "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
}


//实现区块存储的接口程序
type ByIndex []*Block

//计算长度
func (b ByIndex) Len() int            {return len(b)}

//交换区块
func (b ByIndex) Swap(i, j int)       {b[i], b[j] = b[j], b[i]}

//比较索引的大小
func (b ByIndex) Less(i, j int)       { return b[i].Index < b[j].Index }



//区块链完整性校验
func isValidNewBlock(nb, pb *Block) (ok bool) {
  if nb.Hash == calculateHashForBlock(nb) &&
                pb.Index + 1 == nb.Index  &&
                pb.Hash == nb.PreviousHash {
    ok = true
  }

  return
}


































