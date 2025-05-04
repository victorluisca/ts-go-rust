package projector

import "github.com/hellflame/argparse"

type ProjectorOptions struct {
	Pwd       string
	Config    string
	Arguments []string
}

func GetOptions() (*ProjectorOptions, error) {
	parser := argparse.NewParser("projector", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	config := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	args := parser.Strings("a", "arguments", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &ProjectorOptions{
		Pwd:       *pwd,
		Config:    *config,
		Arguments: *args,
	}, nil
}
