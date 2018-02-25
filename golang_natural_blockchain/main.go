package main

import (
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net/http"
  "os"
  "time"
  "github.com/davecgh/go-spew/spew"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)



//区块链的数据模型
type Block struct {
  index             int         //这个块在整个链中的位置
  timestamp         string      //生成块的时间戳
  BPM               int         //每分钟的心跳数， 也就是心率, 这个只是这个范例里面的示例数据而已
  hash              string      //是这个块通过SHA256算法生成的散列值
  prev_hash         string      //代表前一个块的 SHA256 散列值
}


//定义一个区块链的全局变量
var blockchain []Block


//计算给定数据的SHA256 的散列值
func calculateHash(block Block) string {
  record := string(block.index) + block.timestamp + string(block.BPM) + block.prev_hash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}


//生成一个区块
func generateBlock(old_block Block, BPM int) (Block, error) {
  var new_block Block
  t := time.Now()

  new_block.index = old_block.index + 1
  new_block.timestamp = t.String()
  new_block.BPM = BPM
  new_block.prev_hash = old_block.hash
  new_block.hash = calculateHash(new_block)
  return new_block, nil
}



//校验一个区块是否被篡改， 是否正确
func isBlockValid(new_block, old_block Block) bool {
  if old_block.index + 1 != new_block.index {
    return false
  }
  if old_block.hash != new_block.prev_hash {
    return false
  }
  if calculateHash(new_block) != new_block.hash {
    return false
  }
  return true
}


/**
更新某个区块链为最新的链， 因为区块链是会储存在每一个计算节点上的，有一些计算节点因为网络问题或者其他原因，没有
更新到最新的区块链， 所以我们在某个区块链上进行一些操作的时候， 需要把本地的区块链的数据更新到最新，不然会有数据的不同步
**/
func replaceChain(new_blocks []Block) {
  if len(new_blocks) > len(blockchain) {
    blockchain = new_blocks
  }
}


func run() error {
  mux := makeMuxRouter()
  http_addr := os.Getenv("ADDR")
  log.Println("listening on ", os.Getenv("ADDR"))

  s := &http.Server{
    Addr:             ":" + http_addr,
    Handler:          mux,
    ReadTimeout:      10 * time.Second,
    WriteTimeout:     10 * time.Second,
    MaxHeaderBytes:   1 << 20,    //向左位移， 2的20次方
  }

  if err := s.ListenAndServe(); err != nil {
    return err
  }
  return nil
}








