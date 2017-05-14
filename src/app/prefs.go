package app

import (
	"fmt"
	"github.com/kwk-super-snippets/cli/src/app/out"
	"github.com/kwk-super-snippets/types"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var prefs *Preferences
var env *yaml.MapSlice

func NewSetupAlias(name string, ext string, uniquePerHost bool) *types.Alias {
	if uniquePerHost {
		s, _ := os.Hostname()
		//u, _ := user.Current()
		name = fmt.Sprintf("%s-%s", name, strings.ToLower(s))
	}
	return &types.Alias{
		Pouch: types.PouchSettings,
		Name:  name,
		Ext:   ext,
	}
}

func Prefs() *Preferences {
	//if prefs == nil {
	//	panic("Preferences not initialized")
	//	//prefs = DefaultPrefs()
	//}
	return prefs
}

func Env() *yaml.MapSlice {
	return env
}

func SetPrefs(p *Preferences) {
	prefs = p
}

func SetEnv(e *yaml.MapSlice) {
	env = e
}

func DefaultPrefs() *Preferences {
	p := &Preferences{
		Global:  false,
		AutoYes: false,
		Force:   false,
	}
	p.Covert = false
	p.DisableRun = false
	p.RequireRunKeyword = true
	p.WipeTrail = false
	p.SessionTimeout = 15
	p.ListAll = true
	p.RegulateUpdates = true

	p.SlimRows = 3
	p.ExpandedRows = 15
	p.AlwaysExpandRows = false
	p.CommandTimeout = 60
	p.RowSpaces = true
	p.RowLines = false
	p.ListHorizontal = false
	return p
}

type PreferencesHolder struct {
	KwkPrefs    string
	Preferences PersistedPrefs
}

// Preferences is a container for file and flag preferences.
type Preferences struct {
	PersistedPrefs
	Global bool // TODO: Implement in search api SEARCH
	//Quiet   bool // Only display fullNames (for multi deletes etc)
	AutoYes   bool //TODO: Security
	Force     bool // TODO: Security Will squash warning messages e.g. when running third party snippets.
	Encrypt   bool // TODO: Security
	Decrypt   bool // TODO: Security
	LastPouch string
}

// PersistedPrefs are preferences which can be persistent.
type PersistedPrefs struct {
	Covert            bool  // Always opens browser in covert mode, when set to true flag should have no effect. TODO: Update env/darwin.yml
	DisableRun        bool  //Completely disabled running scripts even if using -f TODO: Security
	RequireRunKeyword bool  //If false then `kwk <snipname>` will execute the snippet without the `run|r` parameter. In this case `view|v` command will be required to view the details of a snippet`
	WipeTrail         bool  //deletes the history each time a command is run TODO: Security
	SessionTimeout    int64 // 0 = no timeout, TODO: Implement on api SECURITY
	AutoEncrypt       bool  //Encrypts all snippets when created. TODO: SECURITY
	RegulateUpdates   bool  //Updates based on the recommended schedule. If false get updates as soon as available.
	CommandTimeout    int64
	out.Prefs
}