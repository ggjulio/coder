package scmprovider

type ScmProject struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type GitRepository struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type GitBranch struct {
	Id  string `json:"id" validate:"required"`
	Sha string `json:"sha" validate:"required"`
}
