package mathsets

// Branch holds the root and two leaves
// of a merkel tree node.
type Branch struct {
	Root      string
	LeafLeft  string
	LeafRight string
}

// MerkleTree is a verbose Merkle tree with the
// root, branches and leaves.
type MerkleTree struct {
	Root     string
	Branches []Branch
	Leaves   []string
}
