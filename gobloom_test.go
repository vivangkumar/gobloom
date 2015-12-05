package gobloom

import "testing"

func TestBloomTrue(t *testing.T) {
	bf := NewBloomFilter(10, 3)
	bf.Add("one", "two", "three")
	if !bf.Contains("one") {
		t.Error("Expected to be true, but got false")
	}
}

func TestBloomFalse(t *testing.T) {
	bf := NewBloomFilter(10, 3)
	bf.Add("two")
	if bf.Contains("one") {
		t.Error("Expected to be false, but got true")
	}
}
