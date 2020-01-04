// Package bloomfilter provides a simple bloom filter implementation for the GoLang programming language.
// This package includes the ability to create a new bloom filter from an estimated optimal size and
// optimal number of hash functions for the specified estimated max number of items to be set in the
// bloom filter as well as the maximum tolerable failure rate
package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

var defaultHashFn = fnv.New64a()

// BloomFilter is a probabilistic data structure that will always return false
// if the value is not present but will sometimes return true if the value has not
// been set. BloomFilter is space-efficient, however, items cannot be deleted.
type BloomFilter interface {
	// Add adds a value to the bloom filter
	Add(string)
	// Check checks to see if a value has been added to the bloom filter.
	// Check returns true if it thinks the value has been added and false
	// if it thinks the value has not been added.
	Check(string) bool
}

type bloomFilter struct {
	m int64
	k int64

	table  []bool
	hashFn hash.Hash64
}

// New returns a new instance of BloomFilter  NewBloomFilter accepts
// m and k as arguments where m is the size of the bloom filter and k is the
// number of hashing functions.
func New(m, k int64) BloomFilter {
	bf := &bloomFilter{m: m, k: k}
	bf.table = make([]bool, m)
	bf.hashFn = defaultHashFn
	return bf
}

// NewFromEstimate returns a new instance of BloomFilter from an estimate of the required
// size of the bloom filter as well as an estimate of the optimal hash functions for the estimated
// number of items in the bloom filter
// NewBloomFilterFromEstimate takes in expectedNumberOfItems
// and maxFPRate as arguments where expectedNumberOfItems is the max number of items expected to
// be set in the bloom filter and maxFPRate is the max rate of failure tolerable for the filter
func NewFromEstimate(expectedNumberOfItems int64, maxFPRate float64) BloomFilter {
	optimalSize, optimalHashFns := estimateOptimalBloomSize(expectedNumberOfItems, maxFPRate)
	return New(optimalSize, optimalHashFns)
}

func (bf *bloomFilter) Add(value string) {
	indexes := bf.getHashIndexes(value)
	for _, idx := range indexes {
		bf.table[idx] = true
	}
}

func (bf *bloomFilter) Check(value string) bool {
	indexes := bf.getHashIndexes(value)
	for _, idx := range indexes {
		if !bf.table[idx] {
			return false
		}
	}
	return true
}

func estimateOptimalSize(expectedNumberOfItems int64, maxFPRate float64) (optimalSize int64) {
	numerator := float64(expectedNumberOfItems) * math.Abs(math.Log(maxFPRate))
	denom := math.Pow(math.Log(2), 2)

	optimalSize = int64(math.Ceil(numerator / denom))
	return optimalSize
}

func estimateOptimalHashFns(expectedNumberOfItems, optimalSize int64) (optimalHashFns int64) {
	fr := float64(optimalSize / expectedNumberOfItems)
	optimalHashFns = int64(math.Ceil(fr * math.Log(2)))
	return optimalHashFns
}

func estimateOptimalBloomSize(expectedNumberOfItems int64, maxFPRate float64) (optimalSize, optimalHashFns int64) {
	optimalSize = estimateOptimalSize(expectedNumberOfItems, maxFPRate)
	optimalHashFns = estimateOptimalHashFns(expectedNumberOfItems, optimalSize)
	return optimalSize, optimalHashFns
}

func hashValue(value string, k int64, hashFn hash.Hash64) (uint32, uint32) {
	hashFn.Reset()
	hashFn.Write([]byte(value))
	h64 := hashFn.Sum64()
	h1 := uint32(h64 & ((1 << 32) - 1))
	h2 := uint32(h64 >> 32)
	return h1, h2
}

func (bf *bloomFilter) getHashIndexes(value string) []uint32 {
	var indexes []uint32
	h1, h2 := hashValue(value, bf.k, bf.hashFn)
	for i := 0; int64(i) < bf.k; i++ {
		idx := (h1 + uint32(i)*h2) % uint32(bf.m)
		indexes = append(indexes, idx)
	}
	return indexes
}
