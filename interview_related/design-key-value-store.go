package interview_related

import "sync"

type KVStore struct {
	base map[string]any
	// if len(txns) != 0, means we are under a transaction,
	// right now, all transactions share same global stack, maybe we should create multiple transaction stack to handle nested transaction
	txns []map[string]any

	mu sync.RWMutex
}

func Constructor() KVStore {
	return KVStore{
		base: make(map[string]any),
	}
}

func (kv *KVStore) Get(key string) any {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	// check transaction first
	for _, txn := range kv.txns {
		if val, ok := txn[key]; ok {
			return val
		}
	}

	return kv.base[key]
}

func (kv *KVStore) Set(key string, value any) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if len(kv.txns) == 0 {
		kv.base[key] = value
	} else {
		lastTxn := kv.txns[len(kv.txns)-1]
		lastTxn[key] = value
	}
}

func (kv *KVStore) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if len(kv.txns) == 0 {
		delete(kv.base, key)
	} else {
		lastTxn := kv.txns[len(kv.txns)-1]
		lastTxn[key] = nil // not delete directly, set it to nil
	}
}

func (kv *KVStore) Begin() {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.txns = append(kv.txns, make(map[string]any))
}

func (kv *KVStore) Commit() {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if len(kv.txns) == 0 {
		return
	}

	lastTxn := kv.txns[len(kv.txns)-1]
	kv.txns = kv.txns[:len(kv.txns)-1] // pop last

	if len(lastTxn) == 0 { // merge to base
		for k, v := range lastTxn {
			if v == nil {
				delete(kv.base, k)
			} else {
				kv.base[k] = v
			}
		}
	} else {
		curLastTxn := kv.txns[len(kv.txns)-1]
		for k, v := range lastTxn {
			curLastTxn[k] = v
		}
	}
}

func (kv *KVStore) Rollback() {
	if len(kv.txns) == 0 {
		return
	}
	kv.txns = kv.txns[:len(kv.txns)-1]
}
