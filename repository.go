package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
)

type RepositoryClient interface {
	GetContributedTo(ctx context.Context) ([]RepositoryNode, error)
}

type repositoryClient struct {
	GHClient *githubv4.Client
}

func NewRepositoryClient(GHClient *githubv4.Client) RepositoryClient {
	return &repositoryClient{GHClient: GHClient}
}

type RepositoryNode struct {
	ID            githubv4.ID
	NameWithOwner githubv4.String
}

func (r *repositoryClient) GetContributedTo(ctx context.Context) ([]RepositoryNode, error) {
	var query struct {
		Viewer struct {
			RepositoriesContributedTo struct {
				TotalCount githubv4.Int
				Nodes      []RepositoryNode
				PageInfo   struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"repositoriesContributedTo(first: 100, after: $cursor, contributionTypes: [COMMIT], includeUserRepositories: true)"`
		}
	}
	variables := map[string]interface{}{
		"cursor": (*githubv4.String)(nil), // Null after argument to get first page.
	}

	var allRepositories []RepositoryNode
	for {
		err := r.GHClient.Query(ctx, &query, variables)
		if err != nil {
			fmt.Println(err)
			continue
		}

		allRepositories = append(allRepositories, query.Viewer.RepositoriesContributedTo.Nodes...)
		if !query.Viewer.RepositoriesContributedTo.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(query.Viewer.RepositoriesContributedTo.PageInfo.EndCursor)
	}

	return allRepositories, nil
}
