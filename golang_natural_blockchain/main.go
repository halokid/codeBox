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

type Block struct {
  Index         int
  Timestamp     string
  BPM           int
  Hash          string
  PreHash       string
}


var Blockchain []Block


func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}




func generateBlock(oldBlock Block, BPM int) (Block, error) {
  var newBlock Block

  t := time.Now()

  newBlock.Index = oldBlock.Index + 1
  newBlock.Timestamp = t.String()
  newBlock.BPM = BPM
  newBlock.PreHash = oldBlock.Hash
  newBlock.Hash = calculateHash(newBlock)

  return newBlock, nil
}



func isBlockValid(newBlock, oldBlock Block) bool {
  if oldBlock.Index + 1 != newBlock.Index {
    return false
  }

  if oldBlock.Hash != newBlock.PreHash {
    return false
  }

  if calculateHash(oldBlock) != newBlock.Hash {
    return  false
  }

  return true
}



func replaceChain(newBlocks []Block) {
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
}


func run() error {
  mux := makeMuxRouter()
  httpAddr := os.Getenv("ADDR")
  log.Println("listening on ", os.Getenv("ADDR"))
  s := &http.Server{
    Addr:             ":" + httpAddr,
    Handler:          mux,
    ReadTimeout:      10 * time.Second,
    WriteTimeout:     10 * time.Second,
    MaxHeaderBytes:   1 << 20,
  }

  if err := s.ListenAndServe(); err != nil {
    return err
  }
  return nil
}


func makeMuxRouter() http.Handler {
  muxRouter := mux.NewRouter()
  muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
  muxRouter.HandleFunc("/", handleWriteBlockchain).Methods("POST")
  return muxRouter
}


func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
  bytes, err := json.MarshalIndent(Blockchain, "", " ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  io.WriteString(w, string(bytes))
}


type Message struct {
  BPM int
}


func handleWriteBlockchain(w http.ResponseWriter, r *http.Request) {
  var m Message

  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&m); err != nil {
    respondWithJSON(w, r, http.StatusBadRequest, r.Body)
    return
  }

  defer r.Body.Close()

  newBlock, err := generateBlock(Blockchain[len(Blockchain) - 1], m.BPM)
  if err != nil {
    respondWithJSON(w, r, http.StatusInternalServerError, m)
    return
  }

  if isBlockValid(newBlock, Blockchain[len(Blockchain) - 1]) {
    newBlockchain := append(Blockchain, newBlock)
    replaceChain(newBlockchain)
    spew.Dump(Blockchain)
  }

  respondWithJSON(w, r, http.StatusCreated, newBlock)
}

//参数 payload sinterface{} 的完美应用
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
  response, err := json.MarshalIndent(payload, "", "  ")
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("HTTP 500: Internal Server Error"))
    return
  }
  w.WriteHeader(code)
  w.Write(response)
}




func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  go func() {
    t := time.Now()
    //生成第一个创世块, 其实这里的hash值可以随你自己设置
    genesisBlock := Block{0, t.String(), 0, "", ""}
    spew.Dump(genesisBlock)
    Blockchain = append(Blockchain, genesisBlock)
  }()
  log.Fatal(run())

}
























