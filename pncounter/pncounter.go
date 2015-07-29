package pncounter

import "fmt"

type PNCounter interface {
	Increment(string, int64)
	Decrement(string, int64)
	Value() int64
}

type counter struct {
	positives map[string]int64
	negatives map[string]int64
}

func (c counter) Increment(replicaId string, n int64) {
	c.positives[replicaId] += n
}

func (c counter) Decrement(replicaId string, n int64) {
	c.negatives[replicaId] += n
}

func (c counter) Value() int64 {
	var positives int64
	for _, v := range c.positives {
		positives += v
	}

	var negatives int64
	for _, v := range c.negatives {
		negatives += v
	}

	return positives - negatives
}

func (c counter) Inspect() string {
	return fmt.Sprintf("positives: %v, negatives: %v", c.positives, c.negatives)
}

func New() counter {
	return counter{
		positives: make(map[string]int64),
		negatives: make(map[string]int64),
	}
}

func (c counter) Merge(other counter) counter {
	mergedCounter := New()

	for replicaId, value := range other.negatives {
		mergedCounter.negatives[replicaId] = max(c.negatives[replicaId], value)
	}

	for replicaId, value := range other.positives {
		mergedCounter.positives[replicaId] = max(c.positives[replicaId], value)
	}

	return mergedCounter
}

func max(a int64, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
