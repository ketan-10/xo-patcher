package templates

type TemplateType uint16

const (
	ENUM TemplateType = iota
	TABLE
	REPO
	WIRE
)

func (tt *TemplateType) String() string {
	switch *tt {
	case ENUM:
		return "enum"
	case TABLE:
		return "table"
	case REPO:
		return "repo"
	case WIRE:
		return "wire"
	}

	return ""
}
