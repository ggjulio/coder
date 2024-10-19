package scmprovider

import ()

// type ApiOptions struct {
// 	Void
// }

type ScmProvider interface {
	GetProjects() ([]*ScmProject, error)
	GetRepositories(projectId string) ([]*GitRepository, error)
	GetBranches(projectId string, repositoryId string) ([]*GitBranch, error)
}

// func GetScmProvider(url string) (ScmProvider, error) {
// 	// TODO: may have multiple instances of same kind of providers, check external auth code.
// 	var provider = "bitbucket"
// 	switch provider {
// 	case "bitbucket":
// 		return NewScmProviderBitbucketServer()
// 	default:
// 		return nil, fmt.Errorf("Scm Provider not implemented: %v", provider)
// 	}
// }
