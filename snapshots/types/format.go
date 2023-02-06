package types

// CurrentFormat is the currently used format for snapshots. Snapshots using the same format
// must be identical across all nodes for a given height, so this must be bumped when the binary
// snapshot output changes.

// Format #1 - Assumes all stores in multistore have version == height
// Format #2 - Transmits actual store version within multistore
const CurrentFormat uint32 = 2
