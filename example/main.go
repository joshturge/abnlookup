package main

import (
	"fmt"
	"os"

	abnlookup "github.com/joshturge/ABN-Lookup-Go"
)

func main() {
	client, err := abnlookup.NewClient(os.Getenv("AUTH_GUID"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results, err := client.SearchByABN("49 093 669 660", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("ACN Number: %s\n", results.Response.BusinessEntity.ASICNumber)

	if abnlookup.ValidateACN(results.Response.BusinessEntity.ASICNumber) {
		fmt.Println("ACN is valid")
	} else {
		fmt.Println("ACN is not valid")
	}
}
