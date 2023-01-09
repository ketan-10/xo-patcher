package internal

import (
	"database/sql"
)

type Args struct {
	// DBC is database connection string
	DBC          string     `arg:"-"`
	DB           *sql.DB    `arg:"-"`
	Loader       ILoader    `arg:"-"`
	LoaderType   LoaderType `arg:"-"`
	DatabaseName string     `arg:"-"`
}

func GetDefaultArgs() *Args {
	return &Args{
		DBC: "",
	}
}

func (*Args) ExecuteTemplate(fileName string, obj interface{}) {

}
