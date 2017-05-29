package runtime

import (
	"fmt"
	"github.com/kwk-super-snippets/cli/src/app/out"
	"github.com/kwk-super-snippets/types"
	"os"
	"strings"
	"gopkg.in/yaml.v2"
	"log"
)

func newRuntimeAlias(name string, ext string, uniquePerMachine bool) *types.Alias {
	if uniquePerMachine {
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

func DefaultPrefs() *out.Prefs {
	return &out.Prefs{
		GlobalSearch:      false,
		AutoYes:           false,
		Covert:            false,
		RequireRunKeyword: true,
		ListAll:           true,
		Quiet:             false,
		SnippetThumbRows:  3,
		ExpandedThumbRows: 15,
		CommandTimeout:    60,
		RowSpaces:         true,
		RowLines:          false,
	}
}

func DefaultEnv() *yaml.MapSlice {
	env := &yaml.MapSlice{}
	err := yaml.Unmarshal([]byte(defaultEnv), env)
	if err != nil {
		log.Fatal(err)
	}
	return env
}

type PrefsFile struct {
	KwkPrefs string
	Options  out.Prefs
}

/*
	Road-map:
	WipeTrail         bool  //deletes the history each time a command is run TODO: Security
	SessionTimeout    int64 // 0 = no timeout, TODO: Implement on api SECURITY
	AutoEncrypt       bool  //Encrypts all snippets when created. TODO: SECURITY
	RegulateUpdates   bool  //Updates based on the recommended schedule. If false get updates as soon as available.
	DisableRun        bool  //Completely disabled running scripts even if using -f TODO: Security
	Encrypt   	  bool // TODO: Security
	Decrypt   	  bool // TODO: Security
*/
