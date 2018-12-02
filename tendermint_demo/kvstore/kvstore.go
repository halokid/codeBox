package kvstore

import (
  "bytes"
  "encoding/binary"
  "encoding/json"
  "fmt"
  "github.com/tendermint/abci/example/code"
  "github.com/tendermint/abci/types"
  cmn "github.com/tendermint/tmlibs/common"
  dbm "github.com/tendermint/tmlibs/db"
)

var (
  stateKey            =     []byte("stateKey")
  kvPairPrefixKey     =     []byte("kvPairKey")
)

type State struct {
  db              dbm.DB
  Size            int64 `json:"size"`
  Height          int64 `json:"height"`
  AppHash         []byte `json:"app_hash"`
}

// get state
func loadState(db dbm.DB) State {
  stateBytes := db.Get(stateKey)
  var state State
  if len(stateKey) != 0 {
    // decode json format
    err := json.Unmarshal(stateBytes, &state)
    if err != nil {
      panic(err)
    }
  }
  state.db = db
  return state
}

// save state
func saveState(state State) {
  // encode json format
  stateBytes, err := json.Marshal(state)
  if err != nil {
    panic(err)
  }
  state.db.Set(stateKey, stateBytes)
}

func prefixKey(key []byte) []byte {
  return append(kvPairPrefixKey, key...)
}
// ---------------------------------------

var _ types.Application = (*KVStoreApplication)(nil)

type KVStoreApplication struct {
  types.BaseApplication

  state State
}

func NewKVStoreApplication() *KVStoreApplication {
  state := loadState(dbm.NewMemDB())
  return &KVStoreApplication{state:   state}
}

func (app *KVStoreApplication) Info(req types.RequestInfo) (resinfo types.ResponseInfo) {
  return types.ResponseInfo{Data:   fmt.Sprintf("{\"size\":%v)}", app.state.Size)}
}

func (app *KVStoreApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {
  var key, value []byte
  parts := bytes.Split(tx, []byte("="))
  if len(parts) == 2 {
    key, value = parts[0], parts[1]
  } else {
    key, value = tx, tx
  }

  app.state.db.Set(prefixKey(key), value)
  app.state.Size += 1

  tags := []cmn.KVPair {
    {[]byte("app.creator"), []byte("jae")},
    {[]byte("app.key"), key},
  }
  return types.ResponseDeliverTx{Code:  code.CodeTypeOK, Tags:  tags}
}


func (app *KVStoreApplication) CheckTx(tx []byte) types.ResponseCheckTx {
  return types.ResponseCheckTx{Code:    code.CodeTypeOK}
}


func (app *KVStoreApplication) Commit() types.ResponseCommit {
  // Using a memdb - just return the big endian size of the db
  appHash := make([]byte, 8)
  binary.PutVarint(appHash, app.state.Size)
  app.state.AppHash = appHash
  app.state.Height += 1
  saveState(app.state)
  return types.ResponseCommit{Data: appHash}
}

func (app *KVStoreApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
  if reqQuery.Prove {
    value := app.state.db.Get(prefixKey(reqQuery.Data))
    resQuery.Index = -1 // TODO make Proof return index
    resQuery.Key = reqQuery.Data
    resQuery.Value = value
    if value != nil {
      resQuery.Log = "exists"
    } else {
      resQuery.Log = "does not exist"
    }
    return
  } else {
    value := app.state.db.Get(prefixKey(reqQuery.Data))
    resQuery.Value = value
    if value != nil {
      resQuery.Log = "exists"
    } else {
      resQuery.Log = "does not exist"
    }
    return
  }

}




























