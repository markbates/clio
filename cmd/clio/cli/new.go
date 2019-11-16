package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/markbates/clio"
	"github.com/markbates/clio/cmd/internal/cmdx"
	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/pkging/stdos"
)

func New(ctx context.Context, args []string) error {
	opts := struct {
		help bool
	}{}

	flags := flag.NewFlagSet("clio new", flag.ContinueOnError)
	flags.BoolVar(&opts.help, "h", false, "display help")

	cmdx.Usage(ctx, flags)

	if err := flags.Parse(args); err != nil {
		return err
	}

	if opts.help {
		flags.Usage()
		return nil
	}

	args = flags.Args()

	if len(args) == 0 {
		return fmt.Errorf("no module name specified")
	}

	ip := args[0]
	if len(ip) == 0 {
		return fmt.Errorf("no module name specified")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	name := path.Base(ip)
	dir := filepath.Join(pwd, name)
	info := here.Info{
		ImportPath: ip,
		Name:       name,
		Dir:        dir,
		Module: here.Module{
			Path: ip,
			Dir:  dir,
		},
	}

	pkg, err := stdos.New(info)
	if err != nil {
		return err
	}

	return clio.New(pkg)
}
