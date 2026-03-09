package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func findDifferences(oldHashes, newHashes [][32]byte) {
    changedCount := 0
    totalChunks := len(newHashes)

    for i := 0; i < totalChunks; i++ {
        if i >= len(oldHashes) || oldHashes[i] != newHashes[i] {
            changedCount++
            fmt.Printf("Chunk %d has changed!\n", i)
        }
    }

    fmt.Printf("Change chunks: %d\n", changedCount)
    fmt.Printf("Out of: %d\n", totalChunks)

    percentage := float64(changedCount) / float64(totalChunks) * 100
    fmt.Printf("Change rate: %.2f%%\n", percentage)
}

func main() {
	file1, err1 := os.Open("bootstrap_log.txt")
	file2, err2 := os.Open("bootstrap_log2.txt")
	if err1 != nil || err2 != nil {
		panic("Could not open files")
	}
	defer file1.Close()
	defer file2.Close()

	chunkSize := 0.3
	buffer := make([]byte, chunkSize)

	hashes1 := [][32]byte{}
	hashes2 := [][32]byte{}

	for {
		n, err := file1.Read(buffer)
		if err == io.EOF { break }
		hash := sha256.Sum256(buffer[:n])
		hashes1 = append(hashes1, hash)
	}

	for {
		n, err := file2.Read(buffer)
		if err == io.EOF { break }
		hash := sha256.Sum256(buffer[:n])
		hashes2 = append(hashes2, hash)
	}

	findDifferences(hashes1, hashes2)
}
