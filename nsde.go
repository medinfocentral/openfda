package openfda

type NSElement struct {
	MarketingStartDate string `json:"marketing_start_date"`
	MarketingEndDate   string `json:"marketing_end_date"`
	PackageNDC         string `json:"package_ndc"`
	MarketingCategory  string `json:"marketing_category"`
	ProductType        string `json:"product_type"`

	// ProprietaryName             string `json:"proprietary_name"`
	// ApplicationNumberOrCitation string `json:"application_number_or_citation"`
	// BillingUnit                 string `json:"billing_unit"`
	// PackageNDC11                string `json:"package_ndc11"`
	// DosageForm                  string `json:"dosage_form"`
}
