package service

type ArtifactType string

const (
	ArtifactTypeGithubIssue      ArtifactType = "gh_issue"
	ArtifactTypeGithubRepository ArtifactType = "gh_repository"
)

type Artifact struct {
	Id               string       `json:"id"`
	Connection       string       `json:"connection"`
	UniqueIdentifier string       `json:"unique_identifier"`
	Type             ArtifactType `json:"type"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Version          string       `json:"version"`
	Link             string       `json:"link"`
}

type ArtifactService struct {
}

func (service *ArtifactService) CreateOrUpdateArtifacts(artifacts []*Artifact) error {
	return nil
}
