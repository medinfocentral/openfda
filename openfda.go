package openfda

const (
	// ProductTypeHumanPrescriptionDrug remamed "HUMAN PRESCRIPTION DRUG" to "RX"
	ProductTypeHumanPrescriptionDrug = "RX"

	// ProductTypeHumanOTCDrug remamed "HUMAN OTC DRUG" to "OTC"
	ProductTypeHumanOTCDrug = "OTC"

	ProductTypeHumanPerscriptionDrug = "HUMAN PRESCRIPTION DRUG"
)

func (d Drug) IsHumanPresecriptionDrug() bool {
	for _, tp := range d.ProductType {
		if tp == ProductTypeHumanPerscriptionDrug {
			return true
		}
	}
	return false
}
