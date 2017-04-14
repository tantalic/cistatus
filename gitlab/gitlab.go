package gitlab

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/plouc/go-gitlab-client"

	"tantalic.com/cistatus"
)

const (
	defaultAPIPath = "/api/v3"
)

type Client struct {
	client *gogitlab.Gitlab
}

// NewClient creates a client suitable for fetching the CI status from a GitLab server
func NewClient(baseUrl, token string) *Client {
	return &Client{
		client: gogitlab.NewGitlab(baseUrl, defaultAPIPath, token),
	}
}

func (g Client) FetchStatus() ([]cistatus.Project, error) {
	var results []cistatus.Project

	projects, err := g.client.Projects()
	if err != nil {
		return results, errors.Wrap(err, "unable to fetch projects")
	}

	for _, project := range projects {

		p := cistatus.Project{
			Name: project.Name,
		}

		projectID := strconv.Itoa(project.Id)

		branches, err := g.client.ProjectBranches(projectID)
		if err != nil {
			return results, errors.Wrapf(err, "unable to fetch branches for %s project", p)
		}

		for _, branch := range branches {

			b := cistatus.Branch{
				Name:   branch.Name,
				Commit: branch.Commit.Id,
			}

			statuses, err := g.client.ProjectCommitStatuses(projectID, branch.Commit.Id)
			if err != nil {
				return results, errors.Wrapf(err, "unable to fetch statuses for %s project, %s branch, %s commit", p, b, b.Commit)
			}

			b.Statuses = make([]cistatus.Status, 0)
			for _, status := range statuses {
				s := cistatus.Status{
					Name:    status.Name,
					Status:  status.Status,
					Created: status.CreatedAt,
					Author:  status.Author.Username,
				}

				b.Statuses = append(b.Statuses, s)
			}

			p.Branches = append(p.Branches, b)
		}
		results = append(results, p)
	}

	return results, nil
}
