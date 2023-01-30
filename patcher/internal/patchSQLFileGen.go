package internal

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/elgris/sqrl"
	"github.com/google/wire"
	"github.com/ketan-10/xo-patcher/patcher/utils"
)

type PatchSQLFileGen struct {
	file *os.File
}

type IPatchSQLFileGen interface {
	Write(sqlizer sqrl.Sqlizer) error
	Close()
}

var NewPatchSQLFileGen = wire.NewSet(
	InitPatchSQLFileGen,
	wire.Bind(new(IPatchSQLFileGen), new(PatchSQLFileGen)),
)

func InitPatchSQLFileGen(ctx context.Context) (*PatchSQLFileGen, error) {
	patchName, err := getPatchContext(ctx)
	if err != nil {
		return nil, err
	}
	file, err := os.Create("patches_gen/" + patchName + ".sql")
	if err != nil {
		return nil, err
	}
	return &PatchSQLFileGen{file: file}, nil

}

func (fileGen *PatchSQLFileGen) Write(sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	
	fileGen.file.WriteString(fmt.Sprintf(query, args...))
	fmt.Println(args)
	// fileGen.file.WriteString(args)
	return nil

}

func (fileGen *PatchSQLFileGen) Close() {
	fileGen.file.Close()
}

func getPatchContext(ctx context.Context) (string, error) {
	if value, ok := ctx.Value(utils.PatchName).(string); ok {
		return value, nil
	}
	return "", errors.New("patchName context invalid")
}
