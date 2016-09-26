package main

import (
	"os"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/settings"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/aliases"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/openers"
	"bitbucket.com/sharingmachine/kwkcli/libs/app"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/gui"
	"bitbucket.com/sharingmachine/kwkcli/libs/rpc"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/system"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/users"
)

func main() {
	os.Setenv("version", "v0.0.1")

	conn := rpc.Conn("127.0.0.1:7777");
	defer conn.Close()

	s := system.New()
	t := settings.New(s, "settings")

	u := users.New(conn, t)
	a := aliases.New(conn, t)
	o := openers.New(s, a)
	w := gui.NewTemplateWriter(os.Stdout)
	d := gui.NewDialogues(w)

	kwkApp := app.NewKwkApp(a, s, t, o, u, d, w)
	kwkApp.App.Run(os.Args)
}
