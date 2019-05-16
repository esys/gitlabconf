package glconf

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Gitlab struct {
	Url   string
	Token string
}

type Variable struct {
	Key       string
	Value     string
	Protected bool
	Overwrite bool
}

type Group struct {
	Key       string
	Variables []Variable
}

type RootGroups struct {
	Variables []Variable
}

type Document struct {
	Gitlab     Gitlab
	RootGroups Group `yaml:"rootGroups"`
	Groups     []Group
}

func (doc *Document) Unmarshal(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), doc)
	if err != nil {
		return err
	}
	return nil
}

func ReadDocument(configPath string) (*Document, error) {
	document := Document{}
	err := document.Unmarshal(configPath)
	if err != nil {
		log.Fatalf("error when trying to unmarshal config file '%s': %s", configPath, err)
		return nil, err
	}

	return &document, nil
}
