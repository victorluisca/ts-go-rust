package projector_test

import (
	"reflect"
	"testing"

	projector "github.com/victorluisca/ts-go-rust/pkg/cli"
)

func getOptions(args []string) *projector.ProjectorOptions {
	opts := projector.ProjectorOptions{
		Arguments: args,
		Config:    "",
		Pwd:       "",
	}
	return &opts
}

func TestConfigPrint(t *testing.T) {
	t.Run("shoud print all", func(t *testing.T) {
		options := getOptions([]string{})

		config, err := projector.NewConfig(options)
		if err != nil {
			t.Errorf("expected to get no error %v", err)
		}

		if !reflect.DeepEqual([]string{}, config.Arguments) {
			t.Errorf("expected args to be an empty string array, but got %+v", config.Arguments)
		}

		if config.Operation != projector.Print {
			t.Errorf("expected operation to be %v but got %v", projector.Print, config.Operation)
		}
	})

	t.Run("shoud print key", func(t *testing.T) {
		options := getOptions([]string{"foo"})

		config, err := projector.NewConfig(options)
		if err != nil {
			t.Errorf("expected to get no error %v", err)
		}

		if !reflect.DeepEqual([]string{"foo"}, config.Arguments) {
			t.Errorf("expected args to be {'foo'} but got %+v", config.Arguments)
		}

		if config.Operation != projector.Print {
			t.Errorf("expected operation to be %v but got %v", projector.Print, config.Operation)
		}
	})
}

func TestConfigAdd(t *testing.T) {
	t.Run("should add key", func(t *testing.T) {
		options := getOptions([]string{"add", "foo", "bar"})

		config, err := projector.NewConfig(options)
		if err != nil {
			t.Errorf("expected to get no error %v", err)
		}

		if !reflect.DeepEqual([]string{"foo", "bar"}, config.Arguments) {
			t.Errorf("expected args to be {'foo', 'bar'}, but got %+v", config.Arguments)
		}

		if config.Operation != projector.Add {
			t.Errorf("expected operation to be %v but got %v", projector.Add, config.Operation)
		}
	})
}

func TestConfigDelete(t *testing.T) {
	t.Run("should delete a key", func(t *testing.T) {
		options := getOptions([]string{"del", "foo"})

		config, err := projector.NewConfig(options)
		if err != nil {
			t.Errorf("expected to get no error %v", err)
		}

		if !reflect.DeepEqual([]string{"foo"}, config.Arguments) {
			t.Errorf("expected args to be {'foo'}, but got %+v", config.Arguments)
		}

		if config.Operation != projector.Delete {
			t.Errorf("expected operation to be %v but got %v", projector.Delete, config.Operation)
		}
	})
}
