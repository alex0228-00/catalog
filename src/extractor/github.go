package extractor

import (
	"context"
	"fmt"

	"catalog/src"
	"catalog/src/service"

	"github.com/google/go-github/github"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GitHubExtractor struct {
	connection *service.Connection
	client     *github.Client

	artifactService *service.ArtifactService
}

func NewGitHubExtractor(conn *service.Connection, artifactService *service.ArtifactService) *GitHubExtractor {
	return &GitHubExtractor{
		connection:      conn,
		client:          github.NewClient(conn.Credentials.Auth()),
		artifactService: artifactService,
	}
}

func (ex *GitHubExtractor) Extract(ctx context.Context) error {
	logger := src.DefaultLogger

	logger.Info("Start extraction")
	repos, _, err := ex.client.Repositories.List(ctx, ex.connection.Host, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to query repositories from GitHub")
	}
	logger.Debug(fmt.Sprintf("Fetched %d repositories, preparing fetching details", len(repos)))

	var (
		artifacts     []*service.Artifact
		extractionErr = src.NewMultipleError("Failed to extract repositories")
	)
	for _, repo := range repos {
		repoArtifacts, err := ex.extractRepository(ctx, repo)
		if err != nil {
			extractionErr.Add(err)
		}
		artifacts = append(artifacts, repoArtifacts...)
	}

	logger.Info(fmt.Sprintf("Fetched %d artifacts from Github", len(artifacts)))

	if err := ex.artifactService.CreateOrUpdateArtifacts(artifacts); err != nil {
		extractionErr.Add(errors.Wrap(err, "Failed to create or update artifacts"))
	}

	return extractionErr.ErrOrNil()
}

func (ex *GitHubExtractor) extractRepository(ctx context.Context, repo *github.Repository) ([]*service.Artifact, *ExtractionError) {
	artifact := &service.Artifact{
		Id:               uuid.New().String(),
		UniqueIdentifier: fmt.Sprintf("%d", *repo.ID),
		Connection:       ex.connection.ID,
		Type:             service.ArtifactTypeGithubRepository,
		Name:             *repo.Name,
		Description:      *repo.Description,
		Version:          fmt.Sprintf("%d", repo.UpdatedAt.Unix()),
		Link:             *repo.HTMLURL,
	}
	issues, err := ex.extractIssues(ctx, repo)
	ret := append([]*service.Artifact{artifact}, issues...)
	return ret, err
}

func (ex *GitHubExtractor) extractIssues(ctx context.Context, repo *github.Repository) ([]*service.Artifact, *ExtractionError) {
	var artifacts []*service.Artifact

	issues, _, err := ex.client.Issues.ListByRepo(ctx, *repo.Owner.Login, *repo.Name, nil)
	if err != nil {
		return nil, NewExtractionError(*repo.Name, errors.Wrap(err, "Failed to query issues from GitHub"))
	}

	for _, issue := range issues {
		artifact := &service.Artifact{
			Id:               uuid.New().String(),
			UniqueIdentifier: fmt.Sprintf("%d", *issue.ID),
			Connection:       ex.connection.ID,
			Type:             service.ArtifactTypeGithubIssue,
			Name:             *issue.Title,
			Description:      *issue.Body,
			Version:          fmt.Sprintf("%d", issue.UpdatedAt.Unix()),
			Link:             *issue.HTMLURL,
		}
		artifacts = append(artifacts, artifact)
	}

	return artifacts, nil
}
