package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/carterjones/awsinfo/info/elb"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: elbinfo <search-value>")
		os.Exit(1)
	}

	searchValue := os.Args[1]

	// Tell the SDK to load defaults from your ~/.aws/config file.
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	// Create a new session.
	sess, err := session.NewSession()
	panicIfErr(err)

	var infos elb.InfoSlice
	err = infos.Load(sess)
	panicIfErr(err)

	// Print the matches.
	numMatches := 0
	for _, v := range infos {
		if v.Matches(searchValue) {
			if numMatches > 0 {
				fmt.Println()
			}

			fmt.Print(v)
			numMatches++
		}
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
