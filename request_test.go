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
}

func TestSearchByName(t *testing.T) {
	time.Sleep(3 * time.Second)
	nq := abnlookup.NameQuery{
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

	people, err := client.SearchByName("Glen", &nq)
	if err != nil {
		t.Error(err)
	}

	if len(people) == 0 {
		t.Log("Length of people is 0")
		t.FailNow()
	}

	for _, person := range people {
		switch {
		case person.LegalName != nil:
			t.Logf("Person legal name: %s", person.LegalName.FullName)
		case person.BusinessName != nil:
			t.Logf("Person business name: %s", person.BusinessName.OrganisationName)
		case person.MainName != nil:
			t.Logf("Person main organisation name: %s", person.MainName.OrganisationName)
		default:
			t.FailNow()
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

	abnList, err := client.SearchByABNStatus(&abnStatusQuery)
	if err != nil {
		t.Error(err)
	}

	if len(abnList) == 0 {
		t.Log("length of abn list is 0")
		t.Fail()
	}
}
