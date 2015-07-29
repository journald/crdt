package pncounter_test

import (
	"testing"

	"github.com/journald/crdt/pncounter"
)

func TestPNCounter(t *testing.T) {
	c := pncounter.New()

	c.Increment("foo", 4)
	c.Increment("bar", 1)
	c.Decrement("foo", 2)

	expectedValue := int64(3)

	if c.Value() != expectedValue {
		t.Errorf("Counter value missmatch:\nExpected: %v\nGot: %v", expectedValue, c.Value())
	}
}

func TestMergeTwoPNCounter(t *testing.T) {
	c1 := pncounter.New()

	c1.Increment("foo", 1)
	c1.Increment("bar", 1)
	c1.Increment("baz", 1)

	c2 := pncounter.New()
	c2.Increment("foo", 2)
	c2.Increment("bar", 0)
	c2.Increment("baz", 1)

	c3 := c1.Merge(c2)

	expectedValue := int64(4)

	if c3.Value() != expectedValue {
		t.Errorf("Counter value missmatch:\nExpected: %v\nGot: %v", expectedValue, c3.Value())
	}
}

func TestMergePNCounterIsIdempotent(t *testing.T) {
	c1 := pncounter.New()

	c1.Increment("foo", 1)
	c1.Increment("bar", 1)
	c1.Increment("baz", 1)

	c2 := pncounter.New()
	c2.Increment("foo", 2)
	c2.Increment("bar", 0)
	c2.Increment("baz", 1)

	c3 := c1.Merge(c2)
	c4 := c1.Merge(c2).Merge(c2)

	if c3.Value() != c4.Value() {
		t.Errorf("Counter value missmatch:\nExpected: %v\nGot: %v", c3.Value(), c4.Value())
	}
}

func TestMergePNCounterIsAssociative(t *testing.T) {
	c1 := pncounter.New()

	c1.Increment("foo", 1)
	c1.Increment("bar", 1)
	c1.Increment("baz", 1)

	c2 := pncounter.New()
	c2.Increment("foo", 2)
	c2.Increment("bar", 0)
	c2.Increment("baz", 1)

	c3 := pncounter.New()
	c3.Increment("foo", 2)
	c3.Increment("bar", 0)
	c3.Increment("baz", 1)

	c4 := (c1.Merge(c2)).Merge(c3)
	c5 := c1.Merge(c2.Merge(c3))

	if c4.Value() != c5.Value() {
		t.Errorf("Counter value missmatch:\nExpected: %v\nGot: %v", c4.Value(), c5.Value())
	}
}

func TestMergePNCounterIsCommutative(t *testing.T) {
	c1 := pncounter.New()

	c1.Increment("foo", 1)
	c1.Increment("bar", 1)
	c1.Increment("baz", 1)

	c2 := pncounter.New()
	c2.Increment("foo", 2)
	c2.Increment("bar", 0)
	c2.Increment("baz", 1)

	c3 := c1.Merge(c2)
	c4 := c2.Merge(c1)

	if c3.Value() != c4.Value() {
		t.Errorf("Counter value missmatch:\nExpected: %v\nGot: %v", c3.Value(), c4.Value())
	}
}
