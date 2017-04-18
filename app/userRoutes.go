package app

import (
	"github.com/urfave/cli"
)

func Accounts(a *UserCli) []cli.Command {
	c := []cli.Command{
		{
			Name:    "profile",
			Aliases: []string{"me", "whoami"},
			Action: func(c *cli.Context) error {
				a.Get()
				return nil
			},
		},
		{
			Name:    "signin",
			Aliases: []string{"login", "switch", "cd"},
			Action: func(c *cli.Context) error {
				a.SignIn(c.Args().Get(0), c.Args().Get(1))
				return nil
			},
		},
		{
			Name:    "signup",
			Aliases: []string{"register"},
			Action: func(c *cli.Context) error {
				a.SignUp()
				return nil
			},
		},
		{
			Name:    "signout",
			Aliases: []string{"logout"},
			Action: func(c *cli.Context) error {
				a.SignOut()
				return nil
			},
		},
		{
			Name: "reset-password",
			Action: func(c *cli.Context) error {
				a.ResetPassword(c.Args().First())
				return nil
			},
		},
		{
			Name: "change-password",
			Action: func(c *cli.Context) error {
				a.ChangePassword()
				return nil
			},
		},
	}
	return c
}