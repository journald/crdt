# CRDT

CRDT stand for [Commutative and/or Conflict-free Replicated Data Types](http://arxiv.org/pdf/0907.0929v1.pdf) are a kind of data
structures that you can use to achieve consensus when implementing _distributed
systems_.

This kind of data structures can be updated on each node of the system and then
merged later without hitting any conflict.

This package implements common CRDTs in Go.
