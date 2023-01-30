package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ketan-10/xo-patcher/patcher/utils"
	"github.com/ketan-10/xo-patcher/patcher/wire_app"
)

var patchNameFlag = flag.String("patch-name", "", "name of patch to run")
var commitFlag = flag.Bool("commit", false, "commit update")
var connectionStringFlag = flag.String("connection", "", "database connection string")

func main() {

	flag.Parse()
	fmt.Println("Hello World")

	connectionString := *connectionStringFlag
	patchName := *patchNameFlag
	if connectionString == "" {
		panic("connection flag must be provided")
	}
	if patchName == "" {
		panic("patch-name flag must be provided")
	}

	fmt.Println(connectionString)

	commit := *commitFlag

	ctx := context.WithValue(context.Background(), utils.Connection, connectionString)
	ctx = context.WithValue(ctx, utils.Commit, commit)
	ctx = context.WithValue(ctx, utils.PatchName, patchName)

	app, _, err := wire_app.GetApp(ctx)

	if err != nil {
		panic(err)
	}

	app.PatchManager.Run(ctx, patchName)

}
