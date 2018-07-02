package main

import (
	"fmt"
	"os"

	"github.com/bbozhev/lemp-cli/commands"
	"github.com/urfave/cli"
)

var (
	Version  = "0"
	CommitId = "0"
)

func main() {

	app := cli.NewApp()
	app.Name = "lemp"
	app.Usage = "Create, delete Nginx Vhosts"
	app.Author = "Boz - Bozhidar Bozhev boz@devops-bulgaria.io"
	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)
	app.Commands = commands.Commands
	app.Run(os.Args)
}
