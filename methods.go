package openfda

import (
	"regexp"
)

var labelsIDRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

// IsValidLabelID checks to ensure that a string is a valid id
func IsValidLabelID(id string) bool {
	return labelsIDRegex.MatchString(id)
}
