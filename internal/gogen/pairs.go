package gogen

type namedPair[T any] struct {
	name   string
	schema T
}

type pairList[T any] []namedPair[T]

func (s pairList[T]) Len() int {
	return len(s)
}

func (s pairList[T]) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s pairList[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func pairsFromMap[T any](m map[string]T) pairList[T] {
	pairs := make(pairList[T], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, namedPair[T]{name: k, schema: v})
	}
	return pairs
}
