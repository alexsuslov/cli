package cli

import "log"

type ActionFn func(args ...string)error

type Cli struct{
	Actions map[string]ActionFn
	Helps []string
	logger bool
	err error
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

func New()*Cli{
	return &Cli{
		Actions: map[string]ActionFn{},
		Helps: []string{},
	}
}

func (Cli *Cli)AddAction(command string, action ActionFn, help string)*Cli{
	Cli.Actions[command] = action
	Cli.Helps = append(Cli.Helps, help)
	return Cli
}
