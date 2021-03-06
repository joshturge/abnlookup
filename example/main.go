package main

import (
	"fmt"
	"os"

	"github.com/joshturge/abnlookup"
)

func main() {
	client, err := abnlookup.NewClient(os.Getenv("AUTH_GUID"), abnlookup.LogDebug)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	abnResults, err := client.SearchByABN("49 093 669 660", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("ACN Number: %s\n", abnResults.ASICNumber)

	for _, ABN := range abnResults.ABN {
		fmt.Println(ABN.ReplacedFrom.String())
	}

	if abnlookup.ValidateACN(abnResults.ASICNumber) {
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

	fmt.Println(personResults[0].ABN.Value)

	abnStatusQuery := abnlookup.ABNStatusQuery{
		Postcode:                   "4159",
		ActiveABNsOnly:             true,
		CurrentGSTRegistrationOnly: true,
		EntityTypeCode:             "PRV",
	}

	abnListResults, err := client.SearchByABNStatus(abnStatusQuery)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("ABNList result: %s\n", abnListResults[0])
}
