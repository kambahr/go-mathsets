// References
// Greg Walker 2020, Merkle Root , Learn me a bitcoin. Available at: https://learnmeabitcoin.com/technical/merkle-root (Accessed: 30 May 2021).
// Teemu Kanstrén 2017, Merkle Trees: Concepts and Use Cases, Coinmonks. Available at: https://medium.com/coinmonks/merkle-trees-concepts-and-use-cases-5da873702318 (Accessed: 30 May 2021).
// Chumbley, A., Moore, K. and Khim, J.2016, Merkle Tree, BRILLIANT. Available at: https://brilliant.org/wiki/merkle-tree/ (Accessed: 30 May 2021).

package mathsets

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func getNodeHash(txt string) string {
	itm1, _ := hex.DecodeString(txt)
	s := fmt.Sprintf("%x", Reversebytes(itm1))

	return s
}
func getNodeRoot(sLeft string, sRight string) string {

	// concat the hex from left and right
	s := fmt.Sprintf("%s%s", sLeft, sRight)
	paired, _ := hex.DecodeString(s)

	// Take the double hash and also reverse the result bytes.
	// Add the hash results to the array so that the loop will drill down the
	// tree to come up with a root.
	h := Reversebytes(Hash256Twice(paired))
	s = fmt.Sprintf("%x", h)

	return s
}

// GetBranchFromPlainText gets the hash of the leaevs and
// their root hash.
func GetBranchFromPlainText(txt1 string, txt2 string) Branch {
	var b Branch

	txtb := []byte(txt1)
	t1 := hex.EncodeToString([]byte(txtb))

	txtb = []byte(txt2)
	t2 := hex.EncodeToString([]byte(txtb))

	b.LeafLeft = getNodeHash(t1)
	b.LeafRight = getNodeHash(t2)

	b.Root = getNodeRoot(b.LeafLeft, b.LeafRight)

	return b
}

// GetMerkleTree returns the entire tree; with leaves, all
// the branches and the root of the tree.
func GetMerkleTree(nodes []string) (MerkleTree, error) {

	var mv MerkleTree

	if len(nodes) == 0 {
		// List is empty.
		return mv, nil
	}

	if len(nodes) == 1 {
		mv.Root = nodes[0]
		var mb Branch
		mb.LeafLeft = nodes[0]
		mv.Branches = []Branch{mb}
		// The Merkle root is the first item.
		return mv, nil
	}

	// Fix the size of the array. It must be even.
	if len(nodes)%2 != 0 {
		// Add the last item to the list.
		nodes = append(nodes, nodes[len(nodes)-2])
	}

	gotocountStart := len(nodes)
	gotocountMove := 0

	branchInx := 0
	halfArryCount := 0
	dirty := false

lblAgain:

	// Make sure the goto loop does not execute foreever.
	if gotocountMove > gotocountStart {
		return mv, errors.New("failed to get merkle root")
	}

	var result []string
	var b Branch

	// Split the array into pairs
	// This count reduces by half on every goto move.
	count := len(nodes)
	mod := count % 2
	if mod != 0 {
		// append the last one
		nodes = append(nodes, nodes[count-1])
	}
	count = len(nodes)

	if !dirty {
		dirty = true
		gotocountStart = count
		halfArryCount = count / 2
	}

	for i := 0; i < (count); i = i + 2 {

		// Make sure we have little endian (reverse the bytes).
		sLeft := getNodeHash(nodes[i])
		sRight := getNodeHash(nodes[i+1])

		if len(mv.Leaves) < gotocountStart {
			mv.Leaves = append(mv.Leaves, sLeft)
			mv.Leaves = append(mv.Leaves, sRight)
		}

		// Take the double hash and also reverse the result bytes.
		// Add the hash results to the array so that the loop will drill down the
		// tree to come up with a root.
		paired := getNodeRoot(sLeft, sRight)

		if len(mv.Branches) < halfArryCount {

			b.LeafLeft = sLeft
			b.LeafRight = sRight
			b.Root = paired
			mv.Branches = append(mv.Branches, b)
		}

		result = append(result, paired)
	}

	branchInx = branchInx + 1

	if len(result) == 1 {
		// Reached the root; return the hex;
		mv.Root = result[0]
		return mv, nil
	}

	nodes = result

	// Simulate recursion -- faster and less memory intensive.
	gotocountMove++

	goto lblAgain
}

// GetMerkleRoot returns the Merkle root of an array of string.
func GetMerkleRoot(nodes []string) (string, error) {

	if len(nodes) == 0 {
		// List is empty.
		return "", nil
	}

	if len(nodes) == 1 {
		// The Merkle root is the first item.
		return nodes[0], nil
	}

	// Fix the size of the array. It must be even.
	if len(nodes)%2 != 0 {
		// Add the last item to the list.
		nodes = append(nodes, nodes[len(nodes)-2])
	}

	gotocountStart := len(nodes)
	gotocountMove := 0

lblAgain:

	// Make sure the goto loop does not execute foreever.
	if gotocountMove > gotocountStart {
		return "", errors.New("failed to get merkle root")
	}

	var result []string

	// Split the array into pairs
	// This count reduces by in half on every goto move.
	count := len(nodes)
	mod := count % 2
	if mod != 0 {
		// append the last one
		nodes = append(nodes, nodes[count-1])
	}
	count = len(nodes)

	for i := 0; i < (count); i = i + 2 {
		// Make sure we have little endian (reverse the bytes).
		itm1, _ := hex.DecodeString(nodes[i])
		sLeft := fmt.Sprintf("%x", Reversebytes(itm1))

		itm2, _ := hex.DecodeString(nodes[i+1])
		sRight := fmt.Sprintf("%x", Reversebytes(itm2))

		// concat the hex from itm1 and itm2
		s := fmt.Sprintf("%s%s", sLeft, sRight)
		paired, _ := hex.DecodeString(s)

		// Take the double hash and also reverse the result bytes.
		// Add the hash results to the array so that the loop will drill down the
		// tree to come up with a root.
		h := Reversebytes(Hash256Twice(paired))
		s = fmt.Sprintf("%x", h)

		result = append(result, s)
	}

	if len(result) == 1 {
		// Reached the root; return the hex.
		// This is the same as result[0].
		h, _ := hex.DecodeString(result[0])
		return fmt.Sprintf("%x", h), nil
	}

	// Have not reached the root; keep going.
	nodes = result

	// Simulate recursion -- faster and less memory intensive.
	gotocountMove++

	goto lblAgain
}
