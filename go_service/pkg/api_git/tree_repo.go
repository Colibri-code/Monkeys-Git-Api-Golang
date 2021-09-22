package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"fmt"

	"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/v5"
)

/*Toma todos los files que tiene el tree Head del repositorio y los
retorna en un []string*/
func ListPathFileRepository(repoPath string) ([]string, error) {

	if repoPath != "" {

		repo, err := git.PlainOpen(repoPath)

		var filepath []string

		ref, err := repo.Head()

		commit, err := repo.CommitObject(ref.Hash())

		tree, err := commit.Tree()

		for _, entry := range tree.Entries {

			filepath = append(filepath, entry.Name)
		}

		return filepath, err
	}

	return nil, git.ErrRepositoryNotExists
}

func TreeData(repoPath string, filepath string) ([]string, error) {

	repo, err := git.PlainOpen(repoPath)

	ref, err := repo.Head()

	var EntryPaths []string

	commit, err := repo.CommitObject(ref.Hash())

	tree, err := commit.Tree()

	if err != nil {

	}

	Tree_entry, err := tree.Tree(filepath)

	for _, entry := range Tree_entry.Entries {

		EntryPaths = append(EntryPaths, entry.Name)
	}
	fmt.Println(Tree_entry)

	if err != nil {
		return nil, object.ErrFileNotFound
	}

	return EntryPaths, err

}
