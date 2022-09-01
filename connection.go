package connection_structure

type Connection struct {
	ID              string
	Target1         string
	Target2         string
	T1_T2           string
	T2_T1           string
	Status          string
	ReasonDocuments []ReasonDocument
}
