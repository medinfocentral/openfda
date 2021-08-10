package openfda

const (
	// ProductTypeHumanPrescriptionDrug renamed "HUMAN PRESCRIPTION DRUG" to "RX"
	ProductTypeHumanPrescriptionDrug = "RX"

	// ProductTypeHumanOTCDrug renamed "HUMAN OTC DRUG" to "OTC"
	ProductTypeHumanOTCDrug = "OTC"
)

func (d Drug) IsHumanPrescriptionDrug() bool {
	for _, tp := range d.ProductType {
		if tp == ProductTypeHumanPrescriptionDrug {
			return true
		}
	}
	return false
}
