package main

import (
	"fmt"
	"os"
	"regexp"
	"unidriver/Godeps/_workspace/src/github.com/codegangsta/cli"
	//	"unidriver/Godeps/_workspace/src/github.com/k0kubun/pp"
	"unidriver/commands"
	"unidriver/parsers"
)

func main() {

	cli.AppHelpTemplate = AppHelpTemplate
	app := cli.NewApp()
	app.Name = "unidriver"
	app.Version = Version
	app.HideHelp = true
	app.Usage = "E2E Testing Application"
	app.Author = "aqafiam"
	app.Before = doBefore
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		remoteFlag,
	}
	app.Run(os.Args)
}

func doBefore(c *cli.Context) error {

	args := c.Args()

	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}
	for _, arg := range args {
		ok, _ := regexp.MatchString(".ya?ml$", arg)
		if !ok {
			fmt.Println(arg + "is not yaml file.")
			cli.ShowAppHelp(c)
			os.Exit(2)
		}
	}

	return nil
}

func doMain(c *cli.Context) {

	datas := parsers.ParseYaml(c.Args())

	commands.Validate(datas)
	commands.Do(c.String("remote"), datas)

}

var remoteFlag = cli.StringFlag{
	Name:  "remote, r",
	Value: "http://localhost:4444/wd/hub",
	Usage: "Remote WebDriver Server's URL",
}

const Version string = "0.1.0"

const AppHelpTemplate string = `USAGE: {{.Name}} [options] [testcase.yml...]

OPTIONS:
{{range .Flags}}  {{.}}
{{end}}
`
