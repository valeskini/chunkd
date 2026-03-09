package main

import (
	"fmt"
	"io"
	"os"
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

//	hashes := [][32]byte{}

	for {
		n, err := file.Read(buffer)
		if err == io.EOF { break }

		fileName := fmt.Sprintf("chunks/chunk.%d", counter)

        os.WriteFile(fileName, buffer[:n], 0644)

        fmt.Printf("saved: %s (%d Bytes)\n", fileName, n)
        counter++

//		hash := sha256.Sum256(buffer[:n])
//		hashes = append(hashes, hash)
	}

	rebuild()
}

func rebuild() {
    outFile, _ := os.Create("rebuilt_log.txt")
    defer outFile.Close()

    counter := 0
    for {
        fileName := fmt.Sprintf("chunks/chunk.%d", counter)

        data, err := os.ReadFile(fileName)
        if err != nil {
            fmt.Println("done")
            break
        }

        outFile.Write(data)
        counter++
    }
}
