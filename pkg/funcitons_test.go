package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNds(t *testing.T) {

	au := &Auto{cost: 40.}
	nds := nds(au)

	assert.Equal(t, au.cost, 40.)
	assert.Equal(t, nds, 8.)
}

func TestConcatF(t *testing.T) {
	assert.Equal(t, concatF("aaaaa", "bbbbbb"), "aaaaa bbbbbb")
}

// BenchmarkIntMin-12    	10 609 413	       111.4 ns/op	      48 B/op	       3 allocs/op
func BenchmarkIntMin(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatF("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkPlus-12    	40 900 516	        28.52 ns/op	      16 B/op	       1 allocs/op
func BenchmarkPlus(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatPlus("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkJoin-12    	36 596 022	        30.95 ns/op	      16 B/op	       1 allocs/op
func BenchmarkJoin(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatJoin("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkBuider-12    	43 688 440	        25.64 ns/op	      16 B/op	       1 allocs/op
func BenchmarkBuider(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatBuilder("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkPool-12    	54 303 804	        22.16 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPool(b *testing.B) {
	s := []byte{}

	for b.Loop() {
		s = concatPool("aaaaa", "bbbbbb")
	}

	_ = s
}
