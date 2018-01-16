package openfda

import (
	"strings"
	"sync"
	"unicode/utf8"
)

func (process *DrugLabelsFileProcessor) processLabel(label *Label) *Label {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cleanOpenFDA(label)
	}()

	if !process.Sanitize.SkipAbuseAndOverdosage {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanAbuseAndOverdosage(label)
		}()
	}
	if !process.Sanitize.SkipAdverseEffectsAndInteractions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanAdverseEffectsAndInteractions(label)
		}()
	}
	if !process.Sanitize.SkipClinicalPharmacology {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanClinicalPharmacology(label)
		}()
	}
	if !process.Sanitize.SkipIndicationsUsageAndDosage {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanIndicationsUsageAndDosage(label)
		}()
	}
	if !process.Sanitize.SkipPatientInformation {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanPatientInformation(label)
		}()
	}
	if !process.Sanitize.SkipSpecialPopulations {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanSpecialPopulations(label)
		}()
	}
	if !process.Sanitize.SkipToxicology {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanToxicology(label)
		}()
	}
	if !process.Sanitize.SkipToxicology {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanToxicology(label)
		}()
	}
	if !process.Sanitize.SkipReferences {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanReferences(label)
		}()
	}
	if !process.Sanitize.SkipSupplyStorageAndHandling {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanSupplyStorageAndHandling(label)
		}()
	}
	if !process.Sanitize.SkipWarningsAndPrecautions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanWarningsAndPrecautions(label)
		}()
	}
	if !process.Sanitize.SkipOtherFields {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cleanOtherFields(label)
		}()
	}
	wg.Wait()
	return label
}

func trimRunesFromFront(s string, validSet string) string {
	for i := 0; i < len(s); {
		r, sizer := utf8.DecodeRuneInString(s[i:])
		if strings.IndexRune(validSet, r) >= 0 {
			i += sizer
		} else {
			s = s[i:]
			break
		}
	}
	return s
}

func trimAnyPrefix(s string, prefixes []string) string {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return strings.TrimPrefix(s, prefix)
		}
	}
	return s
}

func permutations(s string) []string {
	p := []string{
		s,
		strings.ToLower(s),
		strings.ToUpper(s),
		strings.ToTitle(s),
	}

	// upercase only first letter of first word
	// TODO: DOES NOT HANDLE UTF8 fix this
	if len(s) > 2 {
		p = append(p, strings.ToUpper(s[:1])+s[1:])
	}

	if strings.Contains(s, "and") {
		p = append(p, permutations(strings.Replace(s, "and", "&", -1))...)
	}
	if strings.Contains(s, " ") {
		p = append(p, permutations(strings.Replace(s, " ", "", -1))...)
		p = append(p, permutations(strings.Replace(s, " ", "-", -1))...)
	}
	return p
}

// attributeFilterMapHelper helps filtering attributes
type attributeFilterMapHelper struct {
	arr  *[]string
	strs []string
}

func basePrefixSanitizer(targets *[]attributeFilterMapHelper) {
	for _, target := range *targets {
		for i, s := range *target.arr {
			s = trimRunesFromFront(s, "01. 23456789")
			for _, str := range target.strs {
				s = trimAnyPrefix(s, permutations(str))
			}
			s = trimRunesFromFront(s, ":01. 23456789")
			(*target.arr)[i] = s
		}
	}
}

func baseSufixSanitizer(targets *[]attributeFilterMapHelper) {
	for _, target := range *targets {
		for i, s := range *target.arr {
			for _, str := range target.strs {
				s = strings.TrimSuffix(s, str)
			}
			(*target.arr)[i] = s
		}
	}
}

func cleanOpenFDA(label *Label) {
	baseSufixSanitizer(&[]attributeFilterMapHelper{
		{&label.OpenFDA.PharmClassEPC, []string{
			" [PE]", " [pe]",
		}},
	})
	baseSufixSanitizer(&[]attributeFilterMapHelper{
		{&label.OpenFDA.PharmClassEPC, []string{
			" [EPC]", " [epc]",
		}},
	})
	baseSufixSanitizer(&[]attributeFilterMapHelper{
		{&label.OpenFDA.PharmClassMOA, []string{
			" [MOA]", " [MoA]",
		}},
	})
	baseSufixSanitizer(&[]attributeFilterMapHelper{
		{&label.OpenFDA.PharmClassCS, []string{
			" [CS]", " [Chemical/Ingredient]",
		}},
	})
}

// sanitizeAbuseAndOverdosage cleans up attributes:
// DrugAbuseAndDependence, ControlledSubstance, Abuse, Dependence, Overdosage
//
// review:
// id: 3dba3387-f213-464d-b75d-29eeacbd888c, has "OVERAGE Acute: ..." prefix
// id: 9edc0343-7c47-42a3-b63e-40295e9ec17, has "Overdosage Section" prefix
// id: 4a26a655-b559-4a11-b487-d6c6b5a4148e, has strange "OVERDOSAGE Human Experience" prefix
func cleanAbuseAndOverdosage(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.Overdosage, []string{"overdosage", "overdosage section"}},
		{&label.Dependence, []string{"dependence"}},
		{&label.Abuse, []string{"abuse"}},
		{&label.DrugAbuseAndDependence, []string{"drug abuse and dependence"}},
		{&label.ControlledSubstance, []string{"controlled substance"}},
	})
}

// sanitizeAdverseEffectsAndInteractions cleans up attributes:
// AdverseReactionsTable, AdverseReactions, DrugInteractions, DrugAndOrLaboratoryTestInteractions
func cleanAdverseEffectsAndInteractions(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.AdverseReactions, []string{"adverse reactions", "Adverse Reactions "}},
		{&label.DrugInteractions, []string{"Drug Interactions"}},
		{&label.DrugAndOrLaboratoryTestInteractions, []string{
			"drug or laboratory test interactions",
			"drug and laboratory test interactions",
			"drug and or laboratory test interactions",
			"Drug/Laboratory Test Interactions",
		}},
	})
}

func cleanClinicalPharmacology(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.ClinicalPharmacology, []string{"Clinical Pharmacology", "Mechanism of Action "}},
		// id: 23d8af55-f00f-4605-85ca-20d076990c88 has some strange long numbers
		{&label.Pharmacokinetics, []string{"pharmacokinetics"}},
		{&label.Pharmacodynamics, []string{"pharmacodynamics"}},
		{&label.MechanismOfAction, []string{"Mechanism of Action"}},
	})
}

func cleanIndicationsUsageAndDosage(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.IndicationsAndUsage, []string{
			"Indications & usage",
			"Indications and Usage",
			"I NDICATIONS AND USAGE ",
			"INDICA TI ONS AND USAGE "}},
		{&label.Contraindications, []string{"Contraindications", "CONTRA I NDICATIONS"}},
		{&label.DosageAndAdministration, []string{"Dosage and Administration"}},
		{&label.DosageFormsAndStrengths, []string{"Dosage Forms and Strengths"}},
		{&label.Purpose, []string{"Purpose"}},
		{&label.Description, []string{"Description"}},
		{&label.ActiveIngredient, []string{"Active Ingredient"}},
		{&label.InactiveIngredient, []string{"Inactive Ingredient"}},
		{&label.SPLProductDataElements, []string{"SPL Product Data Elements"}},
	})
}

/*
 */

func cleanPatientInformation(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.SPLPatientPackageInsert, []string{"SPL Patient Package Insert"}},
		{&label.InformationForPatients, []string{
			"Information for Patients",
			"PATIENT COU N SELING INFORMATION",
			"PATIENT COUNSELING INFORMATION "}},
		{&label.InformationForOwnersOrCaregivers, []string{"Information for Owners or Caregivers"}},
		{&label.InstructionsForUse, []string{"Instructions for Use"}},
		{&label.AskDoctor, []string{"Ask Doctor"}},
		{&label.AskDoctorOrPharmacist, []string{"Ask Doctor or Pharmacist"}},
		{&label.DoNotUse, []string{"DoNotUse"}},
		{&label.KeepOutOfReachOfChildren, []string{"Keep Out of Reach of Children"}},
		{&label.OtherSafetyInformation, []string{"Other Safety Information"}},
		{&label.Questions, []string{"Questions"}},
		{&label.StopUse, []string{"Stop Use"}},
		{&label.WhenUsing, []string{"When Using"}},
		{&label.PatientMedicationInformation, []string{"Patient Medication Information"}},
		{&label.SPLMedguide, []string{"SPL Medguide", "MEDICATION GUIDE "}},
	})
}
func cleanSpecialPopulations(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.UseInSpecificPopulations, []string{"Use in Specific Populations"}},
		{&label.Pregnancy, []string{"Pregnancy"}},
		{&label.TeratogenicEffects, []string{"Teratogenic Effects"}},
		{&label.LaborAndDelivery, []string{"Labor and Delivery"}},
		{&label.NursingMothers, []string{"Nursing Mothers"}},
		{&label.PregnancyOrBreastFeeding, []string{"Pregnancy or Breast Feeding"}},
		{&label.PediatricUse, []string{
			"Pediatric use",
			"Pediatric Use ",
			"Pediatric Use",
		}},
		{&label.GeriatricUse, []string{"Geriatric use", "Geriatric Use "}},
	})
}
func cleanToxicology(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.NonclinicalToxicology, []string{
			"Non-clinical Toxicology",
			"Non Clinical Toxicology",
			"N ONCLINICAL TOXICOLOGY",
			"NONCLINICAL TOXICOLOGY",
			"Carcinogenesis, Mutagenesis, Impairment of Fertility Carcinogenicity"}},
		{&label.CarcinogenesisAndMutagenesisAndImpairmentOfFertility, []string{
			"Carcinogenesis and Mutagenesis and Impairment of Fertility",
			"Carcinogenesis, Mutagenesis, Impairment of Fertility",
			"Carcinogenesis Mutagenesis Impairment Fertility",
			"Carcinogenesis",
			"Carcinogenicity"}},
		{&label.AnimalPharmacologyAndOrToxicology, []string{
			"Animal Pharmacology and or Toxicology",
			"ANIMAL PHARMACOLOGY OR ANIMAL TOXICOLOGY",
			"ANIMAL PHARMACOLOGY AND ANIMAL TOXICOLOGY"}},
	})
}

func cleanReferences(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.ClinicalStudies, []string{"Clinical Studies"}},
		// {&label.ClinicalStudiesTable, []string{"ClinicalStudiesTable"}},
		{&label.References, []string{"References"}},
	})
}
func cleanSupplyStorageAndHandling(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.HowSupplied, []string{
			"How Supplied",
			"/STORAGE AND HANDLING ",
			"How Supplied "},
		},
		{&label.StorageAndHandling, []string{"Storage and Handling", "Storage Conditions"}},
		{&label.SafeHandlingWarning, []string{"Safe Handling Warning"}},
	})

}
func cleanWarningsAndPrecautions(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.BoxedWarning, []string{"Boxed Warning"}},
		{&label.WarningsAndCautions, []string{
			"Warnings and Cautions",
			"Warnings and Precautions",
			"Warnings and Pre-Cautions"},
		},
		{&label.UserSafetyWarnings, []string{"User Safety Warnings"}},
		{&label.Precautions, []string{"Precautions", "Information for Patients "}},
		{&label.Warnings, []string{"Warnings"}},
		{&label.GeneralPrecautions, []string{"General Precautions"}},
	})
}
func cleanOtherFields(label *Label) {
	basePrefixSanitizer(&[]attributeFilterMapHelper{
		{&label.LaboratoryTests, []string{"Laboratory Tests"}},
		{&label.RecentMajorChanges, []string{"Recent Major Changes"}},
		{&label.Microbiology, []string{"Microbiology"}},
		{&label.PackageLabelPrincipalDisplayPanel, []string{"Package Label Principal Display Panel"}},
		{&label.SPLUnclassifiedSection, []string{"SPL Unclassified Section"}},
	})
}
