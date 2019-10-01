package main

import (
	"fmt"
	"os"

	"github.com/joshturge/abnlookup"
)

func main() {
	client, err := abnlookup.NewClient(os.Getenv("AUTH_GUID"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	abnResults, err := client.SearchByABN("49 093 669 660", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("ACN Number: %s\n", abnResults.BusinessEntityResponse.BusinessEntity.ASICNumber)

	if abnlookup.ValidateACN(abnResults.BusinessEntityResponse.BusinessEntity.ASICNumber) {
		fmt.Println("ACN is valid")
	} else {
		fmt.Println("ACN is not valid")
	}

	nameQuery := abnlookup.NameQuery{
		Name:           "Glen",
		Postcode:       "4156",
		LegalName:      true,
		TradingName:    true,
		BusinessName:   true,
		ActiveABNsOnly: false,
		StateCodes:     []string{"QLD"},
		SearchWidth:    "typical",
		MinimumScore:   20,
		MaxResults:     10,
	}

	personResults, err := client.SearchByName(nameQuery)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(personResults.PersonEntityResponse.SearchResults.PersonEntity[0].ABN.IdentiferValue)
}
