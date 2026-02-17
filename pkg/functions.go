package pkg

import (
	"fmt"
	"strings"
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
