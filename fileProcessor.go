package openfda

import (
	"encoding/json"
	"io"
)

// DrugLabelsFile is a struct that encapsulates an array of Labels
type DrugLabelsFile struct {
	Results []Label `json:"results"`
}

// DrugLabelsFileProcessor specifies how the processor should run
type DrugLabelsFileProcessor struct {
	Max int64

	Sanitize struct {
		SkipAbuseAndOverdosage            bool
		SkipAdverseEffectsAndInteractions bool
		SkipClinicalPharmacology          bool
		SkipIndicationsUsageAndDosage     bool
		SkipPatientInformation            bool
		SkipSpecialPopulations            bool
		SkipToxicology                    bool // ? should be NonclinicalToxicology
		SkipReferences                    bool
		SkipSupplyStorageAndHandling      bool
		SkipWarningsAndPrecautions        bool
		SkipOtherFields                   bool
	}

	Labels []Label
}

// NewProcessor creates a new default Processor
func NewProcessor() DrugLabelsFileProcessor {
	proc := DrugLabelsFileProcessor{
		Max: 100,
	}
	return proc
}

// LoadJSON loads the JSON encoded file into memory
func (process *DrugLabelsFileProcessor) LoadJSON(r io.Reader) error {
	var err error

	doc := DrugLabelsFile{}

	dec := json.NewDecoder(r)
	err = dec.Decode(&doc)
	if err != nil {
		return err
	}
	process.Labels = append(process.Labels, doc.Results...)
	return nil
}

// FilterHumanPrescriptionDrug filters out all non HUMAN PRESCRIPTION DRUGs
func (process *DrugLabelsFileProcessor) FilterHumanPrescriptionDrug() {
	filtered := process.Labels[:0]
	for _, label := range process.Labels {
		if len(label.OpenFDA.ProductType) > 0 && label.OpenFDA.ProductType[0] == "HUMAN PRESCRIPTION DRUG" {
			filtered = append(filtered, label)
		}
	}
	process.Labels = filtered
}

// FilterHumanOTCDrug filters out all non HUMAN OTC DRUGs
func (process *DrugLabelsFileProcessor) FilterHumanOTCDrug() {
	filtered := process.Labels[:0]
	for _, label := range process.Labels {
		if len(label.OpenFDA.ProductType) > 0 && label.OpenFDA.ProductType[0] == "HUMAN OTC DRUG" {
			filtered = append(filtered, label)
		}
	}
	process.Labels = filtered
}

// SaveAsJSON writes the labels as an array of labels into the io.Writer specified
func (process *DrugLabelsFileProcessor) SaveAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(DrugLabelsFile{
		Results: process.Labels})
}

// Process runs the label processor for each label
func (process *DrugLabelsFileProcessor) Process() {
	for i, label := range process.Labels {
		process.Labels[i] = *process.processLabel(&label)
	}
}

// func (process *DrugLabelsFileProcessor) Len() int {
// 	return len(process.Labels)
// }
