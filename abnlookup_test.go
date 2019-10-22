package abnlookup_test

import (
	"testing"

	"github.com/joshturge/abnlookup"
)

var (
	validABNs = []string{
		"34-241-177-887",
		"	30613 501 612",
		"//49 093 669?text660",
		"33 531 321 789",
	}
	invalidABNs = []string{
		"35 531 321 789",
		"7609355599",
		"E3 772 093 958",
		"65832766990",
	}
	validACNs = []string{
		"010 499 966",
		"004 085 616",
		"001 250 004",
		"005 749 986",
	}
	invalidACNs = []string{
		"093 555 993",
		"772 093 95",
		" 832 566 990",
		"001 000 000",
	}
)

func TestValidateABN(t *testing.T) {
	for _, validABN := range validABNs {
		if !abnlookup.ValidateABN(validABN) {
			t.Errorf("valid ABN: %s has been flagged as invalid", validABN)
		}
	}
	for _, invalidABN := range invalidABNs {
		if abnlookup.ValidateABN(invalidABN) {
			t.Errorf("invalid ABN: %s has been flagged as valid", invalidABN)
		}
	}
}

func TestValidateACN(t *testing.T) {
	for _, validACN := range validACNs {
		if !abnlookup.ValidateACN(validACN) {
			t.Errorf("valid ACN: %s has been flagged as invalid", validACN)
		}
	}
	for _, invalidACN := range invalidACNs {
		if abnlookup.ValidateACN(invalidACN) {
			t.Errorf("invalid ACN: %s has been flagged as valid", invalidACN)
		}
	}
}
