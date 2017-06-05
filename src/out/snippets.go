package out

import (
	"fmt"
	"github.com/kwk-super-snippets/types"
	"github.com/kwk-super-snippets/types/vwrite"
	"io"
	"strings"
	"github.com/kwk-super-snippets/cli/src/style"
)

func SnippetDescriptionUpdated(uri string, desc string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s description updated:\n%s\n\n\n", uri, desc)
	}))
}

func SnippetClonedAs(newName string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "Cloned as %s\n\n", newName)
	}))
}

func SnippetCreated(s *types.Snippet) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s %s created \n\n", snippetIcon(s), s.Alias.URI())
	}))
}

func SnippetCat(s *types.Snippet) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintln(w, s.Content)
	})
}

func SnippetAmbiguousCat(snippets []*types.Snippet) vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintln(w, "That snippet name is ambiguous, please choose one of the following: ")
	}))
}

func SnippetEdited(s *types.Snippet) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "Successfully updated %s %s\n\n", snippetIcon(s), s.Alias.VersionURI())
	}))
}

func SnippetNoChanges(s *types.Snippet) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "No changes to %s %s\n\n", snippetIcon(s), s.Alias.VersionURI())
	}))
}

func SnippetEditing(s *types.Snippet) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\n%sEditing:  %s %s...\n\n", style.Margin, snippetIcon(s), s.Alias.VersionURI())
		fmt.Fprintf(w, "%sCTRL+C to CANCEL | Any key to COMMIT\n", style.Margin)
	})
}

func SnippetEditNewPrompt(uri string) vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s doesn't exist - would you like create it? [y/n] \n", uri)
	}))
}

func SnippetList(prefs *Prefs, list *types.ListResponse) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		printPouchSnippets(w, prefs, list)
	})
}

func PrintRoot(prefs *Prefs, cli *types.AppInfo, rr *types.RootResponse, u *types.User) vwrite.Handler {
	return vwrite.HandlerFunc(func(w io.Writer) {
		printRoot(w, prefs, cli, rr, u)
	})
}

func Tagged(uri string, tags []string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s tagged with: %s\n", uri, strings.Join(tags, ", "))
	}))
}

func UnTag(uri string, tags []string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "Tags: %s removed from %s\n", strings.Join(tags, ", "), uri)
	}))
}

func SnippetRenamed(originalUri string, newUri string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s renamed to %s", originalUri, newUri)
	}))
}

func SnippetPatched(uri string) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s patched", uri)
	}))
}

func SnippetCheckDelete(snipNames []*types.SnipName) vwrite.Handler {
	return Warn(vwrite.HandlerFunc(func(w io.Writer) {
		printSnipNames(w, snipNames)
		fmt.Fprint(w, "\nAre you sure you want to delete these snippets? [y/n] ")
	}))
}

func SnippetsDeleted(snipNames []*types.SnipName) vwrite.Handler {
	return Success(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprint(w, "Snippets deleted: ")
		printSnipNames(w, snipNames)
		fmt.Fprintln(w, "")
	}))
}

func SnippetsNotDeleted(snipNames []*types.SnipName) vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprint(w, "Snippets NOT deleted: ")
		printSnipNames(w, snipNames)
		fmt.Fprintln(w, "")
	}))
}

func SnippetsMoved(snipNames []*types.SnipName, pouch string) vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		printSnipNames(w, snipNames)
		fmt.Fprintf(w, " moved to pouch %s\n", pouch)
	}))
}

func SnippetPouchCreatePrompt() vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		fmt.Fprint(w, "Would you like to create the snippet in a new pouch? [y/n] ")
	}))
}

func SnippetView(prefs *Prefs, s *types.Snippet) vwrite.Handler {
	return Info(vwrite.HandlerFunc(func(w io.Writer) {
		printSnippetView(w, prefs, s)
	}))
}
