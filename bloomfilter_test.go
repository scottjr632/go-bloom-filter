package bloomfilter

import "testing"

func TestCanMakeBloomFilter(t *testing.T) {
	bf := New(100, 2)
	if bf == nil {
		t.Errorf("bf in nil")
	}
}

func TestCanMakeOptimalBloomFilter(t *testing.T) {
	items, fr := 300, 0.01
	bf := NewFromEstimate(int64(items), fr)
	if bf == nil {
		t.Errorf("Unable to make new bloom filter")
	}
}

func TestCanMakeFilterUnder32(t *testing.T) {
	v := "test"
	bf := New(10, 2)
	bf.Add(v)

	if !bf.Check(v) {
		t.Error("value was set but not found")
	}
}

func TestCanAddToFilter(t *testing.T) {
	v := "test"
	bf := New(100, 2)
	bf.Add(v)
}

func TestCanCheckFilter(t *testing.T) {
	v := "test"
	bf := New(100, 2)
	bf.Add(v)

	if !bf.Check(v) {
		t.Errorf("value was not found but was set")
	}

	if bf.Check("not actually set") {
		t.Errorf("found value that has not been set")
	}
}

func TestCanCheckFilterFromOptimal(t *testing.T) {
	v := "test"

	items, fr := 300, 0.01
	bf := NewFromEstimate(int64(items), fr)
	if bf == nil {
		t.Errorf("Unable to make new bloom filter")
	}

	bf.Add(v)

	if !bf.Check(v) {
		t.Errorf("value was not found but was set")
	}

	if bf.Check("not actually set") {
		t.Errorf("found value that has not been set")
	}
}

func TestEstimateOptimalSize(t *testing.T) {
	items, fr := 3000000, 0.01
	expectedOptimalSize := 28755176
	optimalSize := estimateOptimalSize(int64(items), fr)
	if optimalSize != int64(expectedOptimalSize) {
		t.Errorf("Invalid optimal size. Got %v Exptected %v", optimalSize, expectedOptimalSize)
	}
}

func TestEstimateOptimalHashFns(t *testing.T) {
	items, fr := int64(3000000), 0.01
	expectedOptimalSize := 28755176
	expectedHashFns := int64(7)
	optimalSize := estimateOptimalSize(items, fr)
	if optimalSize != int64(expectedOptimalSize) {
		t.Errorf("Invalid optimal size. Got %v Exptected %v", optimalSize, expectedOptimalSize)
	}

	hashFns := estimateOptimalHashFns(items, optimalSize)
	if hashFns != expectedHashFns {
		t.Errorf("Invalid optimal hash fns. Got %v expected %v", hashFns, expectedHashFns)
	}
}

func TestEstimateOptimalBloomSize(t *testing.T) {
	items, fr := int64(3000000), 0.01
	expectedOptimalSize := 28755176
	expectedHashFns := int64(7)

	optimalSize, hashFns := estimateOptimalBloomSize(items, fr)
	if optimalSize != int64(expectedOptimalSize) {
		t.Errorf("Invalid optimal size. Got %v Exptected %v", optimalSize, expectedOptimalSize)
	}
	if hashFns != expectedHashFns {
		t.Errorf("Invalid optimal hash fns. Got %v expected %v", hashFns, expectedHashFns)
	}
}
