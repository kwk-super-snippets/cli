package app

import (
	"bitbucket.com/sharingmachine/kwkcli/models"
	"bitbucket.com/sharingmachine/kwkcli/settings"
	"bitbucket.com/sharingmachine/kwkcli/users"
	"bitbucket.com/sharingmachine/kwkcli/ui/tmpl"
	"bitbucket.com/sharingmachine/kwkcli/ui/dlg"
)

type AccountCli struct {
	service  users.IUsers
	settings settings.ISettings
	tmpl.Writer
	dlg.Dialogue
}

func NewAccountCli(u users.IUsers, s settings.ISettings, w tmpl.Writer, d dlg.Dialogue) *AccountCli {
	return &AccountCli{service: u, settings: s, Writer: w, Dialogue: d}
}

func (c *AccountCli) Get() {
	u := &models.User{}
	if err := c.settings.Get(models.ProfileFullKey, u); err != nil {
		c.Render("account:notloggedin", nil)
	} else {
		c.Render("account:profile", u)
	}
}

func (c *AccountCli) SignUp() {

	email := c.Field("account:signup:email", nil).Value.(string)
	username := c.Field("account:signup:username", nil).Value.(string)
	password := c.Field("account:signup:password", nil).Value.(string)

	if u, err := c.service.SignUp(email, username, password); err != nil {
		c.Render("error", err)
	} else {
		if len(u.Token) > 50 {
			c.settings.Upsert(models.ProfileFullKey, u)
			c.Render("account:signedup", u)
		}
	}
}

func (c *AccountCli) SignIn(username string, password string) {
	if username == "" {
		username = c.Field("account:usernamefield", nil).Value.(string)
	}
	if password == "" {
		password = c.Field("account:passwordfield", nil).Value.(string)
	}
	if u, err := c.service.SignIn(username, password); err != nil {
		c.Render("error", err)
	} else {
		if len(u.Token) > 50 {
			c.settings.Upsert(models.ProfileFullKey, u)
			c.Render("account:signedin", u)
		}
	}
}

func (c *AccountCli) SignOut() {
	if err := c.service.Signout(); err != nil {
		c.Render("error", err)
		return
	}
	if err := c.settings.Delete(models.ProfileFullKey); err != nil {
		c.Render("error", err)
		return
	}
	c.Render("account:signedout", nil)
}

func (c *AccountCli) ChangeDirectory(username string) {
	if err := c.settings.ChangeDirectory(username); err != nil {
		c.Render("error", err)
	} else {
		c.Render("account:cd", map[string]string{"username": username})
	}
}