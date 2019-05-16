package main

import (
	"testing"

	"github.com/esys/gitlabconf/internal/glconf"
)

const defaultConfigPath = "gitlab-config.yml"
const sampleConfigPath = "../../configs/gitlab-config-sample.yml"

func TestReadDocument(t *testing.T) {
	expected := "http://localhost/api/v4"
	document := glconf.ReadDocument(sampleConfigPath)
	if document.Gitlab.Url != expected {
		t.Errorf("Url is incorrect, got: '%s', expected: '%s'", document.Gitlab.Url, expected)
	}
	if len(document.RootGroups.Variables) != 2 {
		t.Errorf("Variables count is incorrect, got: %d, expected: %d", len(document.RootGroups.Variables), 2)
	}
}

func TestReadVariable(t *testing.T) {
	expected := "ROOT_VAR2"
	document := glconf.ReadDocument(sampleConfigPath)
	if document.RootGroups.Variables[1].Key != expected {
		t.Errorf("Variable is incorrect, got: '%s', expected: '%s'", document.RootGroups.Variables[1].Key, expected)
	}
}
