package projector_test

import (
	"testing"

	projector "github.com/victorluisca/ts-go-rust/pkg/cli"
)

func getData() *projector.ProjectorData {
	return &projector.ProjectorData{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"bar": "baz",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.ProjectorData) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Arguments: []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "",
		},
		data,
	)
}

func test(t *testing.T, projector *projector.Projector, key, value string) {
	v, ok := projector.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}

	if value != v {
		t.Errorf("expected to find \"%v\" but received %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)
	test(t, projector, "foo", "bar3")
	test(t, projector, "bar", "baz")
}

func TestSetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)

	test(t, projector, "foo", "bar3")
	projector.SetValue("foo", "bar4")
	test(t, projector, "foo", "bar4")

	projector.SetValue("bar", "foo")
	test(t, projector, "bar", "foo")

	projector = getProjector("/", data)
	test(t, projector, "bar", "baz")
}

func TestDeleteValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)
	test(t, projector, "foo", "bar3")

	projector.DeleteValue("foo")
	test(t, projector, "foo", "bar2")

	projector.DeleteValue("bar")
	test(t, projector, "bar", "baz")
}
