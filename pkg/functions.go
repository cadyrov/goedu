package pkg

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

type Auto struct {
	cost float64
}

func nds(auto *Auto) float64 {
	result := (auto.cost) / 5

	return result
}

func concatF(l, r string) string {
	return fmt.Sprintf("%s %s", l, r)
}

func concatPlus(l, r string) string {
	return l + " " + r
}

func concatJoin(l, r string) string {
	return strings.Join([]string{l, r}, " ")
}

func concatBuilder(l, r string) string {
	b := strings.Builder{}
	b.Grow(len(l) + len(r) + 1)

	b.WriteString(l)
	b.WriteByte(' ')
	b.WriteString(r)

	return b.String()
}

var syncPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	},
}

func concatPool(l, r string) []byte {
	b := syncPool.Get().(*bytes.Buffer)
	b.Reset()
	b.Grow(len(l) + len(r) + 1)

	b.WriteString(l)
	b.WriteByte(' ')
	b.WriteString(r)

	syncPool.Put(b)

	return b.Bytes()
}
