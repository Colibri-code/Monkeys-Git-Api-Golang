package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

		tree.Files().ForEach(func(f *object.File) error {

			filepath = append(filepath, f.Name)

			fmt.Printf("File Hashe and Path: %s    %s\n", f.Hash, f.Name)

			return nil
		})

		return filepath, err
	}

	return nil, git.ErrRepositoryNotExists
}

/*func TreeData(repoPath string, filepath string) ([]string, error) {

	repo, err := git.PlainOpen(repoPath)

	ref, err := repo.Head()

	var EntryName []string

	commit, err := repo.CommitObject(ref.Hash())

	tree, err := commit.Tree()

	if err != nil {

	}

	entry , err :=  tree.FindEntry(filepath)

	if err != nil{
		return nil, object.ErrFileNotFound
	}

	return entry.Name, err

}*/
