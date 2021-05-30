package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/kambahr/go-mathsets"
)

// generateTestList creates an array of string consisting of test-data.
// The test-data is comprised of random data; in real scenario, this
// could be row from a database, conents of a file, a JSON object,...
func generateTestList(length uint) []string {

	var a []string
	var i uint

	for i = 0; i < length; i++ {
		// This is some random text for testing purposes.
		// Real-life data could be one row of columns,
		// user name + some attributes,...
		s := fmt.Sprintf("Spring, Summer, Fall, Winter;Guitar is the song %d", i)
		b := []byte(s)

		// If hex is not used, there will still be a root, but there
		// will be no leaves. Uncomment the following to see results.
		// usrNamePlain := fmt.Sprintf("Spring, Summer, Fall, Winter;Guitar is the song%d", i)
		// a = append(a, usrNamePlain)
		// continue

		hexStr := hex.EncodeToString(b)
		a = append(a, hexStr)
	}

	return a
}

// demoVerifyWithBicoinBlock uses a public bitcoin record to demonstrate
// a verifcation of the GetMerkleRoot().
func demoVerifyWithBicoinBlock() {

	// The following is from a real-life transaction from one the bitcoin
	// (block-chain) blocks:
	// block id:
	//    000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506
	// published merkle root:
	//    f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766

	publisedMR := "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766"

	// You can alter any of the following values sto see that the merkel root
	// will not match the above.

	// Transactions
	var tx = []string{
		"8c14f0db3df150123e6f3dbbf30f8b955a8249b62ac1d1ff16284aefa3d06d87",
		"fff2525b8931402dd09222c50775608f75787bd2b87e56995a7bdd30f79702c4",
		"6359f0868171b1d194cbee1af2f16ea598ae8fad666d9b012c8ed2b79a236ec4",
		"e9a66845e05d5abc0ad04ec80f774a7e585c6e8db975962d069a522137b80c1d",
	}

	fmt.Println("Transactions from Bitcoin Block:")
	fmt.Println(strings.Repeat(" ", 12), "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506")
	fmt.Println("Publised Merkel Root:")
	fmt.Println(strings.Repeat(" ", 12), publisedMR)

	mr, err := mathsets.GetMerkleRoot(tx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Calculated Merkel Root:\n\t", strings.Repeat(" ", 3), mr)

	if mr != publisedMR {
		fmt.Println("Failed...")
	}
	fmt.Println(strings.Repeat("-", 40))
}

// demoDataChanged demostrates how you can tell if a data-set
// has been modified without examining a specific range or
// set of records.
func demoDataChanged() {
	// Get test data
	list := generateTestList(10000)

	// Get the merkle root only.
	mr1, err := mathsets.GetMerkleRoot(list)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Alter data
	list[(len(list)/2)-1] = hex.EncodeToString([]byte("some updated data"))
	// Get the merkle root only.
	mr2, err := mathsets.GetMerkleRoot(list)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("MR 1:", strings.Repeat(" ", 6), mr1)
	fmt.Println("MR 2:", strings.Repeat(" ", 6), mr2)

	if mr1 != mr2 {
		fmt.Println(strings.Repeat(" ", 12), "---test-data has been modified---")
	} else {
		fmt.Println("list unchanged.")
	}
	fmt.Println(strings.Repeat("-", 40))
}

// getLongString generates  long string.
func getLongString(size uint) string {
	var a []string
	var i uint
	for i = 0; i < size; i++ {
		a = append(a, strings.Repeat(fmt.Sprintf("Gutiar is the song _%d", i), 40))
	}

	return strings.Join(a, "\n")
}

// demoHashVsMerkelRoot compares the time it takes to created
// a has vs merkle root.
func demoHashVsMerkelRoot() {

	lngStr := getLongString(500000)

	b := []byte(lngStr)

	start := time.Now()
	h := mathsets.Reversebytes(mathsets.Hash256Twice(b))
	hash := fmt.Sprintf("%x", h)
	tookh := time.Since(start)

	fmt.Println("Hash:", strings.Repeat(" ", 6), hash)
	fmt.Println("Took:", strings.Repeat(" ", 6), tookh)

	lines := strings.Split(lngStr, "\n")
	start = time.Now()
	mr, _ := mathsets.GetMerkleRoot(lines)
	tookmr := time.Since(start)
	fmt.Println("\nMerkle Root:", mr)
	fmt.Println("Took:", strings.Repeat(" ", 6), tookmr)
}

func demoGetMerkleTree() {
	// Get the tree
	list := generateTestList(50000)
	mt, err := mathsets.GetMerkleTree(list)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Look at the first element in the tree
	fmt.Println("Tree Root : ", mt.Root)
	fmt.Println("\n------ First Branch ------")
	fmt.Println("Root:       ", mt.Branches[0].Root)
	fmt.Println("Left Leaf:  ", mt.Branches[0].LeafLeft)
	fmt.Println("Right Leaf: ", mt.Branches[0].LeafRight)
}

func main() {

	demoVerifyWithBicoinBlock()

	demoDataChanged()

	// Compare performance bweteen Hash and MekleRoot
	demoHashVsMerkelRoot()

	demoGetMerkleTree()

	// Concepts
	// Syncing Sets - a brief summary
	//    Set Changed?
	//      1. Take the Merkel Root of the destination.
	//      2. Take the Merkel Root of the soure.
	//      3. Compare the root values.
	//
	//    Check a specific value?
	//      1. Iternate thru the dest.
	//      2. Find the target and compare.
	//
	//    If indexed
	//      1. Select by array index and compare.
	//

	return
}
