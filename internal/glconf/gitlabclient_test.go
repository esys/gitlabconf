package glconf

import (
	"testing"
)

func TestReadVariable(t *testing.T) {
	expected := "ROOT_VAR2"
	document := ReadDocument("../../configs/gitlab-config-sample.yml")
	if document.RootGroups.Variables[1].Key != expected {
		t.Errorf("Variable is incorrect, got: '%s', expected: '%s'", document.RootGroups.Variables[1].Key, expected)
	}
}
