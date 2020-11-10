package api_git

// change package to main and put it inside the main.go file if u need to see
// the running example

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func ListFilesDirectories(Path string) ([]string, error) {
	repo, err := git.PlainOpen(Path)
	var files []string
	if repo != nil {
		treeIter, errW := repo.TreeObjects()

		if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {
				t.Files().ForEach(func(f *object.File) error {
					files = append(files, f.Name)
					return nil
				})
				return nil
			})
		}
		if errW != nil {
			log.Printf("There is not a tree %s", errW)
		}
	}
	return files, err
}
