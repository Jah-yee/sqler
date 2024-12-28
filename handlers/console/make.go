package console

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alash3al/sqler/contracts"
	"github.com/alash3al/sqler/models"
	"github.com/urfave/cli/v3"
)

func MakeHandler(u contracts.UserService, a contracts.AuthService) *cli.Command {
	return &cli.Command{
		Name:  "make",
		Usage: "helper commands that help you create sysadmin user, tokens, ... etc",
		Commands: []*cli.Command{
			// make sysadmin
			{
				Name: "sysadmin",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Required: true},
					&cli.StringFlag{Name: "email", Required: true},
					&cli.StringFlag{Name: "password", Required: true},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					return u.Create(ctx, models.UserCreateInput{
						Name:       command.String("name"),
						Email:      command.String("email"),
						Password:   command.String("password"),
						IsSysAdmin: true,
					})
				},
			},

			// make token
			{
				Name: "auth-token",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "email", Required: true},
					&cli.StringFlag{Name: "password", Required: true},
					&cli.DurationFlag{Name: "ttl", Required: true, Usage: `Example "24h" Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					details, err := a.AuthenticateViaEmailAndPassword(
						ctx,
						models.AuthLoginWithEmailAndPasswordInput{
							Email:     command.String("email"),
							Password:  command.String("password"),
							TTL:       command.Duration("ttl"),
							UserAgent: "cli",
							IPAddress: "127.0.0.1",
						},
					)

					if err != nil {
						return err
					}

					j, err := json.MarshalIndent(details, "", "  ")
					if err != nil {
						return err
					}

					_, err = fmt.Println(string(j))

					return err
				},
			},
		},
	}
}
