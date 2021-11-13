package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	ipfsclient "kaleido-backend/pkg/infrastructure/ipfs"
	"kaleido-backend/pkg/mediaaccount"
	"os"
)

var (
	notempty = func(input string) bool {
		return input != ""
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := mediaaccount.Application{}
	//add := text(scanner, "address of account", func(input string) bool {
	//	eoa := common.HexToAddress(input)
	//	return eoa != common.Address{}
	//})
	//input.Account = common.HexToAddress(add)
	input.Name = text(scanner, "media name", notempty)
	input.Description = text(scanner, "description", notempty)
	input.URL = text(scanner, "url", notempty)
	input.MailAddress = text(scanner, "mailAddress", notempty)
	b, _ := json.Marshal(&input)
	fmt.Println("uploading json")
	fmt.Println(string(b))
	cid, err := ipfsclient.NewIpfsClient().UploadMedia(context.Background(), input)
	if err != nil {
		fmt.Println("unexpected error occured")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("added " + cid)
}

func text(sc *bufio.Scanner, inputDescription string, matchFunc func(input string) bool) string {
	fmt.Println(inputDescription)
	for sc.Scan() {
		input := sc.Text()
		if matchFunc(input) {
			return input
		}
		fmt.Println("enter " + inputDescription + " again.")
	}
	return ""
}
