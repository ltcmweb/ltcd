package scrypt

import "sync/atomic"

type (
	Hash    struct{ Key, Val []byte }
	hashMap map[string][]byte
)

var cache atomic.Value

func Scrypt(x []byte) []byte {
	if m, ok := cache.Load().(hashMap); ok {
		if x, ok := m[string(x)]; ok {
			return x
		}
	}
	return scrypt(x)
}

func SetCache(hashes []Hash) {
	m := hashMap{}
	for _, hash := range hashes {
		m[string(hash.Key)] = hash.Val
	}
	cache.Store(m)
}
