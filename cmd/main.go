package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/akamensky/argparse"
	"github.com/esys/gitlabconf/internal/glconf"
)

func main() {
	log.SetLevel(log.DebugLevel)
	parser := argparse.NewParser("gilabconf", "Configure a gitlab instance")
	configPath := parser.String("configfile", "string", &argparse.Options{Required: false, Default: "gitlab-config.yml",
		Help: "Config file path"})

	log.Debugf("Config filepath is %s", *configPath)

	document, err := glconf.ReadDocument(*configPath)
	if err != nil {
		log.Fatalf("cannot read config file %s: %s", *configPath, err)
	}
	log.Debugf("Document is  %+v", document)

	client := glconf.NewGitlabClient(document.Gitlab.Url, document.Gitlab.Token)

	// first, create variables at root level
	rootKeys := client.ListRootGroupKeys()
	for _, key := range rootKeys {
		client.CreateVariables(key, document.RootGroups.Variables)
	}

	// then create variables at subgroup level
	for _, group := range document.Groups {

	}
}
