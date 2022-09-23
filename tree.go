package graph

type Trees[K comparable, T any] []struct {
	root *Node[K, T]
}

type Node[K comparable, T any] struct {
	key    K
	value  T
	Childs []Node[K, T]
}

func ToTrees[K comparable, T any](g Graph[K, T]) Trees[K, T] {
	// 从任一点开始，获得所有可能的树

	return nil
}

func toTrees[K comparable, T any](g Graph[K, T], source K) Trees[K, T] {

	return nil
}
