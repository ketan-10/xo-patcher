package templates

type TemplateType uint16

const (
	ENUM TemplateType = iota
	TABLE
	REPO
)

func (tt *TemplateType) String() string {
	switch *tt {
	case ENUM:
		return "enum"
	case TABLE: 
		return "table"
	case REPO:
		return "repo"
	}
	return ""
}
