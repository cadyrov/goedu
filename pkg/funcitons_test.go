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

// BenchmarkIntMin-12    	 5382906	       222.6 ns/op	      48 B/op	       3 allocs/op
func BenchmarkIntMin(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatF("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkPlus-12  16616991	        66.78 ns/op	      16 B/op	       1 allocs/op
func BenchmarkPlus(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatPlus("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkJoin-12    	17572365	        62.85 ns/op	      16 B/op	       1 allocs/op
func BenchmarkJoin(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatJoin("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkBuider-12    	14192539	        81.52 ns/op	      24 B/op	       2 allocs/op
// BenchmarkBuider-12    	21609860	        51.40 ns/op	      16 B/op	       1 allocs/op
func BenchmarkBuider(b *testing.B) {
	s := ""

	for b.Loop() {
		s = concatBuilder("aaaaa", "bbbbbb")
	}

	_ = s
}

// BenchmarkPool-12    	27338060	        43.80 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPool(b *testing.B) {
	s := []byte{}

	for b.Loop() {
		s = concatPool("aaaaa", "bbbbbb")
	}

	_ = s
}
