package maps

import "sync"

type SyncMapBenchmarkAdapter struct {
  m  sync.Map
}

func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
  return &SyncMapBenchmarkAdapter{}
}

func (m *SyncMapBenchmarkAdapter) Set(key interface{}, val interface{}) {
  m.m.Store(key, val)
}

func (m *SyncMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
  return m.m.Load(key)
}

func (m *SyncMapBenchmarkAdapter) Del(key interface{}) {
  m.m.Delete(key)
}


