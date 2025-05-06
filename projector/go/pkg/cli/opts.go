package projector

import "github.com/hellflame/argparse"

type ProjectorOptions struct {
	Arguments []string
	Pwd       string
	Config    string
}

func GetOptions() (*ProjectorOptions, error) {
	parser := argparse.NewParser("projector", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "arguments", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	config := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &ProjectorOptions{
		Arguments: *args,
		Pwd:       *pwd,
		Config:    *config,
	}, nil
}
