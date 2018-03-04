package golang_natural_blockchain

import (
  "bufio"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "io"
  "log"
  "net"
  "os"
  "strconv"
  "time"
  "sync"

  "github.com/davecgh/go-spew/spew"
  "github.com/joho/godotenv"
  "text/scanner"
)

type Block struct {
  Index         int
  Timestamp     string
  BPM           int
  Hash          string
  PreHash       string
}


var Blockchain []Block
var mutex = &sync.Mutex{}


//生成一个hash值
func calculateHash(block Block) string {
  record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
  h := sha256.New()
  h.Write([]byte(record))
  hashed := h.Sum(nil)
  return hex.EncodeToString(hashed)
}


//生成一个区块
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


//验证区块是否正确
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


//bcServer处理进来的并行请求的区块
var bcServer chan []Block



func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  bcServer = make(chan []Block)

  //create genesis block
  t := time.Now()
  genesisBlock := Block{0, t.String(), 0, "", " "}
  spew.Dump(genesisBlock)
  Blockchain = append(Blockchain, genesisBlock)

  //start TCP and server TCP server
  //建立一个server
  server, err := net.Listen("tcp", ":" + os.Getenv("TCPADDR"))
  if err != nil {
    log.Fatal(err)
  }

  defer server.Close()

  //循环处理 server收到的请求
  for {
    conn, err := server.Accept()
    if err != nil {
      log.Fatal(err)
    }
    //处理
    go handleConn(conn)
  }
}


func handleConn(conn net.Conn) {
  defer conn.Close()

  io.WriteString(conn, "Enter a new BPM:")

  scanner = bufio.NewScanner(conn)

  // take in BPM from stdin and add it to blockchain after conducting necessary validation
  go func() {
    for scanner.Scan() {
      bpm, err := strconv.Atoi(scanner.Text())
      if err != nil {
        log.Printf("%v not a number: %v", scanner.Text(), err)
        continue
      }
      newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], bpm)
      if err != nil {
        log.Println(err)
        continue
      }
      if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
        newBlockchain := append(Blockchain, newBlock)
        replaceChain(newBlockchain)
      }

      bcServer <- Blockchain
      io.WriteString(conn, "\nEnter a new BPM:")
    }
  }()


  //simulate reaeiving broadcast
  //为了保证不输出相同的blockchai， 这里需要用 mutex 这个锁， 而且用了一个延时，这里的输出不用那么频繁
  go func() {
    for {
      time.Sleep(30 * time.Second)
      mutex.Lock()
      output, err := json.Marshal(Blockchain)
      if err != nil {
        log.Fatal(err)
      }
      mutex.Unlock()
      io.WriteString(conn, string(output))
    }
  }()

  for _= range bcServer {
    spew.Dump(Blockchain)
  }
}


// make sure the chain we're checking is longer than the current blockchain
func replaceChain(newBlocks []Block) {
  mutex.Lock()
  if len(newBlocks) > len(Blockchain) {
    Blockchain = newBlocks
  }
  mutex.Unlock()
}


