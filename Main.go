package main

import (
    "bufio"
    "encoding/hex"
    "fmt"
    "crypto/sha256"
    "os"
)

func main() {
    // Open the input file
    inputFile, err := os.Open("transactions.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer inputFile.Close()

    // Read the transactions into an array
    var transactions []string
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        transactions = append(transactions, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }

    // Compute the Merkle tree root
    merkleRoot := computeMerkleRoot(transactions)
    fmt.Printf("Merkle root: %x\n", merkleRoot)
}

func computeMerkleRoot(transactions []string) []byte {
    // Convert the transactions from hex to byte arrays
    var byteTransactions [][]byte
    for _, tx := range transactions {
        byteTx, _ := hex.DecodeString(tx)
        byteTransactions = append(byteTransactions, byteTx)
    }

    // Compute the Merkle tree
    tree := make([][]byte, len(byteTransactions))
    for i, tx := range byteTransactions {
        tree[i] = hash(tx)
    }
    for len(tree) > 1 {
        if len(tree)%2 != 0 {
            tree = append(tree, tree[len(tree)-1])
        }
        var level [][]byte
        for i := 0; i < len(tree); i += 2 {
            level = append(level, hash(append(tree[i], tree[i+1]...)))
        }
        tree = level
    }

    return tree[0]
}

func hash(data []byte) []byte {
    hasher := sha256.New()
    hasher.Write(data)
    return hasher.Sum(nil)
}
