package connection_structure

const DocumentTypeText = "text"
const DocumentTypePhoto = "photo"
const DocumentTypeLink = "link"

type ReasonDocument struct {
	ID   string
	Name string
	Data string
	Type string
}
