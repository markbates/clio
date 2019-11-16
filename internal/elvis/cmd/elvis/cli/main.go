package cli

import (
	"context"
	"flag"
	"fmt"

	"github.com/costello/elvis/cmd/internal/cmdx"
)

func Main(ctx context.Context, args []string) error {
	opts := struct {
		version bool
		help    bool
	}{}

	flags := flag.NewFlagSet("elvis", flag.ContinueOnError)
	flags.BoolVar(&opts.version, "v", false, "display version")
	flags.BoolVar(&opts.help, "h", false, "display help")

	cmdx.Usage(ctx, flags)

	if err := flags.Parse(args); err != nil {
		return err
	}

	stdout := cmdx.Stdout(ctx)
	if opts.version {
		fmt.Fprintln(stdout, Version)
		return nil
	}

	if opts.help {
		flags.Usage()
		return nil
	}

	args = flags.Args()
	return nil
}
