package app

import (
	"bitbucket.com/sharingmachine/kwkcli/search"
	"bitbucket.com/sharingmachine/kwkcli/config"
	"bitbucket.com/sharingmachine/kwkcli/system"
	"bitbucket.com/sharingmachine/kwkcli/account"
	"bitbucket.com/sharingmachine/kwkcli/ui/dlg"
	"bitbucket.com/sharingmachine/kwkcli/ui/tmpl"
	"bitbucket.com/sharingmachine/kwkcli/snippets"
	"bitbucket.com/sharingmachine/kwkcli/cmd"
	"gopkg.in/urfave/cli.v1"
	"bitbucket.com/sharingmachine/kwkcli/rpc"
)

type KwkApp struct {
	App            *cli.App
	Snippets       snippets.Service
	System         system.ISystem
	Settings       config.Settings
	AccountManage  account.Manager
	Runner         cmd.Runner
	Dialogue       dlg.Dialog
	TemplateWriter tmpl.Writer
	Search         search.Term
	Api 		rpc.Sys
}

func New(a snippets.Service, s system.ISystem, t config.Settings, r cmd.Runner, u account.Manager,
	d dlg.Dialog, w tmpl.Writer, h search.Term, api rpc.Sys) *KwkApp {

	app := cli.NewApp()
	//cli.HelpPrinter = system.Help

	accCli := NewAccountCli(u, t, w, d)
	app.Commands = append(app.Commands, Accounts(accCli)...)

	sysCli := NewSystemCli(s, api, u, w)
	app.Commands = append(app.Commands, System(sysCli)...)

	snipCli := NewSnippetCli(a, r, s, d, w, t, h)
	app.Commands = append(app.Commands, Snippets(snipCli)...)
	app.CommandNotFound = func(c *cli.Context, fullKey string) {
		snipCli.Run(fullKey, []string(c.Args())[1:])
	}
	searchCli := NewSearchCli(h, w, d)
	app.Commands = append(app.Commands, Search(searchCli)...)

	return &KwkApp{
		App: app,
		System: s,
		Settings: t,
		Runner: r,
		AccountManage: u,
		Dialogue: d,
		Snippets: a,
		TemplateWriter: w,
		Search:h,
		Api:api,
	}
}

func (a *KwkApp) Run(args ...string) {
	params := []string{"[app]"}
	params = append(params, args...)
	a.App.Run(params)
}