package abnlookup

import (
	"fmt"
	"regexp"
)

// ABN and ACN weights used to calculate weighting sums
var (
	acnWeights = []int{8, 7, 6, 5, 4, 3, 2, 1}
	abnWeights = []int{10, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
)

// ValidateABN will check if an ABN is valid.
// For more information on how this works you can
// refer to: https://abr.business.gov.au/Help/AbnFormat
func ValidateABN(abn string) bool {
	// Remove all non-integer characters from the ABN
	abn = cleanNumericString(abn)

	if abn[:1] == "0" {
		return false
	}
	// Subtract 1 from the first check digit of the abn
	abnByte := []byte(abn)
	abnByte[0]--
	abn = string(abnByte)

	if abn == "" || len(abn) != 11 {
		return false
	}

	// If the weightedSum is a multiple of 89 then it's a valid ABN
	return ((calcWeightingSum("ABN", abn) % 89) == 0)
}

// ValidateACN will check if an ACN is valid.
// For more details refer to: https://asic.gov.au/for-business/registering-a-company/steps-to-register-a-company/australian-company-numbers/australian-company-number-digit-check/
func ValidateACN(acn string) bool {
	// Remove all non-integer characters from the ACN
	acn = cleanNumericString(acn)

	if acn == "" || len(acn) != 9 {
		return false
	}

	acnWeightingSum := calcWeightingSum("ACN", acn)

	// Convert the given check digit to int
	givenCheck := int(acn[8]) - 48

	// Calculate the check digit
	calcCheck := (10 - (acnWeightingSum % 10)) % 10

	return calcCheck == givenCheck

}

// cleanNumericString will remove all non-integer characters from a string
func cleanNumericString(str string) string {
	// Make a regexp that only allows numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		// Something went horribly wrong if we reach an error here
		panic(fmt.Sprintf("couldn't compile regex: %s\n", err.Error()))
	}

	// Replace all non-integer characters with an empty string
	return reg.ReplaceAllString(str, "")
}

// calcWeightingSum will calculate the weighting sum for either an ABN or ACN
func calcWeightingSum(validationType string, str string) int {
	var weights []int
	switch validationType {
	case "ACN":
		weights = acnWeights
	case "ABN":
		weights = abnWeights
	}

	var weightingSum int
	var num int
	for i := 0; i <= len(weights)-1; i++ {
		num = int(str[i]) - 48
		weightingSum += num * weights[i]
	}

	return weightingSum
}
