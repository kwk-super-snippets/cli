// Package out provides the UX for kwk super snippets CLI
package out

import (
	"fmt"
	"github.com/lunixbochs/vtclean"
	"github.com/rjarmstrong/kwk-types"
	"github.com/rjarmstrong/kwk-types/vwrite"
	"github.com/rjarmstrong/kwk/src/style"
	"io"
	"text/tabwriter"
	"time"
)

func FreeText(text string) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprint(w, text)
	})
}

func formatTime(millis int64) string {
	t := time.Unix(millis/1000, 0)
	return style.Time(t)
}

func Dashboard(prefs *Prefs, cli *types.AppInfo, rr *types.RootResponse, u *types.User) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		PrintRoot(prefs, cli, rr, u).Write(w)
	})
}

func SignedOut() vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprint(w, "\n")
		fmt.Fprintf(w, "%s %s  ◥     ◤  %s\n",
			style.Margin,
			style.Fmt256(style.ColorPouchCyan, "◤"),
			style.Fmt256(style.ColorMonthGrey, "◤"))

		fmt.Fprintf(w, "%s ◣    %s ◢    ◣\n", style.Margin, style.Fmt256(style.ColorBrightRed, "◣"))
		fmt.Fprint(w, "\n")
		fmt.Fprintf(w, "%s %s\n", style.Margin, "super snippets")
		fmt.Fprint(w, "\n")
		fmt.Fprintf(w, "%skwk signup    To get started with kwk\n", style.Margin)
		fmt.Fprintf(w, "%skwk signin    For existing users\n", style.Margin)
		fmt.Fprint(w, "\n")
	})
}

func printSnipNames(w io.Writer, snipNames []*types.SnipName) {
	for i, v := range snipNames {
		fmt.Fprintf(w, "%s", v.String())
		if i-1 < len(snipNames) {
			fmt.Fprint(w, ", ")
		}
	}
}

func multiChoice(w io.Writer, in interface{}) {
	list := in.([]*types.Snippet)
	fmt.Fprint(w, "\n")
	if len(list) == 1 {
		fmt.Fprintf(w, "%sDid you mean: %s? y/n\n\n", style.Margin, style.Fmt256(style.ColorPouchCyan, list[0].String()))
		return
	}
	t := tabwriter.NewWriter(w, 5, 1, 3, ' ', tabwriter.TabIndent)
	for i, v := range list {
		if i%3 == 0 {
			t.Write([]byte(style.Margin))
		}
		fmt256 := style.Fmt16(style.Cyan, i+1)
		t.Write([]byte(fmt.Sprintf("%s %s", fmt256, v.Alias.FileName())))
		x := i + 1
		if x%3 == 0 {
			t.Write([]byte("\n"))
		} else {
			t.Write([]byte("\t"))
		}
	}
	t.Write([]byte("\n\n"))
	t.Flush()
	fmt.Fprint(w, style.Margin+style.Fmt256(style.ColorPouchCyan, "Please select a snippet: "))
}

func fPreview(in string, prefs *Prefs, wrapAt int, lines int) string {
	if prefs.DisablePreview {
		return ""
	}
	in = vtclean.Clean(in, false)
	return style.FBox(in, wrapAt, lines) + style.End
}
