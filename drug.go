package openfda

import "github.com/globalsign/mgo/bson"

// Drug specifies openFDA fields datastructure as described on open.fda.gov
// with a few minor changes noted below
type Drug struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`

	SPLID    []string `json:"spl_id,omitempty" bson:"spl_id,omitempty"`
	SPLSetID []string `json:"spl_set_id,omitempty" bson:"spl_set_id,omitempty"`

	ApplicationNumber []string `json:"application_number,omitempty" bson:"application_number,omitempty"`
	ManufacturerName  []string `json:"manufacturer_name,omitempty" bson:"manufacturer_name,omitempty"`
	BrandName         []string `json:"brand_name,omitempty" bson:"brand_name,omitempty"`
	GenericName       []string `json:"generic_name,omitempty" bson:"generic_name,omitempty"`

	// different from spec
	ProductType []string `json:"product_type,omitempty" bson:"product_type,omitempty"`

	ProductNDC    []string `json:"product_ndc,omitempty" bson:"product_ndc,omitempty"`
	NUI           []string `json:"nui,omitempty" bson:"nui,omitempty"`
	PackageNDC    []string `json:"package_ndc,omitempty" bson:"package_ndc,omitempty"`
	PharmClassCS  []string `json:"pharm_class_cs,omitempty" bson:"pharm_class_cs,omitempty"`
	PharmClassEPC []string `json:"pharm_class_epc,omitempty" bson:"pharm_class_epc,omitempty"`
	PharmClassMOA []string `json:"pharm_class_moa,omitempty" bson:"pharm_class_moa,omitempty"`
	PharmClassPE  []string `json:"pharm_class_pe,omitempty" bson:"pharm_class_pe,omitempty"`
	Route         []string `json:"route,omitempty" bson:"route,omitempty"`
	RxCUI         []string `json:"rxcui,omitempty" bson:"rxcui,omitempty"`
	SubstanceName []string `json:"substance_name,omitempty" bson:"substance_name,omitempty"`
	UNII          []string `json:"unii,omitempty" bson:"unii,omitempty"`
	UPC           []string `json:"upc,omitempty" bson:"upc,omitempty"`

	// Undocumented Fields Found In Responses
	IsOriginalPackager []bool `json:"is_original_packager,omitempty" bson:"is_original_packager,omitempty"`
}
