package projector

import (
	"encoding/json"
	"maps"
	"os"
	"path"
)

type ProjectorData struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config *Config
	data   *ProjectorData
}

func CreateProjector(config *Config, data *ProjectorData) *Projector {
	return &Projector{
		config: config,
		data:   data,
	}
}

func (p *Projector) GetValue(key string) (string, bool) {
	current := p.config.Pwd
	previous := ""

	output := ""
	found := false

	for current != previous {
		if dir, ok := p.data.Projector[current]; ok {
			if value, ok := dir[key]; ok {
				output = value
				found = true
				break
			}
		}
		previous = current
		current = path.Dir(current)
	}

	return output, found
}

func (p *Projector) GetValueAll() map[string]string {
	output := map[string]string{}
	paths := []string{}
	current := p.config.Pwd
	previous := ""

	for current != previous {
		paths = append(paths, current)
		previous = current
		current = path.Dir(current)
	}

	for i := len(paths) - 1; i >= 0; i-- {
		if dir, ok := p.data.Projector[paths[i]]; ok {
			maps.Copy(output, dir)
		}
	}
	return output
}

func (p *Projector) SetValue(key, value string) {
	pwd := p.config.Pwd

	if _, ok := p.data.Projector[pwd]; !ok {
		p.data.Projector[pwd] = map[string]string{}
	}

	p.data.Projector[pwd][key] = value
}

func (p *Projector) DeleteValue(key string) {
	pwd := p.config.Pwd

	if dir, ok := p.data.Projector[pwd]; ok {
		delete(dir, key)
	}
}

func (p *Projector) Save() error {
	dir := path.Dir(p.config.Config)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	json, err := json.Marshal(p.data)
	if err != nil {
		return err
	}

	os.WriteFile(p.config.Config, json, 0755)
	return nil
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		config: config,
		data: &ProjectorData{
			Projector: map[string]map[string]string{},
		},
	}
}

func NewProjector(config *Config) *Projector {
	if _, err := os.Stat(config.Config); err == nil {
		contents, err := os.ReadFile(config.Config)
		if err != nil {
			return defaultProjector(config)
		}

		var data ProjectorData
		err = json.Unmarshal(contents, &data)
		if err != nil {
			return defaultProjector(config)
		}

		return &Projector{
			data:   &data,
			config: config,
		}
	}

	return defaultProjector(config)
}
