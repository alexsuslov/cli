package cli

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type ActionFn func(args []string) error

type Cli struct {
	title    string
	Values   map[string]string
	Actions  map[string]ActionFn
	Template map[string]*template.Template
	Helps    []string
	logger   bool
	err      error
}

func (Cli *Cli) IsErr(errs ...error) bool {
	if Cli.err != nil {
		return true
	}
	for _, err := range errs {
		if err != nil {
			Cli.err = err
			return true
		}
	}
	return false
}

func (Cli Cli) Println(v ...interface{}) {
	if Cli.logger {
		log.Println(v...)
	}
}

func (Cli Cli) Printf(format string, v ...interface{}) {
	if Cli.logger {
		log.Printf(format, v...)
	}
}

func (Cli Cli) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// New cli with Title in help
func New(name string, title string) *Cli {
	cli := &Cli{
		title:    title,
		Template: map[string]*template.Template{},
		Values: map[string]string{
			"name": name},
		Actions: map[string]ActionFn{},
	}
	return cli.AddAction("help", cli.Help, help)
}

func (Cli *Cli) AddAction(command string, action ActionFn, help string) *Cli {
	Cli.Actions[command] = action
	Cli.Template[command] = template.Must(template.New("command").Parse(help))
	return Cli
}

func (Cli *Cli) Action(args []string) error {
	if len(args) > 1 {
		if action, ok := Cli.Actions[args[1]]; ok {
			return action(args)
		}
	}
	return Cli.Help(args)
}

var help = `> {{.name}} help
return this help
`

func (Cli *Cli) Help(args []string) error {
	fmt.Println(Cli.title)
	fmt.Println("---")
	for _, t := range Cli.Template {
		err := t.Execute(os.Stdout, Cli.Values)
		if err != nil {
			return err
		}
	}
	return nil
}
