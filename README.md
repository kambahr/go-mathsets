# A Go implementation of the Merkle Tree

## Get a Merkle Root or browse an entire merkle tree of a data-set

### Functions
There are two primary functions. One gets the entire tree and other only returns the merkle root (faster).

``` go
// GetMerkleTree returns the entire tree; with leaves, all the branches and the root of the tree.
func GetMerkleTree(nodes []string) (MerkleTree, error)

// GetMerkleRoot returns the Merkle root of an array of string.
func GetMerkleRoot(nodes []string) (string, error)
```

### Demo App
Run the test app to see a demo of the usage.

#### Usage example from a Bitcoin Block 
Verify the calculation of the merkle root by comparing to the published merkle root of block 
``` bash
000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506
```
of the Bitcoin Block Chain.


#### Detect changes in a set
Take a merkle root of a set; change a value in the set; and then tell if the set has changed without having to inspect all data -- demonstrated by demoDataChanged().

#### Performance Comparison with Hash

demoHashVsMerkelRoot() compares the calculation time between merkle root and just getting a hash.


#### References
``` bash
Greg Walker 2020, Merkle Root, Learn me a bitcoin, viewed 30 May 2021, <https://learnmeabitcoin.com/technical/merkle-root>
Teemu Kanstrén 2017, Merkle Trees: Concepts and Use Cases, Coinmonks, viewed 30 May 2021, <https://medium.com/coinmonks/merkle-trees-concepts-and-use-cases-5da873702318>
Chumbley, A., Moore, K. and Khim, J.2016, Merkle Tree, BRILLIANT, viewed 30 May 2021, <https://brilliant.org/wiki/merkle-tree/>
```

#### To run the sample
Clone and run the sample inside your work/src directory or:

- Start a shell window
- cd &lt;to any directory (other than $GOPATH)&gt;
- git clone https:&#47;&#47;github.com&#47;kambahr/go-mathsets.git && cd go-mathsets/test
- go mod init go-mathsets/test
- go mod tidy
- go mod vendor
- go build -o mkdemo && ./mkdemo