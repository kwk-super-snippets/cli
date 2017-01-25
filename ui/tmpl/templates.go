package tmpl

import (
	"bitbucket.com/sharingmachine/kwkcli/models"
	"github.com/dustin/go-humanize"
	_ "github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	_ "github.com/olekukonko/tablewriter"
	"bitbucket.com/sharingmachine/kwkcli/ui/style"
	"encoding/json"
	"text/template"
	"strings"
	"bytes"
	"fmt"
)

var Templates = map[string]*template.Template{}

func init() {
	// Aliases
	add("snippet:delete", "Are you sure you want to delete {{.FullName | yellow }}? [y/n] ", template.FuncMap{"yellow": yellow})
	add("snippet:deleted", "{{.FullName | blue }} deleted.", template.FuncMap{"blue": blue})
	add("snippet:updated", "Description updated:\n{{ .Description | blue }}", template.FuncMap{"blue": blue})
	addColor("api:not-found", "Not found\n", blue)
	add("snippet:cloned", "Cloned as {{.Username}}/{{.FullName | blue}}\n", template.FuncMap{"blue": blue})
	add("snippet:new", "{{.FullName | blue }} created "+style.OpenLock+"\n", template.FuncMap{"blue": blue})
	add("snippet:newprivate", "{{.FullName | blue }} created "+style.Lock+"\n", template.FuncMap{"blue": blue})
	add("snippet:cat", "{{.Snip | blue}}", template.FuncMap{"blue": blue})
	add("snippet:edited", "Successfully updated {{ .FullName | blue }}", template.FuncMap{"blue": blue})
	add("snippet:editing", "{{ \"Editing... \" | blue }}\nPlease edit file and save.\n - NB: After saving, no changes will be saved until running kwk edit <name> again.\n - Ctrl+C to abort.\n", template.FuncMap{"blue": blue})

	add("snippet:ambiguouscat", "That snippet is ambiguous please run it again with the extension:\n{{range .Items}}{{.FullName | blue }}\n{{ end }}", template.FuncMap{"blue": blue})
	add("snippet:list", "{{. | listSnippets }}", template.FuncMap{"listSnippets": listSnippets })
	add("snippet:tag", "{{.FullName | blue }} tagged.", template.FuncMap{"blue": blue})
	add("snippet:untag", "{{.FullName | blue }} untagged.", template.FuncMap{"blue": blue})
	add("snippet:renamed", "{{.fullName | blue }} renamed to {{.newFullName | blue }}", template.FuncMap{"blue": blue})
	add("snippet:madeprivate", "{{.fullName | blue }} made private "+style.Lock, template.FuncMap{"blue": blue})
	add("snippet:patched", "{{.FullName | blue }} patched.", template.FuncMap{"blue": blue})
	add("snippet:notdeleted", "{{.FullName | blue }} was NOT deleted.", template.FuncMap{"blue": blue})
	add("snippet:inspect",
		"\n{{range .Items}}"+"Full name: {{.Username}}/{{.FullName}}"+"\nSnippet: {{.Snip}}"+"\nVersion: {{.Version}}"+"\nTags: {{range $index, $element := .Tags}}{{if $index}}, {{end}} {{$element}}{{ end }}"+"\nWeb: \033[4mhttp://www.kwk.co/{{.Username}}/{{.FullName}}\033[0m"+"\nDescription: {{.Description}}"+"\nRun count: {{.RunCount}}"+"\nClone count: {{.CloneCount}}"+"\n{{ end }}\n\n", nil)

	// System
	add("system:upgraded", "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n   Successfully upgraded!  \n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n", nil)
	add("system:version", "kwk version:\n CLI: {{ .cliVersion | blue}}\n API: {{ .apiVersion | blue}}\n", template.FuncMap{"blue": blue})
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

	// errors
	add("validation:title", "{{. | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("validation:multi-line", " - {{ .Desc | yellow }}\n", template.FuncMap{"yellow": yellow})
	add("validation:one-line", style.Warning+"  {{ .Desc | yellow }} {{ .Code | yellow }}\n", template.FuncMap{"yellow": yellow})

	add("api:not-authenticated", "{{ \"Please login to continue: kwk login\" | yellow }}\n", template.FuncMap{"yellow": yellow})
	addColor("api:error", style.Fire+"  We have a code RED error. \n- To report type: kwk upload-errors \n- You can also try to upgrade: npm update kwkcli -g\n", red)
	addColor("api:not-available", style.Ambulance+"  Kwk is DOWN! Please try again in a bit.\n", yellow)
	add("api:exists", "{{ \"That item already exists.\" | yellow }}\n", template.FuncMap{"yellow": yellow})
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
		options = options + fmt.Sprintf("[%s] %s   ", style.Colour(style.LightBlue, i+1), v.FullName)
	}
	return options
}

func listSnippets(list *models.SnippetList) string {
	buf := new(bytes.Buffer)

	fmt.Fprint(buf, style.Colour(style.LightBlue, "\nkwk.co/"+list.Username+"\n\n"))

	tbl := tablewriter.NewWriter(buf)
	tbl.SetHeader([]string{"Name", "Version", "Snippet", "Tags", "Runs", "Clones", ""})
	tbl.SetAutoWrapText(false)
	tbl.SetBorder(false)
	tbl.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	tbl.SetCenterSeparator("")
	tbl.SetColumnSeparator("")
	tbl.SetAutoFormatHeaders(false)
	tbl.SetHeaderLine(true)
	tbl.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range list.Items {
		var tags = []string{}
		for _, v := range v.Tags {
			if v == "error" {
				tags = append(tags, style.Colour(style.Pink, v))
			} else {
				tags = append(tags, v)
			}
		}

		var snip string
		var name string

		name = style.Colour(style.LightBlue, v.Name) + style.Colour(style.Subdued, "."+v.Extension)
		if v.Private {
			name = style.Colour(style.Subdued, ".") + name
			if v.Role == models.RolePreferences {
				snip = style.Colour(style.Yellow, `(Local prefs) 'kwk edit prefs'`)
			} else if v.Role == models.RoleEnvironment {
				snip = style.Colour(style.Yellow, `(Runtime environment) 'kwk edit env'`)
			} else {
				snip = style.Colour(style.Subdued, `(Private)`)
			}
		} else {
			snip = fmt.Sprintf("%s", uri(v.Snip))
		}

		tbl.Append([]string{
			name,
			fmt.Sprintf("%d", v.Version),
			snip,
			strings.Join(tags, ", "),
			fmt.Sprintf("%d", v.RunCount),
			fmt.Sprintf("%d", v.CloneCount),
			humanize.Time(v.Created),
		})

	}
	tbl.Render()

	if len(list.Items) == 0 {
		fmt.Println(style.Colour(style.Yellow, "Create some snippets to fill this view!\n"))
	}
	fmt.Fprintf(buf, "\n%d of %d records\n\n", len(list.Items), list.Total)
	fmt.Fprint(buf, "\n\n")

	return buf.String()
}

func getConfigName(name string) string {
	return "." + strings.Split(name, "_")[1]
}

func alphaSearchResult(result models.SearchResult) string {
	if result.Highlights == nil {
		result.Highlights = map[string]string{}
	}
	if result.Highlights["snip"] == "" {
		result.Highlights["snip"] = result.Snip
	}
	lines := highlightsToLines(result.Highlights)
	f := ""
	for _, line := range lines {
		f = f + line.Key[:4] + "\u2847  " + line.Line + "\n"
	}
	f = style.Colour(style.Subdued, f)
	f = style.ColourSpan(40, f, "<em>", "</em>", style.Subdued)
	return f
}

func highlightsToLines(highlights map[string]string) []SearchResultLine {
	allLines := []SearchResultLine{}
	for k, v := range highlights {
		lines := strings.Split(v, "\n")
		for _, line := range lines {
			allLines = append(allLines, SearchResultLine{Key: k, Line: line})
		}
	}
	return allLines
}

type SearchResultLine struct {
	Key  string
	Line string
}

type ColorFunc func(text string) string

func blue(text string) string {
	return style.Colour(style.LightBlue, text)
}

func yellow(text string) string {
	return style.Colour(style.Yellow, text)
}

func red(text string) string {
	return style.Colour(style.Red, text)
}

func subdued(text string) string {
	return style.Colour(style.Subdued, text)
}

func uri(text string) string {
	text = strings.Replace(text, "https://", "", 1)
	text = strings.Replace(text, "http://", "", 1)
	text = strings.Replace(text, "www.", "", 1)
	text = strings.Replace(text, "\n", " ", -1)
	if len(text) >= 40 {
		text = text[0:10] + style.Colour(style.Subdued, "...") + text[len(text) - 30:]
	}
	if text == "" {
		text = "<empty>"
	}
	return text
}

func PrettyPrint(obj interface{}) {
	fmt.Println("")
	p, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Print(string(p))
	fmt.Print("\n\n")
}
