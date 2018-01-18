package openfda

import (
	"fmt"
	"regexp"
)

func (label DrugLabel) String() string {
	return fmt.Sprintf("%s \"%s\" \"%s\"", label.SetID, label.OpenFDA.BrandName, label.OpenFDA.GenericName)
}

var labelsIDRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

// IsValidLabelID checks to ensure that a string is a valid id
func IsValidLabelID(id string) bool {
	return labelsIDRegex.MatchString(id)
}
