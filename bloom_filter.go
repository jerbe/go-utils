package goutils

import (
	"hash/fnv"
	"strconv"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/15 21:08
  @describe :
*/

type BloomFilter struct {
	bitmap []uint64

	size int

	hashNum int
}

func (b *BloomFilter) Set(data string) {
	for i := 0; i < b.hashNum; i++ {
		idx := b.hash(data, i)
		b.setBit(idx)
	}
}

func (b *BloomFilter) Contain(data string) bool {
	for i := 0; i < b.hashNum; i++ {
		idx := b.hash(data, i)
		if !b.checkBit(idx) {
			return false
		}
	}
	return true
}

func (b *BloomFilter) hash(data string, i int) uint64 {
	h := fnv.New64()
	h.Write([]byte(data))
	h.Write([]byte(strconv.Itoa(i)))
	return h.Sum64()
}

func (b *BloomFilter) setBit(i uint64) {
	i = i % uint64(b.size)
	idx := i / 64
	pos := int(i % 64)
	isSet := false

	for !isSet {
		oldVal := AtomicLoadInt(&(b.bitmap[idx])).(uint64)
		newVal := IntBitSet(oldVal, pos).(uint64)
		isSet = AtomicCompareAndSwapInt(&(b.bitmap[idx]), oldVal, newVal)
	}
}

func (b *BloomFilter) checkBit(i uint64) bool {
	i = i % uint64(b.size)
	idx := i / 64
	pos := int(i % 64)
	d := AtomicLoadInt(&b.bitmap[idx]).(uint64)
	return IntBitCheck(d, pos)
}

func NewBloomFilter(size int, hashNum int) *BloomFilter {
	numWords := (size + 63) / 64
	bitmap := make([]uint64, numWords)
	return &BloomFilter{bitmap: bitmap, size: size, hashNum: hashNum}
}
