package main

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("api.thegraph.com/ipfs/")
	files := []string{"mediametadata", "periodmetadata", "spacemetadata"}
	for _, f := range files {
		file, err := os.Open("test/" + f + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
		defer file.Close()
		cid, err := sh.Add(file)
		fmt.Printf("added %s", cid)
		err = sh.Pin(cid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
	}

}
