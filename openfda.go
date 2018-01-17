package openfda

type OpenFDA struct {
	SPLID    []string `json:"spl_id,omitempty" bson:"spl_id,omitempty"`
	SPLSetID []string `json:"spl_set_id,omitempty" bson:"spl_set_id,omitempty"`

	ApplicationNumber []string `json:"application_number,omitempty" bson:"application_number,omitempty"`
	ManufacturerName  []string `json:"manufacturer_name,omitempty" bson:"manufacturer_name,omitempty"`

	BrandNameList []string `json:"brand_name,omitempty" bson:"brand_name_list,omitempty"`
	BrandName     string   `json:"brand,omitempty" bson:"brand_name,omitempty"`

	GenericNameList []string `json:"generic_name,omitempty" bson:"generic_name_list,omitempty"`
	GenericName     string   `json:"generic,omitempty" bson:"generic_name,omitempty"`

	ProductNDCList []string `json:"product_ndc,omitempty" bson:"product_ndc_list,omitempty"`
	ProductNDC     string   `json:"ndc,omitempty" bson:"ndc,omitempty"`

	NUI           []string `json:"nui,omitempty" bson:"nui,omitempty"`
	PackageNDC    []string `json:"package_ndc,omitempty" bson:"package_ndc,omitempty"`
	PharmClassCS  []string `json:"pharm_class_cs,omitempty" bson:"pharm_class_cs,omitempty"`
	PharmClassEPC []string `json:"pharm_class_epc,omitempty" bson:"pharm_class_epc,omitempty"`
	PharmClassMOA []string `json:"pharm_class_moa,omitempty" bson:"pharm_class_moa,omitempty"`
	PharmClassPE  []string `json:"pharm_class_pe,omitempty" bson:"pharm_class_pe,omitempty"`
	ProductType   []string `json:"product_type,omitempty" bson:"product_type,omitempty"`
	Route         []string `json:"route,omitempty" bson:"route,omitempty"`
	RxCUI         []string `json:"rxcui,omitempty" bson:"rxcui,omitempty"`
	SubstanceName []string `json:"substance_name,omitempty" bson:"substance_name,omitempty"`
	UNII          []string `json:"unii,omitempty" bson:"unii,omitempty"`
	UPC           []string `json:"upc,omitempty" bson:"upc,omitempty"`

	// Undocumented Fields Found In Responses
	IsOriginalPackager []bool `json:"is_original_packager,omitempty" bson:"is_original_packager,omitempty"`
}
