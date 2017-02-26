package tmpl

import (
	"bitbucket.com/sharingmachine/kwkcli/ui/style"
	"bitbucket.com/sharingmachine/kwkcli/models"
	"github.com/rjarmstrong/go-humanize"
	"encoding/json"
	"text/template"
	"strings"
	"fmt"
	"time"
)

var Templates = map[string]*template.Template{}

//const logo = `
//  ▋                         ▋
//  ▋                         ▋
//  ▋   ◢  ◥◣           ◢◤   ▋   ◢◤
//  ▋ ◤      ◥◣    ◢◤   ◢◤    ▋ ◤
//  ▋ ◣       ◥◣ ◢◤ ◥◣ ◢◤     ▋ ◣
//  ▋   ◥     ◢◤     ◢◤      ▋   ◥◣
//`

const logo = `
`

const(
	MARGIN = "  "
)

func init() {
	// Aliases
	add("dashboard", style.Fmt(style.Cyan, logo)+"{{. | listRoot }}", template.FuncMap{"listRoot": listRoot })

	add("snippet:updated", "Description updated:\n{{ .Description | blue }}", template.FuncMap{"blue": blue})
	add("api:not-found", "{{. | yellow }} Not found\n", template.FuncMap{"yellow": yellow})
	add("snippet:cloned", "Cloned as {{.Username}}/{{.FullName | blue}}\n", template.FuncMap{"blue": blue})
	add("snippet:new", "{{. | blue }} created "+style.OpenLock+"\n", template.FuncMap{"blue": blue})
	add("snippet:newprivate", "{{.FullName | blue }} created "+style.Lock+"\n", template.FuncMap{"blue": blue})
	add("snippet:cat", "{{.Snip | blue}}", template.FuncMap{"blue": blue})
	add("snippet:edited", "Successfully updated {{ .String | blue }}", template.FuncMap{"blue": blue})
	add("snippet:editing", "{{ \"Editing... \" | blue }}\nPlease edit file and save.\n - NB: After saving, no changes will be saved until running kwk edit <name> again.\n - Ctrl+C to abort.\n", template.FuncMap{"blue": blue})
	add("snippet:edit-prompt", "{{ .String | blue }} doesn't exist - would you like create it? [y/n] \n", template.FuncMap{"blue": blue})

	add("snippet:ambiguouscat", "That snippet is ambiguous please run it again with the extension:\n{{range .Items}}{{.FullName | blue }}\n{{ end }}", template.FuncMap{"blue": blue})
	add("snippet:list", "{{. | listPouch }}", template.FuncMap{"listPouch": listPouch })
	add("pouch:list-root", "{{. | listRoot }}", template.FuncMap{"listRoot": listRoot })

	add("snippet:tag", "{{.FullName | blue }} tagged.", template.FuncMap{"blue": blue})
	add("snippet:untag", "{{.FullName | blue }} untagged.", template.FuncMap{"blue": blue})
	add("snippet:renamed", "{{.originalName | blue }} renamed to {{.newName | blue }}", template.FuncMap{"blue": blue})
	add("snippet:madeprivate", "{{.fullName | blue }} made private "+style.Lock, template.FuncMap{"blue": blue})
	add("snippet:patched", "{{.FullName | blue }} patched.", template.FuncMap{"blue": blue})

	add("snippet:check-delete", "Are you sure you want to delete snippet {{. | yellow }}? [y/n] ", template.FuncMap{"yellow": yellow})
	add("snippet:deleted", "Snippets {{. | blue }} deleted.", template.FuncMap{"blue": blue})
	add("snippet:not-deleted", "Snippets {{. | blue }} NOT deleted.", template.FuncMap{"blue": blue})

	add("snippet:moved-root", "{{ .Quant | blue }} snippet(s) moved to root.", template.FuncMap{"blue": blue})
	add("snippet:moved-pouch", "{{ .Quant | blue }} snippet(s) moved to pouch {{ .Pouch | blue }}", template.FuncMap{"blue": blue})
	add("snippet:create-pouch", "{{ \"Would you like to create the snippet in a new pouch? [y/n] \" | yellow }}?  ", template.FuncMap{"yellow": yellow})

	//add("snippet:inspect",
	//	"\n{{range .Items}}"+"Name: {{.String}}"+"\nClone count: {{.CloneCount}}"+"\n{{ end }}\n\n", nil)

	add("snippet:inspect", "" +
		"{{ .String | blue }}\n" +
		"{{ . | status }}\n" +
		"Description: {{.Description}}\n\n"+
		"Run count: {{.RunCount}}  Created: {{.Created | human}}\n\n" +
		"Snippet:\n\n{{ . | snippet }}\n\n" +
		"Preview:\n\n{{ . | preview }}\n\n" +
		"Tags: {{range $index, $element := .Tags}}{{if $index}}, {{end}} {{$element}}{{ end }}\n\n",
		template.FuncMap{"blue": blue,
			"human" : humanTime,
			"status" : statusString,
			"snippet": func(s *models.Snippet) string {
				return FmtSnippet(s, 100, 0)
			},
			"preview" : func(s *models.Snippet) string {
				return FmtOutPreview(s)
			},
	})

	add("pouch:not-deleted", "{{. | blue }} was NOT deleted.", template.FuncMap{"blue": blue})
	add("pouch:deleted", "{{. | blue }} was deleted.", template.FuncMap{"blue": blue})

	add("pouch:check-delete", "Are you sure you want to delete pouch {{. | yellow }}? [y/n] ", template.FuncMap{"yellow": yellow})
	add("pouch:created", "Pouch: {{. | blue }} created.", template.FuncMap{"blue": blue})
	add("pouch:renamed", "Pouch: {{. | blue }} renamed.", template.FuncMap{"blue": blue})
	add("pouch:locked", "🔒  pouch {{. | blue }} locked.", template.FuncMap{"blue": blue})
	add("pouch:unlocked", "🔓  pouch {{. | blue }} unlocked and public.", template.FuncMap{"blue": blue})
	add("pouch:not-locked", "Pouch: {{. | blue }} NOT locked.", template.FuncMap{"blue": blue})
	add("pouch:check-unlock", "Are you sure you want pouch 👝  {{. | blue }} public ? [y/n]", template.FuncMap{"blue": blue})

	// System
	add("system:upgraded", "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n   Successfully upgraded!  \n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n", nil)
	add("system:version", "kwk version:\n CLI: {{ .String | blue }} released {{ .Time | time }}\n API: {{ .Api.String | blue}}\n", template.FuncMap{"blue": blue, "time": humanTime })

	// Account
	add("account:signedup", "Welcome to kwk {{.Username | blue }}!\n You're signed in already.\n", template.FuncMap{"blue": blue})
	addColor("account:usernamefield", "Your Kwk Username: ", blue)
	addColor("account:passwordfield", "Your Password: ", blue)
	add("account:signedin", "Welcome back {{.Username | blue }}!\n", template.FuncMap{"blue": blue})
	addColor("account:signedout", "And you're signed out.\n", blue)
	add("account:profile", "You are: {{.Username | blue}}!\n", template.FuncMap{"blue": blue})

	add("dialog:choose", "{{. | multiChoice }}\n", template.FuncMap{"multiChoice": multiChoice})
	add("dialog:header", "{{.| blue }}\n", template.FuncMap{"blue": blue})

	add("env:changed", style.InfoDeskPerson+"  {{ \"env.yml\" | blue }} set to: {{.| blue }}\n", template.FuncMap{"blue": blue})

	addColor("account:signup:email", "What's your email? ", blue)
	addColor("account:signup:username", "Choose a great username: ", blue)
	addColor("account:signup:password", "And enter a password (1 num, 1 cap, 8 chars): ", blue)

	add("search:alpha", "\n\033[7m  \"{{ .Term }}\" found in {{ .Total }} results in {{ .Took }} ms  \033[0m\n\n{{range .Results}}{{ .Username }}{{ \"/\" }}{{ .Name | blue }}.{{ .Extension | subdued }}\n{{ . | formatSearchResult}}\n{{end}}", template.FuncMap{"formatSearchResult": alphaSearchResult, "blue": blue, "subdued": subdued})
	add("search:alphaSuggest", "\n\033[7m Suggestions: \033[0m\n\n{{range .Results}}{{ .Username }}{{ \"/\" }}{{ .Name | blue }}.{{ .Extension | subdued }}\n{{end}}\n", template.FuncMap{"blue": blue, "subdued": subdued})
	add("search:typeahead", "{{range .Results}}{{ .String }}\n{{end}}", nil)

	// errors
	add("validation:title", "{{. | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("validation:multi-line", " - {{ .Desc | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("validation:one-line", style.Warning+"  {{ .Desc | yellow }}\n", template.FuncMap{"yellow": yellow})

	add("api:not-authenticated", "{{ \"Please login to continue: kwk login\" | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("api:not-implemented", "{{ \"The kwk cli is a greater version than supported by kwk API.\" | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("api:denied", "{{ \"Permission denied\" | yellow }}\n", template.FuncMap{"yellow": yellow})
	addColor("api:error", "\n"+style.Fire+"  We have a code RED error. \n- To report type: kwk upload-errors \n- You can also try to upgrade: npm update kwkcli -g\n", red)
	addColor("api:not-available", style.Ambulance+"  Kwk is DOWN! Please try again in a bit.\n", yellow)
	add("api:exists", "{{ \"That item already exists.\" | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("free-text", "{{.}}", nil)
}

func humanTime(t int64) string {
	return humanize.Time(time.Unix(t, 0))
}

func add(name string, templateText string, funcMap template.FuncMap) {
	t := template.New(name)
	if funcMap != nil {
		t.Funcs(funcMap)
	}
	Templates[name] = template.Must(t.Parse(templateText))
}

func addColor(name string, text string, color ColorFunc) {
	add(name, fmt.Sprintf("{{ %q | color }}", text), template.FuncMap{"color": color})
}

func multiChoice(list []models.Snippet) string {
	var options string
	for i, v := range list {
		options = options + fmt.Sprintf("[%s] %s   ", style.Fmt(style.Cyan, i+1), v.SnipName.String())
	}
	return options
}

type ColorFunc func(int interface{}) string

func blue(in interface{}) string {
	return style.Fmt(style.Cyan, fmt.Sprintf("%v", in))
}

func yellow(in interface{}) string {
	return style.Fmt(style.Yellow, fmt.Sprintf("%v", in))
}

func red(in interface{}) string {
	return style.Fmt(style.Red, fmt.Sprintf("%v", in))
}

func subdued(in interface{}) string {
	return style.Fmt(style.Subdued, fmt.Sprintf("%v", in))
}

func uri(text string) string {
	text = strings.Replace(text, "https://", "", 1)
	text = strings.Replace(text, "http://", "", 1)
	text = strings.Replace(text, "www.", "", 1)
	if len(text) >= 40 {
		text = text[0:10] + "..." + text[len(text)-30:]
	}
	return text
}

func PrettyPrint(obj interface{}) {
	fmt.Println("")
	p, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Print(string(p))
	fmt.Print("\n\n")
}
