package main

import (
	"fmt"
	"os"

	abnlookup "github.com/joshturge/ABN-Lookup-Go"
)

func main() {
	client, err := abnlookup.NewClient("GUID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results, err := client.SearchByABN("78 345 431 247", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Charity Desc: %s\n", results.Response.BusinessEntity.CharityType.CharityTypeDescription)
}
