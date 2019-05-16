package glconf

import (
	"testing"
)

func TestReadDocument(t *testing.T) {
	expected := "http://localhost/api/v4"
	document := ReadDocument("../../configs/gitlab-config-sample.yml")
	if document.Gitlab.Url != expected {
		t.Errorf("Url is incorrect, got: '%s', expected: '%s'", document.Gitlab.Url, expected)
	}
	if len(document.RootGroups.Variables) != 2 {
		t.Errorf("Variables count is incorrect, got: %d, expected: %d", len(document.RootGroups.Variables), 2)
	}
}
