package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/ketan-10/xo-patcher/xo/internal"

	// empty import, so that init method will be called, and drivers will be loaded
	_ "github.com/ketan-10/xo-patcher/xo/loaders"
	"github.com/xo/dburl"
)

func main() {
	var err error
	fmt.Println("Started")

	args := internal.GetDefaultArgs()

	// validate Args
	arg.MustParse(args)

	// open DB and save it to args
	err = openDB(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer args.DB.Close()

	// loaders
	err = getLoaderOfDriver(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	
	err = args.Loader.LoadSchema(args)

}

func openDB(args *internal.Args) error {
	url, err := dburl.Parse(args.DBC)
	if err != nil {
		return err
	}

	if url.Driver != "mysql" {
		// This can be extended for others
		return fmt.Errorf("only mysql is supported for xo-patcher not %s", url.Driver)
	}

	// enum 🤮
	lt := internal.LoaderType(0)
	lt.Unmarshal(url.Driver)
	args.LoaderType = lt

	// open mysql connection
	args.DB, err = sql.Open(url.Driver, url.DSN)
	if err != nil {
		return err
	}
	return nil
}

func getLoaderOfDriver(args *internal.Args) error {
	var ok bool
	args.Loader, ok = internal.AllLoaders[args.LoaderType]
	if !ok {
		return fmt.Errorf("for driver %s, no registered loader found", args.LoaderType.String())
	}
	return nil
}
