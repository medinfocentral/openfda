package openfda

type DrugLabel struct {
	// ID and version
	EffectiveTime string `json:"effective_time" bson:"effective_time,omitempty"`
	SetID         string `json:"set_id" bson:"set_id,omitempty"`
	ID            string `json:"id" bson:"_id"`
	Version       int    `json:"version,string" bson:"version,omitempty"`

	// Abuse and overdosage
	DrugAbuseAndDependence []string `json:"drug_abuse_and_dependence,omitempty" bson:"drug_abuse_and_dependence,omitempty"`
	ControlledSubstance    []string `json:"controlled_substance,omitempty" bson:"controlled_substance,omitempty"`
	Abuse                  []string `json:"abuse,omitempty" bson:"abuse,omitempty"`
	Dependence             []string `json:"dependence,omitempty" bson:"dependence,omitempty"`
	Overdosage             []string `json:"overdosage,omitempty" bson:"overdosage,omitempty"`

	// Adverse effects and interactions
	AdverseReactionsTable               []string `json:"adverse_reactions_table,omitempty" bson:"adverse_reactions_table,omitempty"`
	AdverseReactions                    []string `json:"adverse_reactions,omitempty" bson:"adverse_reactions,omitempty"`
	DrugInteractions                    []string `json:"drug_interactions,omitempty" bson:"drug_interactions,omitempty"`
	DrugAndOrLaboratoryTestInteractions []string `json:"drug_and_or_laboratory_test_interactions,omitempty" bson:"drug_and_or_laboratory_test_interactions,omitempty"`

	// Clinical pharmacology
	ClinicalPharmacology []string `json:"clinical_pharmacology,omitempty" bson:"clinical_pharmacology,omitempty"`
	Pharmacokinetics     []string `json:"pharmacokinetics,omitempty" bson:"pharmacokinetics,omitempty"`
	Pharmacodynamics     []string `json:"pharmacodynamics,omitempty" bson:"pharmacodynamics,omitempty"`
	MechanismOfAction    []string `json:"mechanism_of_action,omitempty" bson:"mechanism_of_action,omitempty"`

	// Indications, usage, and dosage
	IndicationsAndUsage     []string `json:"indications_and_usage,omitempty" bson:"indications_and_usage,omitempty"`
	Contraindications       []string `json:"contraindications,omitempty" bson:"contraindications,omitempty"`
	DosageAndAdministration []string `json:"dosage_and_administration,omitempty" bson:"dosage_and_administration,omitempty"`
	DosageFormsAndStrengths []string `json:"dosage_forms_and_strengths,omitempty" bson:"dosage_forms_and_strengths,omitempty"`
	Purpose                 []string `json:"purpose,omitempty" bson:"purpose,omitempty"`
	Description             []string `json:"description,omitempty" bson:"description,omitempty"`
	ActiveIngredient        []string `json:"active_ingredient,omitempty" bson:"active_ingredient,omitempty"`
	InactiveIngredient      []string `json:"inactive_ingredient,omitempty" bson:"inactive_ingredient,omitempty"`
	SPLProductDataElements  []string `json:"spl_product_data_elements,omitempty" bson:"spl_product_data_elements,omitempty"`

	// Patient information
	SPLPatientPackageInsert          []string `json:"spl_patient_package_insert,omitempty" bson:"spl_patient_package_insert,omitempty"`
	InformationForPatients           []string `json:"information_for_patients,omitempty" bson:"information_for_patients,omitempty"`
	InformationForOwnersOrCaregivers []string `json:"information_for_owners_or_caregivers,omitempty" bson:"information_for_owners_or_caregivers,omitempty"`
	InstructionsForUse               []string `json:"instructions_for_use,omitempty" bson:"instructions_for_use,omitempty"`
	AskDoctor                        []string `json:"ask_doctor,omitempty" bson:"ask_doctor,omitempty"`
	AskDoctorOrPharmacist            []string `json:"ask_doctor_or_pharmacist,omitempty" bson:"ask_doctor_or_pharmacist,omitempty"`
	DoNotUse                         []string `json:"do_not_use,omitempty" bson:"do_not_use,omitempty"`
	KeepOutOfReachOfChildren         []string `json:"keep_out_of_reach_of_children,omitempty" bson:"keep_out_of_reach_of_children,omitempty"`
	OtherSafetyInformation           []string `json:"other_safety_information,omitempty" bson:"other_safety_information,omitempty"`
	Questions                        []string `json:"questions,omitempty" bson:"questions,omitempty"`
	StopUse                          []string `json:"stop_use,omitempty" bson:"stop_use,omitempty"`
	WhenUsing                        []string `json:"when_using,omitempty" bson:"when_using,omitempty"`
	PatientMedicationInformation     []string `json:"patient_medication_information,omitempty" bson:"patient_medication_information,omitempty"`
	SPLMedguide                      []string `json:"spl_medguide,omitempty" bson:"spl_medguide,omitempty"`

	// Special populations
	UseInSpecificPopulations []string `json:"use_in_specific_populations,omitempty" bson:"use_in_specific_populations,omitempty"`
	Pregnancy                []string `json:"pregnancy,omitempty" bson:"pregnancy,omitempty"`
	TeratogenicEffects       []string `json:"teratogenic_effects,omitempty" bson:"teratogenic_effects,omitempty"`
	LaborAndDelivery         []string `json:"labor_and_delivery,omitempty" bson:"labor_and_delivery,omitempty"`
	NursingMothers           []string `json:"nursing_mothers,omitempty" bson:"nursing_mothers,omitempty"`
	PregnancyOrBreastFeeding []string `json:"pregnancy_or_breast_feeding,omitempty" bson:"pregnancy_or_breast_feeding,omitempty"`
	PediatricUse             []string `json:"pediatric_use,omitempty" bson:"pediatric_use,omitempty"`
	GeriatricUse             []string `json:"geriatric_use,omitempty" bson:"geriatric_use,omitempty"`

	// toxicology
	NonclinicalToxicology                                []string `json:"nonclinical_toxicology,omitempty" bson:"nonclinical_toxicology,omitempty"`
	CarcinogenesisAndMutagenesisAndImpairmentOfFertility []string `json:"carcinogenesis_and_mutagenesis_and_impairment_of_fertility,omitempty" bson:"carcinogenesis_and_mutagenesis_and_impairment_of_fertility,omitempty"`
	AnimalPharmacologyAndOrToxicology                    []string `json:"animal_pharmacology_and_or_toxicology,omitempty" bson:"animal_pharmacology_and_or_toxicology,omitempty"`

	// References
	ClinicalStudies      []string `json:"clinical_studies,omitempty" bson:"clinical_studies,omitempty"`
	ClinicalStudiesTable []string `json:"clinical_studies_table,omitempty" bson:"clinical_studies_table,omitempty"`
	References           []string `json:"References,omitempty" bson:"References,omitempty"`

	// Supply, storage, and handling
	HowSupplied         []string `json:"how_supplied,omitempty" bson:"how_supplied,omitempty"`
	StorageAndHandling  []string `json:"storage_and_handling,omitempty" bson:"storage_and_handling,omitempty"`
	SafeHandlingWarning []string `json:"safe_handling_warning,omitempty" bson:"safe_handling_warning,omitempty"`

	// Warnings and precautions
	BoxedWarning        []string `json:"boxed_warning,omitempty" bson:"boxed_warning,omitempty"`
	WarningsAndCautions []string `json:"warnings_and_cautions,omitempty" bson:"warnings_and_cautions,omitempty"`
	UserSafetyWarnings  []string `json:"user_safety_warnings,omitempty" bson:"user_safety_warnings,omitempty"`
	Precautions         []string `json:"precautions,omitempty" bson:"precautions,omitempty"`
	Warnings            []string `json:"warnings,omitempty" bson:"warnings,omitempty"`
	GeneralPrecautions  []string `json:"general_precautions,omitempty" bson:"general_precautions,omitempty"`

	// Other fields
	LaboratoryTests                   []string `json:"laboratory_tests,omitempty" bson:"laboratory_tests,omitempty"`
	RecentMajorChanges                []string `json:"recent_major_changes,omitempty" bson:"recent_major_changes,omitempty"`
	Microbiology                      []string `json:"microbiology,omitempty" bson:"microbiology,omitempty"`
	PackageLabelPrincipalDisplayPanel []string `json:"package_label_principal_display_panel,omitempty" bson:"package_label_principal_display_panel,omitempty"`
	SPLUnclassifiedSection            []string `json:"spl_unclassified_section,omitempty" bson:"spl_unclassified_section,omitempty"`
}
