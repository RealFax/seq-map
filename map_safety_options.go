//go:build safety_map

package odmap

import "cmp"

type Option[K cmp.Ordered, V any] func(m *safetyMap[K, V])

func WithComparer[K cmp.Ordered, V any](comparer func(K, K) int) Option[K, V] {
	return func(m *safetyMap[K, V]) {
		m.compare = comparer
	}
}
