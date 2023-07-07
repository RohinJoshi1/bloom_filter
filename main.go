package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	filter []byte
	size   int
}

var mHasher hash.Hash32

func init() {
	// mHasher = murmur3.New32()
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func mHash(key string, size int) int {
	mHasher.Write([]byte(key))
	result := int(mHasher.Sum32()) % size
	mHasher.Reset()
	return int(result)

}
func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		make([]uint8, size),
		size,
	}
}

func (bf *BloomFilter) Add(key string) {
	num := mHash(key, bf.size) //This is giving me the index
	//Get the index of the byte in the arry
	aidx := num / 8
	bidx := num % 8
	//Set the appropriate bit to 1
	bf.filter[aidx] = bf.filter[aidx] | 1<<bidx

}

func (bf *BloomFilter) Contains(key string) bool {
	idx := mHash(key, bf.size)
	aidx := idx / 8
	bidx := idx % 8
	//Set the appropriate bit to 1
	bit := bf.filter[aidx] & (1 << bidx)
	var res bool
	res = bit > 0 //if 1 return true else return false
	return res
}

func main() {
	bloom := NewBloomFilter(16)
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		bloom.Add(key)
	}

	for _, key := range keys {
		fmt.Println(bloom.Contains(key))
	}
	fmt.Println(bloom.Contains("z"))

}
