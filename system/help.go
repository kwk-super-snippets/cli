package system

import (
	"fmt"
	"io"
	"github.com/fatih/color"
)

func Help(w io.Writer, template string, data interface{}) {
	c := color.New(color.FgCyan).Add(color.Bold)
	c.Printf("\n ===================================================================== ")
	c.Printf("\n ~~~~~~~~~~~~~~~~~~~~~~~~   KWK Power Links.  ~~~~~~~~~~~~~~~~~~~~~~~~ \n\n")
	c.Printf(" The ultimate URI manager. Create short and memorable codes called\n")
	c.Printf(" `kwklinks` to store URLs, computer paths, AppLinks etc.\n\n")
	c.Printf(" Usage: kwk [kwklink|cmd] [subcmd] [args]\n")
	fmt.Print("\n e.g.: `kwk open got-spoilers` to open all G.O.T. spoiler websites.\n")

	c.Printf("\n Commands:\n")
	fmt.Print("    <kwklink,..>                      - Open and navigate to uris in default browser etc.\n")
	fmt.Print("    new        <uri> [name]           - Create a new kwklink, optionally provide a memorable name\n")

	fmt.Print("    list       [tag,..] [and|or|not]  - List kwklinks, filter by tags\n")
	fmt.Print("    search     [term] [tag]           - * Search kwklinks and their metadata by keyword, filter by tags\n")
	fmt.Print("    suggest    <uri>                  - * List suggested kwklinks or tags for the given uri\n")
	fmt.Print("    tag        <kwklink> [tag,..]     - Add tags to a kwklink\n")
	fmt.Print("    open       <tag>,.. [page]        - Open links for given tags, 5 at a time\n")
	fmt.Print("    untag      <kwklink> [tag,..]     - Remove tags from a kwklink\n")
	fmt.Print("    inspect    <kwklink>              - Look at the details of a kwk link\n")
	fmt.Print("    note       <kwklink> [text]       - Add a description to a link\n")
	fmt.Print("    version    <kwklink> [text]       - \n")

	fmt.Print("    update\n")
	fmt.Print("      kwklink  <kwklink> <kwklink>    - Update kwklink name <old> <new>\n")
	fmt.Print("      uri      <kwklink> <uri>        - * Update uri, auto increments the version\n")
	fmt.Print("    delete     <kwklink>              - * Deletes kwklink with warning prompt. Will give 404.\n")
	fmt.Print("    detail     <kwklink>              - Get details and info\n")
	fmt.Print("    covert     <kwklink>              - Open in covert (incognito mode)\n")
	fmt.Print("    get        <kwklink> [page]       - Gets URIs without navigating. (Copies first to clipboard)\n")

	c.Printf("\n Community:\n")
	fmt.Print("    cd    	  [username]             - *Navigate to another users kwklinks. Used when searching or listing \n")
	fmt.Print("    pin    	  [name][kwklink][tag]   - *Pin current version of someone elses kwklink to your own list \n")
	fmt.Print("    comment    [tag][kwklink][text]   - *Comment on tag or kwklink\n")
	fmt.Print("    profile    [username]             - *View a profile ??:\n")
	fmt.Print("    share      [kwklink|tag] [handle] - *Share with someone with a given handle:\n")
	fmt.Print("                                        twitter, email, kwk username\n")

	c.Printf("\n Analytics:\n")
	fmt.Print("    stats      [kwklink][tag]         - *Get statistics and filter by kwklink or tag\n")

	c.Printf("\n Account:\n")
	fmt.Print("    login      <username><password>   - Login with secret key.\n")
	fmt.Print("    me                                - View profile create ascii profile.\n")
	fmt.Print("    logout                            - Clears locally cached secret key.\n")
	fmt.Print("    signup     <email> <password> <username>  - Sign-up with a username.\n")

	fmt.Print("\n\n  * Filter only Tags: today yesterday thisweek lastweek thismonth lastmonth thisyear lastyear")
	fmt.Print("\n ** kwklinks are case sensitive")

	fmt.Print("\n\n More Commands: `kwk [admin|device] help`")

	//Day II: fmt.Print("	lock       <kwklink> <pin>          - Lock a kwklink with a pin\n")
	//Day II: fmt.Print("	subscribe  <domain>	            - Subscribe with custom domain. Free for 30 days.\n")
	//Day II: fmt.Print("	rate <kwklink> 9	            - Subscribe with custom domain. Free for 30 days.\n")
	//Day II: fmt.Print("	note <kwklink> "I like this one	    - Subscribe with custom domain. Free for 30 days.\n")

	//fmt.Printf("\n Admin:\n")
	//fmt.Printf("	cache       ls                  - List locally cached kwklinks.\n")
	//fmt.Printf("	cache       clear               - Clears any locally cached data.\n")
	//fmt.Printf("	upgrade                    	- Downloads and upgrades kwk cli client.\n")
	//fmt.Printf("	config      warn  [on|off]      - Warns if attempting to open dodgy kwklink.\n")
	//fmt.Printf("	config      quiet [on|off]      - Prevents links from being printed to console.\n")
	//fmt.Printf("	version                    	\n")
	c.Printf("\n ===================================================================== \n\n")
}