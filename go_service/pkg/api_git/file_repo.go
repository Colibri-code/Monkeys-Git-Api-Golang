package api_git

import (
	"fmt"

	"github.com/go-git/go-git/v5"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func ListPaths(Path string) ([]string, error) {
	repo, err := git.PlainOpen(Path)
	var file []string

	/*Obtengo el Tree Head del repositorio*/
	TreeHead := TreeCommitHead(repo)

	fmt.Println(TreeHead)

	if repo != nil {
		//	treeIter, errW := repo.TreeObjects()

		/*Obtengo todos los path de los archivos iniciales (PATH GLOBALES)*/
		TreeHead.Files().ForEach(func(f *object.File) error {

			treeEntry, err := TreeHead.FindEntry(f.Name)

			if err != nil {

			}
			fmt.Println(treeEntry)
			file = append(file, f.Name)

			return nil
		})

		/* 	if treeIter != nil {
			treeIter.ForEach(func(t *object.Tree) error {

				for _, pathentries := range t.Entries {

					fmt.Println(pathentries.Name)
				}
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
		}*/
	}
	return file, err
}
