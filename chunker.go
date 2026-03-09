package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"encoding/json"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
	file, err := os.Open("bootstrap_log.txt")
	check(err)
	defer file.Close()

	os.MkdirAll("chunks", 0755)

	chunkSize := 1024 * 4
	buffer := make([]byte, chunkSize)
	counter := 0

	hashes := []string{}

	for {
		n, err := file.Read(buffer)
		if err == io.EOF { break }

		hash := sha256.Sum256(buffer[:n])
		fileName := fmt.Sprintf("chunks/chunk.%x", hash)

        os.WriteFile(fileName, buffer[:n], 0644)

        fmt.Printf("saved: %s (%d Bytes)\n", fileName, n)
        counter++

        hashString := fmt.Sprintf("%x", hash)
        hashes = append(hashes, hashString)
	}
	jsonData, err := json.MarshalIndent(hashes, "", "  ")
	check(err)

	os.WriteFile("manifest.json", jsonData, 0644)
	rebuild()
}

func rebuild() {
    outFile, _ := os.Create("rebuilt_log.txt")
    defer outFile.Close()

    jsonData, _ := os.ReadFile("manifest.json")

    var hashes []string

    err := json.Unmarshal(jsonData, &hashes)
    check(err)

    for _, hash := range hashes {
        fileName := fmt.Sprintf("chunks/chunk.%s", hash)
        data, _ := os.ReadFile(fileName)
        outFile.Write(data)
    }
}
