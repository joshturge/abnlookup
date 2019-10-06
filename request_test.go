package abnlookup_test

import (
	"os"
	"testing"
	"time"

	"github.com/joshturge/abnlookup"
)

var client *abnlookup.Client

func TestClient(t *testing.T) {
	var err error
	client, err = abnlookup.NewClient(os.Getenv("AUTH_GUID"), abnlookup.LogDebug)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchByABN(t *testing.T) {
	time.Sleep(3 * time.Second)
	business, err := client.SearchByABN("12 586 695 715", true)
	if err != nil {
		t.Error(err)
	}

	if business.Type.Code == "" {
		t.Fail()
	}

	if business.ABN[0].Value == "" {
		t.FailNow()
	}

	t.Logf("Current state code: %s", business.MainBusinessPhysicalAddress[0].StateCode)
}

func TestSearchByASIC(t *testing.T) {
	time.Sleep(3 * time.Second)
	business, err := client.SearchByASIC("000 000 019", true)
	if err != nil {
		t.Error(err)
	}

	if business.Type.Code == "" {
		t.Fail()
	}

	if business.ABN[0].Value == "" {
		t.FailNow()
	}

	t.Logf("Current state code: %s", business.MainBusinessPhysicalAddress[0].StateCode)
}

func TestSearchByName(t *testing.T) {
	time.Sleep(3 * time.Second)
	nq := abnlookup.NameQuery{
		Name:           "Glen",
		Postcode:       "4156",
		LegalName:      true,
		TradingName:    true,
		BusinessName:   true,
		ActiveABNsOnly: false,
		StateCodes:     []string{"QLD"},
		SearchWidth:    "typical",
		MinimumScore:   20,
		MaxResults:     5,
	}

	people, err := client.SearchByName(nq)
	if err != nil {
		t.Error(err)
	}

	if len(people) == 0 {
		t.FailNow()
	}

	for _, person := range people {
		if person.LegalName.FullName == "" {
			if person.BusinessOrginisation.Name == "" {
				if person.MainName.Name == "" {
					t.FailNow()
				}
			}
		}
	}
}

// This test will test the filterSearch Method which is used by all filter methods including the one below
func TestSearchByABNStatus(t *testing.T) {
	abnStatusQuery := abnlookup.ABNStatusQuery{
		Postcode:                   "4156",
		ActiveABNsOnly:             false,
		CurrentGSTRegistrationOnly: false,
		EntityTypeCode:             "PUB",
	}

	abnList, err := client.SearchByABNStatus(abnStatusQuery)
	if err != nil {
		t.Error(err)
	}

	if len(abnList) == 0 {
		t.Log("length of abn list is 0")
		t.Fail()
	}

	t.Logf("abn list length %d", len(abnList))
}
