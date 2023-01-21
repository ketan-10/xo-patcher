package templates

type TemplateType uint16

const (
	ENUM TemplateType = iota
)

func (tt *TemplateType) String() string {
	switch *tt {
	case ENUM:
		return "enum"
	}
	return ""
}
