package dsu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/dsu"
)

func TestFind(t *testing.T) {
	dsu := new(dsu.DisjointSet[struct{}])
	key := dsu.Add(struct{}{})

	assert.Equal(t, dsu.Find(key), key)
}

func TestUnion(t *testing.T) {
	dsu := new(dsu.DisjointSet[uint])
	a := dsu.Add(69)
	b := dsu.Add(420)
	c := dsu.Add(96)

	dsu.Union(a, b)

	assert.Equal(t, dsu.Find(a), a)
	assert.Equal(t, dsu.Find(b), a)

	dsu.Union(b, c)

	assert.Equal(t, dsu.Find(a), a)
	assert.Equal(t, dsu.Find(b), a)
	assert.Equal(t, dsu.Find(c), a)
}

func TestValue(t *testing.T) {
	dsu := new(dsu.DisjointSet[string])
	foo := dsu.Add("foo")
	bar := dsu.Add("bar")

	assert.Equal(t, dsu.Value(foo), "foo")
	assert.Equal(t, dsu.Value(bar), "bar")
}
