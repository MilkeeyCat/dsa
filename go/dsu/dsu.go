// A really simple implementation of a disjoint-set data structure.
//
// This puppy has path compression & union by rank optimizations(yoinked from
// wikipedia).
package dsu

type Key int

type node[T any] struct {
	value  T
	parent Key
	rank   uint
}

type DisjointSet[T any] struct {
	nodes []node[T]
}

func (ds *DisjointSet[T]) Add(value T) Key {
	key := Key(len(ds.nodes))

	ds.nodes = append(ds.nodes, node[T]{value: value, parent: key})

	return key
}

func (ds *DisjointSet[T]) Find(key Key) Key {
	node := &ds.nodes[key]

	if node.parent == key {
		return key
	}

	node.parent = ds.Find(node.parent)

	return node.parent
}

func (ds *DisjointSet[T]) Union(keyA, keyB Key) {
	rootA := &ds.nodes[ds.Find(keyA)]
	rootB := &ds.nodes[ds.Find(keyB)]

	if rootA == rootB {
		return
	}

	if rootA.rank < rootB.rank {
		rootA, rootB = rootB, rootA
	}

	rootB.parent = rootA.parent

	if rootA.rank == rootB.rank {
		rootA.rank++
	}
}

func (ds *DisjointSet[T]) Value(key Key) T {
	return ds.nodes[key].value
}
