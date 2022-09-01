package connection_structure

type Person struct {
	ID          string
	Attributes  []Attribute
	Connections []Connection
	CreatedAt   string
}

type Attribute struct {
	Name           string
	Value          string
	Type           string
	Status         string
	ReasonDocument []ReasonDocument
}
