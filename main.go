package main

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type RepositoryStats struct {
	CountCommits      int
	CountLinesAdded   int
	CountLinesRemoved int
	FileStats         map[string]*FileTypeStats
}

type FileTypeStats struct {
	CountLinesAdded   int
	CountLinesRemoved int
}

func main() {
	ctx := context.Background()

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	repositoryClient := NewRepositoryClient(client)
	allRepo, err := repositoryClient.GetContributedTo(ctx)
	fmt.Println(allRepo, err)

	/*acceptedFileExt := []string{"go", "mod"}

	repositories := make(map[string]*RepositoryStats)
	email := "ajanssens@molotov.tv"

	repo, err := git.PlainOpen("/mnt/d/code/Molotov/go-sections-driver")
	if err != nil {
		log.Fatal(err)
	}
	logs, _ := repo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime})
	s := &RepositoryStats{
		FileStats: make(map[string]*FileTypeStats),
	}
	repositories["go-sections-driver"] = s
	err = logs.ForEach(func(commit *object.Commit) error {
		if commit.Author.Email == email && commit.NumParents() == 1 {
			s.CountCommits++
			fstats, err := commit.Stats()
			if err == nil {
				for _, fstat := range fstats {
					split := strings.Split(fstat.Name, ".")
					fileExt := split[len(split)-1]

					if slice.Contains(acceptedFileExt, fileExt) {
						if _, ok := s.FileStats[fileExt]; !ok {
							s.FileStats[fileExt] = &FileTypeStats{}
						}

						s.FileStats[fileExt].CountLinesAdded = s.FileStats[fileExt].CountLinesAdded + fstat.Addition
						s.FileStats[fileExt].CountLinesRemoved = s.FileStats[fileExt].CountLinesRemoved + fstat.Deletion
					}

					s.CountLinesAdded = s.CountLinesAdded + fstat.Addition
					s.CountLinesRemoved = s.CountLinesRemoved + fstat.Deletion

				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repositories)*/
}
