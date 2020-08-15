package internal

import (
	"github.com/realroy/migrator/internal/subcommand/migrate"

	"github.com/realroy/migrator/internal/subcommand/generate"

	"github.com/alexflint/go-arg"
)

var args struct {
	Generate *generate.Args `arg:"subcommand:generate"`
	Migrate  *migrate.Args  `arg:"subcommand:migrate"`
}

// Execute ...
func Execute() {
	p := arg.MustParse(&args)

	switch {
	case args.Generate != nil:
		generate.Execute(args.Generate)
	case args.Migrate != nil:
		migrate.Execute(args.Migrate)
	}

	if p.Subcommand() == nil {
		p.Fail("missing subcommand")
	}
}
