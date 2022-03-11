package main

import (
	"flag"
	"fmt"
	"github.com/ivankf/metro/pkg/parse"
	"io"
	"math/rand"
	"os"
	"time"
)

func (m *Main) Run(args ...string) error {
	name, args := parse.ParseCommandName(args)
	switch name {
	case "add":
		return nil
	case "edit":
		return nil
	case "list":
		return nil
	case "view":
		return nil
	case "grep":
		return nil
	case "update":
		return nil
	case "delete":
		return nil
	case "template":
		return nil
	case "import":
		return nil
	case "export":
		return nil
	case "erase":
		return nil
	case "sync":
		return nil
	case "help":
		fmt.Println(usage)
		return nil
	case "version":
		if err := NewVersionCommand().Run(args...); err != nil {
			return fmt.Errorf("version: %s", err)
		}
	default:
		fmt.Printf("Invalid command: %s\n", name)
		fmt.Println(usage)
	}

	return nil
}

type VersionCommand struct {
	Stdout io.Writer
	Stderr io.Writer
}

func NewVersionCommand() *VersionCommand {
	return &VersionCommand{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (cmd *VersionCommand) Run(args ...string) error {
	// Parse flags in case -h is specified.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Usage = func() {
		_, err := fmt.Fprintln(cmd.Stdout, versionUsage)
		if err != nil {
			return
		}
	}
	if err := fs.Parse(args); err != nil {
		return err
	}

	// Print version info.
	_, err := fmt.Fprintf(cmd.Stdout, "metro v%s (git: %s %s)\n", version, branch, commit)
	if err != nil {
		return err
	}

	return nil
}

var versionUsage = `Displays the metro version, build branch and git commit hash.

Usage: metro version
`

func main() {
	rand.Seed(time.Now().UnixNano())
	m := NewMain()
	if err := m.Run(os.Args[1:]...); err != nil {
		os.Exit(1)
	}

}

var (
	version string
	branch  string
	commit  string
)

func init() {
	if version == "" {
		version = "unknown"
	}
	if branch == "" {
		branch = "unknown"
	}
	if commit == "" {
		commit = "unknown"
	}
}

type Main struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewMain() *Main {
	return &Main{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

var usage = `usage: metro [-h] [--version] {add,edit,list,view,grep,update,delete,template,import,export,erase,sync,version,help} ...

A knowledge base organizer

positional arguments:
{add,edit,list,view,grep,update,delete,template,import,export,erase,sync,version,help}
                        commands
    add                 Add an artifact
    edit                Edit an artifact content
    list                Search for artifacts
    view                View artifacts
    grep                Grep through kb artifacts
    update              Update artifact properties
    delete              Delete artifacts
    template            Manage templates for artifacts
    import              Import a knowledge base
    export              Export the knowledge base
    erase               Erase the entire kb knowledge base
    sync                Synchronize the knowledge base with a remote git repository
    version             show program's version number and exit
    help                Show help of a particular command`
