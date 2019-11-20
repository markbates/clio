package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/markbates/clio"
	"github.com/markbates/clio/cmd/internal/cmdx"
)

func Main(ctx context.Context, args []string) error {
	opts := struct {
		version bool
		help    bool
	}{}

	flags := flag.NewFlagSet("clio", flag.ContinueOnError)
	flags.BoolVar(&opts.version, "v", false, "display version")
	flags.BoolVar(&opts.help, "h", false, "display help")

	cmdx.Usage(ctx, flags)

	if err := flags.Parse(args); err != nil {
		return err
	}

	stdout := cmdx.Stdout(ctx)
	if opts.version {
		fmt.Fprintln(stdout, clio.Version)
		return nil
	}

	args = flags.Args()
	if opts.help || len(args) == 0 {
		flags.Usage()
		return nil
	}

	arg := args[0]
	if len(args) > 0 {
		args = args[1:]
	}

	switch arg {
	case "new":
		return New(ctx, args)
	}

	return nil
}
