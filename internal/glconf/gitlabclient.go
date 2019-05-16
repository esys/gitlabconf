package glconf

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/xanzy/go-gitlab"
)

type GitlabClient struct {
	gitClient *gitlab.Client
	url       string
	token     string
}

func NewGitlabClient(url string, token string) *GitlabClient {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(url)
	return &GitlabClient{gitClient: git, url: url, token: token}
}

func (client *GitlabClient) ListRootGroupKeys() []string {
	groups, _, err := client.gitClient.Groups.ListGroups(nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var rootGroupKeys []string
	for _, group := range groups {
		if !strings.Contains(group.FullPath, "/") {
			rootGroupKeys = append(rootGroupKeys, group.FullPath)
		}
	}

	log.Printf("Root group keys: %+v", rootGroupKeys)
	return rootGroupKeys
}

func (client *GitlabClient) CreateVariables(groupKey string, variables []Variable) error {
	existingVars, _, err := client.gitClient.GroupVariables.ListVariables(groupKey)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	var existingVarKeys []string
	for _, v := range existingVars {
		existingVarKeys = append(existingVarKeys, v.Key)
	}
	log.Printf("Existing variables key for group %s : %+v", groupKey, existingVarKeys)

	for _, v := range variables {
		existing := Contains(existingVarKeys, v.Key)
		if !existing {
			log.Printf("Create non-existing variable %s", v.Key)
			client.createGroupVariable(groupKey, v)
		} else if existing && v.Overwrite {
			log.Printf("Overwrite existing variable %s (new: %+v)", v.Key, v)
			client.updateGroupVariable(groupKey, v)
		} else {
			log.Printf("No update for value %s (overwrite is false)", v.Key)
		}
	}

	return nil
}

func (client *GitlabClient) createGroupVariable(groupKey string, variable Variable) {
	opt := &gitlab.CreateVariableOptions{
		Key:       gitlab.String(variable.Key),
		Value:     gitlab.String(variable.Value),
		Protected: gitlab.Bool(variable.Protected),
	}
	client.gitClient.GroupVariables.CreateVariable(groupKey, opt)
}

func (client *GitlabClient) updateGroupVariable(groupKey string, variable Variable) {
	opt := &gitlab.UpdateVariableOptions{
		Value:     gitlab.String(variable.Value),
		Protected: gitlab.Bool(variable.Protected),
	}
	client.gitClient.GroupVariables.UpdateVariable(groupKey, variable.Key, opt)
}
