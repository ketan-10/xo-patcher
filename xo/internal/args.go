package internal

import (
	"bytes"
	"database/sql"
	"os"
	"text/template"

	"github.com/ketan-10/xo-patcher/xo/templates"
)

type Args struct {
	// DBC is database connection string
	DBC          string               `arg:"-"`
	DB           *sql.DB              `arg:"-"`
	Loader       ILoader              `arg:"-"`
	LoaderType   LoaderType           `arg:"-"`
	DatabaseName string               `arg:"-"`
	GeneratedDir string               `arg:"-"`
	Generated    []*GeneratedTemplate `arg:"-"`
}

func GetDefaultArgs(connection string) *Args {
	return &Args{
		GeneratedDir: "xo_gen",
		DBC:          connection,
	}
}

type GeneratedTemplate struct {
	TemplateType templates.TemplateType
	FileName     string
	Buffer       *bytes.Buffer
}

func (arg *Args) ExecuteTemplate(tt templates.TemplateType, fileName string, obj interface{}) error {
	// v, err := i.ReadFile("templates/" + fileName)

	genTmp := &GeneratedTemplate{
		TemplateType: tt,
		FileName:     fileName,
		Buffer:       new(bytes.Buffer),
	}

	// read template file
	templateFileLocation := "./templates/" + arg.LoaderType.String() + "/" + tt.String() + ".go.tpl"
	file, err := os.ReadFile(templateFileLocation)
	if err != nil {
		return err
	}

	t, err := template.
		New(templateFileLocation).
		Funcs(template.FuncMap(templates.HelperFunc)).
		Parse(string(file))
	if err != nil {
		return err
	}
	err = t.Execute(genTmp.Buffer, obj)
	if err != nil {
		return err
	}
	// save the generated buffer
	arg.Generated = append(arg.Generated, genTmp)
	return nil
}
