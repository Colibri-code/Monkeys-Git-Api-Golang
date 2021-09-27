package api_git

import (
	"log"

	"github.com/go-git/go-git/v5"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func ListPaths(Path string) ([]string, error) {
	repo, err := git.PlainOpen(Path)
	var file []string
	if repo != nil {
		treeIter, errW := repo.TreeObjects()

		if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {

				// ... get the files iterator and print the file

				t.Files().ForEach(func(f *object.File) error {

					file = append(file, f.Name)

					return nil
				})
				return nil
			})
		}
		if errW != nil {
			log.Printf("There is not a tree %s", errW)
		}
	}
	return file, err
}
