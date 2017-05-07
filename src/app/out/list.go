package out

import (
	"bytes"
	"fmt"
	"github.com/kwk-super-snippets/cli/src/models"
	"github.com/kwk-super-snippets/cli/src/style"
	"github.com/kwk-super-snippets/types"
	"github.com/rjarmstrong/tablewriter"
	"golang.org/x/text/unicode/norm"
	"io"
	"sort"
	"strings"
	"time"
)

var mainMarkers = map[string]string{
	"go": "func main() {",
}

type CodeLine struct {
	Margin string
	Body   string
}

func StatusText(s *types.Snippet) string {
	if s.Ext == "url" {
		return "bookmark"
	}
	if s.RunStatus == types.UseStatusSuccess {
		return "success"
	} else if s.RunStatus == types.UseStatusFail {
		return "error"
	}
	return "static"
}

func FStatus(s *types.Snippet, includeText bool) string {
	if s.RunStatus == types.UseStatusSuccess {
		if includeText {
			return style.Fmt256(style.ColorYesGreen, style.IconTick) + "  Success"
		}
		return style.Fmt256(style.ColorYesGreen, style.IconTick)
	} else if s.RunStatus == types.UseStatusFail {
		if includeText {
			return style.Fmt256(style.ColorBrightRed, style.IconBroke) + "  Error"
		}
		return style.Fmt256(style.ColorBrightRed, style.IconBroke)
	}
	return style.Fmt16(style.Subdued, "? ")
}

func printRoot(w io.Writer, r *models.ListView) {
	var all []interface{}
	for _, v := range r.Pouches {
		if v.Name != "" {
			all = append(all, v)
		}
	}
	for _, v := range r.Personal {
		all = append(all, v)
	}

	fmtHeader(w, r.Username, "", nil)
	fmt.Fprint(w, strings.Repeat(" ", 50), style.Fmt16(style.Subdued, "◉  "+models.Principal.Username+"    TLS12"))
	fmt.Fprint(w, style.TwoLines)
	w.Write(listHorizontal(all, &r.UserStats))

	if len(r.Snippets) > 0 {
		fmt.Fprintf(w, "\n%sLast:", style.Margin)
		printSnippets(w, r, true)
	}

	if clientIsNew(r.LastUpgrade) {
		fmt.Fprint(w, style.Fmt16(style.Subdued, fmt.Sprintf("\n\n%skwk auto-updated to %s %s", style.Margin, r.Version, style.Time(time.Unix(r.LastUpgrade, 0)))))
	} else {
		fmt.Fprintln(w, "")
	}
	fmt.Fprint(w, "\n\n")
}

func clientIsNew(t int64) bool {
	if t == 0 {
		return false
	}
	return t > (time.Now().Unix() - 60)
}

//func printCommunity(w *bytes.Buffer) {
//	fmt.Fprint(w, "\n", style.MARGIN, style.Fmt(style.Subdued, "Community"), "\n")
//	com := []interface{}{}
//	com = append(com, &models.Pouch{
//		Name:       style.Fmt(style.Cyan, "/kwk/") + "unicode",
//		Username:   "kwk",
//		PouchStats: models.PouchStats{Runs: 12},
//	}, &models.Pouch{
//		Name:       style.Fmt(style.Cyan, "/kwk/") + "news",
//		Username:   "kwk",
//		PouchStats: models.PouchStats{Runs: 12},
//	},
//		&models.Pouch{
//			Name:       style.Fmt(style.Cyan, "/kwk/") + "devops",
//			Username:   "kwk",
//			PouchStats: models.PouchStats{Runs: 12},
//		})
//	w.Write(listHorizontal(com, nil))
//	w.WriteString("\n")
//}

func printPouchHeadAndFoot(w io.Writer, list *models.ListView) {
	fmtHeader(w, list.Username, list.Pouch.Name, nil)
	fmt.Fprint(w, style.Margin, style.Margin, pouchIcon(list.Pouch, false))
	fmt.Fprint(w, "  ")
	fmt.Fprint(w, locked(list.Pouch.MakePrivate))
	fmt.Fprint(w, " Pouch")
	fmt.Fprint(w, style.Margin, style.Margin, len(list.Snippets), " snippets")
	fmt.Fprint(w, "\n")
}

func locked(locked bool) string {
	if locked {
		return "Locked (Private)"
	}
	return "Public"
}

func printPouchSnippets(w io.Writer, list *models.ListView) {
	if models.Prefs().Naked {
		fmt.Fprint(w, listNaked(list))
	} else {
		//sort.Slice(list.Snippets, func(i, j int) bool {
		//	return list.Snippets[i].Created <= list.Snippets[j].Created
		//})
		if list.Pouch != nil {
			printPouchHeadAndFoot(w, list)
		}
		printSnippets(w, list, false)
		if list.Pouch != nil && len(list.Snippets) > 10 && !models.Prefs().HorizontalLists {
			printPouchHeadAndFoot(w, list)
		}
		fmt.Fprint(w, "\n")
	}
}

const timeLayout = "2 Jan 15:04 06"

func listNaked(list *models.ListView) interface{} {
	w := &bytes.Buffer{}
	tbl := tablewriter.NewWriter(w)
	tbl.SetHeader([]string{"Name", "Username", "Pouch", "Ext", "Private", "Status", "Runs", "Views", "Deps", "LastActive", "Updated"})
	tbl.SetAutoWrapText(false)
	tbl.SetBorder(false)
	tbl.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	tbl.SetCenterSeparator("")
	tbl.SetColumnSeparator("")
	tbl.SetRowLine(false)
	tbl.SetAutoFormatHeaders(false)
	tbl.SetHeaderLine(false)
	tbl.SetColWidth(5)
	tbl.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range list.Snippets {
		var private string
		if v.Private {
			private = "private"
		} else {
			private = "public"
		}
		tbl.Append([]string{
			v.Name,
			v.Username,
			v.Pouch,
			v.Ext,
			private,
			StatusText(v),
			fmt.Sprintf("%d", v.Runs),
			fmt.Sprintf("%d", v.Views),
			fmt.Sprintf("%d", len(v.Dependencies)),
			v.RunStatusTime.Format(timeLayout),
			v.Created.Format(timeLayout),
		})
	}
	tbl.Render()
	return w.String()
}

func printSnippets(w io.Writer, list *models.ListView, fullName bool) {
	if models.Prefs() != nil && models.Prefs().HorizontalLists {
		sort.Slice(list.Snippets, func(i, j int) bool {
			return list.Snippets[i].Name < list.Snippets[j].Name
		})
		l := []interface{}{}
		for _, v := range list.Snippets {
			l = append(l, v)
		}
		fmt.Fprint(w, "\n\n"+string(listHorizontal(l, nil))+"\n\n")
		return
	}

	if len(list.Snippets) == 0 {
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, style.Margin)
		fmt.Fprint(w, style.Fmt16(style.Subdued, "<empty pouch>"))
		fmt.Fprint(w, style.TwoLines)
		fmt.Fprint(w, style.Margin)
		fmt.Fprint(w, style.Fmt16(style.Cyan, "Add new snippets to this pouch: "))
		if list.Pouch != nil {
			fmt.Fprintf(w, "`kwk new <snippet> %s/<name>.<ext>`", list.Pouch.Name)
		}
		fmt.Fprint(w, "\n")
		return
	}

	tbl := tablewriter.NewWriter(w)
	tbl.SetHeader([]string{"", "", ""})
	tbl.SetAutoWrapText(false)
	tbl.SetBorder(false)
	tbl.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	tbl.SetCenterSeparator("")
	tbl.SetColumnSeparator(" ")
	if models.Prefs().RowLines {
		tbl.SetRowLine(true)
	}

	tbl.SetAutoFormatHeaders(false)
	tbl.SetHeaderLine(true)
	tbl.SetColWidth(1)

	tbl.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	for i, v := range list.Snippets {
		// col1
		name := &bytes.Buffer{}
		name.WriteString(snippetIcon(v))
		name.WriteString("  ")
		sn := v.SnipName.String()
		if fullName {
			sn = v.String()
		}
		nt := style.Fmt256(style.ColorBrighterWhite, sn)
		name.WriteString(nt)
		if v.Description != "" {
			name.WriteString("\n\n")
			name.WriteString(style.Fmt256(style.ColorMonthGrey, style.FBox(v.Description, 25, 3)))
		}
		if v.Role == types.RoleEnvironment {
			name.WriteString("\n\n")
			name.WriteString(style.Fmt16(style.Subdued, "short-cut: kwk edit env"))
		}
		// col2
		var lines int
		if models.Prefs().AlwaysExpandRows {
			lines = models.Prefs().ExpandedRows
		} else {
			lines = models.Prefs().SlimRows
		}
		status := &bytes.Buffer{}
		runCount := fmtRunCount(v.Runs)
		status.WriteString(PadRight(runCount, " ", 21))
		status.WriteString(" ")
		if !v.RunStatusTime.IsZero() {
			h := PadLeft(style.Time(v.RunStatusTime), " ", 4)
			t := fmt.Sprintf("%s", style.Fmt256(239, h))
			status.WriteString(t)
		}
		//col3
		snip := FCodeview(v, 60, lines, (i+1)%2 == 0)
		if models.Prefs().RowSpaces {
			snip = snip + "\n"
		}
		if len(v.Preview) >= 10 {
			snip = snip + "\n\n" + style.Margin + style.Fmt256(style.ColorMonthGrey, style.FPreview(v.Preview, 120, 1))
		}

		tbl.Append([]string{
			name.String(),
			status.String(),
			snip,
		})
	}
	tbl.Render()

	//fmt.Fprint(w, style.Start)
	//fmt.Fprintf(w, "%dm", style.Subdued)
	//fmt.Fprint(w, style.MARGIN)
	//fmt.Fprintf(w,"Expand list: `kwk expand %s`", list.Pouch)
	//fmt.Fprint(w, style.MARGIN)
	//fmt.Fprint(w, style.End)
	////fmt.Fprint(w, style.Start)
	//fmt.Fprintf(w, "%dm", style.Subdued)
	//fmt.Fprint(w, style.MARGIN)
	//fmt.Fprintf(w, "%d of max 32 snippets in pouch", len(list.Snippets))
	//fmt.Fprint(w, style.End)
}

func PadRight(str, pad string, length int) string {
	if len(str) < length {
		return str + strings.Repeat(pad, length-len(str))
	}
	return str
}

func PadLeft(str, pad string, length int) string {
	if len(str) < length {
		return strings.Repeat(pad, length-len(str)) + str
	}
	return str
}

func snippetIcon(s *types.Snippet) string {
	icon := style.IconSnippet
	if s.IsApp() {
		icon = style.IconApp
	} else if s.Ext == "url" {
		icon = style.IconBookmark
	}
	if s.RunStatus == types.UseStatusSuccess {
		return style.Fmt256(122, icon)
	} else if s.RunStatus == types.UseStatusFail {
		return style.Fmt256(196, icon)
	}
	return style.Fmt256(style.ColorMonthGrey, icon)
}

func fmtRunCount(count int64) string {
	return fmt.Sprintf(style.Fmt256(247, "↻ %0d"), count)
}

func fmtHeader(w io.Writer, username string, pouch string, s *types.SnipName) {
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, style.Margin)
	fmt.Fprint(w, style.Start)
	fmt.Fprint(w, "7m")
	fmt.Fprint(w, " ❯ ")
	fmt.Fprint(w, types.KwkHost)
	fmt.Fprint(w, "/")
	if pouch == "" && s == nil {
		fmt.Fprint(w, style.Start)
		fmt.Fprint(w, "1m")
		fmt.Fprint(w, username)
		fmt.Fprint(w, " ")
		fmt.Fprint(w, style.End)
		return
	}
	fmt.Fprint(w, username)
	fmt.Fprint(w, "/")
	if s == nil {
		fmt.Fprint(w, style.Start)
		fmt.Fprint(w, "1m")
		fmt.Fprint(w, pouch)
		fmt.Fprint(w, " ")
		fmt.Fprint(w, style.End)
		return
	}
	if pouch != "" {
		fmt.Fprint(w, pouch)
		fmt.Fprint(w, "/")
	}
	fmt.Fprint(w, style.Start)
	fmt.Fprint(w, "1m")
	fmt.Fprint(w, s.String())
	fmt.Fprint(w, " ")
	fmt.Fprint(w, style.End)
}

func pad(width int, in string) *bytes.Buffer {
	buff := &bytes.Buffer{}
	diff := width - len([]rune(in))
	if diff > 0 {
		buff.WriteString(in)
		buff.WriteString(strings.Repeat(" ", diff))
	} else {
		var ia norm.Iter
		ia.InitString(norm.NFKD, in)
		nc := 0
		for !ia.Done() && nc < width {
			nc += 1
			buff.Write(ia.Next())
		}
	}
	return buff
}
