# CRDT

CRDT stand for [Conflict-free Replicated Data Types] are a kind of data
structures that you can use to achieve consensus when implementing _distributed
systems_.

This kind of data structures can be updated on each node of the system and then
merged later without hitting any conflict.

This package implements common CRDTs in Go.
