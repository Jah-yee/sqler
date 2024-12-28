package console

import (
	"context"
	"fmt"
	"github.com/alash3al/sqler/services"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v3"
)

func MigrationHandler(m services.MigrationService) *cli.Command {
	return &cli.Command{
		Name:  "migration",
		Usage: "database migration related subcommands",
		Commands: []*cli.Command{
			// migration status
			{
				Name:  "status",
				Usage: "inspect the migration status",
				Action: func(ctx context.Context, command *cli.Command) error {
					status, err := m.Status(ctx)
					if err != nil {
						return err
					}

					tbl := table.NewWriter()

					tbl.SetStyle(table.StyleLight)
					tbl.AppendHeader(table.Row{"ID", "MigrationService", "Migrated At"})

					for _, item := range status {
						tbl.AppendRow(table.Row{item.ID, item.Filename, item.MigratedAt})
					}

					_, err = fmt.Println(tbl.Render())

					return err
				},
			},

			// migration apply
			{
				Name:  "apply",
				Usage: "execute the available migration files",
				Action: func(ctx context.Context, command *cli.Command) error {
					status, err := m.Apply(ctx)
					if err != nil {
						return err
					}

					_, err = fmt.Println(text.FgGreen.Sprintf("migrated %d file(s)", len(status)))

					return err
				},
			},
		},
	}
}
