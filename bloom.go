package gobloom

import (
	"github.com/reusee/mmh3"
	"github.com/smalllixin/bitarray"
)

// Represents a bloom filter
type BloomFilter struct {
	size      int                // size of the bloom filter
	hashFuncs int                // number of times to hash
	filter    *bitarray.BitArray // bit array that backs the filter
}

// Creates a new bloom filter
func NewBloomFilter(size, hashFuncs int) *BloomFilter {
	return &BloomFilter{
		size,
		hashFuncs,
		bitarray.NewBitArray(uint32(size), 1),
	}
}

// Add new items to the filter
func (bf *BloomFilter) Add(items ...string) {
	for _, item := range items {
		pos := bf.hashItem(item)
		for _, p := range pos {
			bf.filter.SetB(uint32(p), 1)
		}
	}
}

// hash an item `hashFuncs` number of times and return
// the bit positions to be set
func (bf *BloomFilter) hashItem(item string) []int {
	hasher := mmh3.New32()
	positions := make([]int, bf.hashFuncs)

	for i := 0; i < bf.hashFuncs; i++ {
		hasher.Write([]byte(item))
		pos := int(hasher.Sum32()) % bf.size
		positions = append(positions, pos)
	}

	return positions
}

// Check if the filter contains an element
func (bf *BloomFilter) Contains(item string) bool {
	pos := bf.hashItem(item)

	for _, p := range pos {
		if bf.filter.GetB(uint32(p)) == 0 {
			return false
		}
	}

	return true
}
