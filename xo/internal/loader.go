package internal

import (
	"fmt"
	"strings"

	"github.com/ketan-10/xo-patcher/xo/loaders/models"
)

// The loader interface
type ILoader interface {
	LoadSchema(*Args) error
}

var AllLoaders = map[LoaderType]ILoader{}

// The loader implementation
// Drivers like mysql will create object for this.
type LoaderImp struct {
	EnumList      func(db models.XODB, databaseName string) ([]*models.Enum, error)
	DatabaseName  func(db models.XODB) (string, error)
	EnumValueList func(db models.XODB, databaseName string, tableName string, columnName string) (string, error)
}

// Entry point to load everything
func (lt *LoaderImp) LoadSchema(args *Args) error {
	var err error

	database, err := lt.loadDatabaseName(args)
	if err != nil {
		return err
	}
	args.DatabaseName = database

	err = lt.loadEnums(args)
	if err != nil {
		return err
	}

	return nil
}

func (lt *LoaderImp) loadDatabaseName(args *Args) (string, error) {
	if lt.DatabaseName == nil {
		return "", fmt.Errorf("schema name loader is not implemented for %s", args.LoaderType.String())
	}
	return lt.DatabaseName(args.DB)
}

type EnumDTO struct {
	*models.Enum
	DatabaseName string
	Values       []string
}

func (lt *LoaderImp) loadEnums(args *Args) error {
	enums, err := lt.EnumList(args.DB, args.DatabaseName)
	if err != nil {
		return err
	}

	var allEnumDTO []*EnumDTO
	for _, e := range enums {
		// fmt.Printf("%s, %s \n", e.ColumnName, e.TableName)
		enumValues, err := lt.loadEnumValues(args, e)

		if err != nil {
			return err
		}

		allEnumDTO = append(allEnumDTO, &EnumDTO{
			e,
			args.DatabaseName,
			enumValues,
		})
	}

	for _, enum := range allEnumDTO {
		args.ExecuteTemplate("mysql.enum.go.tpl", enum)
	}

	// fmt.Printf(""enums)
	return nil
}

func (lt *LoaderImp) loadEnumValues(args *Args, enum *models.Enum) ([]string, error) {
	if lt.EnumValueList == nil {
		return nil, fmt.Errorf("enumValue loader is not implemented for %s", args.LoaderType.String())
	}

	values, err := lt.EnumValueList(args.DB, args.DatabaseName, enum.TableName, enum.ColumnName)

	if err != nil {
		return nil, err
	}
	// value is in 'A','B','C' we want to convert to a list
	list := strings.Split(values[1:len(values)-1], "','")
	return list, nil
}
