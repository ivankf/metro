package add

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func (cmd *Command) parseFlags(args ...string) error {
	config := NewConfig()
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&cmd.title, "title", "", "Title of the added artifact")
	fs.StringVar(&cmd.category, "category", "", "associated to the artifact")
	fs.StringVar(&cmd.tags, "tags", "", "Tags to associate to the artifact in the form \"tag1;tag2;...;tagN\"")
	fs.StringVar(&cmd.author, "author", config.author, "Author of the artifact")
	fs.StringVar(&cmd.status, "status", "", "Status of the artifact")
	fs.StringVar(&cmd.template, "template", config.template, "Template to apply to the artifact")
	fs.StringVar(&cmd.body, "body", config.body, "Body of the artifact")
	fs.Usage = cmd.printUsage
	fs.SetOutput(cmd.Stdout)

	if err := fs.Parse(args); err != nil {
		return err
	}
	return nil

}

type Command struct {
	StdoutLogger *log.Logger
	StderrLogger *log.Logger

	Stderr io.Writer
	Stdout io.Writer

	title    string
	category string
	tags     string
	author   string
	status   string
	template string
	body     string
}

func NewCommand() *Command {
	return &Command{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

func init() {
}

func (cmd *Command) Run(args ...string) error {
	cmd = NewCommand()
	if err := cmd.parseFlags(args...); err != nil {
		return err
	}

	return nil
}

func (cmd *Command) printUsage() {
	usage := `Usage:
  -author string
        -a AUTHOR, --author AUTHOR
  -body string
        Body of the artifact
  -category string
        associated to the artifact
  -status string

  -tags string
        Tags to associate to the artifact in the form "tag1;tag2;...;tagN"
  -template string
        Template to apply to the artifact
  -title string
        Title of the added artifact`
	_, err := fmt.Fprintf(cmd.Stdout, usage)
	if err != nil {
		return
	}
}
